package day12p2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input", 1)
	if result != 7017 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data", 5)
	if result != 525152 {
		t.Fail()
	}
}

func TestSolutionT1(t *testing.T) {
	result := solution("test_data_1", 5)
	if result != 1 {
		t.Fail()
	}
}

func TestSolutionT2(t *testing.T) {
	result := solution("test_data_2", 3)
	if result != 264342 {
		t.Fail()
	}
}
