package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"sharelock/pkg/sharelockPB"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

const (
	baseUrl  = "http://0.0.0.0:50051"
	clientId = "client-id-1"
)

func main() {
	log.Print("[INFO] starting grpc example")
	globalCtx, cancelGlobalCtx := context.WithCancel(context.Background())

	// ping
	ping(globalCtx)

	wg := sync.WaitGroup{}
	lockKey := "test-key-1"

	// acquire lock
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 100) // for consistent order
		lock(globalCtx, lockKey, 1)
	}()

	// acquire lock again
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500) // for consistent order
		lock(globalCtx, lockKey, 2)
	}()

	// unlock lock
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2) // for consistent order
		unlock(globalCtx, lockKey, 1)
	}()

	// unlock lock again
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3) // for consistent order
		unlock(globalCtx, lockKey, 2)
	}()

	wg.Wait()
	cancelGlobalCtx()
}

func ping(ctx context.Context) {
	log.Print("[INFO] pinging sharelock")

	req := &sharelockPB.ShareLockPingRequest{
		Message: "ping",
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		log.Fatal("[ERROR] marshalling ping request : ", err.Error())
	}

	pingCtx, cancelPingCtx := context.WithTimeout(ctx, time.Second*5)
	defer cancelPingCtx()

	httpReq, err := http.NewRequestWithContext(
		pingCtx,
		http.MethodPost,
		baseUrl+"/ping",
		bytes.NewBuffer(reqJson),
	)
	if err != nil {
		log.Fatal("[ERROR] creating ping request : ", err.Error())
	}
	http.Header.Set(httpReq.Header, "X-Client-Id", clientId)

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		log.Fatal("[ERROR] executing ping request : ", err.Error())
	}

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		log.Fatal("[ERROR] reading ping response body : ", err.Error())
	}

	var pingResp sharelockPB.ShareLockPingResponse
	err = json.Unmarshal(respBody, &pingResp)
	if err != nil {
		log.Fatal("[ERROR] unmarshalling ping response body : ", err.Error())
	}
	if pingResp.GetMessage() != "pong" {
		log.Fatal("[ERROR] pinging sharelock, message ", pingResp.GetMessage())
	}
	log.Print("[INFO] sharelock pinged")
}

func lock(ctx context.Context, lockKey string, attemptId int) {
	log.Printf("[INFO] acquiring lock for key %s by attempt %d", lockKey, attemptId)

	req := &sharelockPB.LockRequest{
		Key:       lockKey,
		TimeoutMs: 10_000,
	}
	reqJson, err := json.Marshal(req)
	if err != nil {
		log.Fatal("[ERROR] marshalling lock request : ", err.Error())
	}

	lockCtx, cancellockCtx := context.WithTimeout(ctx, time.Second*5)
	defer cancellockCtx()
	httpReq, err := http.NewRequestWithContext(
		lockCtx,
		http.MethodPost,
		baseUrl+"/lock",
		bytes.NewBuffer(reqJson),
	)
	if err != nil {
		log.Fatal("[ERROR] creating lock request : ", err.Error())
	}
	http.Header.Set(httpReq.Header, "X-Client-Id", clientId)

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		log.Fatal("[ERROR] executing lock request : ", err.Error())
	}

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		log.Fatal("[ERROR] reading lock response body : ", err.Error())
	}
	var lockResp sharelockPB.LockResponse
	err = json.Unmarshal(respBody, &lockResp)
	if err != nil {
		log.Fatal("[ERROR] unmarshalling lock response body : ", err.Error())
	}
	if lockResp.GetStatus() != sharelockPB.Status_Acquired {
		log.Fatal("[ERROR] acquiring lock with sharelock, status ", lockResp.GetStatus())
	}
	log.Printf("[INFO] lock acquired for key %s by attempt %d", lockKey, attemptId)
}

func unlock(ctx context.Context, lockKey string, attemptId int) {
	log.Printf("[INFO] releasing lock for key %s by attempt %d", lockKey, attemptId)

	req := &sharelockPB.UnlockRequest{
		Key: lockKey,
	}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		log.Fatal("[ERROR] marshalling unlock request : ", err.Error())
	}

	unlockCtx, cancelUnlockCtx := context.WithTimeout(ctx, time.Second*5)
	defer cancelUnlockCtx()

	httpReq, err := http.NewRequestWithContext(
		unlockCtx,
		http.MethodPost,
		baseUrl+"/unlock",
		bytes.NewBuffer(jsonReq),
	)
	if err != nil {
		log.Fatal("[ERROR] creating unlock request : ", err.Error())
	}
	http.Header.Set(httpReq.Header, "X-Client-Id", clientId)

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		log.Fatal("[ERROR] executing unlock request : ", err.Error())
	}

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		log.Fatal("[ERROR] reading unlock response body : ", err.Error())
	}
	var unlockResp sharelockPB.UnlockResponse
	err = json.Unmarshal(respBody, &unlockResp)
	if err != nil {
		log.Fatal("[ERROR] unmarshalling unlock response body : ", err.Error())
	}
	if unlockResp.GetStatus() != sharelockPB.Status_Released {
		log.Fatal("[ERROR] releasing lock with sharelock, status ", unlockResp.GetStatus())
	}
	log.Printf("[INFO] lock released for key %s by attempt %d", lockKey, attemptId)
}
