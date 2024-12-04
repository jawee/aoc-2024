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
	file, err := os.Open(pwd + "/day4/atest.txt")
	// file, err := os.Open(pwd + "/day4/btest.txt")
	// file, err := os.Open(pwd + "/day4/input.txt")

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

	matrix := [][]string{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		matrix = append(matrix, []string{})
		matrix[i] = append(matrix[i], split...)
		i++
	}

	// for _, row := range matrix {
	// 	fmt.Printf("%+v\n", row)
	// }

	for row := range matrix {
		for col := range matrix[row] {
			// fmt.Printf("%s", matrix[row][col])
			if matrix[row][col] == "A" {
				if searchMas(matrix, row, col) {
					sum++;
				}
			}
		}
	}
	return sum
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0

	matrix := [][]string{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		matrix = append(matrix, []string{})
		matrix[i] = append(matrix[i], split...)
		i++
	}

	// for _, row := range matrix {
	// 	fmt.Printf("%+v\n", row)
	// }

	for row := range matrix {
		for col := range matrix[row] {
			fmt.Printf("%s", matrix[row][col])
			if matrix[row][col] == "X" {
				// if isForward(matrix, row, col) {
				if searchWord(matrix, row, col, 1, 0) {
					sum++
				}
				// if isBackward(matrix, row, col) {
				if searchWord(matrix, row, col, -1, 0) {
					sum++
				}
				// if isUp(matrix, row, col) {
				if searchWord(matrix, row, col, 0, -1) {
					sum++
				}
				// if isDown(matrix, row, col) {
				if searchWord(matrix, row, col, 0, 1) {
					sum++
				}
				// if isDiagonalLeftUp(matrix, row, col) {
				if searchWord(matrix, row, col, -1, -1) {
					sum++
				}
				// if isDiagonalLeftDown(matrix, row, col) {
				if searchWord(matrix, row, col, -1, 1) {
					sum++
				}
				// if isDiagonalRightUp(matrix, row, col) {
				if searchWord(matrix, row, col, 1, -1) {
					sum++
				}
				// if isDiagonalRightDown(matrix, row, col) {
				if searchWord(matrix, row, col, 1, 1) {
					sum++
				}
			}
		}
		fmt.Printf("\n")
	}

	return sum
}

func searchMas(matrix [][]string, row, col int) bool {
	if col-1 < 0 || col+1 >= len(matrix[0]) {
		return false
	}
	if row-1 < 0 || row+1 >= len(matrix) {
		return false
	}

	if matrix[row-1][col-1] != "M" && matrix[row-1][col-1] != "S"  {
		return false
	}
	if matrix[row-1][col+1] != "M" && matrix[row-1][col+1] != "S"  {
		return false
	}
	if matrix[row+1][col-1] != "M" && matrix[row+1][col-1] != "S"  {
		return false
	}
	if matrix[row+1][col+1] != "M" && matrix[row+1][col+1] != "S"  {
		return false
	}
	return true
}
func searchWord(matrix [][]string, row, col, dx, dy int) bool {
	if (col+(dx*3)) < 0 || (col+(dx*3)) >= len(matrix[0]) {
		return false
	}

	if (row+(dy*3)) < 0 || (row+(dy*3)) >= len(matrix) {
		return false
	}

	if matrix[row+(1*dy)][col+(1*dx)] != "M" {
		return false
	}

	if matrix[row+(2*dy)][col+(2*dx)] != "A" {
		return false
	}

	if matrix[row+(3*dy)][col+(3*dx)] != "S" {
		return false
	}

	return true
}
