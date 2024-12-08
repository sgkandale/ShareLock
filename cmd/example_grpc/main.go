package main

import (
	"context"
	"log"
	"sync"
	"time"

	"sharelock/pkg/sharelockPB"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var client sharelockPB.ShareLockServiceClient

func main() {
	log.Print("[INFO] starting grpc example")
	globalCtx, cancelGlobalCtx := context.WithCancel(context.Background())

	grpcOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	log.Print("[INFO] creating grpc client")
	cc, err := grpc.NewClient("0.0.0.0:50052", grpcOptions...)
	if err != nil {
		log.Fatalf("[ERROR]: dialing grpc server : %s", err.Error())
	}

	client = sharelockPB.NewShareLockServiceClient(cc)

	mdCtx := metadata.NewOutgoingContext(
		globalCtx,
		metadata.Pairs(
			"X-Client-Id", "client-id-1",
		),
	)

	// ping
	log.Print("[INFO] pinging sharelock")
	pingCtx, cancelPingCtx := context.WithTimeout(mdCtx, time.Second*5)
	defer cancelPingCtx()
	_, err = client.Ping(
		pingCtx,
		&sharelockPB.ShareLockPingRequest{Message: "ping"},
	)
	if err != nil {
		log.Fatal("[ERROR] pinging sharelock : ", err.Error())
	}

	wg := sync.WaitGroup{}
	lockKey := "test-key-1"

	// acquire lock
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 100) // for consistent order
		lock(mdCtx, lockKey, 1)
	}()

	// acquire lock again
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500) // for consistent order
		lock(mdCtx, lockKey, 2)
	}()

	// unlock lock
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2) // for consistent order
		unlock(mdCtx, lockKey, 1)
	}()

	// unlock lock again
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3) // for consistent order
		unlock(mdCtx, lockKey, 2)
	}()

	wg.Wait()
	cancelGlobalCtx()
}

func lock(ctx context.Context, lockKey string, attemptId int) {
	log.Printf("[INFO] acquiring lock for key %s by attempt %d", lockKey, attemptId)
	lockCtx, cancellockCtx := context.WithTimeout(ctx, time.Second*5)
	defer cancellockCtx()
	lockResp, err := client.Lock(
		lockCtx,
		&sharelockPB.LockRequest{
			Key:       lockKey,
			TimeoutMs: 10000, // timeout to wait till lock is available
		},
	)
	if err != nil {
		log.Fatal("[ERROR] acquiring lock with sharelock : ", err.Error())
	}
	if lockResp.GetStatus() != sharelockPB.Status_Acquired {
		log.Fatal("[ERROR] acquiring lock with sharelock, status ", lockResp.GetStatus())
	}
	log.Printf("[INFO] lock acquired for key %s by attempt %d", lockKey, attemptId)
}

func unlock(ctx context.Context, lockKey string, attemptId int) {
	log.Printf("[INFO] releasing lock for key %s by attempt %d", lockKey, attemptId)
	unlockCtx, cancelUnlockCtx := context.WithTimeout(ctx, time.Second*5)
	defer cancelUnlockCtx()
	unlockResp, err := client.Unlock(
		unlockCtx,
		&sharelockPB.UnlockRequest{
			Key: lockKey,
		},
	)
	if err != nil {
		log.Fatal("[ERROR] releasing lock with sharelock : ", err.Error())
	}
	if unlockResp.GetStatus() != sharelockPB.Status_Released {
		log.Fatal("[ERROR] releasing lock with sharelock, status ", unlockResp.GetStatus())
	}
	log.Printf("[INFO] lock released for key %s by attempt %d", lockKey, attemptId)
}
