package main

import "fmt"

func main() {
	//  什么是切片
	//	数组的问题，大小确定，不能修改，切片是动态的数组
	//	var identifier []type
	//var cousers []string // 定义一个切片
	//var cousers = []string{"django", "scrapy", "tornado"}
	//fmt.Println(cousers)
	//fmt.Printf("%T\n", cousers)

	//  第二种方法：切片的另一种初始化方法，make
	//  切片没有长度，make初始化需要传递一个长度
	//courses := make([]string, 5)
	//fmt.Printf("%T\n", courses)
	//	slice对标的就是python中的list
	//	第三种方法：通过数组变成一个切片
	//var courses = [5]string{"django", "scrapy", "tornado", "flask", "fastapi"}
	//subCourse := courses[1:4] // python中取值交切片；在go语言中的切片是一种数据结构
	//fmt.Println(courses)
	//fmt.Println(subCourse)
	//fmt.Printf("%T\n", subCourse)
	////	切片操作会影响原始数组的值
	//subCourse[0] = "asyncio"
	//fmt.Println(courses)
	//fmt.Println(subCourse)
	//fmt.Printf("%T\n", subCourse)

	//	第四种方式：new
	//subCourse := *new([]string)
	//subCourse := new([]string)
	//fmt.Println(subCourse)

	//	数组的传递是值传递，切片是引用传递，会改变原始值
	//	slice是动态数组，可以动态添加值
	var courses = [5]string{"django", "scrapy", "tornado", "flask", "fastapi"}
	subCourse := courses[1:4]
	subCourse2 := subCourse[1:2]
	fmt.Printf("%T\n%v\n", subCourse2, subCourse2)

	//append 可以向切片追加元素
	subCourse2 = append(subCourse2, "django", "scrapy")
	fmt.Println(subCourse2)
	subCourse3 := []string{}
	copy(subCourse3, subCourse2)
	//	拷贝的时候，目标函数的长度需要设置好,否则拷贝不成功
	fmt.Println(subCourse3)
	//	append函数可以追加多个元素
	appendCourses := []string{"java", "scala", "c"}
	subCourse2 = append(subCourse2, appendCourses...)
	fmt.Println(subCourse2)
	//	从切片中删除元素
	deletecourses := [5]string{"django", "scrapy", "tornado", "flask", "fastapi"}
	courseSlice := deletecourses[:] // 数组全部转成切片
	courseSlice = append(courseSlice[:1], courseSlice[2:]...)
	fmt.Println(courseSlice)
	//	判断某个元素是否在切片中，遍历循环

	//	go语言slice和Python的list底层都是基于数组来实现的
	//  slice进行的操作都会影响原来的数组，slice更像是一个指针，本身不存值
	//	slice的原理，因为很多底层的知识相对来说很多时候并不难而是需要花费比较多的时间去慢慢理解
	//	1、第一个现象
	a := make([]int, 0)
	b := []int{1, 2, 3}
	fmt.Println(copy(a, b))
	//  2、第二个现象
	c := b[:]
	c[0] = 80
	fmt.Println(b)
	fmt.Println(c)
	//	3、第三个现象
	c = append(c, 9)
	fmt.Println(b) // append函数没有影响到原来的数组
	fmt.Println(c)
	c[0] = 10
	fmt.Println(b)
	fmt.Println(c) // 为什么append函数之后的再调用c[0]不会影响原来的数组
	//	4、第四个现象
	fmt.Println(len(c))
	fmt.Println(cap(c)) // cap指的是容量，长度和容量不同
	//	切片底层使用数组来实现的，既要使用数组又要满足动态的功能
	//	假设有一个值，实际上申请数组的时候可能有两个，如果后续要增加数据，那么直接添加到数组的结尾，这时候不需要重新申请额外内存空间
	//	1、使用make方法初始化，查看len和cap是多少
	//d := make([]int, 5)
	d := make([]int, 5, 6)
	fmt.Printf("len=%d, cap=%d\n", len(d), cap(d))
	//	2、通过数组取切片
	data := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[2:4]
	for _, value := range slice {
		fmt.Println(value)
	}
	fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice)) // 容量是从起始到结尾
	//	3、直接定义
	slice2 := []int{1, 2, 3}
	fmt.Printf("len=%d, cap=%d\n", len(slice2), cap(slice2))
	//	切片扩容问题，扩容阶段会影响速度
	/*
		首先判断，如果新申请的容量（cap）大于2倍的旧容量（old cap）,最终容量（new cap）就是新申请的容量（cap）
		否则判断，如果旧切片小于1024，则最终容量（new cap)就是旧容量（old cap）的两倍，即（new cap=double cap）
		否则判断，如果旧切片长度大于等于1024，则最终容量（new cap）从旧容量（old cap）开始循环增加原来的1/4
		如果最终容量（cap）计算值溢出，最终容量（cap）就是新申请容量（cap）
	*/
	//  如果小于1024扩容速度是2倍，如果大于1024，扩容速度就是1.25
	oldSlice := make([]int, 0)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))
	oldSlice = append(oldSlice, 1)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))
	oldSlice = append(oldSlice, 2)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))
	oldSlice = append(oldSlice, 3)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))
	oldSlice = append(oldSlice, 4)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))
	oldSlice = append(oldSlice, 5)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))

	//	切片来说，底层是数组，如果基于数组产生，会有一个问题影响原来的数组
	//  2、切片的扩容机制
	//  3、切片是引用传递

	//	当make遇到了append易出现的问题
	s1 := make([]int, 5) // 非空slice，置空可为0
	s1 = append(s1, 10)  // 在5个后面追加
	fmt.Println(s1)
}
