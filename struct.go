package main

import (
	"fmt"
)

//type person struct{
//	Name string
//	Age int
//}
//func main(){
//	a := person{}
//	fmt.Println(a)//{ 0} string默认空字符串 int默认0
//}

//type person struct {
//	Name string
//	Age  int
//}
//func main() {
//	a := person{}
//	a.Name = "joe"
//	a.Age = 19
//	fmt.Println(a) //{joe 19}
//}

//type person struct {
//	Name string
//	Age  int
//}
//func main() {
//	a := person{
//		Name: "joe",
//		Age:  19,
//	}
//	fmt.Println(a) //{joe 19}
//}

//type person struct{
//	Name string
//	Age int
//}
//func main(){
//	a := person{
//		Name:"joe",
//		Age:19,
//	}
//	fmt.Println(a)
//	A(&a)
//	A(&a)
//	fmt.Println(a)
//}
//func A(per *person){//接收产品名字per ，类型persong
//	per.Age = 13
//	fmt.Println("A",per)
//}
//func B(per *person){//接收产品名字per ，类型persong
//	per.Age = 15
//	fmt.Println("A",per)
//}

//type person struct{
//	Name string
//	Age int
//}
//func main(){
//	a := &person{
//		Name:"joe",
//		Age:19,
//	}
//	a.Name="ok"
//	fmt.Println(a)
//	A(a)
//	A(a)
//	fmt.Println(a)
//}
//func A(per *person){//接收产品名字per ，类型persong
//	per.Age = 13
//	fmt.Println("A",per)
//}
//func B(per *person){//接收产品名字per ，类型persong
//	per.Age = 15
//	fmt.Println("A",per)
//}

//func main(){
//	a := struct{
//		Name string
//		Age int
//	}{
//		Name:"job",
//		Age:19,
//	}
//	fmt.Println(a)//{job 19}
//}

//func main(){
//	a := &struct{
//		Name string
//		Age int
//	}{
//		Name:"job",
//		Age:19,
//	}
//	fmt.Println(a)//&{job 19}
//}

//type person struct{
//	Name string
//	Age int
//	Contact struct{
//		Phone,City string
//	}
//}
//func main(){
//	a := person{}
//	fmt.Println(a)//{ 0 { }}
//}

//type person struct {
//	Name    string
//	Age     int
//	Contact struct {
//		Phone, City string
//	}
//}
//func main() {
//	a := person{Name: "joe", Age: 19}
//	a.Contact.Phone = "1235456456"
//	a.Contact.City = "beijing"
//	fmt.Println(a) //{joe 19 {1235456456 beijing}}
//}

//type person struct {
//	string
//	int
//}
//func main() {
//	a := person{"joe", 19}
//	fmt.Println(a) //{joe 19}
//}

type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{Name: "joe", Age: 19, human: human{Sex: 0}}
	b := student{Name: "joe", Age: 20, human: human{Sex: 1}}
	a.Name = "joe2"
	a.Age = 13
	a.human.Sex = 100
	fmt.Println(a, b)//{{100} joe2 13} {{1} joe 20}
}
