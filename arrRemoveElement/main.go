package main

import "fmt"

/* Soal LeetCode - Remove Element
* Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
* The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
* Example:
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).

*
* */

// this will apply remove element in-place
func main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			// kalo element == val, di parameter 1 - kita akan append start dari 0 - index ke i
			// sementara di parameter ke 2 - kita akan pecah slicenya start dari element index ke i + 1 sampai ujung
			// dan terakhir kita pake variadic untuk ngecombine si slicenya
			nums = append(nums[:i], nums[i+1:]...)
			i-- // kenapa ini kok index nya di kurangin lg?
			// supaya kalo misal ada val yang berjejeran, dia iterasinya gk kelewat
		}
	}
	fmt.Println(nums)
}
