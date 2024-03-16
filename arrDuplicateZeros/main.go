package main

import (
	"fmt"
	"slices"
)

/* Soal LeetCode DuplicateZeros
* Given a fixed-length integer array arr,
* duplicate each occurrence of zero,
* shifting the remaining elements to the right.

* Note that elements beyond the length of the original array are not written.
* Do the above modifications to the input array in place and do not return anything.

*
* */

func main() {
	array := []int{1, 0, 2, 3, 0, 4, 5, 0}
	// kalo mau in-place modified array, kita pake len arraynya dlu
	fmt.Println(array, "<-- array root")
	fmt.Println(array[:2], "idx: 1 + 1")
	fmt.Println(array[1:7], "idx: 1 sampai oriLen - 1")
	oriLen := len(array)

	counter := 0
	for i := 0; i < oriLen; i++ {
		// cek kalo elementnya = 0, insert ke current elementnya
		fmt.Printf("iterasi ke: %d\n", counter)
		counter++
		fmt.Println("index ke: ", i)
		if array[i] == 0 {
			fmt.Println(array, "<--- array inside if before append")
			fmt.Println(array[:6], "<-- ")
			fmt.Println(array[5:7], "<--")
			// ohh jadi dengan append kaya dibawah ini, dia akan mecah si slicenya jadi 2 bagian
			// nah terus di params yg kedua kan pake ..., nah dia akan nambahin dari slice bagian 1 + slice bagian 2
			array = append(array[:i+1], array[i:oriLen-1]...)
			fmt.Println(array, "<--- array inside if after append")
			i++
		}
		fmt.Println(array, "<== array inside loop")
	}

	fmt.Println(array)
}

func DuplicateZeroNotModifiedParams(array []int) []int {
	for idx, val := range array {
		if val == 0 {
			array = slices.Insert(array, idx+1, 0)
			idx += 2
		}
	}

	return array[:8]
}
