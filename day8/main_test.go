package main

import (
	"strings"
	"testing"
) 

func Test1(t *testing.T) {
	input := [][]string{
		strings.Split("..........", ""),
		strings.Split("..........", ""), //1, 3
		strings.Split("..........", ""),
		strings.Split("....a.....", ""),
		strings.Split("..........", ""),
		strings.Split(".....a....", ""),
		strings.Split("..........", ""),
		strings.Split("..........", ""), //7, 6
		strings.Split("..........", ""),
		strings.Split("..........", ""),
	}

	res := countAntinodes(input)

	// if res[1][3] != true || res[7][6] != true {
	if res != 2 {
		t.Fatalf("Failed.")
	}
}

func Test2(t *testing.T) {
// 3, 4. Found another antenna at 5, 5.
	antiNodes := getAntinodes(3,4,5,5)

	if len(antiNodes) != 2 {
		t.Fatalf("Failed. %+v", antiNodes)
	}

	// t.Logf("%+v", antiNodes)
}

func Test3(t *testing.T) {
// 3, 4. Found another antenna at 5, 5.
	antiNodes := getAntinodes(5,5,3,4)

	if len(antiNodes) != 2 {
		t.Fatalf("Failed. %+v", antiNodes)
	}

	// t.Logf("%+v", antiNodes)
}

func Test4(t *testing.T) {
// 3, 4. Found another antenna at 5, 5.
	antiNodes := getAntinodes(3,6,5,5)

	if len(antiNodes) != 2 {
		t.Fatalf("Failed. %+v", antiNodes)
	}

	// t.Logf("%+v", antiNodes)
}

func Test5(t *testing.T) {
// 3, 4. Found another antenna at 5, 5.
	antiNodes := getAntinodes(5,5,3,6)

	if len(antiNodes) != 2 {
		t.Fatalf("Failed. %+v", antiNodes)
	}

	// t.Logf("%+v", antiNodes)
}

func Test6(t *testing.T) {
	input := [][]string{
		strings.Split("T.........", ""),
		strings.Split("...T......", ""),
		strings.Split(".T........", ""),
		strings.Split("..........", ""),
		strings.Split("..........", ""),
		strings.Split("..........", ""),
		strings.Split("..........", ""),
		strings.Split("..........", ""),
		strings.Split("..........", ""),
		strings.Split("..........", ""),
	}

	res := countAntinodesB(input)

	// if res[1][3] != true || res[7][6] != true {
	if res != 9 {
		t.Fatalf("Expected %d, got %d.", 9, res)
	}
}

func Test7(t *testing.T) {
// 3, 4. Found another antenna at 5, 5.
	antiNodes := getAntinodesB(0,0,2,1,9,9)

	if len(antiNodes) != 3 {
		t.Fatalf("Failed. %+v", antiNodes)
	}

	// t.Logf("%+v", antiNodes)
}
