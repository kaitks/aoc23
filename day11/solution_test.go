package day11

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input", 1)
	if result != 567 {
		t.Fail()
	}
}

func TestSolutionP2(t *testing.T) {
	result := solution("input", 1000000)
	if result != 702770569197 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data", 1)
	if result != 374 {
		t.Fail()
	}
}

func TestSolutionT1(t *testing.T) {
	result := solution("test_data_1", 1)
	if result != 8 {
		t.Fail()
	}
}

func TestSolutionT2(t *testing.T) {
	result := solution("test_data_2", 1)
	if result != 10 {
		t.Fail()
	}
}
