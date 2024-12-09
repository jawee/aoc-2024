package main

import (
	"testing"
) 

func Test1(t *testing.T) {
	input := "12345"
	
	_ = getCheckSum(input)

	t.Fatalf("Fail")
}
