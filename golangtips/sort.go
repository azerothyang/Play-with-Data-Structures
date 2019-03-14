package main

import "fmt"

func main() {
	var arr = []int{3, 4, 5, 1, 2, 56, 45, 1, 23, 4, 6, 7, 5}
	fmt.Println(quickSort(arr))
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	//冒泡排序的核心是相邻元素比较,判断小的放左边，大的放右边
	for range arr {
		flag := false
		for k := 0; k < length-1; k++ {
			if arr[k+1] < arr[k] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
				flag = true
			}
		}
		if flag == false {
			//当某次循环出现一次交换位置都没有发生，则结束
			break
		}
	}
	return arr
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	num := arr[0]
	var left []int
	var mid []int
	var right []int
	for _, v := range arr {
		if v > num {
			right = append(right, v)
		} else if v == num {
			mid = append(mid, num)
		} else {
			left = append(left, v)
		}
	}
	left = quickSort(left)
	left = append(left, mid...)
	right = quickSort(right)
	left = append(left, right...)
	return left
}
