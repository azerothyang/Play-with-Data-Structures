package main

import (
	"array"
	"fmt"
)

func main() {
	arr := array.NewArray(20)
	arr.AddLast(1)
	arr.AddLast(2)
	arr.AddLast(4)
	arr.Add(0, 9)
	fmt.Println(arr.ToString())
	fmt.Println(arr.Get(3))
	arr.Set(3, 1130)
	fmt.Println(arr.Get(3))
}
