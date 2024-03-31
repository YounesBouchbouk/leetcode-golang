package main

func searchInRetated(nums []int, target int) bool {
	pivot := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			pivot = i + 1
		}
	}

	if searchWithBinarySearch(nums[:pivot], target) || searchWithBinarySearch(nums[pivot:], target) {
		return true
	}

	return false
}

func searchWithBinarySearch(nums []int, target int) bool {

	left, right := 0, len(nums)

	for left < right {
		midd := (left + right) / 2

		if nums[midd] == target {
			return true
		}

		if nums[midd] < target {
			left = midd + 1
		} else {
			right = midd
		}

	}

	return false
}
