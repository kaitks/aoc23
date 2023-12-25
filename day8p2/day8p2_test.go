package day8v2

import "testing"

func TestDay8P2(t *testing.T) {
	result := day8p2("input")
	if result != 22357 {
		t.Fail()
	}
}

func TestDay8P2T(t *testing.T) {
	result := day8p2("test_data")
	if result != 2 {
		t.Fail()
	}
}

func TestDay8P2_1(t *testing.T) {
	result := day8p2("test_data_1")
	if result != 6 {
		t.Fail()
	}
}
