package main

import (
	"testing"
)

func TestIsOk(t *testing.T) {
	pages := []int{75,47,61,53,29}
	pageRules := map[int][]int{
		75: []int{47, 61, 53, 29},
		47: []int{61, 53, 29},
		61: []int{53, 29},
		53: []int{29},
	}

	if !isCorrectOrder(pages, pageRules) {
		t.Fatalf("Expected true, got false")
	}
}
