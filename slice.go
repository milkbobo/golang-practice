package main

import (
	"fmt"
)

//func main() {
//	for i := 0; i < 3; i++ {
//		v := 1
//		fmt.Println(&v)
//	}
//}

//func main(){
//	var s1 []int
//	fmt.Println(s1)
//}
//[]

//func main(){
//	a:=[10]int{}
//	fmt.Println(a)
//	s1 :=a[9]
//	fmt.Println(s1)
//}
//[0 0 0 0 0 0 0 0 0 0]
//0

//func main() {
//	a := [10]int 0，1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Println(a)
//	s1 := a[5:10] //[5 6 7 8 9,] 第5个到第10个
//	fmt.Println(s1)
//}

//func main() {
//	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Println(a)
//	s1 := a[:5] //[0 1 2 3 4] 前5个元素
//	fmt.Println(s1)
//}

//func main() {
//	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Println(a)
//	s1 := a[5:] //[5 6 7 8 9] 最后5个元素
//	fmt.Println(s1)
//}

//func main() {
//	s1 := make([]int, 3, 10)
//	fmt.Println(len(s1),cap(s1))//3 10
//}

//func main(){
//	a := []byte{'a','b','c','d','e','f','g','h','i','j','k'}
//	sa := a[2:5]
//	fmt.Println(string(sa))//cde
//}

//func main(){
//	a := []byte{'a','b','c','d','e','f','g','h','i','j','k'}
//	sa := a[2:5]
//	sb := sa[3:5]
//	fmt.Println(string(sb))//fg
//}

//func main(){
//	a := []byte{'a','b','c','d','e','f','g','h','i','j','k'}
//	sa := a[2:5]
//	fmt.Println(len(sa),cap(sa))//3 9
//	sb := sa[3:5]
//	fmt.Println(len(sb),cap(sb))//2 6
//	fmt.Println(string(sa))//cde
//	fmt.Println(string(sb))//fg
//}

//func main(){
//	a := []byte{'a','b','c','d','e','f','g','h','i','j','k'}
//	sa := a[2:5]
//	fmt.Println(len(sa),cap(sa))//3 9
//	sb := sa[3:5]
//	fmt.Println(len(sb),cap(sb))//2 6
//	fmt.Println(string(sa))//cde
//	fmt.Println(string(sb))//fg
//}

//func main() {
//	a := []int{1, 2, 3, 4, 5}
//	s1 := a[2:5]
//	s2 := a[1:3]
//	fmt.Println(s1, s2)//[3 4 5] [2 3]
//	s2 = append(s2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1)
//	s1[0] = 9
//	fmt.Println(s1, s2)//[9 4 5] [2 3 1 2 1 1 1 1 1 1 1 1]
//}

//func main(){
//	s1 := []int{1,2,3,4,5,6}
//	s2 := []int{7,8,9}

////	copy(s1,s2)
////	fmt.Println(s1)//[7 8 9 4 5 6]
////	fmt.Println(s2)//[7 8 9]

//	copy(s2,s1)
//	fmt.Println(s1)//[1 2 3 4 5 6]
//	fmt.Println(s2)//[1 2 3]
//}

//func main(){
//	s1 := []int{1,2,3,4,5,6}
//	s2 := []int{7,8,9}

//	copy(s1,s2[1:3])
//	fmt.Println(s1)//[8 9 3 4 5 6]
//	fmt.Println(s2)//[7 8 9]

//}

func main(){
	s1 := []int{1,2,3,4,5}
	s2 := s1[:]
	fmt.Println(s2)
}

