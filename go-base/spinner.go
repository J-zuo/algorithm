package main

import (
	"fmt"
	"time"
)

func main() {
	//实现一个在计算斐波拉契第45个元素值时，由于需要一定时间，在此期间我们让用户看到一个可见的标识符来表明程序依然在正常的运行
	go spinner(500 * time.Millisecond) //100毫秒
	const n = 55
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `
main {
console.log(event)
};
-\|/` { //反引号代表
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
