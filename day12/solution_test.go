package day12

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 567 {
		t.Fail()
	}
}

//func TestSolutionP2(t *testing.T) {
//	result := solution("input", 1000000)
//	if result != 702770569197 {
//		t.Fail()
//	}
//}

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
