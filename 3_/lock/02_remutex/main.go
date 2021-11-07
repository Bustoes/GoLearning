package main

import (
	"fmt"
	"sync"
	"time"
)

//读写互斥锁

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()

	wg.Done()
}

func write() {
	rwLock.Lock()
	x++
	time.Sleep(time.Millisecond * 10)
	rwLock.Unlock()

	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++	 {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}
