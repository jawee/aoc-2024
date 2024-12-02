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
	// file, err := os.Open(pwd + "/day2/atest.txt")
	// file, err := os.Open(pwd + "/day2/btest.txt")
	file, err := os.Open(pwd + "/day2/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	res := b(file)
	// res := b(file)
	fmt.Printf("Res: %d\n", res)
}

func b(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	count := 0

	lineC := 0
	for scanner.Scan() {
		isRemoved := false
		isSafe := true
		orderDecided := false
		descending := true
		line := scanner.Text()
		split := strings.Split(line, " ")
		prev := -1
		for i, v := range split {
			val, _ := strconv.Atoi(v)
			if prev == -1 {
				// TODO: need to be able to remove first
				// next, _ := strconv.Atoi(split[i+1])
				// if abs(val, next) > 3 || abs(val, next) == 0 {
				// 	continue
				// }
				prev = val
				continue
			}
			if abs(val, prev) > 3 || abs(val, prev) == 0 {
				if !isRemoved {
					fmt.Printf("%d removed\n", val)
					isRemoved = true
					continue
				}
				fmt.Printf("%d %d abs\n", val, prev)
				isSafe = false
				// unsafe 
				break
			}
			if !orderDecided {
				orderDecided = true
				if val > prev {
					descending = false
				} else {
					descending = true
				}
			}

			if descending && val > prev {
				if !isRemoved {
					fmt.Printf("%d removed\n", val)
					isRemoved = true
					continue
				}
				fmt.Printf("%d > %d descending\n", val, prev)
				isSafe = false
				break
			}

			if !descending && val < prev {
				if !isRemoved {
					fmt.Printf("%d removed\n", val)
					isRemoved = true
					continue
				}
				fmt.Printf("%d < %d ascending\n", val, prev)
				isSafe = false
				break
			}
			prev = val
		}

		if isSafe {
			count++
		}

		fmt.Printf("Line %d\n", lineC)
		lineC++
	}

	return count
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		isSafe := true
		orderDecided := false
		descending := true
		line := scanner.Text()
		split := strings.Split(line, " ")
		prev := -1
		for _, v := range split {
			val, _ := strconv.Atoi(v)
			if prev == -1 {
				prev = val
				continue
			}
			if abs(val, prev) > 3 || abs(val, prev) == 0 {
				fmt.Printf("%d %d abs\n", val, prev)
				isSafe = false
				// unsafe 
				break
			}
			if !orderDecided {
				orderDecided = true
				if val > prev {
					descending = false
				} else {
					descending = true
				}
			}

			if descending && val > prev {
				fmt.Printf("%d > %d descending\n", val, prev)
				isSafe = false
				break
			}

			if !descending && val < prev {
				fmt.Printf("%d < %d ascending\n", val, prev)
				isSafe = false
				break
			}
			prev = val
		}
		if isSafe {
			count++
		}
	}

	return count
}

func abs(a int, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}


