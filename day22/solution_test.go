package day22

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 636 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	if solution("test_data") != 5 {
		t.Fail()
	}
}
