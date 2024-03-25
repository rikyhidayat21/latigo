package main_test

import (
	"net/http"
	"testing"
	"time"
)

func TestHeadTime(t *testing.T) {
	resp, err := http.Head("https://www.time.gov/")
	if err != nil {
		t.Fatal(err)
	}

	// close connection, always!
	_ = resp.Body.Close()

	now := time.Now().Round(time.Second)
	date := resp.Header.Get("date")
	if date == " " {
		t.Fatal("no Date header received from time.gov")
	}

	// parse time
	dt, err := time.Parse(time.RFC1123, date)
	if err != nil {
		t.Fatalf("error parse: %v", err)
	}

	t.Logf("time.gov: %s (skew %s)", dt, now.Sub(dt))
}
