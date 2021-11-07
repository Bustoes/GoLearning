package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性与JSON序列化。
//如果一个标识符首字母是大写的，那么就是对外可见的。
//需要一个内置包里的函数来访问结构体里的字段时，如果字段对外不可见，函数将无法访问到此字段

//结构体标签 tag
//是结构体的元信息，可以在运行时通过反射的机制来读取。

type student struct {
	ID   int
	Name string
}

type class struct {
	Title    string    `json:"title,omitempty" db:"TITLE"`
	Students []student `json:"students,omitempty" xml:"STUDENTS"`
}

func newStudent(id int, name string) student {
	return student{
		ID: id,
		Name: name,
	}
}

func main() {
	//创建一个班级c1
	c1 := class{
		Title: "火箭101",
		Students: make([]student, 0, 20),
	}

	//向班级中添加学生
	for i := 0; i < 10; i++ {
		tempStu := newStudent(i + 1, fmt.Sprintf("stu%d", i + 1))
		c1.Students = append(c1.Students, tempStu)
	}
	fmt.Printf("%#v\n", c1)

	//JSON序列化：将GO语言中的数据转换成满足JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	//JSON反序列化：JSON格式字符串转换为GO语言中的数据
	var c2 class
	err = json.Unmarshal([]byte(data), &c2) //需要修改c2里面的值，所以需要传递一个地址
	if err != nil {
		fmt.Println("json unmarshal failed, err:", err)
		return
	}
	fmt.Printf("%#v\n", c2)



}
