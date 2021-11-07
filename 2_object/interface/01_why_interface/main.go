package main

import (
	"fmt"
)

type dog struct {}

func (d dog) say() {
	fmt.Println("汪汪汪...")
}

type cat struct {}

func (c cat) say() {
	fmt.Println("喵喵喵...")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println(p.name + "啊啊啊...")
}


/*
与java的联想：interface有继承的想法，因为它定义了方法的类型。又有些像局部的泛型。
java中是先实现接口，再重写方法
golang中是先写方法，再实现接口
*/
//实现接口
//一个对象只要实现了接口里的所有方法，那么就实现了这个接口
//定义一个抽象的类型，只要实现了say()这个方法，只要实现了say()方法，都可以称为sayer类型

type sayer interface {
	say()
}

func mo(arg sayer) {
	//不管传进来什么都要执行
	arg.say()
}
func main() {
	c1 := cat{}
	d1 := dog{}
	mo(c1)
	mo(d1)

	p1 := person{name: "Jack"}
	mo(p1)

	//接口型变量
	var s sayer		//相当于同一类型就可以相互赋值
	s = c1
	s.say()
	s = p1
	s.say()


}

