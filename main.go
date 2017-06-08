package main

import (
	"clockwise-matrix/clockwise"
	"fmt"
)

func main() {
	values := [][]int{}

	row1 := []int{2, 3, 4, 8}
	row2 := []int{5, 7, 9, 12}
	row3 := []int{1, 0, 6, 10}
	// row4 := []int{96, 97, 98, 99}

	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)
	// values = append(values, row4)

	result := clockwise.GetClockwiseMatrixString(values)
	fmt.Printf("Result: %s\n", result)
}

// “2, 3, 4, 8, 12, 10, 6, 0, 1, 5, 7, 9”.

// 2, 3, 4, 8
// 5, 7, 9, 12
// 1, 0, 6, 10
