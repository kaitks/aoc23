package day8p2

import "testing"

func TestDay8P2(t *testing.T) {
	result := day8p2("input")
	if result != 22357 {
		t.Fail()
	}
}

func TestDay8P2T(t *testing.T) {
	result := day8p2("test_data")
	if result != 6 {
		t.Fail()
	}
}
