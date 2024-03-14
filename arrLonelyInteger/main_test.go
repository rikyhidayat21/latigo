package main

import "testing"

func TestLonelyIntegerExist(t *testing.T) {
	array1 := []int32{1, 2, 3, 4, 3, 2, 1}
	result := LonelyInteger(array1)
	if result != 4 {
		t.Error("SALAH!, lonelyInteger seharusnya 4")
	}
}

func TestLonelyIntegerNotExists(t *testing.T) {
	array1 := []int32{1, 1, 2, 2, 3, 3}
	result1 := LonelyInteger(array1)
	if result1 != -1 {
		t.Error("SALAH!, tidak ada LonelyInteger seharusnya -1")
	}
}
