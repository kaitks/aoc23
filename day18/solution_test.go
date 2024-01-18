package day18

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 61661 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 62 {
		t.Fail()
	}
}
