package main

import (
	"fmt"
)

const ()

func main() {
	a := 1
	a++
	var p *int = &a
	fmt.Println(*p)
}
