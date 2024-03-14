package main

import "fmt"

func main() {
	// TODO: tulis soalnya
	//
	heightBuilding := []int{2, 3, 4, 5}
	fmt.Println(buildingFacingSun(heightBuilding))
}

func buildingFacingSun(array []int) int {
	countBuilding := 1
	currMaxHeight := array[0]
	for _, v := range array {
		if v > currMaxHeight {
			countBuilding++
			currMaxHeight = v
		}
	}
	return countBuilding
}
