package main

import (
	"fmt"
	"magicbox/routinePool"
	"time"
)

func main() {
	pool := routinePool.NewRoutinePool(10)
	pool.Add(routinePool.Job(aa))
	pool.Add(routinePool.Job(aa))
	pool.Add(routinePool.Job(aa))

	fmt.Println(pool.Size())
	// fixpool := routinePool.NewFixRoutinePool(1, 20)
	// f := routinePool.Job(aa)
	// fixpool.Add(f)
	// fixpool.Add(f)
	// fixpool.Add(f)
	time.Sleep(10000000000000)
}

func aa(...interface{}) interface{} {
	time.Sleep(1000000000)

	fmt.Println("aaaa")

	return nil
}
