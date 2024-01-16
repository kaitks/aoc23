package day17v2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 1195 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 102 {
		t.Fail()
	}
}

func TestSolutionT1(t *testing.T) {
	result := solution("test_data_1")
	if result != 45 {
		t.Fail()
	}
}
