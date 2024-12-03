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
	// file, err := os.Open(pwd + "/day3/atest.txt")
	// file, err := os.Open(pwd + "/day3/btest.txt")
	file, err := os.Open(pwd + "/day3/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	res := b(file)
	fmt.Printf("Res: %d\n", res)
}

type parser struct {
	line []rune
	pos  int
}

func (p *parser) nextChar() (rune, bool) {
	if p.pos >= len(p.line) {
		return 'a', false
	}

	res := p.line[p.pos]
	p.pos++

	return res, true
}

func b(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		parser := parser{line: runes, pos: 0}

		rune, ok := parser.nextChar()
		shouldMul := true
		for ok {
			if rune == 'm' && shouldMul {
				n1, n2, isMul := parser.mul()
				if isMul {
					count += n1 * n2
				}
			}
			if rune == 'd' {
				if parser.isDo() {
					shouldMul = true
				}

				if parser.isDont() {
					shouldMul = false
				}
			}
			rune, ok = parser.nextChar()
		}
	}

	return count
}

func (p *parser) isDo() bool {
	if p.pos+2 >= len(p.line) {
		return false
	}

	if p.line[p.pos] == 'o' && p.line[p.pos+1] == '(' && p.line[p.pos+2] == ')' {
		return true
	}
	
	return false
}

func (p *parser) isDont() bool {
	if p.pos+5 >= len(p.line) {
		return false
	}

	if p.line[p.pos] == 'o' && p.line[p.pos+1] == 'n' && p.line[p.pos+2] == '\'' && p.line[p.pos+3] == 't' && p.line[p.pos+4] == '(' && p.line[p.pos+5] == ')' {
		return true
	}
	return false
}

func (p *parser) mul() (int, int, bool) {
	rune, ok := p.nextChar()
	if !ok {
		return 0, 0, false
	}

	n1 := -1
	n2 := -1

	if rune == 'u' {
		rune = p.line[p.pos]
		if rune == 'l' {
			rune = p.line[p.pos+1]
			if rune == '(' {
				rune = p.line[p.pos+2]
				number1 := ""
				c := 2
				for unicode.IsDigit(rune) {
					number1 += string(rune)
					c++
					rune = p.line[p.pos+c]
				}
				if rune == ',' && len(number1) > 0 {
					c++
					rune = p.line[p.pos+c]
					number2 := ""
					for unicode.IsDigit(rune) {
						number2 += string(rune)
						c++
						rune = p.line[p.pos+c]
					}
					if rune == ')' && len(number2) > 0 {
						n1, _ = strconv.Atoi(number1)
						n2, _ = strconv.Atoi(number2)
					}
				}
			}
		}
	}
	if n1 != -1 && n2 != -1 {
		return n1, n2, true
	}
	return 0, 0, false
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		parser := parser{line: runes, pos: 0}

		rune, ok := parser.nextChar()
		shouldMul := true
		for ok {
			if rune == 'm' && shouldMul {
				n1, n2, isMul := parser.mul()
				if isMul {
					count += n1 * n2
				}
			}
			rune, ok = parser.nextChar()
		}
	}

	return count
}
