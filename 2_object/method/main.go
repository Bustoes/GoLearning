package main

import "fmt"

type person struct {
	name string
	age int8
}

func newPerson(name string, age int8) *person {
	return &person{
		name: name,
		age: age,
	}
}

//定义方法	若使用type，可以实现int的定义方法
//相当于给person添加函数
func (p person) dream() { //方法属于此接收者，理解为调用者
	fmt.Printf("method.name: %s\n", p.name)
}
func (p *person) setAge1(newAge int8) { //方法属于此接收者
	p.age = newAge	//p是指针类型
}
func (p person) setAge2(newAge int8) { //方法属于此接收者
	p.age = newAge	//p是基本类型， 这里修改的p和main函数中的p1不是一个对象了
}

func main() {
	p1 := newPerson("Jack", 18)

	p1.dream()
	fmt.Printf("%#v\n", p1)
	p1.setAge1(80)
	fmt.Printf("%#v\n", p1)
	p1.setAge2(40) //不发生改变
	fmt.Printf("%#v\n", p1)

}
