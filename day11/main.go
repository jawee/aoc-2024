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
	// file, err := os.Open(pwd + "/day11/atest.txt")
	// file, err := os.Open(pwd + "/day11/btest.txt")
	file, err := os.Open(pwd + "/day11/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	// res := a(file)
	res := b(file)
	fmt.Printf("%d\n", res)
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	// sum := 0

	input := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		for _, v := range split {
			v, _ := strconv.Atoi(v)
			input = append(input, v)
		}

	}
	result := blink(input)
	for range 24 {
		result = blink(result)
	}

	return len(result)
}

type stone struct {
	number int
	depth  int
}

func b(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0

	input := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		for _, v := range split {
			v, _ := strconv.Atoi(v)
			input = append(input, v)
		}

	}
	cache := map[stone]int{}
	for _, v := range input {
		sum += blinkBDepth(v, 75, cache)
	}

	return sum
}

func blinkBDepth(v, depth int, cache map[stone]int) int {
	if val, ok := cache[stone{number: v, depth: depth}]; ok {
		return val
	}

	left, right := blinkB(v)

	if depth == 1 {
		if right == -1 {
			cache[stone{number: v, depth: depth}] = 1
			return 1
		}
		cache[stone{number: v, depth: depth}] = 2
		return 2
	}

	output := blinkBDepth(left, depth-1, cache)
	cache[stone{number: left, depth: depth-1}] = output 
	if right != -1 {
		tmp := blinkBDepth(right, depth-1, cache)
		cache[stone{number: right, depth: depth-1}] = tmp 
		output += tmp
	}
	cache[stone{number: v, depth: depth}] = output 

	return output
}

func blinkB(v int) (int, int) {
		strV := fmt.Sprintf("%d", v)
		if v == 0 {
			return 1, -1
		} else if len(strV)%2 == 0 {
			//split
			v1, err := strconv.Atoi(strV[:len(strV)/2])
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			v2, err := strconv.Atoi(strV[len(strV)/2:])
			if err != nil {
				fmt.Printf("%s\n", err)
			}
		return v1, v2
		} else {
			// multiply with 2024
			return v*2024, -1
		}
}

func blink(input []int) []int {
	result := []int{}

	for _, v := range input {
		strV := fmt.Sprintf("%d", v)
		if v == 0 {
			result = append(result, 1)
		} else if len(strV)%2 == 0 {
			//split
			v1, err := strconv.Atoi(strV[:len(strV)/2])
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			v2, err := strconv.Atoi(strV[len(strV)/2:])
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			result = append(result, v1)
			result = append(result, v2)

		} else {
			// multiply with 2048
			result = append(result, v*2024)
		}
	}
	return result
}
