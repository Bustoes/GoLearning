package main

import (
	"fmt"
	"sync"
)

// map不支持并发修改
// fatal error: concurrent map writes

var m = make(map[int]int)
var (
	wg = sync.WaitGroup{}
	sm = sync.Map{}
)

func get(key int) int {
	return m[key]
}

func set(key int, value int) {
	m[key] = value
}

func main() {
	//for i := 0; i < 20; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		set(i, i + 100)		// 设置键值对
	//		fmt.Printf("key: %v, value: %v\n", i, get(i))	// 获取键值对
	//		wg.Done()
	//	}(i)
	//}
	//
	//wg.Wait()


	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			sm.Store(i, i + 100)		// 设置键值对
			value, _ := sm.Load(i)		// 获取键值对
			fmt.Printf("key: %v, value: %v\n", i, value)
			wg.Done()
		}(i)
	}

	wg.Wait()





}
