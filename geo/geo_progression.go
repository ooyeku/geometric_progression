package main

import (
	"fmt"
)

// Geo set sum implementation
var r = 5 // ratio
var f = 4 // first in set
var f1 = nextInSet(f, r)
var f2 = nextInSet(f1, r)
var f3 = nextInSet(f2, r)

// GeoProSum returns the sum of all elements in a geometric progression slice
func GeoProSum(start int, last int, ratio int) int {
	return ((last * ratio) - start) / (ratio - 1)
}

// nextInSet returns next element in geometric progression set given an element and the set ratio
func nextInSet(x int, ratio int) int {
	// input previous number in set, returns next
	return x * ratio
}

// BuildGeoProgression returns a geometric progression set where each element is a ratio of the previous and next number
func BuildGeoProgression(start int, ratio int, size int) []int {
	var b []int
	b = append(b, start)
	for i := 1; i <= (size - 1); i++ {
		var last = b[len(b)-1] // getting the last element in slice
		b = append(b, nextInSet(last, ratio))
	}
	return b
}

// Last returns the last element in the slice
func Last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}

// First returns the first element in the slice
func First[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[0], true
}

// Swap accepts two integer values and returns the two values in reversed order (thereby swapping the position of the two variables)
func Swap(x int, y int) (int, int) {
	var newX = y
	var newY = x
	return newX, newY
}

// BubbleSort accepts a slices of any length and sorts the integer variables from smallest to biggest
func BubbleSort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = Swap(array[j], array[j+1])
			}
		}
	}
	return array
}

// FindSortedIndex prints the index of a specified value (v) and  integer array (array) and returns the index of that value
func FindSortedIndex(v int, array []int) {
	n := len(array)
	a := 0
	b := n - 1
	for a < b {
		k := (a + b) / 2
		if array[k] == v {
			fmt.Println(k)
		}
		if array[k] > v {
			b = k - 1
		} else {
			a = k + 1
		}
	}
}

// ListElement listElement returns the element of the specified array, one element per line.
func ListElement(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	return
}

// AppendListElement append a value to the end of the array
func AppendListElement(array []int, value int) []int {
	var length = len(array)
	var temp = make([]int, length+1)

	for i := 0; i < length; i++ {
		temp[i] = array[i]
	}
	temp[length] = value
	return temp
}

func InsertEleIndex(array []int, length int, temp []int, value int, insertIndex int) []int {
	for i := 0; i < length; i++ {
		if i < insertIndex {
			temp[i] = array[i]
		} else {
			temp[i+1] = array[i]
		}
	}
	temp[insertIndex] = value // Insert value to insertIndex of new array
	return temp
}

func deleteEleIndex(array []int, index int) []int {
	var length = len(array)
	var temp = make([]int, length-1) // New array with size 1 smaller
	for i := 0; i < length; i++ {
		if i < index {
			// values before index to be deleted go get copied to new array
			temp[i] = array[i]
		}
		if i > index {
			// values after index to be deleted get copied to new array  index - 1
			temp[i-1] = array[i]
		}
	}
	return temp
}

// shuffleOneLeft accepts an integer array and returns all elements shifted one index to the left
func shuffleOneLeft(array []int) []int {
	// shuffles all values 1 position to the left
	for i := 0; i < len(array)-1; i++ {
		array[i], array[i+1] = Swap(array[i], array[i+1])
	}
	return array
}

func main() {
	geo := BuildGeoProgression(1, 9, 31555)
	fmt.Println(GeoProSum(geo[0], geo[29], 7))
	fmt.Println(shuffleOneLeft(geo))
	del := deleteEleIndex(geo, 0)
	fmt.Println(del)
	geo = AppendListElement(del, 77)
	ListElement(geo)
	fmt.Println(First(geo))
	fmt.Println(Last(geo))

	f, _ := First(geo)
	l, _ := Last(geo)
	f, l = Swap(f, l)
	fmt.Println(f, l)
}
