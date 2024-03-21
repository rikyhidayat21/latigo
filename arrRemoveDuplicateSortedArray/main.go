package main

import "fmt"

/* Soal LeetCode Remove duplicate element
* Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
* The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
* Example:
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

*/

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(len(nums), "<-- before loop")
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	fmt.Println(len(nums), "<-- after loop")
	fmt.Println(nums)
}
