package main

import "fmt"

func add(params ...int) (sum int) {
	//	解决不确定参的的传递
	for _, v := range params {
		sum += v
	}
	return sum
}

func filter(score []int) []int {
	reSlice := make([]int, 0)
	for _, v := range score {
		if v >= 60 {
			reSlice = append(reSlice, v)
		}
	}
	return reSlice
}

//  函数中以函数作为参数
func filter2(score []int, f func(int) bool) []int {
	reSlice := make([]int, 0)
	for _, v := range score {
		if f(v) {
			reSlice = append(reSlice, v)
		}
	}
	return reSlice
}

// 省略号用法
func main() {
	//	通过省略号去动态设置多个参数值
	fmt.Println(add(1, 3, 5, 6))
	//	这种效果区别于slice，[]int，slice是引用传递，不是值传递
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(add(slice1...)) // 将slice打散
	// 省略号的用途：1、函数不定长传参；2、将slice打散
	arr := [...]int{1, 2, 3}
	fmt.Printf("%T", arr)
	//	go语言中中非常重要的特性，函数的一些特性，函数是一等公民特性，可以作为参数、返回值、赋值给另一个变量
	myfunc := add
	fmt.Printf("%T\n", myfunc)
	// 匿名函数，函数内部定义函数
	myfunc2 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("%T\n", myfunc2)
	fmt.Println(myfunc2(1, 2))
	// 函数定义的同时调用
	myfunc3 := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(myfunc3)
	//  函数当做参数
	fmt.Println(func(a, b int) int {
		return a + b
	}(1, 2))
	//	将函数作为另一个函数的参数
	//  写一个函数用于过滤一部分数据
	score := []int{10, 50, 80, 90, 85}
	fmt.Println(filter(score))
	fmt.Println(filter2(score, func(a int) bool {
		if a >= 60 {
			return true
		} else {
			return false
		}
	}))
}
