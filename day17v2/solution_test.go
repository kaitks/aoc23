package day17v2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 7741 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 51 {
		t.Fail()
	}
}

func TestSolutionT1(t *testing.T) {
	result := solution("test_data_1")
	if result != 6 {
		t.Fail()
	}
}
