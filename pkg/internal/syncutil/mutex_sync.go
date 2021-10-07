// +build !deadlock

package syncutil

import "sync"

// A Mutex is a mutual exclusion lock.
type Mutex struct {
	sync.Mutex
}

// An RWMutex is a reader/writer mutual exclusion lock.
type RWMutex struct {
	sync.RWMutex
}

type Once struct {
	sync.Once
}

type WaitGroup struct {
	sync.WaitGroup
}

type Map struct {
	sync.Map
}
