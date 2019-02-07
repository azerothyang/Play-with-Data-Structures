package main

import (
	"array"
	"fmt"
)

func main() {
	arr := array.NewArray(4)
	arr.AddLast(1)
	arr.AddLast(2)
	arr.AddLast(4)
	arr.Add(0, 9)
	fmt.Println(arr.Get(3))
	arr.Set(3, 1130)
	fmt.Println(arr.ToString())
	fmt.Println(arr.Get(3))

	fmt.Println(arr.Contains(11))

	fmt.Println(arr.Find(1130))

	fmt.Println(arr.ToString())
	arr.RemoveElement(13130)
	fmt.Println(arr.ToString())
	arr.AddLast(9420)
	fmt.Println(arr.ToString())
	arr.RemoveElement(9420)
	fmt.Println(arr.ToString())
	arr.RemoveElement(9)
	fmt.Println(arr.ToString())
}
