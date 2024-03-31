package main

import (
	"fmt"
	"sort"
)

// fmt.Println("Map-2: ", topKFrequent([]int{2 , 2, 3 , 3 , 3 , 4 , 6}, 1))

func topKFrequent(nums []int, k int) []int {

	// start with our hash map
	mymap := make(map[int]int)

	// fill the hash map with each value in arraylist as key and it frequent as value
	for _, i := range nums {
		mymap[i] = mymap[i] + 1
	}

	fmt.Println(mymap)

	// if len of map 1 return it
	if len(nums) == 1 {
		return nums
	}

	keys := make([]int, 0, len(mymap))

	//extract key from map
	for key := range mymap {
		keys = append(keys, key)
	}

	fmt.Println(keys)

	// filter keys by their value in map
	sort.SliceStable(keys, func(i, j int) bool {
		return mymap[keys[i]] > mymap[keys[j]]
	})

	// return top k user required
	result := keys[0:k]

	return result
}
