package main

func searchInsert(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			return i
		}

		if i == len(nums)-1 {
			return len(nums)
		}
	}
	return -1
}
