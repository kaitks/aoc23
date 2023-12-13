package main

import "testing"

func TestDay5(t *testing.T) {
	result := day5("input")
	if result != 331445006 {
		t.Fail()
	}
	r := day5("test_data")
	if r != 35 {
		t.Fail()
	}
	r1 := day5("test_data_1")
	if r1 != 82 {
		t.Fail()
	}
	r2 := day5("test_data_2")
	if r2 != 43 {
		t.Fail()
	}
}
