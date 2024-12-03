package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func main() {
	pwd, _ := os.Getwd()
	// file, err := os.Open(pwd + "/atest.txt")
	// file, err := os.Open(pwd + "/day3/btest.txt")
	file, err := os.Open(pwd + "/day3/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	res := a(file)
	fmt.Printf("Res: %d\n", res)
}

type parser struct {
	line string
	pos  int
}

func (p *parser) nextChar() (rune, bool) {
	if p.pos >= len(p.line) {
		return 'a', false
	}

	res := p.line[p.pos]
	p.pos++

	return rune(res), true
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		// split := strings.Split(line, " ")
		// parser := parser{line: line, pos: 0}
		runes := []rune(line)
		for i, v := range runes {
			rune := v
			if rune == 'm' {
				rune = runes[i+1]
				if rune == 'u' {
					rune = runes[i+2]
					if rune == 'l' {
						rune = runes[i+3]
						if rune == '(' {
							rune = runes[i+4]
							number1 := ""
							c := 4
							for unicode.IsDigit(rune) {
								number1 += string(rune)
								c++
								rune = runes[i+c]
							}
							if rune == ',' && len(number1) > 0 {
								c++
								rune = runes[i+c]
								number2 := ""
								for unicode.IsDigit(rune) {
									number2 += string(rune)
									c++
									rune = runes[i+c]
								}
								if rune == ')' && len(number2) > 0 {
									n1, _ := strconv.Atoi(number1)
									n2, _ := strconv.Atoi(number2)
									count += n1 * n2
								}
							}
						}
					}
				}
			}
		}
		// rune, ok := parser.nextChar()
		// for ok {
		// 	if rune == 'm' {
		// 		rune, ok = parser.nextChar()
		// 		if rune == 'u' {
		// 			rune, ok = parser.nextChar()
		// 			if rune == 'l' {
		// 				rune, ok = parser.nextChar()
		// 				if rune == '(' {
		// 					rune, ok = parser.nextChar()
		// 					number1 := ""
		// 					for unicode.IsDigit(rune) {
		// 						number1 += string(rune)
		// 						rune, ok = parser.nextChar()
		// 					}
		// 					if rune == ',' && len(number1) > 0 {
		// 						rune, ok = parser.nextChar()
		// 						number2 := ""
		// 						for unicode.IsDigit(rune) {
		// 							number2 += string(rune)
		// 							rune, ok = parser.nextChar()
		// 						}
		// 						if rune == ')' && len(number2) > 0 {
		// 							n1, _ := strconv.Atoi(number1)
		// 							n2, _ := strconv.Atoi(number1)
		// 							count += n1 * n2
		// 							break
		// 						}
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// }
	}

	return count
}
