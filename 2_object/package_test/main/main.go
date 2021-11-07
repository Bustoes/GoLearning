package main

import (
	"fmt"
	cal "package_test/calc"	//给导入的包起别名
	//只想导入包，而不想使用包里面的内容时，使用匿名导入	常用于导入某个包内数据库的驱动



)

func main() {
	ret := cal.Add(10, 20)
	fmt.Println(ret)
	fmt.Println(cal.Name)


}
