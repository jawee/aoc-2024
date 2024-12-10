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
	file, err := os.Open(pwd + "/day10/atest.txt")
	// file, err := os.Open(pwd + "/day10/btest.txt")
	// file, err := os.Open(pwd + "/day10/input.txt")

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

	theMap := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		str := strings.Split(line, "")
		lineArr := []int{}
		for _, s := range str {
			num, _ := strconv.Atoi(s)
			lineArr = append(lineArr, num)
		}
		theMap = append(theMap, lineArr)
	}

	// fmt.Printf("theMap: %+v\n", theMap)
	sum = calculateTrailheads(theMap)

	return sum
}

func calculateTrailheads(theMap [][]int) int {
	sum := 0
	for y := range theMap {
		for x := range theMap[y] {
			if theMap[y][x] == 0 {
				sum += checkIfTrailhead(theMap, x, y)
			}
		}
	}
	return sum
}

func checkIfTrailhead(theMap [][]int, x, y int) int {
	count := 0

	if y-1 >= 0 {
		upVal := theMap[y-1][x]
		if isPath(theMap, y-1, x, upVal) {
			count++
		}
	}

	//right
	if x+1 < len(theMap[0]) {
		rightVal := theMap[y][x+1]
		if isPath(theMap, y, x+1, rightVal) {
			count++
		}
	}

	//down
	if y+1 < len(theMap) {
		downVal := theMap[y+1][x]
		if isPath(theMap, y+1, x, downVal) {
			count++
		}
	}

	//left
	if x-1 >= 0 {
		leftVal := theMap[y][x-1]
		if isPath(theMap, y, x-1, leftVal) {
			count++
		}
	}

	return count
}

func isPath(theMap [][]int, x, y int, val int) bool {
	if val == 9 {
		return true
	}

	if x < 0 || x >= len(theMap[0]) || y < 0 || y >= len(theMap) {
		return false
	}

	// up
	if y-1 >= 0 {
		upVal := theMap[y-1][x]
		if upVal-1 == val && isPath(theMap, y-1, x, upVal) {
			return true
		}
	}

	//right
	if x+1 < len(theMap[0]) {
		rightVal := theMap[y][x+1]
		if rightVal-1 == val && isPath(theMap, y, x+1, rightVal) {
			return true
		}
	}

	//down
	if y+1 < len(theMap) {
		downVal := theMap[y+1][x]
		if downVal-1 == val && isPath(theMap, y+1, x, downVal) {
			return true
		}
	}

	//left
	if x-1 >= 0 {
		leftVal := theMap[y][x-1]
		if leftVal-1 == val && isPath(theMap, y, x-1, leftVal) {
			return true
		}
	}

	return false
}
