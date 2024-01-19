package day19p2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 432788 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 19114 {
		t.Fail()
	}
}
