package main

import (
	"fmt"
	"strconv"
)

func main() {
	//	printf 格式化
	name := "xiaoming"
	age := 25
	fmt.Println("name:" + name + ",age:" + strconv.Itoa(age))
	fmt.Printf("%v\n", age)   //  原样输出
	fmt.Printf("%#v\n", name) //  字符串带引号
	fmt.Printf("%T\n", age)   //  输出类型
	fmt.Printf("name:%s, age:%+d\n", name, age)
	fmt.Printf("name:%s, age:%4d\n", name, age) // 补空格
	fmt.Printf("name:%s, age:%b\n", name, age)  // 二进制换算
	desc := fmt.Sprintf("name:%s, age:%x\n", name, age)
	fmt.Println(desc)
	data := 65
	fmt.Printf("%c\n", data)
	fmt.Printf("%q\n", data)

	//	输入
	//var info string
	//fmt.Scanln(&info) // 阻塞等待获取键盘录入的数据
	//fmt.Println(info)
	// 通过scanf输入
	var n string
	var a int
	fmt.Println("请输入你的姓名和年龄：")
	fmt.Scanf("%s %d", &n, &a)
	fmt.Println(n, a)

}
