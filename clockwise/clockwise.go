package clockwise

import (
	"fmt"
	"strconv"
)

func GetClockwiseMatrixString(values [][]int) string {
	var result string
	length := len(values)
	for i := 0; i < length; i++ {
		if i == 0 || i%2 == 0 {
			values, result = printTopDown(values, result)
		} else {
			values, result = printBottomUp(values, result)
		}
	}
	return result
}

func printTopDown(values [][]int, result string) (v [][]int, r string) {
	for _, val := range values[0] {
		result += fmt.Sprintf("%s, ", strconv.Itoa(val))
	}
	values = values[1:][:]
	if len(values) > 1 {
		for i := 0; i < len(values); i++ {
			for j := len(values[i]) - 1; j < len(values[i]); j++ {
				result += fmt.Sprintf("%s, ", strconv.Itoa(values[i][j]))
				values[i] = append(values[i][:j], values[i][j+1:]...)
			}
		}
	}
	return values, result
}

func printBottomUp(values [][]int, result string) (v [][]int, r string) {
	lastRowIndex := len(values) - 1
	for i := len(values[lastRowIndex]); i > 0; i-- {
		result += fmt.Sprintf("%s, ", strconv.Itoa(values[lastRowIndex][i-1]))
	}
	values = values[:lastRowIndex][:]

	if len(values) > 1 {
		for i := len(values) - 1; i >= 0; i-- {
			for j := 0; j < 1; j++ {
				result += fmt.Sprintf("%s, ", strconv.Itoa(values[i][j]))
				values[i] = append(values[i][j+1:], values[i][:j]...)
			}
		}
	}
	return values, result
}
