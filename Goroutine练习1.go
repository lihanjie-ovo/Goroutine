package main

import "fmt"

// 拆分成10个线程 求出1-10000的和 格式:1-1000 1001-2000 2001-3000 ....
func main() {

	mk := make(chan int)

	for i := 0; i < 10; i++ {
		go func(min int, max int) {
			sum := 0
			for j := min; j <= max; j++ {
				sum += j
			}
			fmt.Println(min, max, ":", sum)
			mk <- sum
		}(1+i*1000, (1+i)*1000)
	}

	sum := 0
	counter := 1
	for v := range mk {
		sum += v
		if counter == 10 {
			close(mk)
		}
		counter++
	}
	fmt.Println("1-10000总和是:", sum)

}
