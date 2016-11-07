package main

import (
	"errors"
	"fmt"
)

//定义节点
type Node struct {
	Data int
	Next *Node
}

func main() {

	// 随机生成数字出来

	nums := []int{1, 2, 3, 4, 5, 11, 12, 13, 14, 15, 21, 22, 23}

	//第一个节点
	firstData := Node{}
	//上一个节点
	lastData := &Node{}

	for singleKey, singeData := range nums {

		nowData := Node{}
		nowData.Data = singeData
		lastData.Next = &nowData

		if singleKey == 0 {
			firstData.Next = &nowData
		}

		lastData = &nowData

	}

	//遍历查看
	look(firstData)

	//查
	targerData, err := find(firstData, 12)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查12:%+v\n", targerData)

	//删除
	err = del(&firstData, 11)
	if err != nil {
		panic(err)
	}

	//增
	err = add(&firstData, 88, 3)
	if err != nil {
		panic(err)
	}
	look(firstData)
	fmt.Printf("增88\n")

	//改
	err = mod(&firstData, 101, 9)
	if err != nil {
		panic(err)
	}
	look(firstData)
	fmt.Printf("改101\n")

}

func look(data Node) {
	i := data
	fmt.Printf("%+v\n", i)
	for i.Next != nil {
		// fmt.Printf("%+v\n", i)
		i = *i.Next
		fmt.Printf("%+v\n", i)
	}
}

//查询方法
func find(data Node, findNum int) (*Node, error) {
	nowData := data
	for nowData.Next != nil {
		if nowData.Data == findNum {
			return &nowData, nil
		} else {
			if nowData.Next == nil {
				return nil, errors.New("找不到该数字")
			}
			nowData = *nowData.Next
		}
	}
	return &Node{}, errors.New("找不到该数字")
}

//增加方法
//among 放在第几个后面
func add(data *Node, addNum int, among int) error {
	nowData := data
	nums := 1
	// lastData := &Node{}
	for &nowData.Next != nil {
		if nums == among {
			newNode :=
				Node{
					Data: addNum,
					Next: nowData.Next,
				}
			nowData.Next = &newNode
			// lastData.Next = &newNode

			return nil
		} else {
			nowData = nowData.Next
		}

		if nowData.Next == nil {
			return errors.New("没有该among位置")
		}

		nums++
	}
	return errors.New("没有该among位置")
}

//修改方法
//among 修改第几个
func mod(data *Node, modNum int, among int) error {
	nowData := data
	nums := 1
	// lastData := &Node{}
	for &nowData.Next != nil {
		if nums == among {
			nowData.Data = modNum

			return nil
		} else {
			nowData = nowData.Next
		}

		if nowData.Next == nil {
			return errors.New("没有该among位置")
		}

		nums++
	}
	return errors.New("没有该among位置")
}

//删除方法
//among 修改第几个
func del(data *Node, DelAmong int) error {
	nowData := data
	lastData := &Node{}
	nums := 1

	for &nowData.Next != nil {
		if nums == DelAmong {
			lastData.Next = nowData.Next

			return nil
		} else {
			nowData = nowData.Next
		}

		lastData = nowData

		if nowData.Next == nil {
			return errors.New("没有该among位置")
		}

		nums++
	}

	return errors.New("没有该among位置")

}
