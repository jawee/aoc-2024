package main

import (
	"testing"
)

func Test1(t *testing.T) {
	numbers:= []int{81,40,27}

	if !isValid(3267, numbers) {
		t.Fatalf("Got false, expected true for %+v\n", numbers)
	}
}

func Test2(t *testing.T) {
	numbers := []int{3,1,2}
// 6: 2 3
	// 2: 3 1 2
// 6: 2 3 3
	if isValid(2, numbers) {
		t.Fatalf("Got true, expected false for %+v\n", numbers)
	}
}
