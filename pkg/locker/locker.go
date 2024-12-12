package locker

import (
	"context"
	"time"
)

type Locker struct {
	keys          map[string]*KeyHandler
	deleteKeyChan chan string
	lockChan      chan *Client
	unlockChan    chan *Client
}

func NewLocker() *Locker {
	return &Locker{
		keys:          make(map[string]*KeyHandler, 10_000),
		deleteKeyChan: make(chan string, 10_000),
		lockChan:      make(chan *Client, 10_000),
		unlockChan:    make(chan *Client, 10_000),
	}
}

func (l *Locker) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case toRemoveKey := <-l.deleteKeyChan:
			delete(l.keys, toRemoveKey)
		case client := <-l.lockChan:
			if client == nil {
				continue
			}
			keyHandler, exist := l.keys[client.LockKey]
			if exist {
				keyHandler.clientsChan <- client
			} else {
				newKeyHandler := &KeyHandler{
					key:           client.LockKey,
					releaseChan:   make(chan bool, 1),
					clientsChan:   make(chan *Client, 10_000),
					holdingId:     "",
					deleteKeyChan: l.deleteKeyChan,
				}
				newKeyHandler.releaseChan <- true
				l.keys[client.LockKey] = newKeyHandler
				go newKeyHandler.Handle()
				newKeyHandler.clientsChan <- client
			}
		case client := <-l.unlockChan:
			if client == nil {
				continue
			}
			keyHandler, exist := l.keys[client.LockKey]
			if exist {
				if keyHandler.holdingId == client.Id {
					keyHandler.holdingId = ""
					client.StatusChan <- Status_Unlocked
					if len(keyHandler.clientsChan) == 0 {
						l.deleteKeyChan <- keyHandler.key
					}
					keyHandler.releaseChan <- true
					continue
				}
			}
			client.StatusChan <- Status_UnknownLock
		}
	}
}

func (l *Locker) Lock(client *Client) {
	if client == nil {
		return
	}
	if client.StatusChan == nil ||
		client.Id == "" || client.LockKey == "" {
		client.StatusChan <- Status_InvalidData
		return
	}
	l.lockChan <- client
}

func (l *Locker) Unlock(client *Client) {
	if client == nil || client.StatusChan == nil ||
		client.Id == "" || client.LockKey == "" {
		return
	}
	l.unlockChan <- client
}

type KeyHandler struct {
	key               string
	releaseChan       chan bool
	clientsChan       chan *Client
	holdingId         string
	deleteKeyChan     chan string
	clientTimeout     *time.Ticker
	clientTimeoutChan <-chan time.Time
}

func (k *KeyHandler) Handle() {
	defaultTimeTicker := time.NewTicker(time.Minute)

	for {
		select {
		case <-defaultTimeTicker.C:
			if len(k.clientsChan) == 0 {
				k.deleteKeyChan <- k.key
				return
			}
		case <-k.clientTimeoutChan:
			k.releaseChan <- true
		case <-k.releaseChan:
			if k.clientTimeout != nil {
				k.clientTimeout.Stop()
			}
			if len(k.clientsChan) == 0 {
				k.deleteKeyChan <- k.key
				return
			}
			client := <-k.clientsChan
			if client == nil {
				k.releaseChan <- true
				continue
			}
			if client.Ctx.Err() != nil {
				k.releaseChan <- true
				continue
			}
			k.holdingId = client.Id
			k.clientTimeout = time.NewTicker(time.Minute)
			k.clientTimeoutChan = k.clientTimeout.C
			client.StatusChan <- Status_Locked
		}
	}
}
