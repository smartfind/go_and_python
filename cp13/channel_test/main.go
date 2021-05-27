package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func consumer(queue chan int) {
	defer wg.Done()
	// 第一种取法
	//data := <-queue
	//fmt.Println(data)

	//for range取法
	//for data := range queue {
	//	fmt.Println(data)
	//}

	//	for取法
	for {
		data, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(data)
		time.Sleep(time.Second)
	}
}

func main() {
	/*
		channel提供了一种通信机制，定向，类似Python java消息队列
	*/
	//定义一个channel
	var msg chan int
	//	初始化这个channel,两种初始化方式
	//msg = make(chan int)    // 第一种方式：无缓冲，需要先开启消费才能放数据
	msg = make(chan int, 1) // 第二种方式：有缓冲空间
	//	在go语言中，使用make初始化的有三种：1、slice；2、map；3、channel
	msg <- 1 //将1放到channel中
	//msg <- 2      //管道看起来像空间数组，如果满了没被取出再放值就会出现死锁
	//data := <-msg //将channel中的值取出一个
	//fmt.Println(data)

	// 直接内部调用
	//go func(queue chan int) {
	//	data := <-queue
	//	fmt.Println(data)
	//}(msg)
	//time.Sleep(time.Second)

	//	外部函数调用
	wg.Add(1)
	go consumer(msg)
	msg <- 2
	//关闭channel,1、已经关闭的channel不能再发送数据了；关闭消费者可以取到数据
	close(msg)
	wg.Wait()
}
