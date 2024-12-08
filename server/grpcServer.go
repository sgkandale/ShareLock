package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"sharelock/config"
	"sharelock/pkg/helpers"
	"sharelock/pkg/locker"
	"sharelock/pkg/sharelockPB"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type GrpcServer struct {
	sharelockPB.UnimplementedShareLockServiceServer
	listener net.Listener
	port     int
	srv      *grpc.Server
	locker   *locker.Locker
}

func NewGrpcServer(ctx context.Context, cfg *config.Server, locker *locker.Locker) Server {
	mck := NewMockServer()
	if cfg == nil {
		log.Printf("[ERROR] config is nil in server.NewGrpcServer")
		return mck
	}
	if !cfg.Enable {
		return mck
	}
	if cfg.Port == 0 {
		log.Printf("[ERROR] port is 0 in server.NewGrpcServer")
		return mck
	}
	if cfg.TLS {
		log.Printf("[INFO] tls is enabled for grpc server")
		if cfg.CertPath == "" {
			log.Printf("[ERROR] cert path is empty in server.NewGrpcServer")
			return mck
		}
		if cfg.KeyPath == "" {
			log.Printf("[ERROR] key path is empty in server.NewGrpcServer")
			return mck
		}
	}

	var srv *grpc.Server

	if cfg.TLS {
		creds, err := credentials.NewServerTLSFromFile(
			cfg.CertPath,
			cfg.KeyPath,
		)
		if err != nil {
			log.Print("[ERROR] creating grpc server tls : ", err)
			return mck
		}
		srv = grpc.NewServer(grpc.Creds(creds))
	} else {
		srv = grpc.NewServer()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	if err != nil {
		log.Printf("[ERROR] creating tcp listener for grpc on port %d : %s", cfg.Port, err)
		return mck
	}

	grpcServer := &GrpcServer{
		listener: lis,
		port:     cfg.Port,
		srv:      srv,
		locker:   locker,
	}

	sharelockPB.RegisterShareLockServiceServer(srv, grpcServer)

	return grpcServer
}

func (g *GrpcServer) Start() {
	log.Printf("[INFO] starting grpc server on port %d", g.port)
	err := g.srv.Serve(g.listener)
	if err != nil {
		log.Printf("[ERROR] starting grpc server on port %d : %s", g.port, err)
		return
	}
}

func (g *GrpcServer) Stop() {
	log.Print("[INFO] stopping grpc server")
	g.srv.Stop()
}

func (g *GrpcServer) Ping(ctx context.Context, r *sharelockPB.ShareLockPingRequest) (*sharelockPB.ShareLockPingResponse, error) {
	return &sharelockPB.ShareLockPingResponse{
		Message: "pong",
	}, nil
}

func (g *GrpcServer) Lock(ctx context.Context, r *sharelockPB.LockRequest) (*sharelockPB.LockResponse, error) {
	if r == nil {
		return nil, helpers.Err_Srv_NilRequest
	}
	if len(r.Key) == 0 {
		return nil, helpers.Err_Srv_Request_KeyMissing
	}
	md := GetGrpcMetadata(ctx)

	expiryDeadline := time.Now().Add(time.Second * 10)
	if r.TimeoutMs <= 0 {
		expiry, ok := ctx.Deadline()
		if ok {
			expiryDeadline = expiry
		}
	}

	lockerCtx, cancelLockerCtx := context.WithDeadline(ctx, expiryDeadline)
	defer cancelLockerCtx()

	newClient := locker.Client{
		Ctx:        lockerCtx,
		Id:         md.ClientId,
		LockKey:    r.Key,
		StatusChan: make(chan locker.Status),
	}
	g.locker.Lock(&newClient)

listenerLoop:
	for {
		select {
		case status := <-newClient.StatusChan:
			switch status {
			case locker.Status_Locked:
				return &sharelockPB.LockResponse{
					Status: sharelockPB.Status_Acquired,
				}, nil
			case locker.Status_Timeout:
				break listenerLoop
			}
		case <-ctx.Done():
			break listenerLoop
		}
	}

	return &sharelockPB.LockResponse{
		Status: sharelockPB.Status_Timeout,
	}, nil
}

func (g *GrpcServer) Unlock(ctx context.Context, r *sharelockPB.UnlockRequest) (*sharelockPB.UnlockResponse, error) {
	if r == nil {
		return nil, helpers.Err_Srv_NilRequest
	}
	if len(r.Key) == 0 {
		return nil, helpers.Err_Srv_Request_KeyMissing
	}
	md := GetGrpcMetadata(ctx)

	newClient := locker.Client{
		Ctx:        ctx,
		Id:         md.ClientId,
		LockKey:    r.Key,
		StatusChan: make(chan locker.Status),
	}
	g.locker.Unlock(&newClient)

listenerLoop:
	for {
		select {
		case status := <-newClient.StatusChan:
			switch status {
			case locker.Status_Unlocked:
				return &sharelockPB.UnlockResponse{
					Status: sharelockPB.Status_Released,
				}, nil
			case locker.Status_Timeout:
				break listenerLoop
			case locker.Status_UnknownLock:
				return &sharelockPB.UnlockResponse{
					Status: sharelockPB.Status_UnknownLock,
				}, nil
			}
		case <-ctx.Done():
			break listenerLoop
		}
	}

	return &sharelockPB.UnlockResponse{
		Status: sharelockPB.Status_Timeout,
	}, nil
}

type GrpcMetadata struct {
	ClientId string
}

func GetGrpcMetadata(ctx context.Context) GrpcMetadata {
	grpcMetadata := GrpcMetadata{}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		clientId := md.Get("X-Client-Id")
		if len(clientId) > 0 {
			grpcMetadata.ClientId = clientId[0]
		}
	}
	return grpcMetadata
}
