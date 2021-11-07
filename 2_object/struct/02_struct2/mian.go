package main

import "fmt"

//结构体的匿名字段
type Person0 struct {
	string
	int
}

//命名嵌套相当于组合结构体；匿名嵌套相当于继承结构体。

//嵌套结构体
type Address struct {
	Province string
	City string
}
type Person struct {
	Name string
	Gender string
	Age int
	//Address Address	//可以把这个变为匿名嵌套
	Address
}

func (p Person) use() {
	fmt.Println(p.City)
}


func main() {
	var _ = Person0{
		string: "Jack",
		int: 18,
	}.string

	p1 := Person{
		Name: "Jack",
		Gender: "男",
		Age: 18,
		Address: Address{
			Province: "陕西",
			City: "西安",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p1.Province)
	fmt.Printf("%#v\n", p1.Address.Province)
	//可以通过p1再通过Address来访问Province, City。
	//使用匿名嵌套，才使其可以直接通过p1来访问Province, City。




}
