package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	x := make(map[int]int, len(nums))
	for i, j := range nums {
		x[target-j] = i
	}
	for index, item := range nums {
		index2, ok := x[item]
		if index == index2 {
			continue
		}
		if ok {
			return []int{index, index2}
		}
	}
	return []int{}
}

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}
