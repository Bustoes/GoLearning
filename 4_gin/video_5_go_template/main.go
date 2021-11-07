package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}

	u1 := User{
		Name:   "Jack",
		gender: "男",
		Age:    18,
	}

	err = t.Execute(w, u1)
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
