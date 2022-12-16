package main

type Column struct {
	size   int         // how many rows the colum has
	header string      // title of each column
	Values []RowValues // array of rows
}

type RowValues struct {
	position int // index of value in row
	value    any
}

type Table struct {
	columns   []Column
	rowValues []RowValues
}

// makeRowValues turns an array of int into a RowValue class
func makeRowValues(array []int) []RowValues {
	rowValues := make([]RowValues, len(array))
	for i := 0; i < len(array); i++ {
		rowValues[i].position = i + 1 // start position from 1 , assume header at position 0
		rowValues[i].value = array[i]
	}
	return rowValues
}
