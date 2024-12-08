package locker

import "context"

type Status int

const (
	Status_Unlocked Status = iota
	Status_Locked
	Status_Timeout
	Status_UnknownLock
)

type Client struct {
	Ctx        context.Context
	Id         string
	LockKey    string
	StatusChan chan Status
}
