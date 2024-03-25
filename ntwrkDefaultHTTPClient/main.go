package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	resp, err := http.Head("https://www.time.gov/")
	if err != nil {
		log.Fatal(err)
	}

	// close connection, always
	defer resp.Body.Close()

	now := time.Now().Round(time.Second)
	date := resp.Header.Get("date")
	if date == " " {
		log.Fatal("no date header received from time.gov")
	}

	log.Println(date, "<-- date before parse")

	// parse time
	dt, err := time.Parse(time.RFC1123, date)
	if err != nil {
		log.Fatalf("Error parse date: %v", err)
	}

	log.Printf("time.gov: %s (skew %s)", dt, now.Sub(dt))
}
