package day8

import "testing"

func TestDay8P1(t *testing.T) {
	result := day8("input")
	if result != 22357 {
		t.Fail()
	}
}

func TestDay8P1T(t *testing.T) {
	result := day8("test_data")
	if result != 2 {
		t.Fail()
	}
}

func TestDay8P1T1(t *testing.T) {
	result := day8("test_data_1")
	if result != 6 {
		t.Fail()
	}
}
