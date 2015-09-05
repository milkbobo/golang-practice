package main

import (
	"fmt"
)

// func main() {
// 	a := 1
// 	switch a {
// 	case 0:
// 		fmt.Println("a=0")
// 	case 1:
// 		fmt.Println("a=1")
// 	default:
// 		fmt.Println("None")
// 	}
// }

// func main() {
// 	a := 1
// 	switch {
// 	case a >= 0:
// 		fmt.Println("a=0")
// 		fallthrough //继续往下执行
// 	case a >= 1:
// 		fmt.Println("a=1")
// 	default:
// 		fmt.Println("None")
// 	}
// }

func main() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough //继续往下执行
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
	// fmt.Println(a) //a是局部变量，不能读取出来
}
