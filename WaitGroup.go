package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go f1(wg)

	wg.Add(1)
	go f2(wg)

	fmt.Println("main方法中的")

	wg.Add(1)
	go f3(wg)

	// 等待
	wg.Wait()
}

func f1(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("f1", i)
		if i%2 == 0 {
			// 让出执行，优先执行效率更加高的，在回头调用效率较低的
			runtime.Gosched()
		}
	}
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("f2中的:", i)
		if i%2 == 1 {
			// 让出执行，优先执行效率更加高的，在回头调用效率较低的
			runtime.Gosched()
		}
	}
	wg.Done()
}

func f3(wg *sync.WaitGroup) {
	result := 5
	for i := 0; ; i++ {
		fmt.Println("f3", i)
		if result == i {
			wg.Done()
			runtime.Goexit() //提前终止，达到某个业务逻辑就终止
		}
	}
}
