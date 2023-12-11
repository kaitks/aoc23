package main

import "testing"

func TestDay5(t *testing.T) {
	result := day5("input")
	if result != 13 {
		t.Fail()
	}
	result1 := day5("test_data")
	if result1 != 27845 {
		t.Fail()
	}
}
