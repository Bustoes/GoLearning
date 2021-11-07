package main

import "fmt"

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

//学员管理的类型
type studentMgr struct {
	allStudents []*student

}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

//调用者是studentMgr
//因为要使用 studentMgr 里面的字段，所以需要 studentMgr 作调用者
//相当于给 studentMgr 添加函数

//1.添加学生
func (s *studentMgr) addStudent(newStu *student) {

	s.allStudents = append(s.allStudents, newStu)
}

//2.编辑学生
func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		//若id相同，则找到学生，并直接取代
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			return
		}
	}
	//如果没有找到
	fmt.Println("输入的学生信息有误")
}
//3.展示学生
func (s *studentMgr) showStudent() {
	for _, stu := range s.allStudents {
		fmt.Printf("%#v\n", stu)
	}
}






