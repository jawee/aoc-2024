package main

import (
	"testing"
)

func TestRight(t *testing.T) {
	matrix := [][]string{
		{"X", "M", "A", "S"},
	}

	if !searchWord(matrix, 0, 0, 1, 0) {
		t.Fatalf("Expected true, got false")
	}
}

func TestB(t *testing.T) {
	matrix := [][]string{
		{"M", "B", "S"},
		{"B", "A", "B"},
		{"M", "B", "S"},
	}

	if !searchMas(matrix, 1, 1) {
		t.Fatalf("Expected true, got false")
	}
}

func TestB2(t *testing.T) {
	matrix := [][]string{
		{"M", "B", "M"},
		{"B", "A", "B"},
		{"S", "B", "S"},
	}

	if !searchMas(matrix, 1, 1) {
		t.Fatalf("Expected true, got false")
	}
}

func TestB3(t *testing.T) {
	matrix := [][]string{
		{"M", "B", "S"},
		{"B", "A", "B"},
		{"S", "B", "S"},
	}

	if searchMas(matrix, 1, 1) {
		t.Fatalf("Expected false, got true")
	}
}
