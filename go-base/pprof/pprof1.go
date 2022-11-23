package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

//来测试一下pprof性能分析
func main() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}

func generate(n int) []int {
	//生成一组长度为n整型切片
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(1000))
	}
	return nums
}

func bubbleSort(nums []int) {
	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-1; j++ {
			//fmt.Println(nums[j], nums[j-1])
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
