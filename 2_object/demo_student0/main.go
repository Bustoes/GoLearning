package main

import (
	"fmt"
	"os"
)

//学员信息管理系统

func main() {
	sm := newStudentMgr()

	//1.打印菜单
	showMenu()
	var _ []student = []student{}
	for true {
		//2.等待用户选择要实行的选项
		var choice int
		fmt.Printf("请输入你要操作的序号：")
		fmt.Scanf("%d", &choice)

		//3.执行用户的选择
		//通过studentMgr 来调用方法
		switch choice {
		case 1:
			sm.addStudent(getStuInfo())
		case 2:
			sm.modifyStudent(getStuInfo())
		case 3:
			sm.showStudent()
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

//获取输入
func getStuInfo() *student {
	fmt.Println("请输入学员信息")
	var id int
	var name string
	var class string

	fmt.Print("请输入学员id：")
	fmt.Scanf("%d", &id)
	fmt.Print("请输入学员姓名：")
	fmt.Scanf("%s", &name)
	fmt.Print("请输入学员班级：")
	fmt.Scanf("%s", &class)
	newStu := newStudent(id, name, class)
	return newStu
}
