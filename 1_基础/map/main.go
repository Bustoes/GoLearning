package main

import (
	"fmt"
)

func main() {
	//声明
	var a map[string]int
	fmt.Println(a == nil)
	//初始化
	a = make(map[string]int, 8)	//容量可以不写，说明程序可以自己判断长度？
	fmt.Println(a == nil)

	//添加键值对
	a["name"] = 100
	a["age"] = 200
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("a : %#v\n", a)

	//声明map的同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b : %#v\n", b)

	var scoreMap = make(map[string]int, 8)
	scoreMap["Jack"] = 100
	scoreMap["Mike"] = 200

	//判断某个键存不存在
	value, ok := scoreMap["Jack"]
	fmt.Printf("%v, %t\n", value, ok)

	//遍历
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//删除
	delete(scoreMap, "Jack")
	delete(scoreMap, "Mike")
	fmt.Println(scoreMap)

	//元素类型为map的切片
	mapSlice := make([]map[string]int, 0, 8)	//[nil nil nil nil nil nil nil nil]
	//只完成了切片的初始化，并没有完成map的初始化
	//完成map的初始化
	mapSlice = append(mapSlice, make(map[string]int, 0))//给切片扩容
	//mapSlice[0] = make(map[string]int, 0)	//?容量？map可以自己扩容？
	mapSlice[0]["Jack"] = 100		//自己扩容
	mapSlice[0]["Mike"] = 100
	mapSlice[0]["Marry"] = 100
	fmt.Println(mapSlice)

	//值是切片类型的map
	sliceMap := make(map[string][]int, 8)	//只完成了map的初始化
	sliceMap["中国"] = make([]int, 0, 8)
	sliceMap["中国"] = append(sliceMap["中国"], 11)
	sliceMap["中国"] = append(sliceMap["中国"], 22)
	sliceMap["中国"] = append(sliceMap["中国"], 33)
	fmt.Println(sliceMap)

	//统计字符串中每个单词出现的次数
	//初始化之后，可以随便往里面添加键值对


/*
	//有序输出map
	for i := 1; i <= 50; i++ {
		key := fmt.Sprint("stu", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	scores := make([]string, 0, 100)
	for k := range scoreMap {
		scores = append(scores, k)
	}

	for _, score := range scores {
		fmt.Println(score, scoreMap[score])
	}
*/

}
