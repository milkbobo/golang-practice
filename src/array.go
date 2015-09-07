package main

import (
	"fmt"
)

// func main() {
// 	var a [2]int
// 	var b [2]int
// 	b = a
// 	fmt.Println(b)
// }

// func main() {
// 	var a [2]string
// 	fmt.Println(a)
// }

// func main() {
// 	a := [2]int{1}
// 	fmt.Println(a)
// }

// func main() {
// 	a := [2]int{1, 1}
// 	fmt.Println(a)
// }

// func main() {
// 	a := [20]int{1}
// 	fmt.Println(a)
// }
// [1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

// func main() {
// 	a := [20]int{19: 1}
// 	fmt.Println(a)
// }
//[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]

// func main() {
// 	a := [...]int{1, 2, 3, 4, 5}
// 	fmt.Println(a)
// }
//[1 2 3 4 5]

// func main() {
// 	a := [...]int{0: 1, 1: 2, 2: 3}
// 	fmt.Println(a)
// }
//[1 2 3]

// func main() {
// 	a := [...]int{19: 1}
// 	fmt.Println(a)
// }
// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]

//func main() {
//	a := [...]int{99: 1}
//	var p *[100]int = &a
//	fmt.Println(p)
//}
//&[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]

//func main() {
//	x,y := 1,2
//	a := [...]*int{&x,&y}
//	fmt.Println(a)
//}
//[0x11f2a220 0x11f2a230]

//func main() {
//	a := [2]int{1, 2}
//	b := [2]int{1, 3}
//	fmt.Println(a == b)
//}
//false

//func main() {
//	a := [2]int{1, 2}
//	b := [1]int{1}
//	fmt.Println(a == b)
//}
//不同长度的数据不能比较

//func main() {
//	p := new([10]int)
//	fmt.Println(p)
//}
//&[0 0 0 0 0 0 0 0 0 0]

//func main() {
//	a := new([10]int)
//	a[1] = 2
//	fmt.Println(a)
//	p := new([10]int)
//	p[1] = 2
//	fmt.Println(p)
//}
//&[0 2 0 0 0 0 0 0 0 0]
//&[0 2 0 0 0 0 0 0 0 0]

//func main() {
//	a := [2][3]int{
//		{1, 1, 1},
//		{2, 2, 2}}
//	fmt.Println(a)
//}
//[[1 1 1] [2 2 2]]

//func main() {
//	a := [...][3]int{
//		{1, 1},
//		{2, 2}}
//	fmt.Println(a)
//}
//[[1 1 0] [2 2 0]]

//func main() {
//	a := [...]int{5, 2, 9, 11, 3}
//	item := len(a)
//	for i := 0; i < item; i++ {
//		for j := i + 1; j < item; j++ {
//			if a[i] < a[j] {
//				temp := a[i]
//				a[i] = a[j]
//				a[j] = temp
//			}
//		}
//	}
//	fmt.Println(a);
//}
//[11 9 5 3 2]
