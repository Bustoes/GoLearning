package main

import "fmt"

// 实现数据共享可以使用全局变量，但会发生数据竞态的问题。这个属于通过共享内存来实现通信

// go语言提倡通过通信共享内存而不是通过共享内存而实现通信。
// 如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。
// channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

func main() {
	// 有缓冲区通道	又称为一步通道
	var ch1 chan int	// 通道是一种引用类型，需要使用make()函数来进行初始化
	ch1 = make(chan int, 1)
	ch1 <- 10
	x := <- ch1
	fmt.Printf("%d\n", x)

	// 无缓冲区通道	又称为同步通道
	ch2 := make(chan int)
	ch2 <- 20	// 发送值
	// fatal error: all goroutines are asleep - deadlock!
	// 通道被锁死了。因为通道没有缓冲区，那么发送值之后没有位置来暂存值，数据会发生阻塞。只有另外一个goroutine从通道中取值时，才会真正发送进去。
	y := <- ch2
	fmt.Printf("%d\n", y)



}

