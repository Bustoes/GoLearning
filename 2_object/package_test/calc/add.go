package calc

import "fmt"

// 标识符首字母大写时对外可见
// 通常标识符首字母小写，使用驼峰法

//全局声明 -> init() -> main()
//若嵌套导入包，则最深层被导入的包的init()函数会被最先执行

// Name 定义一个测试的全局变量
var Name = "Jack"

//init()函数在导入包时自动执行，没有参数也没有返回值
func init() {
	fmt.Println("init()...")
	fmt.Println(Name)
}

// Add 是一个计算两个int数据类型和的函数
func Add(x int, y int) int {
	return x + y
}
