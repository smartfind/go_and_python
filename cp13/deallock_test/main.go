package main

import "fmt"

func main() {
	var msg chan int
	msg = make(chan int)

	//channel是多个goroutine之间线程安全，使用锁保证
	//如果是没有缓冲的channel，在没有启动一个消费者之前放数据就会报错
	msg <- 1 //放数据到msg中的时，此时会阻塞，阻塞之前会获取一把锁，等到数据被消费后才能释放
	data := <-msg
	fmt.Println(data)
}
