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
	link.Add(2, "x")
	link.AddLast("z")
	fmt.Println(link.ToString())
}
