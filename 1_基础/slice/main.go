package main

import "fmt"

func main() {
	//切片	是一个引用类型
	var _ []string
	var a = []int32{1, 2, 3, 4, 5}
	fmt.Println(a)
	//基于数组得到切片
	b := [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}
	c := b[1:4] //左闭右开
	fmt.Println(c)
	//基于切片再次得到切片
	d := c[0:len(c)]
	fmt.Println(d)
	//make()函数得到切片
	e := make([]int, 5, 10) //长度, 容量
	fmt.Println(e)
	fmt.Println(len(c), cap(c))
	fmt.Println("==============================")

	//多个切片表示同一个数组的片段时，它们可以共享数据
	f := b[1:5]
	c[2] = 89
	fmt.Println(f)
	fmt.Printf("%T\n", f)

	//切片的扩容
	var aa []int
	bb := aa
	bb = append(bb, 10)
	fmt.Printf("%v, len:%d, cap:%d, %p\n", aa, len(aa), len(aa), aa)
	fmt.Printf("%v, len:%d, cap:%d, %p\n", bb, len(bb), cap(bb), bb)

	for i := 0; i < 10; i++ {
		aa = append(aa, i)
		fmt.Printf("%v, len:%d, cap:%d, %p\n", aa, len(aa), len(aa), aa)
	}

	//切片的复制
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5, 5)
	copy(s2, s1)
	s3 := s1
	s2[0] = 100

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	//切片的删除
	//删除 i = 2
	str := []string{"北京", "上海", "广州", "深圳"}


	i := 1
	for _, substr := range str[i + 1:] {
		str = append(str[0:i], substr)
		i++
	}
	fmt.Println(str)


}
