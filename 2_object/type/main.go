package main

import "fmt"

//自定义类型
type MyInt int	//相当于包装类？Integer Character
//类型别名
type NewInt = int

func main() {
	var i1 MyInt
	fmt.Printf("%v, %T\n", i1, i1)
	var i2 NewInt
	fmt.Printf("%v, %T\n", i2, i2)







}
