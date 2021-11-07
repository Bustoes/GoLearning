package main

import (
	"fmt"
	"reflect"
)

/*
反射：程序在运行期间，对程序本身访问和修改的能力
程序在编译时，变量被转换为内存地址，变量名不会被编译器写入可执行文件中。运行程序时，程序无法获取自身的信息
反射使程序可以在运行时获取数据的反射信息，并判断它们。
 */

// Elem() 用来取指针存储的数据

type Cat struct {}

type Dog struct {}

func reflectType(x interface{}) {
	//不知道在调用函数时，会传入什么类型的参数
	// 方式1：通过类型断言
	// 方式2：通过反射
	obj := reflect.TypeOf(x)	// *reflect.rtype
	fmt.Printf("%v, %v, %v\n", obj, obj.Name(), obj.Kind())

}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)		// // reflect.Value
	fmt.Printf("%v, %T\n", v, v)

	// 得到传入时类型的变量
	k := v.Kind()
	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("%v, %T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v, %T\n", ret, ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(v.Elem().Int() + 10)
	case reflect.Float32:
		v.Elem().SetFloat(3.21)
	}
}


func main() {
	//var a float32 = 1.234
	//reflectType(a)
	//var b int8 = 10
	//reflectType(b)
	//
	//var c Cat
	//reflectType(c)
	//var d Dog
	//reflectType(d)
	//
	//var e []int
	//reflectType(e)
	//var f []string
	//reflectType(f)

	////value
	//var aa int32 = 100
	//reflectValue(aa)
	//var bb float32 = 1.234
	//reflectValue(bb)

	// set value
	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Printf("%v\n", aaa)

	// IsNil()
	// 报告v的值是否为nil。v的类型必须是通道、函数、接口、映射、指针、切片之一
	// IsValid()
	// 返回v是否有值。如果v是零值或v是引用类型nil时，会返回true，此时v除了IsValid/String/Kind之外的方法都会导致panic
	var p *int
	var i int
	fmt.Printf("var p *int is nil: %v\n", reflect.ValueOf(p).IsNil())
	fmt.Printf("var p *int is nil: %v\n", reflect.ValueOf(p).IsValid())
	//fmt.Printf("var p *int is nil: %v\n", reflect.ValueOf(i).IsNil())
	fmt.Printf("var p *int is nil: %v\n", reflect.ValueOf(i).IsValid())




}


