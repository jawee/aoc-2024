package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/day9/atest.txt")
	// file, err := os.Open(pwd + "/day9/btest.txt")
	// file, err := os.Open(pwd + "/day9/input.txt")

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

		sum = getCheckSum(line)
	}

	return sum
}

func getCheckSum(input string) int {
	str := ""
	for i := 0; i < len(input); i+=2 {
		v := input[i]
		no, _ := strconv.Atoi(string(v))
		for range no {
			str += fmt.Sprintf("%d", i/2)
		}
		if i+1 >= len(input) {
			break
		}
		v = input[i+1]
		no, _ = strconv.Atoi(string(v))
		for range no {
			str += "."
		}
	}

	// resultStr := ""
	for i, v := range str {
		if v == '.' {
			var char rune
			for j := len(str)-1; j > 0; j-- {
				if str[j] != '.' {
					char = rune(str[j])
					str = replaceAtIndex(str, '.', j)
					break
				}
			}
			str = replaceAtIndex(str, char, i)
		}
	}

	fmt.Printf("%s\n", str)
	return 0
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}
