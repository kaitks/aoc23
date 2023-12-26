package day9

import "testing"

func TestP1(t *testing.T) {
	result := part1("input")
	if result != 1921201931 {
		t.Fail()
	}
}

func TestP1T(t *testing.T) {
	result := part1("test_data")
	if result != 114 {
		t.Fail()
	}
}
