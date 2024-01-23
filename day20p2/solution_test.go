package day20p2

import "testing"

func TestSolution(t *testing.T) {
	result := solution("input")
	if result != 232605773145467 {
		t.Fail()
	}
}
