package main

import "fmt"

func main() {
	// TODO: tulis soalnya
	// array sparse -> kita punya 2 input, string slice dan queries slice
	// kita punya queries yang isinya element, kita akan cek berdasarkan queries tersebut ke strings slice
	// kalau element di strings ada based on queries, kita akan hitung terus di jadiin result slice
	strings := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "bc"}
	fmt.Println("arraySparse")
	fmt.Println(arraySparse(strings, queries)) // expected result [2 1 0]
}

func arraySparse(strings []string, queries []string) []int {
	// approach, buat sebuah maps untuk nanti jadi pembandingnya
	strMap := make(map[string]int)
	// loop strings, cek ke maps, kalo ada key maka valuenya di tambah 1, kalo gaada key, dibuat key baru dengan value 1
	for _, v := range strings {
		_, ok := strMap[v]
		if ok {
			strMap[v]++
		} else {
			strMap[v] = 1
		}
	}
	// setelah dapat map nya, kita buat variable slice untuk resultnya

	var ans []int
	// loop queries, cek ke map, ambil valuenya, kalo keynya kosong, kasih value 0
	for _, v := range queries {
		val, ok := strMap[v]
		if ok {
			ans = append(ans, val)
		} else {
			ans = append(ans, 0)
		}
	}

	return ans
}
