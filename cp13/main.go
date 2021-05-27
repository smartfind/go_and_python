package main

import (
	"fmt"
	"time"
)

//goroutine
//python java c++多线程和多进程编程
//go 诞生的比较晚，web2.0开发主键主流，高并发
//多线程-每个线程占用的内存比较多，而且系统切换的开销比较大
//go语言一开始的时候就没打算让我们去实例化一个线程-协程中 go没有历史包袱
//nodejs python3.6中都引入了协程
//python中有两种编程模式：1.多线程和多进程来进行并发编程；2、使用协程进行并发编程，很多库是基于多线程和多进程开发的
//除非某一天所有库，大部分的都支持协程才能和go 相比

func p() {
	fmt.Println("协程测试函数")
}

func main() {
	//go p()
	//匿名函数
	//go func() {
	//	fmt.Println("协程测试函数")
	//}()
	//for i := 0; i < 100; i++ {
	//	闭包
	//	go func() {
	//		for {
	//			fmt.Println(i)
	//			time.Sleep(time.Second)
	//		}
	//	}()
	//}

	for i := 0; i < 100; i++ {
		go func(n int) {
			for {
				fmt.Println(n)
				time.Sleep(time.Second)
			}
		}(i)
	}
	//主死从随
	time.Sleep(time.Second * 2)
	//fmt.Println("哈哈")

}
