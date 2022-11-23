package main

import (
	"github.com/pkg/profile"
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {

	//1 使用 + 拼接字符串
	//s := ""
	//for i := 0; i < n; i++ {
	//	s += randomString(n)
	//}
	//return s
	//2 使用stings.Builder拼接

	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(randomString(n))
	}
	return s.String()

}

func main() {
	//defer profile.Start().Stop()
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	concat(1000)
}
