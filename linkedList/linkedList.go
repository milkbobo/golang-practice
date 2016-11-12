package main

import (
	"errors"
	"fmt"
)

//定义节点
type Node struct {
	Data int
	Prev *Node
	Next *Node
}

func main() {

	firstData := &Node{}
	lastOneData := &Node{}
	var err error

	// firstData := Initialize()

	//增
	for index := 0; index < 10; index++ {
		firstData, lastOneData, err = add(firstData, lastOneData, index)
		if err != nil {
			panic(err)
		}
	}

	//遍历查看
	Look(firstData)
	fmt.Printf("firstData:%+v\n", firstData)
	fmt.Printf("lastOneData:%+v\n", lastOneData)

	// //查
	targerData, err := Find(*firstData, 9)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查9:%+v\n\n", targerData)

	//删除
	err = Del(firstData, 8)
	if err != nil {
		panic(err)
	}
	Look(firstData)
	fmt.Printf("删除下标为8\n\n")

	//改
	err = Mod(firstData, 101, 9)
	if err != nil {
		panic(err)
	}
	Look(firstData)
	fmt.Printf("改101\n")

}

func Initialize() *Node {
	// 随机生成数字出来

	nums := []int{1, 2, 3, 4, 5, 11, 12, 13, 14, 15, 21, 22, 23}

	//第一个节点
	firstData := &Node{}
	//上一个节点
	lastData := &Node{}

	for singleKey, singeData := range nums {

		if singleKey == 0 {
			firstData.Data = singeData
			lastData = firstData
			continue
		}

		nowData := Node{
			Prev: lastData,
		}
		nowData.Data = singeData

		lastData.Next = &nowData

		// if len(nums)-1 == singleKey {
		// 	firstData.Prev = &nowData
		// 	nowData.Next = &firstData
		// }

		lastData = &nowData

	}

	return firstData

}

func Look(data *Node) {

	i := data

	for {
		fmt.Printf("%+v\n", i)
		if i.Next == nil {
			return
		}
		// fmt.Printf("%+v\n", i)
		i = i.Next
		// fmt.Printf("%+v\n", startNode)
	}
}

//查询方法
func Find(data Node, findNum int) (*Node, error) {
	nowData := data
	for {

		if nowData.Data == findNum {
			return &nowData, nil
		}

		if nowData.Next == nil {
			return nil, errors.New("找不到该数字")
		}

		nowData = *nowData.Next

	}

	return nil, errors.New("找不到该数字")

}

// //增加方法
// //among 放在第几个后面
// func add(data *Node, addNum int, among int) error {
// 	nowData := data
// 	nums := 1
// 	// lastData := &Node{}
// 	for {
// 		if nums != among {
// 			if &nowData.Next == nil {
// 				return errors.New("没有该among位置")
// 			}
// 			nums++
// 			nowData = nowData.Next
// 			continue
// 		}
//
// 		newNode :=
// 			Node{
// 				Data: addNum,
// 				Prev: nowData.Prev,
// 				Next: nowData.Next,
// 			}
//
// 		nowData.Next = &newNode
//
// 		return nil
//
// 	}
//
// 	return errors.New("找不到该数字")
//
// }

//增加方法
func add(firstData *Node, lastOneData *Node, addNum int) (*Node, *Node, error) {

	// fmt.Printf("%+v", firstData)

	nowData := &Node{
		Data: addNum,
	}

	if firstData.Prev == nil && firstData.Next == nil {
		nowData.Next = nowData
		return nowData, nowData, nil
	}

	lastOneData.Next = nowData
	nowData.Prev = lastOneData
	lastOneData = nowData
	// lastOneData.Prev =

	return firstData, lastOneData, nil

}

//修改方法
//among 修改第几个
func Mod(data *Node, modNum int, among int) error {
	nowData := data
	nums := 1
	// lastData := &Node{}
	for &nowData.Next != nil {
		if nums != among {
			if nowData.Next == nil {
				return errors.New("没有该among位置")
			}
			nums++
			nowData = nowData.Next
			continue
		}

		nowData.Data = modNum
		return nil

	}
	return errors.New("没有该among位置")
}

//删除方法
//among 修改第几个
func Del(data *Node, DelAmong int) error {
	nowData := data
	nums := 1

	for &nowData.Next != nil {
		if nums != DelAmong {
			if nowData.Next == nil {
				return errors.New("没有该among位置")
			}

			nums++
			nowData = nowData.Next
			continue
		}

		*nowData.Prev.Next = *nowData.Next

		return nil

	}

	return errors.New("没有该among位置")

}
