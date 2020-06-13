package main

import (
	"fmt"
	"math/rand"
)

// 基本生成器

func main() {

	//使用随机数
	rand := RandIntW()
	// 从通道中读，通道为空的时候，是阻塞的
	fmt.Println(<-rand)
	fmt.Println(<-rand)
	for i := 1; i <= 20; i++ {
		fmt.Println(<-rand)

	}
}

// 其中一个随机数生成函数
func RandIntW() chan int {
	// 带有缓冲的通道，添加通讯的效率，添加异步
	ch := make(chan int, 20)
	// 使用并发函数，快速生成随机数
	go func() {
		// 保证生成器可以一直被调用，可以生成多个随机数
		for {
			// ch未满，则rand.int()可以进入通道，否则会阻塞在这里
			ch <- rand.Int()
		}
	}()
	return ch
}
