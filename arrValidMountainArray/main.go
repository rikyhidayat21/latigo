package main

import "fmt"

/* Soal LeetCode valid mountain array
* Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
*
* */

func main() {
	arr := []int{0, 3, 2, 1}
	fmt.Println(validMountain(arr))
}

func validMountain(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	peakIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[peakIndex] == arr[i] {
			return false
		}

		if arr[peakIndex] < arr[i] {
			peakIndex++
		}
	}

	// handle if the peak is end or first of element
	if (arr[peakIndex] == arr[len(arr)-1]) || (arr[peakIndex] == arr[0]) {
		return false
	}

	// handle from right to left
	for i := len(arr) - 1; i > peakIndex; i-- {
		if arr[i] == arr[i-1] {
			return false
		}
	}

	return true
}
