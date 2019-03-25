package linknode

import "fmt"

type LinkedList struct {
	DummyHead *Node
	size int
}

type Node struct {
	E    interface{}
	Next *Node
}

func (ll *LinkedList) ToString() string {
	var str string
	node := ll.DummyHead
	for node.Next != nil {
		str += fmt.Sprintf("%v -> ", node.Next.E)
		node = node.Next
	}
	return str[:len(str) - 4]
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