package main

import (
	"fmt"
	"math"
)

func binarySearchAlgorithm(arraylist []int, vl int) int {

	end := len(arraylist) - 1

	beg := 0

	for beg <= end {
		middle := math.Floor((float64(beg) + float64(end)) / 2)

		fmt.Println(middle)
		if arraylist[int(middle)] == vl {
			return int(middle)
		}

		if arraylist[int(middle)] < vl {
			beg = int(middle) + 1
		} else {
			end = int(middle)
		}

	}

	return -1
}
