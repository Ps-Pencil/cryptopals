package tools

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	input := [][]byte{
		{0, 1, 2, 3, 4},
		{5, 6, 7, 8, 9},
		{10, 11, 12},
	}

	want := [][]byte{
		{0, 5, 10},
		{1, 6, 11},
		{2, 7, 12},
		{3, 8},
		{4, 9},
	}

	got := transposeBytes(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transpose got %v; want %v.", got, want)
	}
}
