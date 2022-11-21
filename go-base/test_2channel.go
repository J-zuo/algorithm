package main

import (
	"fmt"
)

func main() {

	//var ch chan int
	//close(ch) //panic:close of nil channel

	//var ch1 = make(chan int)
	//close(ch1)
	//ch1 <- 12 //向已关闭的ch1中写数据：panic: send on closed channel

	//尝试向一个无缓冲的通道中写入两次
	ch2 := make(chan int)
	//ch2 <- 1
	go func() {
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3
	}()

	//for {
	//	v, ok := <-ch2
	//	fmt.Println(v, ok)
	//}
	for {
		select {
		case v, ok := <-ch2:
			fmt.Println(v, ok)
		}
	}

}
