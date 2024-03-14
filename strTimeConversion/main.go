package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// TODO: tulis soalnya
	fmt.Println(timeConversion("07:05:45PM")) // -> expect to be: 19:05:45
}

func timeConversion(str string) string {
	am := strings.HasSuffix(str, "AM")
	pm := strings.HasSuffix(str, "PM")
	withoutSuffix := str[:len(str)-2]
	timeSlice := strings.Split(withoutSuffix, ":")
	hh, mm, ss := timeSlice[0], timeSlice[1], timeSlice[2]
	hhInt, err := strconv.Atoi(hh) // convert hh to integer
	if err != nil {
		panic("cannot convert string to integer")
	}

	if am && hhInt == 12 {
		hhInt = 0
	}

	if pm && hhInt != 12 {
		hhInt += 12
	}

	hh = fmt.Sprint(hhInt)
	return fmt.Sprintf("%02s:%02s:%02s", hh, mm, ss) // %02s supaya kalau nilai nya cuman 5, dia akan nambahin 0 didepannya jadi 05
}
