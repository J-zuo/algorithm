package main

import "fmt"

//select的使用
//什么是select?它是关键字之一,select能监控多个channel的状态

func main() {
	ch1 := make(chan int, 0)

	go func() {
		fmt.Println("goroutine结束")
		fmt.Println(<-ch1)
	}()

	//ch1 <- 666

	select {
	case ch1 <- 6666:
		fmt.Println("llll")
		return
	}

}
