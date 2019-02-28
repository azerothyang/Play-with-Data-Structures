package test

import (
	"fmt"
	"loop-queue"
	"queue"
	"testing"
	"time"
)

const operations = 100000

func TestQueue(t *testing.T) {
	startTime := time.Now()
	q := queue.New()
	for i := 0; i < operations; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < operations; i++ {
		q.Dequeue()
	}
	fmt.Printf("queue cost time %v\n", time.Since(startTime))
}

func TestLoopQueue(t *testing.T) {
	startTime := time.Now()
	q := lqueue.New(2)
	for i := 0; i < operations; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < operations; i++ {
		q.Dequeue()
	}
	fmt.Printf("loop queue cost time %v\n", time.Since(startTime))
}
