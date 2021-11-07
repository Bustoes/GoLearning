package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
在一些前后端不分离的Web架构中，我们通常需要在后端将一些数据渲染到HTML文档中，
从而实现动态的网页（网页的布局和样式大致一样，但展示的内容并不一样）效果。

我们这里说的模板可以理解为事先定义好的HTML文档文件，模板渲染的作用机制可以简单理解为
文本替换操作–使用相应的数据去替换HTML文档中事先准备好的标记。

Go语言模板引擎的使用可以分为三部分：定义模板文件、解析模板文件和模板渲染.

*/

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1.定义模板
	// 2.解析模板
	t, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}

	// 3.渲染模板
	name := "Jack"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("execute template failed, err: %v\n", err)
		return
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}

}
