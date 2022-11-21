package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 64)
	go Productor(5, ch)

	go Productor(3, ch)
	go Consumer(ch)

	//sig := make(chan os.Signal, 1)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//fmt.Printf("quit(%v)\n", <-sig)

	time.Sleep(1 * time.Millisecond)

}

//生产者，写进管道
func Productor(num int, out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i + num
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		//fmt.Println(v)
		fmt.Printf("%v-", v)
	}
}
