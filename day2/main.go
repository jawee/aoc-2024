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
	fmt.Printf("Res: %d\n", res)
}

func isSafeB(vals []int) bool {
	isSafe := true
	orderDecided := false
	descending := true
	for i, val := range vals {
		if i == 0 {
			continue
		}

		if abs(val, vals[i-1]) > 3 || abs(val, vals[i-1]) == 0 {
			isSafe = false
			break
		}

		if !orderDecided {
			orderDecided = true
			if val > vals[i-1] {
				descending = false
			} else {
				descending = true
			}
		}

		if descending && val > vals[i-1] {
			isSafe = false
			break
		}

		if !descending && val < vals[i-1] {
			isSafe = false
			break
		}
	}
	return isSafe
}

func b(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		isSafe := true
		line := scanner.Text()
		split := strings.Split(line, " ")
		list := []int{}
		for _, v := range split {
			val, _ := strconv.Atoi(v)
			list = append(list, val)
		}

		fmt.Printf("Handling %+v\n", list)
		isSafe = isSafeB(list)

		if !isSafe {
			for i, _ := range list {
				if i+1 == len(list) {
					start := deepCopy(list)
					start = start[:len(start)-1]
					fmt.Printf("start: %+v %d\n", start, i)
					isSafe = isSafeB(start)
					if isSafe {
						break;
					}
					break;
				}
				start := deepCopy(list)
				fmt.Printf("start: %+v %d\n", start, i)
				newList := append(start[:i], start[i+1:]...)
				fmt.Printf("%+v: Trying %+v\n", list, newList)
				isSafe = isSafeB(newList)
				if isSafe {
					break;
				}
			}
		}
		if isSafe {
			count++
		}
	}

	return count
}

func deepCopy(a []int) []int {
	new := []int{}
	for _, v := range a {
		new = append(new, v)
	}
	return new
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
