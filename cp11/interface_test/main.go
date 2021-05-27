package main

import "fmt"

// 捷克偶是一个协议-程序员-1、写代码；2、解决bug；就是一组方法的结合
type Programmer interface {
	Coding() string
	Debug() string
}

type Designer interface {
	Design() string
}

type Manager interface {
	Programmer
	Designer
	Manage() string
}

type UIDesigner struct {
}

func (d UIDesigner) Design() string {
	fmt.Println("我会UI设计")
	return "我会UI设计"
}

// java里面一种类型只要集成了一个接口才行，如果你继承了这个接口的话，那么这个接口里面的全部方法必须都要实现

type Pythoner struct {
	UIDesigner
}

func (p Pythoner) Coding() string {
	fmt.Println("Python开发者")
	return "Python开发者"
}

func (p Pythoner) Debug() string {
	fmt.Println("我会Python的debug")
	return "我会Python的debug"
}

//func (p Pythoner) Design() string {
//	fmt.Println("我是一个Python开发者，但是UI设计我也会")
//	return "我是一个Python开发者，但是UI设计我也会"
//}

func (p Pythoner) Manage() string {
	fmt.Println("我是一个Python开发者，管理我也懂")
	return "我是一个Python开发者，管理我也懂"
}

type Go struct {
}

func (p Go) Coding() string {
	fmt.Println("Go开发者")
	return "Go开发者"
}

func (p Go) Debug() string {
	fmt.Println("我会Go的debug")
	return "我会Go的debug"
}

// 对于Pythoner这个结构体来说，实现任何方法都可以，只要不全部实现coding和Debug的话，那么Pythoner就不是一个Programmer
// 1、Pythoner本身自己就是一个类型，不必在意是不是Programmer
// 2、多态
// 3、在讲解多态之前，需要对interface做一点说明：在go语言中接口是一种类型，是一种抽象类型

//开发中经常会遇到的问题
//开发一个点上网站，支付环节：微信、支付宝、银行卡，你的系统支持各种类型的支付，每一种支付类型都有统一的支付接口
//定一个协议：1、创建订单；2、支付；查询支付状态；4、退款
//支付发起了
//type Alipay struct {
//}
//type Weipay struct {
//}
//type Bank struct {
//}
//
//var a Bank
//var b Weipay
//var c Alipay

// 多态，什么类型的时候你申明的类型是一种兼容类型，但是实际赋值的时候是另一种类型
//接口的类型是强制的
//比如有一个缓存需求，当前使用是redis，后期技术改造需要使用memcache
//这种多态特性在Python中不需要，动态语言直接赋值即可
//go语言中并不支持继承
//如果后期接入一种新的支付，或者取消已有的支付很麻烦

func main() {
	var pro Programmer = Pythoner{}
	pro.Coding()
	var pros []Programmer
	pros = append(pros, Pythoner{}, Go{})

	//	接口虽然是一种类型纪念馆，但是和其他类型不太一样，接口是一种抽象类型，struct是具象
	p := Pythoner{}
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", pro)

	//	1、go struct组合，组合一起实现所有接口的方法也是可以的
	//	2、接口本身也支持组合
	var m Manager = Pythoner{}

	//	struct组合完成接口,接口支持组合-1、继承；2、结构体组合实现了所有接口方法也没问题
	m.Design()
	//	go语言本身也推荐鸭子类型-error

}
