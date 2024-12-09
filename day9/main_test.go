package main

import (
	"testing"
) 

func Test1(t *testing.T) {
	input := "2333133121414131402"
	
	res := getCheckSum(input)

	if 1928 != res {
		t.Fatalf("Fail %d\n", res)
	}
}
