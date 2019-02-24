package main

import (
	"fmt"
	"queue"
)

func main() {
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue("x")
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
