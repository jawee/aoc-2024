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

// func Test1(t *testing.T) {
//     testCases := []struct {
//         line    string
//         expected int
//     }{
//         {
//             line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
//             expected: 8,
//         },
//         {
//             line: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
//             expected: 2,
//         },
//         {
//             line: "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
//             expected: 1,
//         },
//         // {
//         //     line: "Card 4: 41  5 73 84 69 | 59 84 76 51 58  5 54 83",
//         //     expected: 2,
//         // },
//     }
//
//
//     for i, v := range testCases {
//         // res := calculateASum(v.line)
//         //
//         // if res != v.expected {
//         //     t.Fatalf("Case %d: Got '%d' expected '%d'\n", i, res, v.expected)
//         // }
//     }
// }
