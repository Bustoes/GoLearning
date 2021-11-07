package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadFile("E:\\HTML\\常用标签\\表单标签.html")

	_, _ = fmt.Fprintf(w, string(bytes))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err: %v", err)
		return
	}

}
