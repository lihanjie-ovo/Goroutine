package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	rlk := &sync.RWMutex{}

	go func() {
		rlk.RLock()
		defer rlk.RUnlock()
		fmt.Println("A 开始执行")
		time.Sleep(time.Second)
		fmt.Println("A 结束执行")
	}()

	go func() {
		rlk.RLock()
		defer rlk.RUnlock()
		fmt.Println("B 开始执行")
		time.Sleep(time.Second)
		fmt.Println("b 结束执行")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("main 结束~")
}
