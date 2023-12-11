package main

import "testing"

func TestDay4(t *testing.T) {
	result := day4("test_data")
	if result != 13 {
		t.Fail()
	}
	result1 := day4("input")
	if result1 != 27845 {
		t.Fail()
	}
}

func TestDay4P2(t *testing.T) {
	result := day4p2("test_data")
	if result != 30 {
		t.Fail()
	}
	result1 := day4p2("input")
	if result1 != 9496801 {
		t.Fail()
	}
}
