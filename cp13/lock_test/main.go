package main

import (
	"fmt"
	"sync"
)

/*
锁-资源竞争
1、理论上结果是0
2、实际上情况：1、结果不是0；2、每次结果不相同
*/

var total int
var wg sync.WaitGroup
var lock sync.Mutex

//互斥锁，读写锁，同步数据；能不用锁就不用，性能下降
//绝大多数的web系统来说，都是读多写少
//有1w个人同时读取数据库，A读的时候B能读？为什么要加锁呢 保证数据一致性读和写上加同一把锁
//并发严重下降，B读了一个数据，不会对C读取数据产生影响，一定是写和读之间造成的
func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//先锁上
		lock.Lock()
		total += 1
		lock.Unlock() // 释放
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//先锁上
		lock.Lock()
		total -= 1
		lock.Unlock() // 释放
	}
}
func main() {
	wg.Add(1)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
