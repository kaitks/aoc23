package day21p2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input", 26501365)
	if result != 584211423220706 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	if solution("test_data", 6) != 16 {
		t.Fail()
	}
	if solution("test_data", 10) != 50 {
		t.Fail()
	}
	if solution("test_data", 50) != 1594 {
		t.Fail()
	}
	if solution("test_data", 100) != 6536 {
		t.Fail()
	}
	if solution("test_data", 500) != 167004 {
		t.Fail()
	}
	if solution("test_data", 1000) != 668697 {
		t.Fail()
	}
	if solution("test_data", 5000) != 16733044 {
		t.Fail()
	}
}
