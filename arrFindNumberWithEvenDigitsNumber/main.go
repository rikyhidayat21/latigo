package main

import "fmt"

/* Soal Find Number With Even Digits Number
* Given an array nums of integers,
* return how many of them contain an even number of digits.
* Example:
Input: nums = [555,901,482,1771]
Output: 1
Explanation:
Only 1771 contains an even number of digits.
*
* */

func main() {
	fmt.Println("Find Number With Even Digits Number")
	nums := []int{555, 901, 482, 1771}
	fmt.Println(FindNumberWithEvenDigitsNumber(nums))
}

func FindNumberWithEvenDigitsNumber(nums []int) int {
	count := 0
	temp := ""
	for _, v := range nums {
		temp = fmt.Sprintf("%d", v)
		if len(temp)%2 == 0 {
			count++
		}
	}

	return count
}
