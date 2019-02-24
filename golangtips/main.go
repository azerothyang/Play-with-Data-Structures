package main

import (
	"fmt"
	"strings"
)

func main() {
	rd := strings.NewReader("你hello world!")
	fmt.Println(rd.ReadRune())
	fmt.Println(rd.ReadRune())
	fmt.Println(rd.ReadRune())
}
