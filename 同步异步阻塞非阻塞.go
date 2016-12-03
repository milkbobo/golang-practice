package main

import (
	"fmt"
	// "runtime"
	// "sync"
	"time"
)

func main() {

	startTime := time.Now()
	BlockingSynchronization()
	blockingSynchronizationTime := time.Now()

	// NonBlockingSynchronization()
	NonBlockingSynchronization2()
	nonBlockingSynchronizationTime := time.Now()

	// BlockingAsynchronous()
	// blockingAsynchronousTime := time.Now()
	//
	// NonBlockingAsynchronous()
	// nonBlockingAsynchronousTime := time.Now()

	// fmt.Printf("同步阻塞:%+v\n", time.Duration(blockingSynchronizationTime.UnixNano()-startTime.UnixNano()))
	//
	// fmt.Printf("同步非阻塞:%+v\n", time.Duration(nonBlockingSynchronizationTime.UnixNano()-blockingSynchronizationTime.UnixNano()))

	// fmt.Printf("异步阻塞:%+v\n", time.Duration(blockingAsynchronousTime.UnixNano()-nonBlockingSynchronizationTime.UnixNano()))
	//
	// fmt.Printf("异步非阻塞:%+v\n", time.Duration(nonBlockingAsynchronousTime.UnixNano()-blockingAsynchronousTime.UnixNano()))

}

//煲水
func heatWater() {
	time.Sleep(10 * time.Millisecond)
}

var xiang3 = make(chan int)

func heatWater4() {
	time.Sleep(10 * time.Millisecond)
	xiang3 <- 1
	println("ok")
}

//计算数学题 1加到100
func count(num *int) {

	time.Sleep(10 * time.Millisecond)
	*num = *num + 1

}

//同步阻塞
func BlockingSynchronization() {
	for index := 0; index < 100; index++ {
		heatWater()
	}

	var num *int
	temp := 1
	num = &temp

	for index := 1; index < 100; index++ {
		count(num)
	}
	fmt.Printf("结果结果%+v\n", *num)
}

//同步非阻塞
func calculate() {
	k := 0
	for index := 0; index < 100; index++ {
		k++
	}
	fmt.Printf("结果结果%+v\n", k)
}

func NonBlockingSynchronization2() {

	go heatWater4()

	for {
		select {
		case <-xiang3:
			return
		default:
			calculate()
		}
	}
}

// //同步非阻塞
// func NonBlockingSynchronization() {
//
// 	go func() {
// 		for index := 0; index < 100; index++ {
// 			taoge()
// 			<-xiang
// 		}
// 	}()
//
// 	var num *int
// 	temp := 1
// 	num = &temp
//
// 	for index := 1; index < 100; index++ {
// 		count(num)
// 	}
// 	<-ok
// 	fmt.Printf("结果结果%+v\n", *num)
// }

//异步非阻塞
//涛哥伟哥两个人一起煲水
//目标一个100壶水
var mubiao = 100

//现在已煲数量
var nowNums = 0

//第一壶水的响声
var xiang = make(chan int, 1)

//第二壶水的响声
var xiang2 = make(chan int, 1)

// 第一壶水煲好
var ok = make(chan int, 1)

// 第二壶水煲好
var ok2 = make(chan int, 1)

//水壶1
func heatWater1() {
	time.Sleep(10 * time.Millisecond)
	// println("11")
	xiang <- 1
}

//水壶2
func heatWater2() {
	time.Sleep(10 * time.Millisecond)
	// println("22")
	xiang2 <- 1
}

//涛哥
func taoge() {
	go heatWater1()

	<-xiang
	if mubiao != nowNums {
		nowNums++
		go taoge()
	} else {
		// ok <- 1
		finish = true
	}

}

//伟哥
func weige() {
	go heatWater2()

	<-xiang2
	if mubiao != nowNums {
		nowNums++
		go weige()
	} else {
		// ok2 <- 1
		finish = true
	}

}

//运行方法
var finish bool

func NonBlockingAsynchronous() {
	nowNums = 0
	go taoge()
	go weige()

	// <-ok
	// <-ok2
	// fmt.Printf("煲水:%+v\n", nowNums)

}

//异步阻塞
var num3 = 0

func heatWater3() {
	time.Sleep(10 * time.Millisecond)
	num3++
}

var xiang6 = make(chan int, 1)
var xiang7 = make(chan int, 1)

func taoge2() {
	time.Sleep(time.Second)
	xiang6 <- 1
}

func weige2() {
	time.Sleep(time.Second)
	xiang7 <- 1
}

func BlockingAsynchronous() {
	nowNums = 0
	go taoge2()
	<-xiang6
	go weige2()
	<-xiang7
}
