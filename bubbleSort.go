package main

func bubbleSortAlgorithm(arraylist []int) []int {

	for i := 0; i < len(arraylist); i++ {
		for j := 0; j < len(arraylist)-1-i; j++ {
			if arraylist[j] > arraylist[j+1] {
				tmp := arraylist[j]
				arraylist[j] = arraylist[j+1]
				arraylist[j+1] = tmp
			}
		}
	}

	return arraylist
}
