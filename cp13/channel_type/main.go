package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//queue chan <- int 只能放数据进去
func consumer(queue <-chan int) { //queue <- chan int 只可以取值，不能放；
	defer wg.Done()
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
	//缓冲：1、有缓冲；2、无缓冲
	//双向的还是单向的，为了安全还提供了单向channel
	//var msg chan <- int // 只能放值进去
	//var msg <- chan int // 只能取值
	var msg chan int
	//	初始化这个channel,两种初始化方式
	msg = make(chan int, 1) // 第二种方式：有缓冲空间
	//	在go语言中，使用make初始化的有三种：1、slice；2、map；3、channel
	msg <- 1 //将1放到channel中

	wg.Add(1)
	go consumer(msg) //普通的channel可以直接转换成单向的channel
	msg <- 2
	//关闭channel,1、已经关闭的channel不能再发送数据了；关闭消费者可以取到数据
	close(msg)
	wg.Wait()
}
