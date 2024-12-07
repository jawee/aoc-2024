package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/day7/atest.txt")
	// file, err := os.Open(pwd + "/day7/btest.txt")
	// file, err := os.Open(pwd + "/day7/input.txt")

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
		split := strings.Split(line, ":")

		wantedResult, err := strconv.Atoi(split[0])

		numbersStr := strings.Split(strings.Trim(split[1], " "), " ")

		numbers := []int{}
		for _, v := range numbersStr {

			val, _ := strconv.Atoi(v)
			numbers = append(numbers, val)
		}
	}

	return sum
}
