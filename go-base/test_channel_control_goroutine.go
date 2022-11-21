package main

import (
	"fmt"
	"time"
)

func main() {
	channels := make([]chan int, 10)

	for i := 0; i < cap(channels); i++ {
		channels[i] = make(chan int) //切片中放入一个channel
		go Process(channels[i], i)
	}

	for i, ch := range channels {
		<-ch
		fmt.Println("routine ", i, " quit|")
	}
}

func Process(ch chan int, i int) {
	switch i {
	case 0:
		fmt.Println("执行任务0")
		time.Sleep(time.Second)
		ch <- 1
	case 1:
		fmt.Println("执行任务1")
		time.Sleep(time.Second)
		ch <- 1
	case 2:
		fmt.Println("执行任务2")
		time.Sleep(time.Second)
		ch <- 1
	case 3:
		fmt.Println("执行任务3")
		time.Sleep(5 * time.Second)
		ch <- 1
	default:
		ch <- 1
	}

}
