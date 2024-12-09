package main

import (
	"testing"
)

func Test1(t *testing.T) {
	input := "12345"

	res := getCheckSumA(input)

	if res != 60 {
		t.Fatalf("Fail %d\n", res)
	}
}

func Test2(t *testing.T) {
	input := "12345"

	res := getCheckSumB(input)

	if res != 132 {
		t.Fatalf("Fail %d\n", res)
	}
}
