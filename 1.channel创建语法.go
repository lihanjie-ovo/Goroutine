package main

import "fmt"

func main() {

	// 创建的语法
	// 无缓冲通道
	ch := make(chan int)
	ch2 := make(chan struct{})

	ch3 := make(chan int, 10)
	ch4 := make(chan struct{}, 10)

	fmt.Printf("%T,%v\n", ch, ch)
	fmt.Printf("%T,%v\n", ch2, ch2)
	fmt.Printf("%T,%v\n", ch3, ch3)
	fmt.Printf("%T,%v\n", ch4, ch4)
}
