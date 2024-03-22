package main

import (
	"fmt"
)

func main() {
	nilai := 0
	kelasA := true
	if kelasA {
		nilai, err := jadiSepuluh()
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(nilai)
	} else {
		nilai, err := jadiSebelas()
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(nilai)
	}

	fmt.Println(nilai)
}

func jadiSepuluh() (int, error) {
	return 10, nil
}

func jadiSebelas() (int, error) {
	return 11, nil
}

func jadiDuapuluh() int {
	return 11
}
