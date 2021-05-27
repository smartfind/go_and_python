package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//	定义布尔类型
	//var gender bool = true
	//fmt.Println(gender)
	//gender_2 := false
	//fmt.Println(gender_2)
	//	相比python而言，很多场景下数字有上限，可以选择合适的数据类型来降低内存的占用
	//	如果直接使用int,int是一种动态类型，取决于机器本身多少位，64位机器上运行那么int就是int64，如果是32及其那么就是4个字节
	//var age int = 18
	//fmt.Println(unsafe.Sizeof(age))

	//	float类型
	//var weight float32 = 55.88
	//fmt.Println(weight)
	// float类型的最大长度
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	//	为什么64为的float最大值远大于int64，float底层存储和int的存储不一样
	//	默认情况选择的float64
	weight := 69.65
	fmt.Printf("%T", weight)
	//	byte和rune
	//  静态语言中中文处理容易出错，byte处理ASCII码，rune处理中文
	// byte本质是uint8类型，rune本质是int32类型
	var a byte = 18
	fmt.Println(a)
	b := 'b'
	//	1、b+1可以和数字计算；2、b+1的类型是32；3、int类型可以只变成字符
	fmt.Println(reflect.TypeOf(b + 1))
	fmt.Printf("%T\n", b+1)
	fmt.Printf("b+1=%c", b+1)
}
