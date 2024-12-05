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

func TestB(t *testing.T) {
	pages := []int{75,97,47,61,53}
	pageRules := map[int][]int{
		75: []int{47, 61, 53, 29},
		47: []int{61, 53, 29},
		61: []int{53, 29},
		53: []int{29},
		97: []int{13,61,47,29,53,75},
	}

	if getMiddlePages(pages, pageRules) != 47 {
		t.Fatalf("Expected 47, got %d", getMiddlePages(pages, pageRules))
	}
}

func TestB2(t *testing.T) {
	pages := []int{61,13,29}
	pageRules := map[int][]int{
		29: []int{13},
		75: []int{47, 61, 53, 29},
		47: []int{61, 53, 29},
		61: []int{53, 29},
		53: []int{29},
		97: []int{13,61,47,29,53,75},
	}

	if getMiddlePages(pages, pageRules) != 29 {
		t.Fatalf("Expected 29, got %d", getMiddlePages(pages, pageRules))
	}
}

func TestB3(t *testing.T) {
	pages := []int{97,13,75,29,47}
	pageRules := map[int][]int{
		29: []int{13},
		75: []int{47, 61, 53, 29},
		47: []int{61, 53, 29},
		61: []int{53, 29},
		53: []int{29},
		97: []int{13,61,47,29,53,75},
	}

	if getMiddlePages(pages, pageRules) != 47 {
		t.Fatalf("Expected 47, got %d", getMiddlePages(pages, pageRules))
	}
}
