package main

import (
	"fmt"
	"unsafe"
)

type Course struct {
	Name  string
	Price int
	Url   string
}

func (c Course) printCourseInfo() {
	fmt.Printf("课程名称：%s，课程价格：%d，课程地址：%s\n", c.Name, c.Price, c.Url)
}

func (c Course) setPrice(price int) {
	c.Price = price
}

func (c *Course) setPrice2(price int) {
	c.Price = price
}

// 1、结构体的方法只能和结构体在同一个包中，可以不同文件名
// 2、内置的int类型不能加方法

func main() {
	//	go语言支持面向对象
	//	面向对象的三大特征：1、封装；2、继承；3、多态；4、方法重载；5、抽象基类
	//  定义struck go语言没有class这个概念，所以可以少理解很多面向对象的抽象概念
	//  1、实例化-kv形式
	var c Course = Course{
		Name:  "scrapy",
		Price: 10000,
		Url:   "www.google.com",
	}
	//	访问
	fmt.Println(c.Name, c.Price, c.Url)
	//	大小写在go语言中的重要性
	//	结构体定义的名称以及属性首字母大写很重要，方便其他包调用

	//  2、第二种实例化方式
	c2 := Course{
		"django",
		10000,
		"www.google.com",
	}
	fmt.Println(c2.Name, c2.Price, c2.Url)

	//	3、如果一个指向结构体的指针
	c3 := &Course{
		"tornado",
		10000,
		"www.google.com",
	}
	fmt.Printf("%T\n", c3)
	fmt.Println((*c3).Name, (*c3).Price, (*c3).Url)
	//  另一个根本的原因-go语言的指针是受限的
	fmt.Println(c3.Name, c3.Price, c3.Url) // 这里go语言的一个语法糖，go语言会将c3转换成(*c3)
	//  4、零值，如果不给它结构体赋值，go语言会给每个字段采用默认值，即类型的默认值
	c4 := Course{}
	fmt.Println(c4.Price)

	//	5、多种方式零值初始化结构体
	var c5 Course = Course{}
	var c6 Course
	var c7 *Course = new(Course)
	//var c8 *Course
	// 为什么c6和c8表现出来的结果不一样，指针如果只申明不赋值，默认值是nil；c6不是指针，是结构体类型
	//  slice和map默认值是nil

	fmt.Println("零值初始化")
	fmt.Println(c5.Price, c6.Price, c7.Price)
	//fmt.Println(c8.Price)

	//	6、结构体是值类型，传递时会拷贝
	c8 := Course{
		"tornado",
		10000,
		"www.google.com",
	}
	c9 := c8
	fmt.Println(c8)
	fmt.Println(c9)
	c8.Price = 222
	fmt.Println(c8)
	fmt.Println(c9)

	//	7、结构体的大小，占用内存的大小，可以使用sizeof来查看对象占用的类型
	//  go语言中struct无处不在
	fmt.Println(unsafe.Sizeof(1))
	//  go语言中string的本质是一个结构体
	//type string struct {
	//	Data uintptr // 指针站8个长度
	//	Len  int     // 长度，64位系统站8个长度
	//}
	fmt.Println(unsafe.Sizeof("afafagfagdgdhshhh")) // 不管多长都是16
	fmt.Println(unsafe.Sizeof(c8))                  // 16+8+16

	//	8、slice的大小，24（64位系统）
	type slice struct {
		array unsafe.Pointer // 底层数组的地址
		len   int            // 长度
		cap   int            // 容量
	}
	s1 := []string{"tornado", "www.google.com"}
	fmt.Println("切片占用的内存：", unsafe.Sizeof(s1))

	m1 := map[string]string{
		"c1": "django",
		"c2": "tornado",
	}
	fmt.Println(unsafe.Sizeof(m1))

	//	结构体方法,达到了封装数据和封装方法的效果
	c10 := Course{
		"tornado",
		10000,
		"www.google.com",
	}
	Course.printCourseInfo(c10)
	c10.printCourseInfo()

	//  为什么价格为发生变化?函数的参数怎么传递的？结构体是值传递
	c10.setPrice(10)
	Course.setPrice(c10, 10) // 上面的调用实际转化为该行调用
	fmt.Println(c10.Price)
	//	如果想改值，需要改为指针类型
	c10.setPrice2(10)
	//(&c10).setPrice2(10)
	fmt.Println(c10.Price)

	//结构体的接受者有两种形式：1、值传递；2、指针传递，当需要修改结构体的值时，或者结构体的数据量很大时
	//	go语言不支持继承，但是有方法能达到同样的效果，组合
}
