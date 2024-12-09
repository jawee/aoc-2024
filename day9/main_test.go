package main

import (
	"testing"
) 

func Test1(t *testing.T) {
	input := "12345"
	
	res := getCheckSum(input)

	if res != 60 {
		t.Fatalf("Fail %d\n", res)
	}
}
