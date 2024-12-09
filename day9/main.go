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
	// res := a(file)
	res := b(file)
	fmt.Printf("%d\n", res)
}

func b(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		sum = getCheckSumB(line)
	}

	return sum
}

type blockb struct {
	id     int
	empty  bool
	length int
}

func getCheckSumB(input string) int {
	structs := []blockb{}
	for i := 0; i < len(input); i += 2 {
		v := input[i]
		no, _ := strconv.Atoi(string(v))
		structs = append(structs, blockb{id: i / 2, empty: false, length: no})
		if i+1 >= len(input) {
			break
		}
		v = input[i+1]
		no, _ = strconv.Atoi(string(v))
		structs = append(structs, blockb{id: 0, empty: true, length: no})
	}

	printStructs(structs)

	maxId := -1
	for _, v := range structs {
		if maxId < v.id && !v.empty {
			maxId = v.id
		}
	}
	for i := maxId; i >= 0; i-- {
		// fmt.Printf("looking for id: %d\n", i)
		idx := -1
		var block blockb
		for id, v := range structs {
			if v.id == i && !v.empty {
				idx = id
				block = v
			}
		}
		if idx == -1 {
			continue
		}
		// fmt.Printf("idx: %d\n", idx)
		for j := range len(structs) {
			if j > idx {
				break
			}
			empty := structs[j]
			if empty.empty {
				if block.length == empty.length {
					// fmt.Printf("simple swap\n")
					printStructs(structs)
					//simple swap
					structs[idx], structs[j] = structs[j], structs[idx]
					printStructs(structs)
					break
				} else if block.length < empty.length {
					// fmt.Printf("complex swap\n")
					clonedStructs := deepCopy(structs)
					printStructs(clonedStructs)
					before := clonedStructs[:j]
					before = append(before, block)
					before = append(before, blockb{id: 0, empty: true, length: empty.length - block.length})
					after := deepCopy(structs)[j+1 : idx]
					after = append(after, blockb{id: 0, empty: true, length: block.length})
					lastPart := clonedStructs[idx+1:]

					structs = append(deepCopy(before), deepCopy(after)...)
					structs = append(structs, deepCopy(lastPart)...)
					printStructs(structs)
					break
				}
			}
		}

	}

	printStructs(structs)

	sum := 0
	i := 0
	for _, v := range structs {
		if v.empty {
			i += v.length
			continue
		}
		for range v.length {
			sum += (i * v.id)
			i++
		}

	}
	return sum
}
func deepCopy(a []blockb) []blockb {
	new := []blockb{}
	for _, v := range a {
		new = append(new, v)
	}
	return new
}

func printStructs(structs []blockb) {
	for _, v := range structs {
		for range v.length {
			if v.empty {
				// fmt.Printf(".")
			} else {
				// fmt.Printf("%d", v.id)
			}
		}
	}
	// fmt.Printf("\n")
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		sum = getCheckSumA(line)
	}

	return sum
}

type blocka struct {
	id    int
	empty bool
}

func getCheckSumA(input string) int {
	structs := []blocka{}
	for i := 0; i < len(input); i += 2 {
		v := input[i]
		no, _ := strconv.Atoi(string(v))
		for range no {
			structs = append(structs, blocka{id: i / 2, empty: false})
		}
		if i+1 >= len(input) {
			break
		}
		v = input[i+1]
		no, _ = strconv.Atoi(string(v))
		for range no {
			structs = append(structs, blocka{id: i / 2, empty: true})
		}
	}

	// resultStr := ""
	for i, v := range structs {
		if v.empty == true {
			// var b block
			for j := len(structs) - 1; j > 0; j-- {
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
		sum += (i * v.id)
	}
	return sum
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
