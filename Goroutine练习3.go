package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chRand := RandInt()
	fmt.Println(chRand)
	fmt.Println(<-chRand) //从管道中读，在通道为空的时候，是阻塞状态
	for i := 1; i <= 20; i++ {
		fmt.Println(<-chRand)
	}
}

// 增强型生成
func RandInt() chan int {
	ch := make(chan int, 6)
	//	并发的多路复用随机选择
	go func() { // 并发执行，不会阻塞主Main goroutine
		for {
			select { // 若case都没发生，处于阻塞状态
			case ch <- <-RandIntA(): //先将RanIntA中读取的管道数据,写入ch管道
			case ch <- <-RandIntB(): //先将RanIntB中读取的管道数据,写入ch管道
			case ch <- <-RandIntC(): //先将RanIntC中读取的管道数据,写入ch管道
			}
		}
	}()

	return ch
}

// 其中一个随机数生成函数
func RandIntA() chan int {
	// 带有缓冲的通道，添加通讯的效率，添加异步
	ch := make(chan int, 2)
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

// 其中一个随机数生成函数
func RandIntB() chan int {
	// 带有缓冲的通道，添加通讯的效率，添加异步
	ch := make(chan int, 2)
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

// 其中一个随机数生成函数
func RandIntC() chan int {
	// 带有缓冲的通道，添加通讯的效率，添加异步
	ch := make(chan int, 2)
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
