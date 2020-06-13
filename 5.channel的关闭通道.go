package main

import (
	"fmt"
	"time"
)

// 关闭通道

func main() {
	ch := make(chan int, 10)
	go cha1(ch)
	go cha2(ch)

	time.Sleep(20 * time.Second)
}

func cha1(ch chan int) {
	// 做完写操作之后关闭通道,
	defer close(ch)
	for i := 0; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond) // 1000 ms = 1s  100ms = 0.1s
		data := 50 + i
		fmt.Println("f1产生的数:", data)
		ch <- data
	}
}

func cha2(ch chan int) {
	for i := 0; i <= 10; i++ {
		data := <-ch
		fmt.Println("f2处理的数：", data)
		time.Sleep(time.Second)
	}
}
