package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//type Info struct {
//	// 能表述的信息有限的
//	Name      string // neme映射成mysql中的char类型还是varchar类型还是text类型，即使能够说明，但是二外的max_length不能表述
//	AgeDetail int    `json:"age,omitempty"` // 不加tag会有默认值
//	Gender    string
//}

type Info struct {
	Name   string `orm:"name, max_length=17, min_length=5"`
	Age    int    `orm:"age,min=18, max=70"`
	Gender string `orm:"gender,required"`
}

func main() {
	//结构体标签
	/*
		结构体除了面子和类型外，还可以有一个可选的标签（tag）：
		它属于一个附属于字段的字符串，可以是文档或者其他的重要标记。比如在解析json或生成json文件时，常用岛的encoding/json包，
		它提供一些默认的标签，例如：omitempty标签可以在序列化的时候忽略0值或者空值；
		而标签的作用是不进行序列化，其效果和直接将结构体中的字段写成大小写的效果是一样的
	*/

	info := Info{
		Gender: "男",
		Name:   "xiaoming",
	}
	re, _ := json.Marshal(info)
	fmt.Println(string(re))
	//	通过反射包去识别这些tag
	t := reflect.TypeOf(info)
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i) // 获取结构体的每一个字段
		tag := filed.Tag.Get("orm")
		fmt.Printf("%d %v (%v), tag:'%v'\n", i+1, filed.Name, filed.Type.Name(), tag)
		//	具体的应用绝大部分情况是不需要使用到反射，实际开发项目中会用到
	}
}
