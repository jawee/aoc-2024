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
	res := a(file)
	// res := b(file)
	fmt.Printf("%d\n", res)
}

func a(file io.Reader) uint64 {
	scanner := bufio.NewScanner(file)
	var sum uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")

		wantedResult, err := strconv.ParseUint(split[0], 10, 64)
		if err != nil {
			panic(err)
		}

		numbersStr := strings.Split(strings.Trim(split[1], " "), " ")

		numbers := []uint64{}
		for _, v := range numbersStr {
			val, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, val)
		}

		if isValid(wantedResult, numbers) {
			sum += uint64(wantedResult)
		}	
	}

	return sum
}

func createPossibleOperators(count int) [][]string {
	possibleCombinations := [][]string{}
	first := []string{}
	last := []string{}
	for _ = range count {
		first = append(first, "+")
		last = append(last, "*")
	}
	possibleCombinations = append(possibleCombinations, first)
	for i := range count {
		generated := deepCopy(first)
		generated[i] = "*"
		possibleCombinations = append(possibleCombinations, generated)
		for j := range count {
			if i == j {
				continue
			}
			newGenerated := deepCopy(generated)
			newGenerated[j] = "*"
			possibleCombinations = append(possibleCombinations, newGenerated)
		}
	}
	possibleCombinations = append(possibleCombinations, last)
	return possibleCombinations
}
func deepCopy(a []string) []string {
	new := []string{}
	for _, v := range a {
		new = append(new, v)
	}
	return new
}

func isValid(wantedResult uint64, numbers []uint64) bool {
	// possibleCombinations := [][]string{
	// 	{"+", "+"},
	// 	{"*", "+"},
	// 	{"*", "*"},
	// 	{"+", "*"},
	// }
	possibleCombinations := createPossibleOperators(len(numbers)-1)

	for _, operators := range possibleCombinations {
		var sum uint64 = 1
		fmt.Printf("Op: %+v\n", operators)
		for i, v := range numbers {
			fmt.Printf("%d: Handling %d\n", i, v)
			if i == 0 {
				sum *= uint64(v)
				continue
			}
			if operators[i-1] == "+" {
				sum += uint64(v)
			}
			if operators[i-1] == "*" {
				sum *= uint64(v)
			}
		}

		if sum == uint64(wantedResult) {
			return true
		}
	}

	return false
}
