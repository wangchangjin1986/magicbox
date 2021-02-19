package main

import (
	"fmt"
	"magicbox/routinePool"
	"time"
)

func main() {
	pool := routinePool.NewFixRoutinePool(1, 20)
	f := routinePool.Job(aa)
	pool.Add(f)
	pool.Add(f)
	pool.Add(f)
	time.Sleep(10000000000000)
}

func aa(...interface{}) interface{} {
	fmt.Println("aaaa")
	return nil
}
