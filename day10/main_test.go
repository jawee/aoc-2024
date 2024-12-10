package main

import (
	"testing"
)

func Test2(t *testing.T) {
	input := [][]int{
		{9, 9, 9, 0, 9, 9, 9},
		{9, 9, 9, 1, 9, 9, 9},
		{9, 9, 9, 2, 9, 9, 9},
		{6, 5, 4, 3, 4, 5, 6},
		{7, 1, 1, 1, 1, 1, 7},
		{8, 1, 1, 1, 1, 1, 8},
		{9, 1, 1, 1, 1, 1, 9},
	}

	res := calculateTrailheads(input)

	if res != 2 {
		t.Fatalf("Failed. Expected 2, got %d.\n", res)
	}
}

func Test3(t *testing.T) {
	input := [][]int{
		{1, 1, 9, 0, 5, 5, 9},
		{5, 5, 5, 1, 5, 9, 8},
		{7, 7, 7, 2, 4, 3, 7},
		{6, 5, 4, 3, 4, 5, 6},
		{7, 6, 5, 1, 9, 8, 7},
		{8, 7, 6, 4, 4, 4, 4},
		{9, 8, 7, 2, 2, 2, 2},
	}

	res := calculateTrailheads(input)

	if res != 4 {
		t.Fatalf("Failed. Expected 4, got %d.\n", res)
	}
}
