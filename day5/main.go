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
	// res := a(file)
	res := b(file)
	fmt.Printf("%d\n", res)
}

func b(file io.Reader) int {
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

		sum += getMiddlePages(pages, pageRules)
	}

	return sum
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
				break
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

func getMiddlePages(pages []int, pageRules map[int][]int) int {
	if isCorrectOrder(pages, pageRules) {
		return 0
	}

	new := []int{}
	used := map[int]bool{}
	for range pages {
		for _, page := range pages {
			if _, ok := used[page]; ok {
				continue
			}
			breaksRule := false
			for key, rules := range pageRules {
				keyUsed := false
				for k := range used {
					if k == key {
						keyUsed = true
					}
				}
				if keyUsed {
					continue
				}
				ignoreRules := true
				for _, v := range pages {
					if key == v {
						ignoreRules = false
						break

					}
				}
				if ignoreRules {
					continue
				}
				for _, rule := range rules {
					if rule == page {
						breaksRule = true
						break
					}
				}
				if breaksRule {
					break
				}
			}
			if !breaksRule {
				new = append(new, page)
				used[page] = true
				break
			}
		}
	}
	return new[len(new)/2]
}

func deepCopy(a []int) []int {
	new := []int{}
	for _, v := range a {
		new = append(new, v)
	}
	return new
}
