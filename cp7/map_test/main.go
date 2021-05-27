package main

import "fmt"

func main() {
	//	go语言中的map，类似Python中的dict
	//	go语言中的map的key和value类型申明时需要指明
	//  1、字面值
	m1 := map[string]string{
		"m1": "v1",
	}
	fmt.Println(m1)
	//	2、使用make函数，make可以创建slice和map
	m2 := make(map[string]string)
	m2["m2"] = "v2"
	fmt.Println(m2)
	//	3、定义一个空的map
	m3 := map[string]string{}
	fmt.Println(m3)
	//	map中的key不是所有的类型可以支持，该类型需要支持==或者!=判断操作
	//  int,rune
	//  切片不可以
	mp := map[string]string{
		"a": "va",
		"b": "vb",
	}
	//	1、进行增加和修改
	mp["c"] = "vc"
	mp["b"] = "vbr"
	fmt.Println(mp)
	//	查询, 返回空的字符串到底是没有获取到还是本身就是空字符串
	v := mp["d"]
	fmt.Println(v)
	//  可以返回两个值
	v1, ok := mp["d"]
	if ok {
		fmt.Println("找到了", v1)
	} else {
		fmt.Println("很抱歉，没找到")
	}
	//	删除
	delete(mp, "a")
	fmt.Println(mp)
	//	遍历
	for k, v := range mp {
		fmt.Println(k, v)
	}
	//	go语言中也有一个list，就是数据结构中提到的链表

}
