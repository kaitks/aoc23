package day12

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 7017 {
		t.Fail()
	}
}

func TestSolutionT(t *testing.T) {
	result := solution("test_data")
	if result != 21 {
		t.Fail()
	}
}

func TestSolutionT1(t *testing.T) {
	result := solution("test_data_1")
	if result != 1 {
		t.Fail()
	}
}

func TestSolutionT2(t *testing.T) {
	result := solution("test_data_2")
	if result != 10 {
		t.Fail()
	}
}

func TestSolutionT3(t *testing.T) {
	result := solution("test_data_3")
	if result != 1 {
		t.Fail()
	}
}

func TestSolutionT4(t *testing.T) {
	result := solution("test_data_4")
	if result != 2 {
		t.Fail()
	}
}

func TestSolutionT5(t *testing.T) {
	result := solution("test_data_5")
	if result != 1 {
		t.Fail()
	}
}

func TestSolutionT6(t *testing.T) {
	result := solution("test_data_6")
	if result != 5 {
		t.Fail()
	}
}
