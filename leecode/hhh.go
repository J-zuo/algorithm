package main

import (
	"algorithm/sortalgorithm"
	"fmt"
)

func main() {
	//生成一个随机切片
	sq := sortalgorithm.InitSlice(20)
	fmt.Println(sq, sq.Slice, sq.Lenth)
}
