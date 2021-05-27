package main

import "fmt"

//func main() {
//	// defer只能是函数调用不能是表达式，比如a++
//	//f, err := os.Open("xxx.text")
//	//f, _ := os.Open("xxx.text")
//	////if err != nil
//	//defer f.Close()
//	defer fmt.Println("defer test1")
//	defer fmt.Println("defer test2")
//	defer fmt.Println("defer test3")
//	/*defer 语句是go语言的一种用于注册延迟调用的机制，它可以让当前函数执行完毕之后执行
//	对于Python的with语句来说，有类似的功能
//	*/
//	//	当同时出现多个defer调用时，每一个defer语句都是先压入一个栈中，先进后出的原则
//
//}

//func main() {
//	//	defer语句执行时的拷贝机制
//	test := func() {
//		fmt.Println("test1")
//	}
//	defer test()
//	test = func() {
//		fmt.Println("test2")
//	}
//}

//func main() {
//	//	defer语句执行时的拷贝机制
//	x := 10
//	defer func(a int) {
//		fmt.Println(a)
//	}(x)
//	x++
//	fmt.Println(x)
//}

//func main() {
//	//	defer语句执行时的拷贝机制
//	x := 10
//	defer func(a *int) {
//		fmt.Println(*a)
//	}(&x)
//	x++
//	fmt.Println(x)
//}

//func main() {
//	//	defer语句执行时的拷贝机制
//	x := 10
//	defer func() {
//		fmt.Println(x)  // defer 调用函数内的变量指向了函数外部的变量
//	}()
//	x++
//	fmt.Println(x)
//}

func f1() int {
	x := 10
	defer func() {
		x++
	}()
	return x
}

func f2() *int {
	x := 10
	y := &x
	defer func() {
		*y++
	}()
	return y
}

func main() {
	fmt.Println(f1())
	fmt.Println(*f2())
	//	defer本质上是注册了一个延迟函数，defer函数的执行顺序已经确定
	//defer没有嵌套，defer的机制是取代（try except finally）
}
