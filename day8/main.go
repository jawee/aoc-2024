package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/day8/atest.txt")
	// file, err := os.Open(pwd + "/day8/btest.txt")
	// file, err := os.Open(pwd + "/day8/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	res := a(file)
	// res := b(file)
	fmt.Printf("%d\n", res)
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	strs := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		strs = append(strs, split)
	}

	for _, v := range strs {
		fmt.Printf("%+v\n", v)
	}

	sum = countAntinodes(strs)

	return sum
}

func countAntinodes(strs [][]string) int {
	sum := 0
	antiNodes := [][]bool{}
	for range strs {
		antiNodes = append(antiNodes, make([]bool, len(strs)))
	}
	for row := range strs {
		for col := range strs[row] {
			if strs[row][col] != "." {
			}

		}
	}

	return sum
}

