package main

import "fmt"

func main() {
	//  go语言中的数组和Python中的list对相起来理解，slice和Python中的list更像
	//	静态语言中的数组：1、大小确定；2、类型一致
	//	数组声明
	//	var courses [10] string
	//	var courses [5] string = [5]string{"django","flask","scrapy", "tornado", "fastapi"}
	//var courses  = [5]string{"django", "flask", "scrapy", "tornado", "fastapi"}
	courses := [5]string{"django", "flask", "scrapy", "tornado", "fastapi"}
	fmt.Println(courses)
	//	1、修改值，取值；不支持删除和添加值，数组一开始声明大小
	// 取值，修改值
	fmt.Println(courses[0])
	//	修改值
	courses[0] = "django3"
	fmt.Println(courses)

	//	数组的另一种创建方式
	//var a [4] float32
	var a = [4]float32{1.0}
	fmt.Println(a)
	var b = [5]int{'A', 'B'}
	fmt.Println(b)
	//	缺省声明
	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(c)
	d := [5]int{3: 60} // 将下标为3的初始化为60
	fmt.Println(d)
	e := [...]int{1: 10, 5: 10, 6: 99}
	fmt.Println(e)

	//	数组的第一种场景：求长度
	fmt.Println(len(e))
	// 数组的操作第二种场景：遍历数组
	//for index, value := range courses {
	//	println(index, value)
	//}
	//	使用for range求和
	sum := 0
	for _, value := range e {
		sum += value
	}
	fmt.Println(sum)
	//	使用for语句也可以遍历数组
	sum = 0
	for i := 0; i < len(e); i++ {
		sum += e[i]
	}
	fmt.Println(sum)
	//	数组的值类型
	//  在go语言中，courseA和courseB都是数组，但是不是同一种类型
	courseA := [3]string{"django", "flask", "scrapy"}
	courseB := [4]string{"flask", "scrapy", "tornado", "fastapi"}
	fmt.Printf("%T\n", courseA)
	fmt.Printf("%T\n", courseB)
}
