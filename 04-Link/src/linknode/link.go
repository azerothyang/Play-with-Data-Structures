package linknode

import "fmt"

type LinkedList struct {
	DummyHead *Node
	size      int
}

type Node struct {
	E    interface{}
	Next *Node
}

func (link *LinkedList) ToString() string {
	var str string
	node := link.DummyHead
	for node.Next != nil {
		str += fmt.Sprintf("%v -> ", node.Next.E)
		node = node.Next
	}
	return str[:len(str)-4]
}

func New() *LinkedList {
	return &LinkedList{
		&Node{
			nil,
			nil,
		},
		0,
	}
}

func (link *LinkedList) GetSize() int {
	return link.size
}

func (link *LinkedList) IsEmpty() bool {
	return link.size == 0
}

func (link *LinkedList) AddFirst(ele interface{}) {
	newNode := &Node{
		ele,
		link.DummyHead.Next,
	}
	link.DummyHead.Next = newNode
	link.size++
}

// add element at index position
func (link *LinkedList) Add(index int, ele interface{}) {
	if index < 0 || index > link.size {
		fmt.Println("add fail, index is illegal")
		return
	}

	prev := link.DummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	newNode := &Node{
		ele,
		prev.Next,
	}
	prev.Next = newNode
	link.size++
}

func (link *LinkedList) AddLast(ele interface{}) {
	link.Add(link.GetSize(), ele)
}

// get the index node value
func (link *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= link.size {
		return nil
	}
	cur := link.DummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.E
}

// get first node value
func (link *LinkedList) GetFirst() interface{} {
	return link.Get(0)
}

// get last node value
func (link *LinkedList) GetLast() interface{} {
	return link.Get(link.size - 1)
}

// update the index node value
func (link *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= link.size {
		panic("index is out of range")
	}
	cur := link.DummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.E = e
}

// check linkList whether contains e
func (link *LinkedList) Contains(e interface{}) bool {
	cur := link.DummyHead.Next
	for cur != nil {
		if cur.E == e {
			return true
		}
		cur = cur.Next
	}
	return false
}

// remove the index node
func (link *LinkedList) Remove(index int) {
	if index < 0 || index >= link.size {
		panic("index is out of range")
	}
	prev := link.DummyHead
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	link.size--
}
