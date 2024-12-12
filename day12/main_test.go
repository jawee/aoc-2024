package main

import (
	"testing"
)

func Test1(t *testing.T) {
	input := [][]string{
		{"A","A","A","A"},
		{"B","B","C","D"},
		{"B","B","C","C"},
		{"E","E","E","C"},
	}

	res := findFencingCost(input)
	if res != 140 {
		t.Fatalf("Expected %d, got %d\n", 140, res)
	}
}

func Test2(t *testing.T) {
	input := [][]string{
		{"O","O","O","O","O"},
		{"O","X","O","X","O"},
		{"O","O","O","O","O"},
		{"O","X","O","X","O"},
		{"O","O","O","O","O"},
	}

	res := findFencingCost(input)
	if res != 772 {
		t.Fatalf("Expected %d, got %d\n", 772, res)
	}
}

func Test3(t *testing.T) {
	input := [][]string{
		{"A","A","A","A"},
		{"B","B","C","D"},
		{"B","B","C","C"},
		{"E","E","E","C"},
	}

	res := floodFill("A", []coordinate{{x: 0, y: 0}}, []coordinate{}, input)
	if len(res) != 4 {
		t.Fatalf("Expected 4, got %d\n", len(res))
	}
}
