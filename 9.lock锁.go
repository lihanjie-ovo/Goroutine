package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	mt := &sync.Mutex{}

	go func() {
		//锁定该routine
		mt.Lock()
		// 解锁该routine
		defer mt.Unlock()

		fmt.Println("A 开始执行")
		time.Sleep(2 * time.Second)
		fmt.Println("A 结束执行")

	}()
	go func() {
		// 上锁
		mt.Lock()
		// 解锁
		defer mt.Unlock()

		fmt.Println("b开始执行")
		time.Sleep(2 * time.Second)
		fmt.Println("b结束执行")

	}()
	time.Sleep(5 * time.Second)

	fmt.Println("main 结束！")
}
