package main

import (
	"fmt"
//	"sort"
)

//func main(){
//	var m map[int]string
//	m = map[int]string{}
//	fmt.Println(m)//map[]
//}

//func main(){
//	var m map[int]string = make(map[int]string)
//	fmt.Println(m)//map[]
//}

//func main(){
//	m := make(map[int]string)
//	fmt.Println(m)//map[]
//}

//func main(){
//	m := make(map[int]string)
//	m[1] = "OK"
//	fmt.Println(m)//map[1:OK]
//}

//func main(){
//	m := make(map[int]string)
//	m[1] = "OK"
//	a := m[1]
//	fmt.Println(a)//OK
//}

//func main(){
//	m := make(map[int]string)
//	m[1] = "OK"
//	delete(m,1)
//	a := m[1]
//	fmt.Println(a)//空白
//}

//func main(){
//	var m map[int]map[int]string
//	m = make(map[int]map[int]string)
//	m[1] = make(map[int]string)
//	m[1][1] = "OK"
//	a := m[1][1]
//	fmt.Println(a)//OK
//}

//func main() {
//	var m map[int]map[int]string
//	m = make(map[int]map[int]string)
//	m[1] = make(map[int]string)
//	a, ok := m[2][1]
//	fmt.Println(a,ok) //false
//}

//func main() {
//	var m map[int]map[int]string
//	m = make(map[int]map[int]string)
//	m[1] = make(map[int]string)
//	a, ok := m[2][1] //多返回值 ，第二个：是否给初始化
//	if !ok {
//		m[2] = make(map[int]string)
//	}
//	m[2][1] = "GOOD"
//	a = m[2][1]
//	fmt.Println(a,ok) //GOOD false
//}

//func main() {
//	var m map[int]map[int]string
//	m = make(map[int]map[int]string)
//	m[1] = make(map[int]string)
//	a, ok := m[2][1] //多返回值 ，第二个：是否给初始化
//	if !ok {
//		m[2] = make(map[int]string)
//	}
//	m[2][1] = "GOOD"
//	a, ok = m[2][1]
//	fmt.Println(a,ok) //GOOD true
//}

//func main(){
//	sm := make([]map[int]string,5)
//	for _,v :=range sm{
//		v = make(map[int]string,1)
//		v[1] = "OK"
//		fmt.Println(v)
//	}
//	fmt.Println(sm)
//}
//map[1:OK]
//map[1:OK]
//map[1:OK]
//map[1:OK]
//map[1:OK]
//[map[] map[] map[] map[] map[]]

//func main(){
//	sm := make([]map[int]string,5)
//	for i :=range sm{
//		sm[i] = make(map[int]string,1)
//		sm[i][0] = "OK"
//		fmt.Println(sm[i])
//	}
//	fmt.Println(sm)
//}
//map[0:OK]
//map[0:OK]
//map[0:OK]
//map[0:OK]
//map[0:OK]
//[map[0:OK] map[0:OK] map[0:OK] map[0:OK] map[0:OK]]

//func main() {
//	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
//	s := make([]int, len(m))
//	i := 0
//	for k, _ := range m {
//		s[i] = k
//		i++
//	}
//	sort.Ints(s)//排序
//	fmt.Println(s)//[1 2 3 4 5]
//}

func main(){
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m1)//map[1:a 2:b 3:c 4:d 5:e]
	
	m2 := make(map[string]int)
	for k,v := range m1{
		m2[v] = k
	}
	fmt.Println(m2)//map[e:5 a:1 b:2 c:3 d:4]
}