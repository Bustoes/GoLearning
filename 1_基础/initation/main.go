package main

import (
	"fmt"
	"strings"
)

func main() {
	var _ int  //声明未初始化
	var s2 = 1 //声明且初始化
	var s3 int32 = 3
	s1 := "一串.字符串.被分开"
	fmt.Println(s1, s2, s3)
	fmt.Println("\t制表符和\r换行符")
	fmt.Println(`反引号
\n`)
	fmt.Println(len(s1))
	fmt.Println(strings.Split(s1, "."))
	fmt.Printf("%T\n", strings.Split(s1, "."))
	fmt.Println(strings.Contains(s1, "."))

	s4 := []string{"how", "do", "you", "do"}
	fmt.Println(s4)
	fmt.Println(strings.Join(s4, "+"))

	c1 := 97
	fmt.Printf("%c, %d\n", c1, c1)

	s := "hello茶烧"
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%c, %d\n", s[i], s[i])
	//}

	for _, i2 := range s { //增强for循环
		fmt.Printf("%c, %d\n", i2, i2)
	}

	ss := "abcde"
	i := 0
	for ; i < len(ss); i++ {
		fmt.Println(ss[i])
	}

	//数组的声明
	var _ [3]int
	var _ [3][2]string
	//编译器推导数组的长度
	var w = [...]int32{1, 2, 4}
	var q = [...]string{3: "11", 4: "13", 5: "12"}
	w = [...]int32{1, 2, 3}
	fmt.Println(q)
	fmt.Println(w)
	//二维数组
	array := [3][2]string {
		{"北京", "西安"},
		{"上海", "重庆"},
		{"杭州", "成都"},
	}
	fmt.Println(array[2])

}
