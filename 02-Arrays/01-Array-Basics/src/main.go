package main

import "fmt"

func main()  {
	var arr [10]int
	for i:=0; i<len(arr); i++ {
		arr[i] = i
	}

	var scores = [3]int{100, 99, 66}
	for i:=0; i<len(scores); i++ {
		fmt.Println(scores[i])
	}

	for i := range scores {
		fmt.Println(scores[i])
	}

	scores[0] = 96

	for i:=0 ;i<len(scores); i++ {
		fmt.Println(scores[i])
	}
}
