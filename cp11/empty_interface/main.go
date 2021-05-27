package main

import "fmt"

type Course struct {
	Name  string
	Price int
	url   string
}
type Printer interface {
	printInfo() string
}

func (c Course) printInfo() string {
	return "课程信息"
}

func printvalue(i interface{}) {
	fmt.Printf("%v\n", i)
}

func main() {
	// 空接口
	var i interface{} //空接口
	//	空接口类似java和Python中的object元类
	i = Course{}
	fmt.Println(i)
	// 空接口的第一个用法，可以把任何类型赋值给空接口变量
	//	2、参数传递
	i = 100
	printvalue(i)
	i = "哈哈"
	printvalue(i)
	i = []string{"scrapy", "tornado"}
	printvalue(i)
	//	空接口可以作为map的值
	var teacherInfo = make(map[string]interface{})
	teacherInfo["name"] = "小明"
	teacherInfo["age"] = 29
	teacherInfo["weight"] = 70.5
	fmt.Println(teacherInfo)

	//	类型断言
	//接口的一个坑
	//接口有一个默认的规范，接口的名称一般以er结尾
	//c := &Course{}
	//var c Printer = Course{}
	//c.printInfo()
	//

}
