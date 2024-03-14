package main

import (
	"fmt"
	"strconv"
)

func main() {
	// INFO: find lonelyInteger
	// Given an array of integers, where all elements but one occur twice, find the unique element.
	// []int{1,2,3,4,3,2,1} => expect 4

	aoi := []int32{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(LonelyInteger(aoi))
}

// approach
// karena ini nyari integer yang 'single'
// kita bisa gunain maps buat nyimpen valuenya
//
// step 1. init map
// step 2. loop array -> check based on iterasi udah ada key nya belom?
//
//	kalo udah tambah value, kalo belom bikin key baru di map dan assign value 1
//
// step 3. loop maps -> check kalo ada yang value nya 1, langsung return key nya
func LonelyInteger(array []int32) int32 {
	// step 1
	maps := make(map[string]int32)

	// step 2
	for _, v := range array {
		iToStr := fmt.Sprintf("%d", v)
		_, ok := maps[iToStr]
		if ok {
			maps[iToStr]++
		} else {
			maps[iToStr] = 1
		}
	}

	var ans int32
	for k, v := range maps {
		if v == 1 {
			s, _ := strconv.Atoi(k)
			ans = int32(s)
		} else {
			ans = -1
		}
	}

	return ans
}
