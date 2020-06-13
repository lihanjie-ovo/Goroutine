package main

import (
	"fmt"
	"time"
)
// 遍历 range

func main() {
	ch := make(chan int, 10)
	go set(ch)
	go get(ch)

	time.Sleep(11 * time.Second)

}

func set(ch chan int) {
	// 记得释放写的操作
	defer close(ch)
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}
}

func get(ch chan int) {
	//遍历通道中的数据
	for v := range ch {
		fmt.Println("get函数中的:", v)
	}

}
