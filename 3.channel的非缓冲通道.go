package main

import (
	"fmt"
	"time"
)

// 非缓冲通道
func main() {
	ch := make(chan int)

	go ff(ch)
	// 阻塞读取该通道
	<-ch // 等待中，读取无缓冲通道，处于阻塞状态
	fmt.Println("Goroutine执行结束")

}

func ff(ch chan int) {
	fmt.Println("f1开始执行")

	// 处理
	time.Sleep(1 * time.Second)

	ch <- 1
	fmt.Println("f1执行结束")
}
