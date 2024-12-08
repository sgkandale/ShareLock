package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"sharelock/config"
	"sharelock/pkg/locker"
	"sharelock/pkg/sharelockPB"
)

type HttpServer struct {
	port   int
	mux    *http.ServeMux
	locker *locker.Locker
}

func NewHttpServer(ctx context.Context, cfg *config.Server, locker *locker.Locker) Server {
	log.Print("[INFO] creating http server")
	mck := NewMockServer()
	if cfg == nil {
		log.Printf("[ERROR] config is nil in server.NewHttpServer")
		return mck
	}
	if !cfg.Enable {
		return mck
	}
	if cfg.Port == 0 {
		log.Printf("[ERROR] port is 0 in server.NewHttpServer")
		return mck
	}

	httpServer := &HttpServer{
		port:   cfg.Port,
		locker: locker,
	}

	srv := http.NewServeMux()
	srv.HandleFunc("/ping", httpServer.Ping)
	srv.HandleFunc("/lock", httpServer.Lock)
	srv.HandleFunc("/unlock", httpServer.Unlock)
	httpServer.mux = srv
	return httpServer
}

func (h *HttpServer) Start() {
	log.Printf("[INFO] starting http server on port %d", h.port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), h.mux)
	if err != nil {
		log.Printf("[ERROR] starting http server on port %d : %s", h.port, err)
		return
	}
}

func (h *HttpServer) Stop() {
	log.Print("[INFO] stopping http server")
}

func (h *HttpServer) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := &sharelockPB.ShareLockPingResponse{
		Message: "pong",
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Print("[ERROR] json marshalling ping response in http server : ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(jsonResp)
}

func (h *HttpServer) Lock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r == nil || r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	respEncoder := json.NewEncoder(w)

	req := sharelockPB.LockRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print("[ERROR] reading request body in httpServer.Lock : ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expiryDeadline := time.Now().Add(time.Second * 10)
	if req.TimeoutMs <= 0 {
		expiry, ok := r.Context().Deadline()
		if ok {
			expiryDeadline = expiry
		}
	}

	lockerCtx, cancelLockerCtx := context.WithDeadline(r.Context(), expiryDeadline)
	defer cancelLockerCtx()

	newClient := locker.Client{
		Ctx:        lockerCtx,
		Id:         r.Header.Get("X-Client-Id"),
		LockKey:    req.Key,
		StatusChan: make(chan locker.Status),
	}
	h.locker.Lock(&newClient)

listenerLoop:
	for {
		select {
		case status := <-newClient.StatusChan:
			switch status {
			case locker.Status_Locked:
				resp := &sharelockPB.LockResponse{
					Status: sharelockPB.Status_Acquired,
				}
				respEncoder.Encode(resp)
				return
			case locker.Status_Timeout:
				break listenerLoop
			}
		case <-r.Context().Done():
			break listenerLoop
		}
	}

	w.WriteHeader(http.StatusRequestTimeout)
	respEncoder.Encode(
		&sharelockPB.UnlockResponse{Status: sharelockPB.Status_Timeout},
	)
}

func (h *HttpServer) Unlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r == nil || r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	respEncoder := json.NewEncoder(w)

	req := sharelockPB.UnlockRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print("[ERROR] reading request body in httpServer.Unlock : ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newClient := locker.Client{
		Ctx:        r.Context(),
		Id:         r.Header.Get("X-Client-Id"),
		LockKey:    req.Key,
		StatusChan: make(chan locker.Status),
	}
	h.locker.Unlock(&newClient)

listenerLoop:
	for {
		select {
		case status := <-newClient.StatusChan:
			switch status {
			case locker.Status_Unlocked:
				respEncoder.Encode(
					&sharelockPB.UnlockResponse{
						Status: sharelockPB.Status_Released,
					},
				)
				return
			case locker.Status_Timeout:
				break listenerLoop
			case locker.Status_UnknownLock:
				respEncoder.Encode(
					&sharelockPB.UnlockResponse{
						Status: sharelockPB.Status_UnknownLock,
					},
				)
				w.WriteHeader(http.StatusNotFound)
				return
			}
		case <-r.Context().Done():
			break listenerLoop
		}
	}

	w.WriteHeader(http.StatusRequestTimeout)
	respEncoder.Encode(
		&sharelockPB.UnlockResponse{Status: sharelockPB.Status_Timeout},
	)
}
