package main

import (
	"fmt"
	"math"
	"slices"
)

/* Soal LeetCode Squares of a Sorted Array
* Given an integer array nums sorted in non-decreasing order,
* return an array of the squares of each number sorted in non-decreasing order.
*
* Example:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
*
* */

/* Approach
* loop arraynya, terus elementnya di pangkat 2 in pake package math
* setelah semua element di pangkat 2, kemudian di sort pake bantuan package slices
* */

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println("Squares of a Sorted Array")
	fmt.Println(SquaresOfSortedArray(nums))
}

func SquaresOfSortedArray(array []int) []int {
	var ans []int
	for _, v := range array {
		ans = append(ans, int(math.Pow(float64(v), 2)))
	}
	slices.Sort(ans)
	return ans
}
