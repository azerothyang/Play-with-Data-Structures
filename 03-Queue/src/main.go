package main

import (
	"fmt"
	"loop-queue"
)

func main() {
	q := lqueue.New(4)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue("x")
	q.Enqueue("y")
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

}
