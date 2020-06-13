package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go step1(ch)
	go step2(ch)

	time.Sleep(2 * time.Second)
}

func step1(ch chan int) {
	fmt.Println("步骤1开始执行")
	time.Sleep(1 * time.Second)
	data := 50
	ch <- data
}

func step2(ch chan int) {
	fmt.Println("步骤2开始执行")
	data2 := <-ch
	fmt.Println("step1的数据传输:", data2)
}
