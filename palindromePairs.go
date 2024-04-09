package main

func palindromePairs(words []string) [][]int {
	var result [][]int
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j {
				if isPalindromeNewOne(words[i] + words[j]) {
					result = append(result, []int{i, j})
				}
			}
		}
	}
	return result
}

func isPalindromeNewOne(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
