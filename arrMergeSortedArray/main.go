package main

import (
	"fmt"
	"slices"
)

/* Soal LeetCode Merge Sorted Array
*
* You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
* and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
* Merge nums1 and nums2 into a single array sorted in non-decreasing order.
*
* Example:
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
*
* */

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	// n := 3 // kita gaperluin si parameter n ini buat execute logicnya
	oriLenNums1 := len(nums1)

	counterNums2 := 0
	// kita loop start dari value m, kenapa? supaya element nums1 sepanjang m masih tetap ada
	for i := m; i < oriLenNums1; i++ {
		// disini kita akan nambahin current index dengan value dari nums2
		// pake variable counter, karena di loop ini kita start dari m, bukan dari 0
		nums1 = append(nums1[:i], nums2[counterNums2])
		counterNums2++
	}

	slices.Sort(nums1)
	fmt.Println(nums1)
}
