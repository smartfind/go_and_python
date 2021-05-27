package main

import (
	"fmt"
	"sync"
	"time"
)

/*
锁-资源竞争
1、理论上结果是0
2、实际上情况：1、结果不是0；2、每次结果不相同
*/

var wg sync.WaitGroup
var rwlock sync.RWMutex

//互斥锁，读写锁，同步数据；能不用锁就不用，性能下降
//绝大多数的web系统来说，都是读多写少
//有1w个人同时读取数据库，A读的时候B能读？为什么要加锁呢 保证数据一致性读和写上加同一把锁
//并发严重下降，B读了一个数据，不会对C读取数据产生影响，一定是写和读之间造成的
func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println("现在开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("数据读取成功")
	rwlock.RUnlock()
}
func write() {
	defer wg.Done()
	rwlock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 10)
	fmt.Println("数据修改成功")
	rwlock.Unlock()
}
func main() {
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go read()
	}
	for i := 0; i < 1; i++ {
		go write()
	}
	wg.Wait()

}
