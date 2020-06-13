package main

import (
	"fmt"
	"time"
)

// 单向通道
// 并不会影响通道本身,通道还是双向通道，只是函数定义的参数会让函数成为只读或者只写的函数

func main() {
	ch := make(chan int, 10)
	go set1(ch)
	go get1(ch)

	time.Sleep(11 * time.Second)

}

// 写操作    // ch chan<-int  该函数只能执行写操作
func set1(ch chan<- int) {
	// 记得释放写的操作
	defer close(ch)
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}
}

// 读操作   ch <-chan int  该函数只能执行读操作
func get1(ch <-chan int) {
	//遍历通道中的数据
	for v := range ch {
		fmt.Println("get函数中的:", v)
	}

}
