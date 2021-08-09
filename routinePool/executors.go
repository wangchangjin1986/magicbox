package routinePool

import (
	"magicbox/msync"
	"magicbox/util"
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
		wg:   msync.New(),
	}
	pool.start()
	return pool
}
func NewFixRoutinePool(size int, bufferSize int) FixPool {
	pool := FixPool{
		size:        size,
		chanSize:    bufferSize,
		wg:          msync.New(),
		job2runChan: make(chan Job, bufferSize),
		commandChan: make(chan int),
	}
	pool.start()
	return pool
}

func NewCachedRoutinePool() Pool {
	pool := Pool{
		size: util.INT_MAX,
		wg:   msync.New(),
	}
	pool.start()
	return pool
}
