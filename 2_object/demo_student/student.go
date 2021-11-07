package main

type student struct {
	id int //唯一
	name string
	class string
}

//student的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id: id,
		name: name,
		class: class,
	}
}


