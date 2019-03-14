package main

import "fmt"

func main() {
	var arr = []int{3, 4, 5, 1, 2, 56, 45, 1, 23, 4, 6, 7, 5}
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	num := arr[0]
	var left []int
	var mid int
	var right []int
	for _, v := range arr {
		if v > num {
			right = append(right, v)
		} else if v == num {
			mid = num
		} else {
			left = append(left, v)
		}
	}
	left = quickSort(left)
	left = append(left, mid)
	right = quickSort(right)
	left = append(left, right...)
	return left
}
