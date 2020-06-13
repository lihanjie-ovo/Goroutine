package main

import (
	"fmt"
	"time"
)

// 缓冲通道

func main() {
	ch := make(chan int,10)
	go chf1(ch)
	go chf2(ch)

	time.Sleep(20 * time.Second)
}

func chf1(ch chan int) {
	for i := 0; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond) // 1000 ms = 1s  100ms = 0.1s
		data := 50 + i
		fmt.Println("f1产生的数:", data)
		ch <- data
	}
}

func chf2(ch chan int) {
	for i := 0; i <= 10; i++ {
		data := <-ch
		fmt.Println("f2处理的数：", data)
		time.Sleep(time.Second)
	}
}
