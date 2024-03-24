package main

import "fmt"

/* Soal Leet Code
* Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.

Return any array that satisfies this condition.

Input: nums = [3,1,2,4]
Output: [2,4,3,1]
Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
*
* */

func main() {
	nums := []int{3, 1, 2, 4}
	// tempVar := 0
	// counterOddElement := 0
	// for i := 0; i < len(nums); i++ {
	// 	if i == len(nums)-counterOddElement {
	// 		break
	// 	}
	// 	if nums[i]%2 != 0 {
	// 		tempVar = nums[i]
	// 		nums = append(nums[:i], nums[i+1:]...)
	// 		counterOddElement++
	// 		nums = append(nums, tempVar)
	// 		i--
	// 	}
	// }
	// fmt.Println(nums)

	// another approach
	fmt.Println(nums, "<-- before loop")
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]%2 == 0 {
			i++
			fmt.Println(i, "<-- i inner loop even")
		}
		for i < j && nums[j]%2 == 1 {
			j--
			fmt.Println(j, "<-- j inner loop odd")
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println(nums)
}
