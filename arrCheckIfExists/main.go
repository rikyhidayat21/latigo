package main

import "fmt"

/* Soal LeetCode
Given an array arr of integers, check if there exist two indices i and j such that :

i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]
*
* */

func main() {
	arr := []int{10, 2, 5, 3}
	myMap := map[int]bool{}
	flag := false
	for _, v := range arr {
		if myMap[v*2] || (v%2 == 0 && myMap[v/2]) {
			flag = true
		}
		myMap[v] = true
	}

	fmt.Println(flag)
}
