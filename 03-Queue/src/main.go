package main

import (
	"fmt"
	"loop-queue"
)

func main() {
	q := lqueue.New(2)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue("x")
	q.Enqueue("y")
	q.Enqueue("z")
	q.Enqueue("f")
	q.Enqueue("g")
	fmt.Println(q.ToString())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

}
