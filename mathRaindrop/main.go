package main

import "fmt"

func main() {
	// TODO: tulis soalnya
	fmt.Println(raindrops(30))
}

func raindrops(num int) string {
	sounds := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	result := ""

	for key, value := range sounds {
		if num%key == 0 {
			result += value
		}
	}

	if result == "" {
		return fmt.Sprint(result)
	}

	return result
}
