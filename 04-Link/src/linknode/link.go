package linknode

import "fmt"

type LinkedList struct {
	Node *Node
	Head *Node
	size int
}

type Node struct {
	E    interface{}
	Next *Node
}

func (nd *Node) ToString() string {
	return fmt.Sprintln(nd.E)
}

func New() *LinkedList {
	return &LinkedList{
		nil,
		nil,
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
	link.Node = &Node{
		ele,
		link.Head,
	}
	link.Head = link.Node
	link.size++
}

// add element at index position
func (link *LinkedList) Add(index int, ele interface{}) {
	if index < 0 || index > link.size {
		fmt.Println("add fail, index is illegal")
		return
	}

	if index == 0 {
		link.AddFirst(ele)
		return
	}

	prev := link.Head
	for i := 1; i < index; i++ {
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