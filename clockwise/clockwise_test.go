package clockwise

import (
	"testing"
)

func TestProvidedExample(t *testing.T) {
	row1 := []int{2, 3, 4, 8}
	row2 := []int{5, 7, 9, 12}
	row3 := []int{1, 0, 6, 10}

	values := [][]int{}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	expected := "2, 3, 4, 8, 12, 10, 6, 0, 1, 5, 7, 9"

	result := GetClockwiseMatrixString(values)

	if have, want := result, expected; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}

func TestClockwiseOneRowOneColumn(t *testing.T) {
	row1 := []int{1}

	values := [][]int{}
	values = append(values, row1)

	expected := "1"

	result := GetClockwiseMatrixString(values)

	if have, want := result, expected; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}

func TestClockwiseTwoRowsOneColumn(t *testing.T) {
	row1 := []int{1}
	row2 := []int{2}

	values := [][]int{}
	values = append(values, row1)
	values = append(values, row2)

	expected := "1, 2"

	result := GetClockwiseMatrixString(values)

	if have, want := result, expected; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}

func TestClockwiseTwoRowsTwoColumns(t *testing.T) {
	row1 := []int{1, 2}
	row2 := []int{4, 3}

	values := [][]int{}
	values = append(values, row1)
	values = append(values, row2)

	expected := "1, 2, 3, 4"

	result := GetClockwiseMatrixString(values)

	if have, want := result, expected; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}

func TestClockwiseOneRowFourColumns(t *testing.T) {
	row1 := []int{1, 2, 3, 4}

	values := [][]int{}
	values = append(values, row1)

	expected := "1, 2, 3, 4"

	result := GetClockwiseMatrixString(values)

	if have, want := result, expected; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}

func TestClockwiseTwoRowsFourColumns(t *testing.T) {
	row1 := []int{1, 2, 3, 4}
	row2 := []int{8, 7, 6, 5}

	values := [][]int{}
	values = append(values, row1)
	values = append(values, row2)

	expected := "1, 2, 3, 4, 5, 6, 7, 8"

	result := GetClockwiseMatrixString(values)

	if have, want := result, expected; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}

func TestClockwiseThreeRowsFourColumns(t *testing.T) {
	row1 := []int{1, 2, 3, 4}
	row2 := []int{10, 11, 12, 5}
	row3 := []int{9, 8, 7, 6}

	values := [][]int{}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	expected := "1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12"

	result := GetClockwiseMatrixString(values)

	if have, want := result, expected; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}

func TestClockwiseFourRowFourColumns(t *testing.T) {
	row1 := []int{1, 2, 3, 4}
	row2 := []int{12, 13, 14, 5}
	row3 := []int{11, 16, 15, 6}
	row4 := []int{10, 9, 8, 7}

	values := [][]int{}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)
	values = append(values, row4)

	expected := "1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16"

	result := GetClockwiseMatrixString(values)

	if have, want := result, expected; have != want {
		t.Errorf("expected %s but have %s", want, have)
	}
}
