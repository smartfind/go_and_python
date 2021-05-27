package main

import "fmt"

func main() {
	// goto能不用则不用，goto过多，label则过多，整个程序后期维护很麻烦
	// 最容易理解的代码朱行的执行，哪怕多一个函数的调用也是理解上的负担
	// 定义局部变量
	var a int = 10

	//	循环
LOOP:
	for a < 20 {
		if a == 15 {
			//	跳过迭代
			a = a + 1
			goto LOOP
		}
		fmt.Printf("%d\n", a)
		a++
	}
}
