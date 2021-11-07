package main

import "fmt"

func main() {
	a()
	b()
	c()

}

func a() {
	fmt.Println("func a")
}

func b() {
	//可能触发panic
	//在其之前使用defer来调用一个函数
	defer func() {
		err := recover()	//收集错误信息
		if err != nil {
			fmt.Println("func b error")
		}
	} ()
	panic("panic in b")
	//本函数之后的语句不会被执行本函数

}

func c() {
	fmt.Println("func c")
}
