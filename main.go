package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	// fmt.Println("Map-2: ", topKFrequent([]int{2, 2, 3, 3, 3, 4, 6}, 2))
	// fmt.Println("searchInsert", searchInsert([]int{1, 3, 5, 6}, 7))
	// fmt.Println("searchRange", searchRange([]int{8, 8, 8, 4, 2, 6, 5, 4}, 8))
	//fmt.Println("searchInRetated of [2,5,6,0,0,1,2] , 3 ", searchInRetated([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	//fmt.Println("bubblesort algo of [2,1,5,3,9]", bubbleSortAlgorithm([]int{2, 1, 5, 3, 9}))
	// fmt.Println("binary search algorithm", binarySearchAlgorithm([]int{2, 7, 8, 9}, 2))
	//fmt.Println("topKFrequent: ", topKFrequent([]int{2}, 1))
	// fmt.Println("productExceptSelf: ", ProductExceptSelf([]int{1, 2, 3, 4}))
	// fmt.Println("longest : ", longestConsecutive([]int{0, 0, -1}))
	// fmt.Println("lengthOfLongestSus[left]string : ", lengthOfLongestSus[left]string("pwwkew"))
	// fmt.Println("is isPalindrome : ", isPalindrome("A man, a plan, a canal: Panama"))
	// fmt.Println("threeSum : ", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println("nb trap : ", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	// fmt.Println("isValid : ", isValid("({})"))
	// fmt.Println("evalRPN : ", evalRPN([]string{"4", "13", "5", "/", "+"}))
	// fmt.Println("searchMatrix : ", searchMatrix([][]int{
	// 	{1, 3, 5, 7},
	// 	{10, 11, 16, 20},
	// 	{23, 30, 34, 60},
	// }, 3))
	// fmt.Println("the min value : ", searchRotatedSorted([]int{1, 3}, 0))
	// fmt.Println("max profit : ", maxProfit([]int{2, 7, 1, 11, 4}))
	// fmt.Println("MaxSlidingWindow : ", MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// fmt.Println("checkInclusion : ", checkInclusion("hello", "ooolleoooleh"))
	// fmt.Println("climbStairs : ", ClimbStairs(5))
	// fmt.Println("MaxSlidingWindow : ", minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	// fmt.Println("coinChange : ", coinChange([]int{186, 419, 83, 408}, 6249))
	// // rootList := tree.New(4)
	// var smallTe tree.Tree
	// tree.Insert(&smallTe, 4)
	// tree.Insert(&smallTe, 2)
	// tree.Insert(&smallTe, 4)
	// treeStr := smallTe.String()
	// fmt.Println(treeStr)

	fmt.Println("merge interval [[1,3],[2,6],[8,10],[15,18]] ", merge([][]int{
		{1, 4},
		{5, 6},
	}))

}
