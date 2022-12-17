package geo_workshop

import (
	//"fmt"
	//"regexp"
	"testing"
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
// reuturns the summary of the array
func TestBuildGeoProgression(t *testing.T) {
	s := 1
	r := 2
	size := 5
	p := BuildGeoProgression(s, r, size)
	if GeoProSum(s, p[len(p)-1], r) != 31 {
		t.Errorf("Expected %d, got %d", 31, GeoProSum(s, p[len(p)-1], r))
	}
}

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
