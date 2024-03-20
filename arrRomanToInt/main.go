package main

import "fmt"

/* Soal EnjoyAlgorithm
* Convert Roman to Integer
* */

func main() {
	myMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	// romans := "XL" // done
	// romans := "MCMIV" // done
	romans := "LVIII"
	// romans := "MCMXCIV" // done
	count := 0

	for i := 0; i < len(romans); i++ {
		curr := string(romans[i])
		if i+1 >= len(romans) {
			count += myMap[curr]
			break
		}
		next := string(romans[i+1])
		if myMap[curr] < myMap[next] {
			count = count + (myMap[next] - myMap[curr])
			i++
		} else if myMap[curr] >= myMap[next] {
			count = count + myMap[curr]
		}
	}

	fmt.Println(count)
}
