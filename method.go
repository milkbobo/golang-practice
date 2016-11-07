package main

import (
	"fmt"
)

//type A struct {
//	Name string
//}
//type B struct {
//	Name string
//}
//func main() {
//	a := A{}
//	a.Print()
//	fmt.Println(a.Name)
//	b := B{}
//	b.Print()
//	fmt.Println(b.Name)
//}
//func (a *A) Print() {
//	a.Name = "AA"
//	fmt.Println("A")
//}
//func (b B) Print() {
//	b.Name = "BB"
//	fmt.Println("B")
//}
//A
//AA
//B

//type TZ int
//func main(){
//	var a TZ
//	a.Print()
//}
//func (a *TZ)Print(){
//	fmt.Println("TZ")
//}

//type TZ int
//func main(){
//	var a TZ
//	a.Print()
//	(*TZ).Print(&a)
//}
//func (a *TZ)Print(){
//	fmt.Println("TZ")
//}

type A struct{
	name string
}
func main(){
	a := A{}
	a.Print()
	fmt.Println(a.name)
}
func (a *A)Print(){
	a.name="123"
	fmt.Println(a.name)
}