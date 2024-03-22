package main

import "fmt"

/* Soal LeetCode Move Zeroes
* Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
*
* */

// TODO: fix this to can be handle [0,0,1]
func main() {
	// nums := []int{0, 1, 0, 3, 12}
	nums := []int{0, 0, 1}
	j := 0
	// loop array, cek kalo nums[i] == 0, maka biarin
	// kalo nums[i] != 0, maka swap antara nums[i], nums[j] = nums[j], nums[i]
	// notes, j disini sebagai identifier untuk index ya!
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	fmt.Println(nums)
}
