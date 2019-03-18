package main

import "fmt"

func main() {
	var sum int64 = 0
	var i int64 = 0
	for i <= 4000000000 {
		sum += i
		i++
	}
	fmt.Println(sum)
	//link := linknode.New()
	//link.AddFirst("1")
	//link.AddFirst("2")
	//link.AddFirst("3")
	//link.AddFirst("4")
	//link.Add(1, "x")
	//link.AddLast("z")
	//fmt.Println(link.Node.Next.ToString())
	//fmt.Println(222)
}
