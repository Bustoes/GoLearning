package main

import "fmt"

//使用值接收者实现接口和使用指针接收者实现接口的差异


type mover interface {
	move()
}

type person struct {
	name string
	age int8
}

func (p person) say() {
	fmt.Printf("%s在说话...", p.name)
}

//使用值接收者实现接口：类型的值和类型的指针都可以保存在接口变量里
func (p person) move() {
	fmt.Printf("%s在跑...\n", p.name)
}

////使用指针接收者实现接口：只有类型的指针可以保存在接口变量里
//func (p *person) move() {
//	fmt.Printf("%s在跑...\n", p.name)
//}

//=====================================

type sayer interface {		//实现多个接口
	say()
}


type animal interface {
	//move()
	//say()
	//也可以嵌套接口
	mover
	sayer
}

func main() {
	//接口型变量只能实现其接口自身拥有的方法

	var m mover
	p1 := person{		//p1是person类型的值	person
		name: "Jack",
		age: 18,
	}
	m = p1				//err 指针接收者时无法赋值，因为p1是person类型的值，没有实现mover接口
	m.move()

	p2 := &person{		//p2是person类型的指针 *person
		name: "Mike",
		age: 20,
	}
	m = p2
	m.move()

	var s sayer
	s = p2
	s.say()
	fmt.Println()

	var a animal
	a = p2
	a.say()
	a.move()
}



