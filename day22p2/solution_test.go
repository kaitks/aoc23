package day22p2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 79042 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	if solution("test_data") != 7 {
		t.Fail()
	}
}
