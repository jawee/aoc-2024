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
	// file, err := os.Open(pwd + "/day7/atest.txt")
	// file, err := os.Open(pwd + "/day7/btest.txt")
	file, err := os.Open(pwd + "/day7/input.txt")

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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")

		wantedResult, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		numbersStr := strings.Split(strings.Trim(split[1], " "), " ")

		numbers := []int{}
		for _, v := range numbersStr {
			val, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, val)
		}

		if isValidA(wantedResult, numbers) {
			sum += wantedResult
		}	
	}

	return sum
}

func b(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")

		wantedResult, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		numbersStr := strings.Split(strings.Trim(split[1], " "), " ")

		numbers := []int{}
		for _, v := range numbersStr {
			val, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, val)
		}

		if isValidB(wantedResult, numbers) {
			sum += wantedResult
		}	
	}

	return sum
}

func isValidB(wantedResult int, numbers []int) bool {
	res := checkNextB(wantedResult, numbers[0], numbers, 1) 

	return res
}

func checkNextB(wantedResult int, current int, numbers []int, index int) bool {
	if index == len(numbers) {
		return current == wantedResult
	}

	if checkNextB(wantedResult, current+numbers[index], numbers, index+1) {
		return true
	}

	if checkNextB(wantedResult, current*numbers[index], numbers, index+1) {
		return true
	}

	current, _ = strconv.Atoi(fmt.Sprintf("%d%d", current, numbers[index]))
	if checkNextB(wantedResult, current, numbers, index+1) {
		return true
	}

	return false 
}

func isValidA(wantedResult int, numbers []int) bool {
	res := checkNextA(wantedResult, numbers[0], numbers, 1) 

	return res
}

func checkNextA(wantedResult int, current int, numbers []int, index int) bool {
	if index == len(numbers) {
		return current == wantedResult
	}

	if checkNextA(wantedResult, current+numbers[index], numbers, index+1) {
		return true
	}

	if checkNextA(wantedResult, current*numbers[index], numbers, index+1) {
		return true
	}

	return false 
}

