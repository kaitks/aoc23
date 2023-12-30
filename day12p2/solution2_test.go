package day12p2

import "testing"

func TestSolution2(t *testing.T) {
	result := solution2("input", 5)
	if result != 7017 {
		t.Fail()
	}
}

func TestSolution2T(t *testing.T) {
	result := solution2("test_data", 5)
	if result != 525152 {
		t.Fail()
	}
}

func TestSolution2T1(t *testing.T) {
	result := solution2("test_data_1", 5)
	if result != 1 {
		t.Fail()
	}
}

func TestSolution2T2(t *testing.T) {
	result := solution2("test_data_2", 4)
	if result != 264342 {
		t.Fail()
	}
}
