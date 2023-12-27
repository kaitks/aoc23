package day10

import "testing"

func TestP1(t *testing.T) {
	result := part1("input")
	if result != 6754 {
		t.Fail()
	}
}

func TestP1T(t *testing.T) {
	result := part1("test_data")
	if result != 4 {
		t.Fail()
	}
}

func TestP1T1(t *testing.T) {
	result := part1("test_data_1")
	if result != 8 {
		t.Fail()
	}
}
