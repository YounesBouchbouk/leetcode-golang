package main

import "fmt"

func searchRange(nums []int, target int) []int {

	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = append(result, i)
		}
	}

	fmt.Println(result)

	if len(result) == 1 {
		return []int{result[0], result[0]}
	}

	if len(result) == 0 {
		return []int{-1, -1}
	}

	return []int{result[0], result[len(result)-1]}
}
