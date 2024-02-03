package day23p2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 6418 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	if solution("test_data") != 154 {
		t.Fail()
	}
}
