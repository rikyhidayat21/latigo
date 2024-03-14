package main

import "testing"

func TestLonelyInteger(t *testing.T) {
	array1 := []int32{1, 2, 3, 4, 3, 2, 1}
	result := LonelyInteger(array1)
	if result != 4 {
		t.Error("SALAH!, lonelyInteger seharusnya 4")
	}
}
