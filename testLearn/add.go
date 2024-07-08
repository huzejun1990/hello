// @Author huzejun 2024/7/7 20:47:00
package main

import (
	"math/rand"
	"time"
)

func add(a, b int) int {
	return a + b
}

func newSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	//注意，这里在生成切片的时候并没有指定容量
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func newSliceWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	//注意，这里生成切片的时候指定了容量
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
