package main

import (
	"fmt"
	"sync"
)

// 并发：同一时间段内执行多个任务
// 并行：同一时刻执行多个任务

// goroutine 和 channel是Go语言CSP并发模式的重要时间基础

// 创建一个等待对象
var wg sync.WaitGroup

func hello(i int) {
	fmt.Printf("hello Jack %d\n", i)
	wg.Done()	// 发出信号，表示线程已结束，使计数器-1
}

func main() {		// 开启一个主goroutine来执行main()函数
	wg.Add(10000)	// 加入一个信号 计数器+
	//for i := 0; i < 10000; i++ {
	//	go hello(i + 1)
	//}

	for i := 0; i < 10000; i++ {
		go func() {
			fmt.Printf("hello Mike %d\n", i)
			// 这时，可能会出现很多相同的值。这是因为此时函数内部是一个闭包，引用的单纯是外部的i，而没有什么限定。
			// 此线程脱离了主线程。主线程的i循环完毕后，此线程才执行，导致相同的i被很多线程引用
			// 解决：传入参数i
			wg.Done()
		}()
	}

	fmt.Printf("hello main\n")

	// 等待所有goroutine执行完毕再结束main线程
	wg.Wait()


}
