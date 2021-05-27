package main

import (
	"fmt"
	"time"
)

func main() {
	/*
			go语言中提供了一个select的功能，作用于channel之上的，多路复用
		select 会随机公平地选择一个case语句执行
		select 的应用场景：1、timeout的超时机制
	*/
	//chan1 := make(chan int, 1)
	//chan2 := make(chan int, 1)
	//chan1 <- 1
	//chan2 <- 2
	//select {
	//case data := <-chan1:
	//	fmt.Println(data)
	//case data := <-chan2:
	//	fmt.Println(data)
	//}

	//timeout := false
	//go func() {
	//	//	该goroutine如果执行时间超过1s，那么就报告给主的goroutine
	//	time.Sleep(time.Second * 2)
	//	timeout = true
	//}()
	//
	//for {
	//	if timeout {
	//		fmt.Println("结束")
	//		break
	//	}
	//	time.Sleep(time.Millisecond * 10)
	//}

	//timeout := make(chan bool, 1)
	//go func() {
	//	time.Sleep(time.Second * 2)
	//	timeout <- true
	//}()
	//fmt.Println(<-timeout)

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 5)
		timeout <- true
	}()

	timeout2 := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 1)
		timeout2 <- true
	}()

	select {
	case <-timeout:
		fmt.Println("timeout超时了")
	case <-timeout2:
		fmt.Println("timeout2超时了")
	default:
		fmt.Println("继续下一次") // 防止被阻塞

	}

	//fmt.Println(<-timeout)
}
