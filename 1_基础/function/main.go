package main

import (
	"fmt"
	"strings"
)

func main() {
	x, y := intOp1(100, 200)
	fmt.Println(x, y)

	//函数可以作为参数、返回值、变量
	test1 := test
	fmt.Printf("%T\n", test1)


	//而在java中，方法参数可以为一个实现类对象，使用匿名内部类或lambda表达式的方式来使函数作为参数。
	calc(100, 200, add)

	//匿名函数
	sayHallo := func() {
		fmt.Println("匿名函数")
	}
	sayHallo()

	func() {
		fmt.Println("立即执行匿名函数")
	}()

	//闭包
	r := a("Jack")
	r()

	rr := makeSuffixFunc(".txt")
	ret := rr("Jack")
	fmt.Println(ret)

	m, n := cal(100)
	result1 := m(200)	//base = 100 + 200
	result2 := n(200)	//base = 300 - 200

	fmt.Println(result1, result2)

}

//支持具有多个返回值的函数
func intOp1(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}
func intOp2(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

//可变参数本质上是一个切片
//切片要放在最后
func intSum(a int, b ...int) int {
	ret := a
	for _, num := range b {
		ret += num
	}
	return ret
}

//函数可以作为参数也可以作为变量

func test() {}

func add(x int, y int) int {
	return x + y
}

func calc(x int, y int, op func(x int, y int) int) int {
	return op(x, y)
}

//defer
//defer语句会对其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，会将延迟的语句按defer定义的逆序进行执行。



//闭包
//函数+外层变量的引用
func a(name string) func() {
	return func() {
		fmt.Println("Hello", name)
	}
}

func makeSuffixFunc(suffix string) func(string2 string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}

func cal(base int) (func(int) int, func(int) int) {
	sum := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return sum, sub
}


