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
	// file, err := os.Open(pwd + "/day3/atest.txt")
	// file, err := os.Open(pwd + "/day3/btest.txt")
	file, err := os.Open(pwd + "/day3/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	res := a(file)
	fmt.Printf("Res: %d\n", res)
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
	}

	return count
}
