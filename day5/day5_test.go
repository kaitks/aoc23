package main

import "testing"

func TestDay5(t *testing.T) {
	result := day5("input")
	if result != 331445006 {
		t.Fail()
	}
	r := day5("test_data")
	if r != 35 {
		t.Fail()
	}
	r1 := day5("test_data_1")
	if r1 != 82 {
		t.Fail()
	}
	r2 := day5("test_data_2")
	if r2 != 43 {
		t.Fail()
	}
}

func TestDay5P2(t *testing.T) {
	result := day5p2("input")
	if result != 331445006 {
		t.Fail()
	}
	r := day5p2("test_data")
	if r != 46 {
		t.Fail()
	}
}

func TestDay5P2V2(t *testing.T) {
	r := day5p2v2("test_data_4")
	if r != 14 {
		t.Fail()
	}
	//r1 := day5p2v2("test_data_4")
	//if r1 != 14 {
	//	t.Fail()
	//}
}

func TestOverLapByRange(t *testing.T) {
	rang := Range{10, 140}
	//resourcesRanges := []Range{{0, 10}, {11, 130}, {131, 139}}
	//r := getOverlapByRange(rang, resourcesRanges)
	//if len(r) != 4 {
	//	t.Fail()
	//}

	resourcesRanges1 := []Range{{0, 2}, {11, 20}}
	r1 := getOverlapByRange(rang, resourcesRanges1)
	if len(r1) != 4 {
		t.Fail()
	}
}

func TestMakeRange(t *testing.T) {
	result := makeRange(1, 10)
	if len(result) != 10 {
		t.Fail()
	}
}
