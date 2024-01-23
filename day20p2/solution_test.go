package day20p2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 869395600 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 32000000 {
		t.Fail()
	}
}

func TestSolutionT1(t *testing.T) {
	result := solution("test_data_1")
	if result != 11687500 {
		t.Fail()
	}
}
