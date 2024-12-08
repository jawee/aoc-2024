package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	// file, err := os.Open(pwd + "/day8/atest.txt")
	// file, err := os.Open(pwd + "/day8/btest.txt")
	file, err := os.Open(pwd + "/day8/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	// res := a(file)
	res := b(file)
	fmt.Printf("%d\n", res)
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	strs := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		strs = append(strs, split)
	}

	// for _, v := range strs {
	// 	fmt.Printf("%+v\n", v)
	// }
	//
	sum = countAntinodes(strs)

	return sum
}
func b(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	strs := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		strs = append(strs, split)
	}

	// for _, v := range strs {
	// 	fmt.Printf("%+v\n", v)
	// }
	//
	sum = countAntinodesB(strs)

	return sum
}

type antinode struct {
	row int
	col int
}

func getAntinodesB(row1, col1, row2, col2, rowMax, colMax int) []antinode {
	antiNodes := []antinode{}
	dr := row2 - row1
	dc := col2 - col1

	if (dr > 0 && dc > 0) || (dr < 0 && dc < 0) {
		//down right
		i := 1
		tmpNodes := []antinode{}
		for true {
			a1 := antinode{row: row2 + (i * (row2 - row1)), col: col2 + (i * (col2 - col1))}
			i++
			if a1.row < 0 || a1.row >= rowMax {
				break
			}
			if a1.col < 0 || a1.col >= colMax {
				break
			}

			tmpNodes = append(tmpNodes, a1)
		}

		antiNodes = append(antiNodes, tmpNodes...)

		i = 1
		tmpNodes = []antinode{}
		for true {
			a2 := antinode{row: row1 - (i * (row2 - row1)), col: col1 - (i * (col2 - col1))}
			i++
			if a2.row < 0 || a2.row >= rowMax {
				break
			}
			if a2.col < 0 || a2.col >= colMax {
				break
			}
			tmpNodes = append(tmpNodes, a2)
		}
		antiNodes = append(antiNodes, tmpNodes...)
	}
	if (dr < 0 && dc > 0) || (dr > 0 && dc < 0) {
		//up right
		i := 1
		tmpNodes := []antinode{}
		for true {
			a1 := antinode{row: row2 + (i * (row2 - row1)), col: col2 - (i * (col1 - col2))}
			i++
			if a1.row < 0 || a1.row >= rowMax {
				break
			}
			if a1.col < 0 || a1.col >= colMax {
				break
			}
			tmpNodes = append(tmpNodes, a1)
		}

		antiNodes = append(antiNodes, tmpNodes...)

		i = 1
		tmpNodes = []antinode{}
		for true {
			a2 := antinode{row: row1 - (i * (row2 - row1)), col: col1 + (i * (col1 - col2))}
			i++
			if a2.row < 0 || a2.row >= rowMax {
				break
			}
			if a2.col < 0 || a2.col >= colMax {
				break
			}
			tmpNodes = append(tmpNodes, a2)
		}
		antiNodes = append(antiNodes, tmpNodes...)
	}

	return antiNodes
}

func countAntinodesB(strs [][]string) int {
	sum := 0
	antiNodes := [][]bool{}
	for range strs {
		antiNodes = append(antiNodes, make([]bool, len(strs)))
	}
	for row := range strs {
		for col := range strs[row] {
			if strs[row][col] != "." {
				antenna := strs[row][col]
				for row2 := range strs {
					for col2 := range strs[row2] {
						if row == row2 && col2 == col {
							continue
						}

						if strs[row2][col2] == antenna {
							// fmt.Printf("%d, %d. Found another antenna at %d, %d.\n", row, col, row2, col2)

							res := getAntinodesB(row, col, row2, col2, len(strs), len(strs[0]))
							// fmt.Printf("AntiNodes: %+v\n", antiNodes)
							antiNodes[row][col] = true
							antiNodes[row2][col2] = true

							for _, v := range res {
								if v.row < 0 || v.row >= len(antiNodes) {
									continue
								}
								if v.col < 0 || v.col >= len(antiNodes[0]) {
									continue
								}

								antiNodes[v.row][v.col] = true
							}
						}
					}
				}
			}

		}
	}

	for i := range antiNodes {
		for j := range antiNodes[i] {
			if antiNodes[i][j] == true {
				sum++
			}
		}
	}

	return sum
}

func countAntinodes(strs [][]string) int {
	sum := 0
	antiNodes := [][]bool{}
	for range strs {
		antiNodes = append(antiNodes, make([]bool, len(strs)))
	}
	for row := range strs {
		for col := range strs[row] {
			if strs[row][col] != "." {
				antenna := strs[row][col]
				for row2 := range strs {
					for col2 := range strs[row2] {
						if row == row2 && col2 == col {
							continue
						}

						if strs[row2][col2] == antenna {
							// fmt.Printf("%d, %d. Found another antenna at %d, %d.\n", row, col, row2, col2)

							res := getAntinodes(row, col, row2, col2)
							// fmt.Printf("AntiNodes: %+v\n", antiNodes)

							for _, v := range res {
								if v.row < 0 || v.row >= len(antiNodes) {
									continue
								}
								if v.col < 0 || v.col >= len(antiNodes[0]) {
									continue
								}

								antiNodes[v.row][v.col] = true
							}
						}
					}
				}
			}

		}
	}

	for i := range antiNodes {
		for j := range antiNodes[i] {
			if antiNodes[i][j] == true {
				sum++
			}
		}
	}

	return sum
}

func getAntinodes(row1, col1, row2, col2 int) []antinode {
	antiNodes := []antinode{}
	dr := row2 - row1
	dc := col2 - col1

	if (dr > 0 && dc > 0) || (dr < 0 && dc < 0) {
		//down right
		a1 := antinode{row: row2 + (row2 - row1), col: col2 + (col2 - col1)}
		antiNodes = append(antiNodes, a1)
		a2 := antinode{row: row1 - (row2 - row1), col: col1 - (col2 - col1)}
		antiNodes = append(antiNodes, a2)
	}
	if (dr < 0 && dc > 0) || (dr > 0 && dc < 0) {
		//up right
		a1 := antinode{row: row2 + (row2 - row1), col: col2 - (col1 - col2)}
		antiNodes = append(antiNodes, a1)
		a2 := antinode{row: row1 - (row2 - row1), col: col1 + (col1 - col2)}
		antiNodes = append(antiNodes, a2)
	}

	return antiNodes
}

func absOne(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
func absTwo(a int, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
