package main

import "fmt"

func div(a, b int) (int, error) {
	if b == 0 {
		panic("被除数不能为0")
	}
	return a / b, nil
}

func main() {
	////	错误，可预知的可能出现的情况，这些情况导致你的代码出问题，参数检查，数据库访问不了
	//data := 12
	//strconv.Itoa(data)           //go语言认为Itoa的函数不可能出错，没有返回error，内部代码出错时应该抛出panic
	//i, err := strconv.Atoi("12") //Atoi这个函数认为函数内部出现一些预知错误
	//if err != nil {
	//	fmt.Println(i)
	//}
	//	异常，go语言中抛出异常和捕获异常
	a := 6
	b := 0
	defer func() {
		err := recover() // 恢复执行主函数后面的代码
		if err != nil {
			fmt.Println("异常被捕获到")
		}
		fmt.Println("哈哈")
	}()
	fmt.Println(div(a, b))
}
