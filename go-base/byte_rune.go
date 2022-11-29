package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "helloworld"
	str1 := "hello世界"

	fmt.Println("str = ", len(str))
	fmt.Println("str1 = ", len(str1))

	fmt.Println("runes =", utf8.RuneCountInString(str))
	fmt.Println("runes =", utf8.RuneCountInString(str1))

}
