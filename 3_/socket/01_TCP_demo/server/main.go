package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()	// 关闭连接

	// 针对当前连接做数据的发送和接受操作
	for {
		// 接收
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read conn failed, err: %v\n", err)
			break
		}

		recv := string(buf[:n])
		fmt.Printf("接收到的数据：%s\n", recv)

		// 发送
		conn.Write([]byte("ok"))


	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}

	// 2.等待客户端链接
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err: %v\n", err)
			continue
		}

		// 3.启动一个单独的goroutine去处理连接
		go process(accept)


	}



}
