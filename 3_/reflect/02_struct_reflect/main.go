package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

// TypeOf()只拿到定义信息，不能拿到实际的值。ValueOf()可以拿到

type student struct {
	Name string	`json:"name" ini:"n"`
	Score int 	`json:"score" ini:"s"`
}

func (s student) Study() string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}
func (s student) Sleep() string {
	msg := "按时睡觉"
	fmt.Println(msg)
	return msg
}

func printField(x interface{}) {
	// 通过反射获取字段信息
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("%v, %T\n", t.Name(), t.Kind())

	// 遍历所有字段
	for i := 0; i < t.NumField(); i++ {
		// 根据结构体字段的索引获取字段
		fieldObj1 := t.Field(i)
		var _ = v.Field(i)
		fmt.Printf("name: %v, type: %v, tag.json: %v\n", fieldObj1.Name, fieldObj1.Type, fieldObj1.Tag.Get("json"))
	}

	// 根据名字获取字段
	fieldObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name: %v, type: %v, tag.json: %v\n", fieldObj2.Name, fieldObj2.Type, fieldObj2.Tag.Get("json"))
	}

}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())		// 获得方法数量

	// v 能操作方法
	// t 能获取信息
	for i := 0; i < v.NumMethod(); i++ {
		// 根据结构体方法的索引获取方法
		methodValue := v.Method(i)
		methodType := t.Method(i)
		fmt.Printf("method name: %s\n", methodType.Name)
		fmt.Printf("method: %s\n", methodValue.Type())

		//通过反射调用方法是，传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}

	// 通过名字访问方法
	t.MethodByName("Sleep")	// (Method, bool)
	methodObj := v.MethodByName("Sleep")	// Value
	fmt.Println(methodObj.Type())


}

func main() {
	stu1 := student{
		Name: "Jack",
		Score: 19,
	}

	printField(stu1)

	printMethod(stu1)

}
