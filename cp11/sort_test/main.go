package main

import (
	"fmt"
	"sort"
)

type Course struct {
	Name  string
	Price int
	Url   string
}

type Courses []Course

func (c Courses) Len() int {
	return len(c)
}
func (c Courses) Less(i, j int) bool {
	return c[i].Price < c[j].Price
}

func (c Courses) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	//	通过sort来排序
	//	排序算法是否能应付各种类型的排序
	courses := Courses{
		Course{"django", 388, ""},
		Course{"tornado", 499, ""},
		Course{"flask", 699, ""},
		Course{"fastapi", 377, ""},
	}
	sort.Sort(courses) // 协议，你的目的不是要告诉具体的类型，重要的是你的类型必须提供具体的方法
	for _, v := range courses {
		fmt.Println(v)
	}
}
