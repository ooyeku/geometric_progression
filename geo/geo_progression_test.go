package main

import (
	//"fmt"
	//"regexp"
	"testing"
	//    "os"
)

// TestGeoProgSum create first last and ratio to
// ensure GeoProSum returns summary of described geometric progression array
func TestGeoProgSum(t *testing.T) {
	f := 1
	l := 16
	r := 2
	s := GeoProSum(f, l, r)
	if s != 31 {
		t.Errorf("Expected %d, got %d", 31, s)
	}
}

// TestBuildGeoProgression creates a geometric progression based on size
// ratio and starting variable and uses GeoProSum to test that return value
// returns the summary of the array
func TestBuildGeoProgression(t *testing.T) {
	s := 1
	r := 2
	size := 5
	p := BuildGeoProgression(s, r, size)
	if GeoProSum(s, p[len(p)-1], r) != 31 {
		t.Errorf("Expected %d, got %d", 31, GeoProSum(s, p[len(p)-1], r))
	}
}

// TestLast tests that last value in array is returned when Last function is called
func TestLast(t *testing.T) {
	s := 1
	r := 2
	size := 5 // using the same array build as last two tests
	a := BuildGeoProgression(s, r, size)
	l, _ := Last(a)
	if l != 16 {
		t.Errorf("Expected %d, got %d", 16, l)
	}
}

// TestFirst tests that first value in array is returned when First function is called
func TestFirst(t *testing.T) {
	s := 1
	r := 2
	size := 5
	p := BuildGeoProgression(s, r, size)
	f, _ := First(p)
	if f != s {
		t.Errorf("Expected %d, got %d", s, f)
	}
}

// TestBubbleSort tests that BubbleSort function returns correct first and last variables
func TestBubbleSort(t *testing.T) {
	a := make([]int, 0, 8)
	a = append(a, 34, 888, 7, 56, 467, 73, 3, 1)
	sort := BubbleSort(a)
	first, _ := First(sort)
	last, _ := Last(sort)
	if first != 1 {
		t.Errorf("Expected %d, got %d", 1, first)
	}
	if last != 888 {
		t.Errorf("Expected %d, got %d", 888, last)
	}
}

//func TestFindSortedIndex(t *testing.T){
//    a := make([]int, 0, 8)
//    a = append(a, 55, 43, 8, 5, 22, 5)
//     use os.Pipe to transer println output to other input
//    reader, writer, err := os.Pipe()
//    if err != nil {
//        panic(err)
//    }
//    os.Stdout = writer
//    os.Stderr = writer
//    log.SetOutput(writer)
//    out := make(chan string)
//    go func()
//    FindSortedIndex(43, a)
//
//}

// func TestListElement(array[]int)

// TestAppendListElement tests AppendListElement function to insert and retrieve value into the END of array
func TestAppendListElement(t *testing.T) {
	a := make([]int, 0, 8)
	a = append(a, 34, 888, 7, 56, 467, 73, 3, 1)
	an := AppendListElement(a, 12)
	last, _ := Last(an)
	if last != 12 {
		t.Errorf("Expected %d, got %d", 12, an)
	}
}

func TestInsertEleIndex(t *testing.T) {
	geo := BuildGeoProgression(1, 2, 10)
	geo = InsertEleIndex(geo, 44, 3)
	check := geo[3]
	if check != 44 {
		t.Errorf("Expected %d, got %d", 44, check)
	}
}

func TestDeleteEleIndex(t *testing.T) {
	geo := BuildGeoProgression(1, 2, 10)
	geo = deleteEleIndex(geo, 0)
	check := geo[0]
	if check == 1 {
		t.Errorf("Expected element at first index to not be 1")
	}
}

func TestShuffleOneLeft(t *testing.T) {
	geo := BuildGeoProgression(1, 4, 3)
	geo = ShuffleOneLeft(geo)
	check := geo[0]
	if check != 4 {
		t.Errorf("Expected %d, got %d", 4, check)
	}
}

func TestManualProgSum(t *testing.T) {
	geo := BuildGeoProgression(1, 2, 18)
	l, _ := Last(geo)
	auto := GeoProSum(1, l, 2)
	check := ManualProgSum(geo)
	if check != auto {
		t.Errorf("Expected %d, got %d", auto, check)
	}
}
