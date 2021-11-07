package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// 让worker进程在jobs中拿任务。任务做完后继续获取任务。
	for job := range jobs {
		fmt.Printf("worker: %d, started job: %d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker: %d, finished job: %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//开启三个goroutine
	go worker(1, jobs, results)
	go worker(2, jobs, results)
	go worker(3, jobs, results)

	//发送五个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	//输出结果
	//for ret := range results {	// for range 会自动检测通道，如果关闭就会自动退出。但没有关闭。
	for j := 0; j < 5; j++ {
		ret := <- results
		fmt.Println(ret)
	}


}
