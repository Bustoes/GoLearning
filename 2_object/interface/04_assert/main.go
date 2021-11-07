package main

import "fmt"

type xxx interface {}

func main() {
	var x xxx

	x = 1

	//返回结果。
	//类型不对时，ok = false, ret = 类型的零值
	ret, ok := x.(bool)
	if ok {
		fmt.Println(ret)
	}else {
		fmt.Println("不是布尔类型")
	}

	//使用switch进行断言
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型, value: %v\n", v)
	case bool:
		fmt.Printf("是布尔类型, value: %v\n", v)
	case int:
		fmt.Printf("是int类型, value: %v\n", v)
	default:
		fmt.Println("猜不到了")
	}


}