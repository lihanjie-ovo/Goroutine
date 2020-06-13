package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	chTimeout := make(chan struct{})

	//他最多等待 5s
	timeout := 5
	go func() {
		//转换
		time.Sleep(time.Duration(timeout) * time.Second)
		chTimeout <- struct{}{}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 55
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 22
	}()

	select {
	case <-ch1:
		fmt.Println("从ch1中读取了数据")
	case <-ch2:
		fmt.Println("从ch2中读取了数据")
	case <-chTimeout:
		fmt.Println("到时间了~")
	}
}
