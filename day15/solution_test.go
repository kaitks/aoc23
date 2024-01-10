package day15

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 239484 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 145 {
		t.Fail()
	}
}

func TestSolutionT1(t *testing.T) {
	result := solution("test_data_1")
	if result != 0 {
		t.Fail()
	}
}

func TestSolutionT2(t *testing.T) {
	result := solution("test_data_2")
	if result != 0 {
		t.Fail()
	}
}

func TestSolutionT3(t *testing.T) {
	result := solution("test_data_3")
	if result != 0 {
		t.Fail()
	}
}
