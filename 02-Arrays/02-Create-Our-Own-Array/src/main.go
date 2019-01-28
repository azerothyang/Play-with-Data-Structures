package main

import (
	"array"
	"fmt"
)

func main() {
	arr := array.NewArray(4)
	arr.AddLast(1)
	arr.AddLast(2)
	arr.AddLast(3)
	arr.Add(0, 9)
	fmt.Println(arr.ToString())
}
