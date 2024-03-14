package main

import "fmt"

func main() {
	// TODO: tulis soalnya
	//
	leaders := []int{16, 17, 4, 3, 5, 2}
	fmt.Println(findLeaders(leaders))
}

func findLeaders(array []int) []int {
	ans := []int{}
	currMaxLeader := 0

	for i := len(array) - 1; i >= 0; i-- {
		if currMaxLeader < array[i] {
			currMaxLeader = array[i]
			ans = append([]int{currMaxLeader}, ans...)
		}
	}

	return ans
}
