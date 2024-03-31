package main

import "sort"

func merge(intervals [][]int) interface{} {
	if len(intervals) <= 1 {
		return intervals
	}
	sortIntervals(intervals)

	mergedIntervals := make([][]int, 0, len(intervals))
	mergedIntervals = append(mergedIntervals, intervals[0])

	for _, interval := range intervals[1:] {
		if top := mergedIntervals[len(mergedIntervals)-1]; interval[0] > top[1] {
			mergedIntervals = append(mergedIntervals, interval)
		} else if interval[1] > top[1] {
			top[1] = interval[1]
		}
	}
	return mergedIntervals
}

func sortIntervals(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}

// func merge(intervals [][]int) interface{} {

// 	result := [][]int{}

// 	if len(intervals) == 1 {
// 		return intervals
// 	}

// 	for i := 1; i < len(intervals); i++ {
// 		newInterval := []int{}

// 		if intervals[i][0] <= intervals[i-1][1] || intervals[i][0] == intervals[i-1][1]+1 {
// 			newInterval = []int{intervals[i-1][0], intervals[i][1]}
// 		} else {
// 			newInterval = intervals[i]
// 		}
// 		result = append(result, newInterval)
// 	}

// 	if len(result) == 1 {
// 		return result[0]
// 	}
// 	return result
// }
