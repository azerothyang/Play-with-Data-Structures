package main

import (
	"fmt"
	"linknode"
)

func main() {
	link := linknode.New()
	link.AddFirst("1")
	link.AddFirst("2")
	link.AddFirst("3")
	link.AddFirst("4")
	link.Add(1, "x")
	link.AddLast("z")
	fmt.Println(link.Node.ToString())
	//fmt.Println(222)
}
