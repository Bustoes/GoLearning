package main

import "fmt"

//实现结构体的继承，要使用匿名嵌套

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) bark() {
	fmt.Printf("%s会汪汪汪\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "Mike",
		},
	}
	d1.bark()
	d1.move()


}


