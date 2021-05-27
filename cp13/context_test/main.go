package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//1、监控全局变量
//var stop bool

//2、通过chan

//var stop chan bool = make(chan bool)

//两种方案不统一，在go1.7中引入了context
//父context
//父context被取消，那么这个父context生成的子context也会被取消
//ctx,cancel := context.WithCancel(context.Background())

//func cpuInfo() {
func cpuInfo(ctx context.Context) {
	defer wg.Done()
	//for {
	//if stop {
	//	break
	//}

	//select {
	//case <-stop:
	//	fmt.Println("退出CPU监控")
	//default:
	//	time.Sleep(time.Second * 2)
	//	fmt.Println("CPU信息读取完成")

	//select {
	//case <-ctx.Done():
	//	fmt.Println("监控退出")
	//	return
	//default:
	//	time.Sleep(time.Second)
	//	fmt.Println("CPU信息读取完成")
	//}

	//父子context测试
	ctx2, _ := context.WithCancel(ctx)
	go memoryInfo(ctx2)
	select {
	case <-ctx.Done():
		fmt.Println("监控退出")
		return
	default:
		time.Sleep(time.Second)
		fmt.Println("CPU信息读取完成")
	}
}

func memoryInfo(ctx context.Context) {
	defer wg.Done()
	for {

		select {
		case <-ctx.Done():
			fmt.Println("监控内存退出")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("内存信息读取完成")
		}
	}
}

func main() {
	//启动一个goroutine去监控某台服务器的CPU的使用情况
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go cpuInfo(ctx)
	//现在希望可以中断CPU的信息读取
	time.Sleep(time.Second * 6)
	//stop = true
	//stop <- true
	cancel()
	wg.Wait()
	fmt.Println("信息监控完成")

}
