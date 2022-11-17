package main

import "fmt"

func main() {
	str := "hello,世界"

	fmt.Printf("%#v\n", []byte(str))
	for i, c := range []byte(str) {
		fmt.Println(i, c)
	}

	for i := 0; i < len(str); i++ {
		fmt.Printf("%d, %x\n", i, str[i])
	}
}
