package split

import "strings"

// 测试
/*
go test命令是一个按照一定约定和组织的测试代码的驱动程序。
在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。
在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。
go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。
*/

// Split 将s按照sep分割，并返回一个字符串切片
func Split(s, sep string) (ret []string) {
	//添加了一行代码，将运行速度提升了大约一倍。
	ret = make([]string, 0, strings.Count(s, sep) + 1)

	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx + len(sep):]
		idx = strings.Index(s, sep)
	}

	// 追加剩下的字符串
	ret = append(ret, s)

	return
}







