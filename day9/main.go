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
	// file, err := os.Open(pwd + "/day9/atest.txt")
	// file, err := os.Open(pwd + "/day9/btest.txt")
	file, err := os.Open(pwd + "/day9/input.txt")

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

type block struct {
	id int
	empty bool
}

func getCheckSum(input string) int {
	structs := []block{}
	for i := 0; i < len(input); i+=2 {
		v := input[i]
		no, _ := strconv.Atoi(string(v))
		for range no {
			structs = append(structs, block{id: i/2, empty: false})
		}
		if i+1 >= len(input) {
			break
		}
		v = input[i+1]
		no, _ = strconv.Atoi(string(v))
		for range no {
			structs = append(structs, block{id: i/2, empty: true})
		}
	}

	// resultStr := ""
	for i, v := range structs {
		if v.empty == true {
			// var b block
			for j := len(structs)-1; j > 0; j-- {
				if j < i {
					break
				}
				// if str[j] != '.' {
				if !structs[j].empty {
					// char = rune(str[j])
					// str = replaceAtIndex(str, '.', j)
					// b = structs[j]
					structs[i], structs[j] = structs[j], structs[i]
					// i[0], i[1] = i[1], i[0]
					

					break
				}
			}
			// str = replaceAtIndex(str, char, i)
			// fmt.Printf("%s\n", str)
		}
	}
	sum := 0
	for i, v := range structs {
		if v.empty {
			continue
		}
		sum += (i*v.id)
	}
	return sum
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}
