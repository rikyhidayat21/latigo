package main

import "fmt"

/* Soal LeetCode Replace Elements with Greatest Element on Right Side

* Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

After doing so, return the array.

Example:
Input: arr = [17,18,5,4,6,1]
Output: [18,6,6,6,1,-1]
* */

// TODO: fix this
func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(arr)
}

func replaceElement(arr []int) []int {
	max := 0

	// edge case
	if len(arr) == 1 {
		arr = append(arr[:1], -1)
	}

	for i := len(arr) - 1; i > 0; i-- {
		if i == len(arr)-1 {
			arr = append(arr[:i], -1)
			continue
		}

		if arr[i] > arr[i+1] {
			max = arr[i]
			arr = append(arr[:i], max)
		}
	}

	return arr
}
