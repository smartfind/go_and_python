package main

import "fmt"

type Teacher struct {
	Name  string
	Age   int
	Title string
}

func (t Teacher) teacherInfo() {
	fmt.Printf("姓名:%s，年龄:%d，职称：%s\n", t.Name, t.Age, t.Title)
}

type Course struct {
	Teacher Teacher // 将另一个结构体的变量放进来
	Name    string
	Price   int
	Url     string
}

// 匿名嵌套，是语法糖，不算继承
type Course2 struct {
	Teacher // 将另一个结构体的变量放进来
	Name    string
	Price   int
	Url     string
}

func (c Course) courseInfo() {
	fmt.Printf("课程名：%s,价格:%d,讲师信息:%s %d %s\n", c.Name, c.Price, c.Teacher.Name, c.Teacher.Age, c.Teacher.Title)
}

func (c Course2) courseInfo() {
	fmt.Printf("课程名：%s,价格:%d,讲师信息:%s %d %s", c.Name, c.Price, c.Teacher.Name, c.Age, c.Title) // 当内嵌结构体字段冲突时需要指明
}

func main() {
	//	go 语言的继承，组合
	t := Teacher{
		Name:  "小明",
		Age:   28,
		Title: "Python开发工程师",
	}
	c := Course{
		Teacher: t,
		Price:   100,
		Url:     "www.google.com",
		Name:    "tornado",
	}
	c1 := Course2{
		Teacher: t,
		Price:   100,
		Url:     "www.google.com",
		Name:    "tornado",
	}
	c.courseInfo()
	c1.courseInfo()

}
