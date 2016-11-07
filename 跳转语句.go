package main

import (
	"fmt"
)

// func main() {
// 	//可以跳出到指定层
// LABEL1:
// 	for {
// 		for i := 0; i < 10; i++ {
// 			if i > 3 {
// 				break LABEL1
// 			}
// 		}
// 	}
// 	fmt.Println("OK")
// }

// func main() {
// 	//可以跳出到指定层

// 	for {
// 		for i := 0; i < 10; i++ {
// 			if i > 3 {
// 				goto LABEL1
// 			}
// 		}
// 	}
// LABEL1:
// 	fmt.Println("OK")
// }

// func main() {
// 	//跳出单次层循环
// LABEL1:
// 	for i := 0; i < 10; i++ {
// 		for {
// 			continue LABEL1 //执行外层有限循环
// 		}
// 	}
// 	fmt.Println("OK")
// }

func main() {
	//这样会死循环,for一次又跳到LABEL1再重新for
LABEL1:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			break LABEL1
		}
	}
}
