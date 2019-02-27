package queue

import "array"

type QueueI interface {
	GetIndex() int
	Enqueue(interface{})
	Dequeue() interface{}
	IsEmpty() bool
	GetFront() interface{}
}

func New() *Queue {
	return &Queue{
		Data: array.NewArray(2),
	}
}

type Queue struct {
	Data *array.Array
}

func (q *Queue) GetIndex() int {
	return q.Data.GetSize()
}

func (q *Queue) Enqueue(ele interface{}) {
	q.Data.AddLast(ele)
}

func (q *Queue) Dequeue() interface{} {
	return q.Data.RemoveFirst()
}

func (q *Queue) IsEmpty() bool {
	return q.Data.GetSize() == 0
}

func (q *Queue) GetFront() interface{} {
	return q.Data.Get(0)
}
