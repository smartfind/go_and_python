package main

import "fmt"

func main() {
	//	make和new函数
	//	new函数用法
	//	var p *int  // 申明了一个变量p，但是没有初始值，没有内存
	//	*p = 10
	//	有默认值的数据类型，int/byte/rune/float/bool/string这些类型都有默认值
	//	指针/切片/map/接口，这些默认值都nil，类似Python中none
	//	针对指针来说或者其他的默认值是0的情况来说，如何一开始申明的时候分配内存
	var p *int = new(int) // go的编译器就知道先申请一个内存空间，这里的内存中的值全部设置为0
	*p = 1000
	fmt.Println(p)
	//	除了new函数可以申请内存意外，还有一个函数时make，更加常用的是make函数，make函数一般用于切片和map
	var info map[string]string = make(map[string]string)
	info["class"] = "python"
	fmt.Println(info)
	//	new函数返回的是这个值的地址，指针，make函数返回的是指定类型的实例
	//	使用指针会占用额外的空间
	//	map和slice的默认值是nil
	//	go语言中nil是唯一可以用来表示部分类型的零值的标识符，它可以代表许多不同内存布局的值

}
