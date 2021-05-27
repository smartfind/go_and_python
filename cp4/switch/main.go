package main

import "fmt"

func main() {
	score := 85
	/*
		> 90 A
		80-89 B
		70-79 C
		60-69 D
		< 60 E
	*/
	// 使用if条件判断
	grade := "A"
	//if score >= 90 {
	//	grade = "A"
	//} else if score >= 80 && score <= 89 {
	//	grade = "B"
	//} else if score >= 70 && score <= 79 {
	//	grade = "C"
	//} else if score >= 60 && score <= 69 {
	//	grade = "D"
	//} else {
	//	grade = "E"
	//}
	//fmt.Println(grade)
	//	使用switch
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80 && score <= 89:
		grade = "B"
	case score >= 70 && score <= 79:
		grade = "C"
	case score >= 60 && score <= 69:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Println(grade)

	//	不同case表达式使用逗号分隔，一分支多值
	var name = "xiaoming"
	switch name {
	case "xiaoming", "xiaohua":
		fmt.Println("We are classmate")
	}
}
