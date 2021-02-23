package main

import (
	"fmt"
	"magicbox/routinePool"
	"time"
)

func simplePool() int64 {
	pool := routinePool.NewRoutinePool(10)
	pool.Add(routinePool.Job(aa))
	pool.Add(routinePool.Job(aa))
	pool.Add(routinePool.Job(aa))
	return pool.Size()
}
func fixPool() {
	fixpool := routinePool.NewFixRoutinePool(1, 20)
	f := routinePool.Job(aa)
	fixpool.Add(f)
	fixpool.Add(f)
	fixpool.Add(f)
}

func aa(...interface{}) interface{} {
	time.Sleep(1000000000)

	fmt.Println("aaaa")

	return nil
}
