package main

import (
	"fmt"
	"stack"
)

func main() {
	stacker := stack.New(4)
	stacker.Push("1")
	stacker.Push("2")
	stacker.Push("3")
	//for i := 0; i < stacker.Index; i++ {
	//	fmt.Println(stacker.Data[i])
	//}
	//fmt.Println(stacker.Peek())

	fmt.Println(stack.CheckValidParentheses("{}[]"))
}
