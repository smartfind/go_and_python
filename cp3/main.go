package main

import (
	"fmt"
	"strings"
)

func main() {
	//	字符串的基本操作
	//	1、求解字符串的长度
	//  返回的是字节的长度
	//  涉及中文问题产生了变化
	//  Unicode字符集，存储的时候需要编码utf8编码规则，utf8编码是一个动态的编码规则
	//  utf8编码是动态长度的，还能够用一个字节表示英文
	var info string = "xiaoming，你在家吧"
	fmt.Println(len(info))
	// 类型转换
	info_arr := []rune(info)
	fmt.Println(len(info_arr))
	//	转移字符
	date := "2020\\02\\07"
	fmt.Println(date)
	//	2、判断字符串中是否包含某个子串
	fmt.Println(strings.Contains(info, "你在"))
	fmt.Println(strings.Index(info, "你在"))
	//	3、统计字符出现的次数
	fmt.Println(strings.Count(info, "i"))
	//	4、判断前缀和后缀
	fmt.Println(strings.HasPrefix(info, "x"))
	fmt.Println(strings.HasSuffix(info, "x"))
	//	5、大小写转换
	fmt.Println(strings.ToUpper("hcsjk"))
	fmt.Println(strings.ToLower("TVGHJKK"))
	//	6、字符串的比较
	fmt.Println(strings.Compare("a", "b")) // 字符的比较就是ascii码的比较，小返回-1，大返回1，相等0
	fmt.Println(strings.Compare("b", "a")) // 字符的比较就是ascii码的比较，小返回-1，大返回1，相等0
	//	7、去掉空格和指定字符串
	//fmt.Println(strings.TrimSpace("gdhh hfjkj")) // 空格只能在开头或者结尾
	//fmt.Println(strings.TrimSpace(" gdhh hfjkj"))
	fmt.Println(strings.TrimLeft("jafajklf", "j"))
	fmt.Println(strings.TrimRight("jafajklf", "j"))
	fmt.Println(strings.Trim("jafajklf", "j"))
	//	8、字符串按指定字符分割
	fmt.Println(strings.Split("zxcvbbnn", "b"))
	//  9、合并，join方法将字符串数组连接起来
	arrs := strings.Split("zxcvbbnn", "b")
	fmt.Println(strings.Join(arrs, ""))
	//	10、字符串替换
	fmt.Println(strings.Replace("zxcvbbnnnnnn", "nn", "qwer", 2))

}
