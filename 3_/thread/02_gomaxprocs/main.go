package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
	wg.Done()
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(4)	// 只占用一个CPU核心
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
