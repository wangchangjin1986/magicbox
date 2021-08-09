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

	m := routinePool.Job(bb)
	pool.Add(m)

	time.Sleep(1 * time.Second)
	pool.Close()
}

func aa(...interface{}) interface{} {
	fmt.Println("aaaa")
	return nil
}
func bb(...interface{}) interface{} {
	time.Sleep(5 * time.Second)
	fmt.Println("bbbb")
	return nil
}
