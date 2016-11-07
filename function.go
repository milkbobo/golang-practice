package main

import (
	"fmt"
)

//func main() {
//	A(1, 2, 3, 4, 5, 6, 7)
//}
//func A(a ...int) {
//	fmt.Println(a)//[1 2 3 4 5 6 7]
//}

//func main() {
//	A("1",2, 3, 4, 5, 6, 7)
//}
//func A(b string,a ...int) {
//	fmt.Println(a)//[2 3 4 5 6 7]
//	fmt.Println(b)//1
//}

//func main() {
//	A("1",2, 3, 4, 5, 6, 7)
//}
//func A(b string,a ...int) {
//	fmt.Println(a)//[2 3 4 5 6 7]
//	fmt.Println(b)//1
//}

//func main() {
//	s1 := []int{1, 2, 3, 4}
//	A(s1)
//	fmt.Println(s1)//[5 6 7 8]
//}

//func A(s []int) {
//	s[0] = 5
//	s[1] = 6
//	s[2] = 7
//	s[3] = 8
//	fmt.Println(s)//[5 6 7 8]
//}

//func main() {
//	a := 1
//	A(a)
//	fmt.Println(a)//1
//}
//func A(a int) {
//	a = 2
//	fmt.Println(a)//2
//}

//func main() {
//	a := 1
//	A(&a)
//	fmt.Println(a)//2
//}
//func A(a *int) {
//	*a = 2
//	fmt.Println(*a)//2
//}

//func main() {
//	//匿名函数
//	a := func() {
//		fmt.Println("Func A")
//	}
//	a()
//}

//func main() {
//	f := closure(10)
//	fmt.Println(f(1))//11
//	fmt.Println(f(2))//12
//}
//func closure(x int) func(int) int{//x int是参数,func(int) int 是返回值
//	return func (y int)int {
//		return x + y
//	}
//}

//func main() {
//	f := closure(10)
//	fmt.Println(f(1)) //11
//	fmt.Println(f(2)) //12
//}
//func closure(x int) func(int) int { //x int是参数,func(int) int 是返回值
//	fmt.Println("%p/n", &x)
//	return func(y int) int {
//		fmt.Println("%p/n", &x)
//		return x + y
//	}
//}

//func main(){
//	A()
//	B()
//	C()
//}
//func A(){
//	fmt.Println("Func A")
//}
//func B(){
//	defer func(){
//		if err := recover(); err != nil{
//			fmt.Println("Recover in B")
//		}
//	}()
//	panic("Panic in B")
//}

//func C(){
//	fmt.Println("Func C")
//}

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i =", i)
		defer func() { fmt.Println("defer_closure i +", i) }()
		fs[i] = func() { fmt.Println("closure i =", i) }
	}
	for _, f := range fs {
		f()
	}
}
