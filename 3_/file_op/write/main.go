package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func write() {
	fileObj, err := os.OpenFile("file_op/xxx.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()

	str := "追加文本"
	fileObj.Write([]byte(str))
	fileObj.WriteString("hello")
	fmt.Fprintln(fileObj, str)
}

func writeByBufio() {
	fileObj, err := os.OpenFile("file_op/xxx.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()

	writer := bufio.NewWriter(fileObj)

	writer.WriteString("bufio write\n")	//将内容写入缓冲区
	fmt.Fprintln(writer, "hello")

	writer.Flush()	//将缓冲区内容写入磁盘
}

func writeByIoutil() {
	str := "我命由我不由天"
	err := ioutil.WriteFile("file_op/xxx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
}

func main() {
	//write()
	//writeByBufio()
	//writeByIoutil()

	fmt.Fprintln(os.Stdout, "hello")


}
