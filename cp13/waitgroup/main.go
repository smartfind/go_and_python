package main

import (
	"fmt"
	"sync"
)

//如何解决主的goroutine在子协程结束后自动结束
var wg sync.WaitGroup

//WaitGroup提供了三个很有用的函数
/*
Add
Done
Wait
*/
func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
			//wg.Done()
		}(i)
	}
	//time.Sleep(time.Second * 3)
	wg.Wait()
}
