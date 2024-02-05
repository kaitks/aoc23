package day24

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input", 200000000000000, 400000000000000)
	if result != 15107 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	if solution("test_data", 7, 27) != 2 {
		t.Fail()
	}
}
