package day6

import "testing"

func TestDay6P1(t *testing.T) {
	result := day6("input")
	if result != 220320 {
		t.Fail()
	}
}

func TestDay6P1T(t *testing.T) {
	result := day6("test_data")
	if result != 288 {
		t.Fail()
	}
}

func TestDay6P2(t *testing.T) {
	result := day6p2("input")
	if result != 34454850 {
		t.Fail()
	}
}

func TestDay6P2T(t *testing.T) {
	result := day6p2("test_data")
	if result != 71503 {
		t.Fail()
	}
}
