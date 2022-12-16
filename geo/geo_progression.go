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
var l = nextInSet(f3, r)

// geoProSum returns the sum of all elements in a geometric progression slice
func geoProSum(start int, last int, ratio int) int {
	return ((last * ratio) - start) / (ratio - 1)
}

// nextInSet returns next element in geometric progression set given an element and the set ratio
func nextInSet(x int, ratio int) int {
	// input previous number in set, returns next
	return x * ratio
}

// buildGeoProgression returns a geometric progression set where each element is a ratio of the previous and next number
func buildGeoProgression(start int, ratio int, size int) []int {
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

// listElement returns the element of the specified array, one element per line.
func listElement(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	return
}

// appendListElement append a value to the end of the array
func appendListElement(array []int, value int) []int {
	var length = len(array)
	var temp = make([]int, length+1)

	for i := 0; i < length; i++ {
		temp[i] = array[i]
	}
	temp[length] = value
	return temp
}

func insertEleIndex(array []int, length int, temp []int, value int, insertIndex int) []int {
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

// TODO create arguments for main function to make program callable
func main() {
	fmt.Println("Go get it kiddo")
	var x int
	x = 4
	fmt.Println(x)
	fmt.Println(geoProSum(f, l, r))
	var geoPro = buildGeoProgression(3, 4, 12)
	fmt.Println(Last(geoPro))
	var l, _ = Last(geoPro)
	var sumOfSlice = geoProSum(4, l, 3)
	fmt.Println(sumOfSlice) // sum of all numbers in slice
	fmt.Println(geoPro)     // original slice
	fmt.Println(First(geoPro))

	var a = 3
	var b = 7
	a, b = Swap(a, b)
	fmt.Println(a) // should now be 7
	fmt.Println(b) // should now be 3

	sl := make([]int, 0, 8)
	sl = append(sl, 34, 8, 7, 56, 467, 73, 3, 1)
	var sorted = BubbleSort(sl)
	fmt.Println(sorted)
	fmt.Println(len(sl))
	FindSortedIndex(73, sorted) // binary search for the number 6 in sorted array
	listElement(sorted)
	sorted = appendListElement(sorted, r*len(sorted))
	listElement(sorted)
	fmt.Println(geoProSum(sorted[0], sorted[len(sorted)-1], r)) // sorted is no longer a geometric progression so sum is not accurate
	sortedLen := len(sorted)
	var empty = make([]int, sortedLen+1)                     // initialied empy list created for insertEleIndex must be one bigger
	sorted = insertEleIndex(sorted, sortedLen, empty, 12, 3) // add 12 for first index of sorted
	listElement(sorted)
	fmt.Println(sorted)
	sorted = BubbleSort(sorted)
	fmt.Println(sorted)
	sorted = deleteEleIndex(sorted, 3)
	sorted = BubbleSort(sorted)
	fmt.Println(sorted)

	row := makeRowValues(sorted)
	fmt.Println(row[0].value)
	fmt.Println(row[0].position, "\n", len(row))

	//empty = make([]int, len(sl)+1)
	sl = append(sl, 55)
	fmt.Println(sl)
	row2 := makeRowValues(sl)
	fmt.Println(row2[8].value)
	fmt.Println(row2[8].position, "\n", len(row2))
}
