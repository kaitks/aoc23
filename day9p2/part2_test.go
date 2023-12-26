package day9

import "testing"

func TestP2(t *testing.T) {
	result := part2("input")
	if result != 1921201931 {
		t.Fail()
	}
}

func TestP2T(t *testing.T) {
	result := part2("test_data")
	if result != 114 {
		t.Fail()
	}
}
