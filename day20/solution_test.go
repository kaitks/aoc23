package day20

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 142863718918201 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 167409079868000 {
		t.Fail()
	}
}
