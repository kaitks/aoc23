package day18p2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 111131796939729 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 952408144115 {
		t.Fail()
	}
}
