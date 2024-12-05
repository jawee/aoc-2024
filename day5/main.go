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
	// file, err := os.Open(pwd + "/day5/atest.txt")
	// file, err := os.Open(pwd + "/day5/btest.txt")
	file, err := os.Open(pwd + "/day5/input.txt")

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

	pageRules := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "|")
		ruleNum, _ := strconv.Atoi(split[0])
		ruleVal, _ := strconv.Atoi(split[1])
		pageRules[ruleNum] = append(pageRules[ruleNum], ruleVal)
	}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		pages := []int{}
		for _, num := range nums {
			val, _ := strconv.Atoi(num)
			pages = append(pages, val)
		}

		if isCorrectOrder(pages, pageRules) {
			sum += pages[len(pages)/2]
		}
	}

	return sum
}

func isCorrectOrder(pages []int, pageRules map[int][]int) bool {
	isOk := true
	for i := range pages {
		if i == 0 {
			continue
		}

		rules := pageRules[pages[i]]
		for j, page := range pages {
			if j == i {
				continue
			}

			if j > i {
				break;
			}
			for _, rule := range rules {
				if page == rule {
					isOk = false
					break
				}
			}
		}
	}
	return isOk
}

func correctOrder(pages []int, pageRules map[int][]int) int {
	// isOk := true
	for i := range pages {
		if i == 0 {
			continue
		}

		rules := pageRules[pages[i]]
		for j, page := range pages {
			if j == i {
				continue
			}

			if j > i {
				break;
			}
			for _, rule := range rules {
				if page == rule {
					// isOk = false
					break
				}
			}
		}
	}
	return 0
}
