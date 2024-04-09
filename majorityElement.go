package main

func majorityElement(nums []int) []int {

	frHash := make(map[int]int)
	arrResult := []int{}

	for _, num := range nums {
		frHash[num]++
	}

	for key, value := range frHash {
		if value > len(nums)/3 {
			arrResult = append(arrResult, key)
		}
	}
	return arrResult
}
