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
	link.Set(2, "3")
	fmt.Println(link.ToString())
	fmt.Println("remove element:" + link.Remove(1).E.(string))
	fmt.Println(link.ToString())
	fmt.Println("remove element:" + link.RemoveMethod2(2).E.(string))
	fmt.Println(link.ToString())
	fmt.Println(link.GetLast())
	fmt.Println(link.GetFirst())
	fmt.Println(link.Contains("hei yo"))
}
