package day20

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input", 6)
	if result != 869395600 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data", 6)
	if result != 16 {
		t.Fail()
	}
}
