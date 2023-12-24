package day7

import "testing"

func TestDay7P2(t *testing.T) {
	result := day7p2("input")
	if result != 249515436 {
		t.Fail()
	}
}

func TestDay7P2T(t *testing.T) {
	result := day7p2("test_data")
	if result != 5905 {
		t.Fail()
	}
}
