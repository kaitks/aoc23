package main

import "testing"

func TestDay3(t *testing.T) {
	result := day3("test_data")
	if result != 467835 {
		t.Fail()
	}
	result1 := day3("test_data_1")
	if result1 != 1025421 {
		t.Fail()
	}
}
