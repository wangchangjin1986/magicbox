package main

import (
	"container/heap"
	"fmt"
	"magicbox/structure"
)

func priorityQueueInit() {
	queue := structure.NewPriorityQueue(10)

	queue.Push(&structure.Item{
		Value:    nil,
		Priority: 0,
	})
	queue.Push(&structure.Item{
		Value:    nil,
		Priority: 2,
	})

	queue.Push(&structure.Item{
		Value:    nil,
		Priority: 4,
	})
	queue.Push(&structure.Item{
		Value:    nil,
		Priority: 1,
	})
	queue.Push(&structure.Item{
		Value:    nil,
		Priority: 3,
	})
	heap.Init(&queue)

	fmt.Println(queue.PeekAndShift(10))
	fmt.Println(queue.PeekAndShift(10))
	fmt.Println(queue.PeekAndShift(10))
	fmt.Println(queue.PeekAndShift(10))
	fmt.Println(queue.PeekAndShift(10))

}
func priorityQueuePush() {
	queue := structure.NewPriorityQueue(10)

	heap.Push(&queue, &structure.Item{
		Value:    nil,
		Priority: 0,
	})
	heap.Push(&queue, &structure.Item{
		Value:    nil,
		Priority: 2,
	})
	heap.Push(&queue, &structure.Item{
		Value:    nil,
		Priority: 4,
	})
	heap.Push(&queue, &structure.Item{
		Value:    nil,
		Priority: 1,
	})
	heap.Push(&queue, &structure.Item{
		Value:    nil,
		Priority: 3,
	})
	fmt.Println(queue.PeekAndShift(10))
	fmt.Println(queue.PeekAndShift(10))
	fmt.Println(queue.PeekAndShift(10))
	fmt.Println(queue.PeekAndShift(10))
	fmt.Println(queue.PeekAndShift(10))

}
func main() {
	priorityQueuePush()
}
