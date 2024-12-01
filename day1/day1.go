package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	// file, err := os.Open(pwd + "/day1/atest.txt")
	// file, err := os.Open(pwd + "/day1/btest.txt")
	file, err := os.Open(pwd + "/day1/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	// res := a(file)
	res := b(file)
	fmt.Printf("%d\n", res)
}

func b(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(split[0])
		rightNum, _ := strconv.Atoi(split[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	for _, iv := range left {
		c := 0
		for _, v := range right {
			if iv == v {
				c += 1
			}
		}

		sum += iv * c
	}

	return sum
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(split[0])
		rightNum, _ := strconv.Atoi(split[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	left = sortSlice(left)
	right = sortSlice(right)

	for i := 0; i < len(left); i++ {
		sum += abs(left[i], right[i])
	}

	return sum
}

func sortSlice(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	return a
}

func abs(a int, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
