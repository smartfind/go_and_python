package main

import "fmt"

func main() {
	//  for init;condition;post{} 计算1-10的和
	//  1、执行i := 1
	//  2、是否i <= 10
	//  3、sum=1
	//  4、i++ -> i = 2
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//	死循环
	//for {
	//	fmt.Println("今天是2月8日")
	//}

	//i := 0
	//for i < 10 {
	//	fmt.Println("今天是2月8日")
	//	i++
	//}

	//	循环字符串
	name := "xiaoming"
	//for index, value := range name {
	//	// 字符串遍历是ASCII码值，数字
	//	fmt.Println(index, value)
	//	fmt.Printf("%c\n", value)
	//}
	for _, value := range name {
		// 字符串遍历是ASCII码值，数字
		fmt.Printf("%c\n", value)
	}
	info := "今天是星期二"
	//	1、处理中文; 2、字符串时字符的数组
	fmt.Printf("%c\n", name[0])
	fmt.Printf("%c\n", info[0]) // 中文会异常

	//  在做字符串遍历的时候尽可能使用rune
	info_arr := []rune(info)
	for i := 0; i < len(info_arr); i++ {
		fmt.Printf("%c\n", info_arr[i])
	}

}
