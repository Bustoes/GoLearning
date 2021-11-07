package main

import "fmt"

//定义结构体	是一个值类型
type Person struct {
	name string
	city string
	age int8
}

//自己实现构造函数
//如果结构体比较复杂，使用值拷贝对于性能的开销会比较大。所以使用传地址的方式来进行构造。
func newPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age: age,
	}
}

func main() {
	//实例化
	var p1 Person	//只声明
	fmt.Printf("%T, %#v\n", p1, p1)
	p1.name = "Jack"
	p1.city = "xian"
	p1.age = 18
	var pp1 Person = Person{}	//声明且初始化

	fmt.Printf("%T, %#v\n", p1, p1)
	fmt.Printf("%T, %#v\n", pp1, pp1)

	//匿名结构体
	var user struct {
		name string
		married bool
	} = struct {
		name string
		married bool
	}{name: "11", married: true}

	user.name = "小王子"
	user.married = false

	fmt.Printf("%T, %#v\n", user, user)

	//创建结构体指针
	p2 := new(Person)
	p2.name = "Mike"
	p2.city = "上海"
	p2.age = 20
	fmt.Printf("%T, %#v\n", p2, p2)
	p3 := &Person{"Merry", "北京", 19}
	fmt.Printf("%T, %#v\n", p3, p3)




}


