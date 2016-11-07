package main

import (
	"fmt"
	"strconv"
)

type (
	byte int8
	rune int32
	文本   string
)

func main() {
	var a int = 65
	b := strconv.Itoa(a)
	fmt.Println(b)
}
