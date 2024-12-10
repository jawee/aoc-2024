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
	file, err := os.Open(pwd + "/day11/atest.txt")
	// file, err := os.Open(pwd + "/day11/btest.txt")
	// file, err := os.Open(pwd + "/day11/input.txt")

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

	for scanner.Scan() {
		line := scanner.Text()
		str := strings.Split(line, "")
	}

	return sum
}
