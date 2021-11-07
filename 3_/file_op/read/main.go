package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

func readAll() {
	//打开文件
	fileObj1, err := os.Open("file_op/xx.txt")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj1.Close()

	//读取文件
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj1.Read(tmp)

		if err == io.EOF {		//End Of File
			//把当前读到的字节打印后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}

		fmt.Printf("read %d bytes from file.\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

func readByBufio() {
	//打开文件
	fileObj1, err := os.Open("file_op/xx.txt")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj1.Close()

	reader := bufio.NewReader(fileObj1)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {		//End Of File
			return
		}
		if err != nil {
			fmt.Printf("read file by bufio failed, err:%v\n", err)
			return
		}

		fmt.Print(line)
	}
}

// ioutil
func readByIoutil() {
	content, err := ioutil.ReadFile("file_op/xx.txt")
	if err != nil {
		fmt.Printf("read file by ioutil failed, err:%v\n", err)
	}

	fmt.Printf("%s\n", string(content))
}

func main() {
	//readAll()
	//readByBufio()
	readByIoutil()

}

