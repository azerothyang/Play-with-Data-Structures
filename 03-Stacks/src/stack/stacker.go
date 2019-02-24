package stack

type stacker interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	GetIndex() int
	IsEmpty() bool
}
