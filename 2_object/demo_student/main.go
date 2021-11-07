package main

import (
	"fmt"
	"os"
)

//学员信息管理系统

func main() {
	studentList := make([]student, 0, 8)

	//1.打印菜单
	showMenu()
	var _ []student = []student{}
	for true {
		//2.等待用户选择要实行的选项
		var choice int
		fmt.Printf("请输入你要操作的序号：")
		fmt.Scanf("%d", &choice)

		//3.执行用户的选择
		switch choice {
		case 1:
			addStu(&studentList)
		case 2:
			editStu(&studentList)
		case 3:
			showStu(&studentList)
		case 4:
			os.Exit(0)
		}
	}

}

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示学员信息")
	fmt.Println("4.退出")
}

func addStu(studentList *[]student) {
	var id int
	var name string
	var class string

	fmt.Printf("请输入学员id：")
	fmt.Scanf("%d", &id)
	fmt.Printf("请输入学员姓名：")
	fmt.Scanf("%s", &name)
	fmt.Printf("请输入学员班级：")
	fmt.Scanf("%s", &class)

	*studentList = append(*studentList, *newStudent(id, name, class))

}

func editStu(studentList *[]student) {

}

func showStu(studentList *[]student) {
	for _, student := range *studentList {
		fmt.Printf("%#v\n", student)
	}
}

