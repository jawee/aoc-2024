package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{0, 1, 10, 99, 999,}
	expected := []int{1, 2024, 1, 0, 9, 9, 2021976,}

	res := blink(input)

	if len(res) != len(expected) {
		fmt.Printf("Res: %+v\n", res)
		fmt.Printf("Exp: %+v\n", expected)
		t.Fatalf("Failed on len")
	}
	for i := range expected {
		if res[i] != expected[i] {
			t.Fatalf("Failed on idx %d\n", i)
		}
	}
}

func Test2(t *testing.T) {
	input := []int{125, 17}
	expected := []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2,}

	res := blink(input)
	for range 5 {
		res = blink(res)
	}

	if len(res) != len(expected) {
		fmt.Printf("Res: %+v\n", res)
		fmt.Printf("Exp: %+v\n", expected)
		t.Fatalf("Failed on len")
	}
	for i := range expected {
		if res[i] != expected[i] {
			t.Fatalf("Failed on idx %d\n", i)
		}
	}
}
