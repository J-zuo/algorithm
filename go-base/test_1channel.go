package main

import "fmt"

func main() {

	//运行结果是不固定的，不知道哪个协程先哪个协程后

	quit := make(chan struct{})
	data := make(chan int)

	go func() {
		fmt.Println("编号为1的goroutine")
		data <- 11
	}()

	go func() {
		fmt.Println("编号为2的goroutine")
		data <- 22
	}()

	go func() {
		defer close(quit)
		fmt.Println("编号为3的goroutine")
		fmt.Println(<-data)
		fmt.Println(<-data)
	}()

	data <- 0
	<-quit

}
