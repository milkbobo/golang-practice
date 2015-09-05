package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	GB
)

func main() {
	fmt.Println(KB)
}
