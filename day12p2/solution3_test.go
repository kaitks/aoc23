package day12p2

import "testing"

func TestSolution3(t *testing.T) {
	result := solution3("input", 5)
	if result != 527570479489 {
		t.Fail()
	}
}

func TestSolution3T(t *testing.T) {
	result := solution3("test_data", 5)
	if result != 525152 {
		t.Fail()
	}
}

func TestSolution3T1(t *testing.T) {
	result := solution3("test_data_1", 5)
	if result != 1 {
		t.Fail()
	}
}

func TestSolution3T2(t *testing.T) {
	result := solution3("test_data_2", 4)
	if result != 30871702 {
		t.Fail()
	}
}
