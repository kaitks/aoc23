package day13

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 26957 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 405 {
		t.Fail()
	}
}

func TestSolutionT1(t *testing.T) {
	result := solution("test_data_1")
	if result != 5 {
		t.Fail()
	}
}

func TestSolutionT2(t *testing.T) {
	result := solution("test_data_2")
	if result != 400 {
		t.Fail()
	}
}
