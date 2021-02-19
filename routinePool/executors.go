package routinePool

import (
	"sync"
)

const (
	CLOSE = iota
)
const (
	DEFAULT_SIZE = 50
)

func NewRoutinePool(size int) Pool {
	pool := Pool{
		size: size,
		wg:   &sync.WaitGroup{},
	}
	pool.start()
	return pool
}
func NewFixRoutinePool(size int, bufferSize int) FixPool {
	pool := FixPool{
		size:        size,
		chanSize:    bufferSize,
		wg:          sync.WaitGroup{},
		job2runChan: make(chan Job, bufferSize),
		commandChan: make(chan int),
	}
	pool.start()
	return pool
}
