// +build deadlock

package syncutil

import (
	deadlock "github.com/sasha-s/go-deadlock"
)

// A Mutex is a mutual exclusion lock.
type Mutex struct {
	deadlock.Mutex
}

// An RWMutex is a reader/writer mutual exclusion lock.
type RWMutex struct {
	deadlock.RWMutex
}

type Once struct {
	deadlock.Once
}

type WaitGroup struct {
	deadlock.WaitGroup
}

type Map struct {
	deadlock.Map
}

type Cond struct {
	deadlock.Cond
}
