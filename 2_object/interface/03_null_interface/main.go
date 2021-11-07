package main

import "fmt"

//空接口	object
//接口中没有定义任何方法
//任意类型都实现了空接口 --> 空接口变量可以存储任意类型的值
type xxx interface {}

// 空接口的应用
// 1.空接口作为函数的参数
// 2.空接口作为map的value

//一个空接口内部分为了两部分：动态类型、动态值


func main() {
	var x xxx	//定义一个空接口变量

	//可以接受任何类型的值
	x = "hello"
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	x = false
	fmt.Println(x)
	x = struct {}{}
	fmt.Println(x)

	//空接口作map的value
	var m = make(map[string]interface{}, 16)

	m["name"] = "Jack"
	m["age"] = 18
	m["hobby"] = []string{"篮球", "足球", "双色球"}

	fmt.Println(m)


}
