package main

import "fmt"

func main() {
	//ch := make(chan int)
	////ch <- 1
	//go func() {
	//	fmt.Println(<-ch)
	//}()
	//ch <- 1

	c := make(chan int)
	go read(c)
	write(c)
}

func read(c <-chan int) {
	fmt.Println("read:", <-c)
}

func write(c chan<- int) {
	c <- 0
}
