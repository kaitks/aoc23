package day14

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 110128 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 136 {
		t.Fail()
	}
}

func TestSolutionT1(t *testing.T) {
	result := solution("test_data_1")
	if result != 2 {
		t.Fail()
	}
}

func TestSolutionT2(t *testing.T) {
	result := solution("test_data_2")
	if result != 2 {
		t.Fail()
	}
}
