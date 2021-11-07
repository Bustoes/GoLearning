package main

import "fmt"

/*
两个goroutine
	1. 生成0 - 100数字，发送到ch1中
	2. 从ch1中取出数据计算它的平方，把结果发送到ch2中
 */

func producer(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch1 <-chan int, ch2 chan<- int) {
	// 从通道中取值的方式1
	for {
		temp, ok := <- ch1
		if !ok {
			break
		}
		ch2 <- temp * temp
	}
	close(ch2)

}


func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go producer(ch1)
	go consumer(ch1, ch2)

	// 从通道中取值的方式2
	for ret := range ch2 {
		fmt.Println(ret)
	}




}
















func a(ch2 chan int) {
	ch3 := make(chan int, 100)
	for ret := range ch2 {
		t := <- ch2
		sub := t - ret
		ch3 <- t
		for ch2s := range ch2 {
			ch3 <- ch2s
		}
		close(ch2)
		fmt.Printf("%d, %d\n", ret, sub)
		a(ch3)
	}

}
