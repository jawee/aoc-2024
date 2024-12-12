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
	// file, err := os.Open(pwd + "/day12/atest.txt")
	// file, err := os.Open(pwd + "/day12/btest.txt")
	file, err := os.Open(pwd + "/day12/input.txt")

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

	grid := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		grid = append(grid, split)
	}
	sum = findFencingCost(grid)
	return sum
}

type coordinate struct {
	x int
	y int
}

type area struct {
	letter string
	area   []coordinate
}

func findFencingCostB(grid [][]string) int {
	areas := []area{}
	visitedArea := [][]bool{}
	for range grid {
		visitedArea = append(visitedArea, make([]bool, len(grid[0])))
	}

	for row := range grid {
		for col := range grid[row] {
			if visitedArea[row][col] {
				continue
			}
			letter := grid[row][col]
			visited := []coordinate{}
			queue := []coordinate{}
			queue = append(queue, coordinate{x: col, y: row})

			visited = floodFill(letter, queue, visited, grid)

			for _, v := range visited {
				visitedArea[v.y][v.x] = true
			}

			areas = append(areas, area{letter: letter, area: visited})
		}
	}

	sum := 0
	for _, v := range areas {
		// fmt.Printf("Area %s: %d\n", k, len(v))
		sides := calculateSides(v)
		sum += (len(v.area) * sides)
	}

	return sum
}

func calculateSides(v area) int {
	if len(v.area) > 0 && len(v.area) < 3 {
		return 4;
	}
	theMap := map[coordinate]bool{}
	for _, v := range v.area {
		theMap[v] = true
	}

	for k, v := range theMap {
		top := theMap[coordinate{x: k.x, y: k.y-1}]
		bottom := theMap[coordinate{x: k.x, y: k.y+1}]
		left := theMap[coordinate{x: k.x-1, y: k.y}]
		right := theMap[coordinate{x: k.x+1, y: k.y}]

	}
	return 0
}
func findFencingCost(grid [][]string) int {
	areas := []area{}
	visitedArea := [][]bool{}
	for range grid {
		visitedArea = append(visitedArea, make([]bool, len(grid[0])))
	}

	for row := range grid {
		for col := range grid[row] {
			if visitedArea[row][col] {
				continue
			}
			letter := grid[row][col]
			visited := []coordinate{}
			queue := []coordinate{}
			queue = append(queue, coordinate{x: col, y: row})

			visited = floodFill(letter, queue, visited, grid)

			for _, v := range visited {
				visitedArea[v.y][v.x] = true
			}

			areas = append(areas, area{letter: letter, area: visited})
		}
	}

	sum := 0
	for _, v := range areas {
		// fmt.Printf("Area %s: %d\n", k, len(v))
		fences := 0
		for _, c := range v.area {
			row := c.y
			col := c.x

			//up
			if row-1 < 0 || grid[row-1][col] != v.letter {
				fences += 1
			}

			//right
			if col+1 >= len(grid[0]) || grid[row][col+1] != v.letter {
				fences += 1
			}

			// down
			if row+1 >= len(grid) || grid[row+1][col] != v.letter {
				fences += 1
			}

			// left
			if col-1 < 0 || grid[row][col-1] != v.letter {
				fences += 1
			}
		}

		// fmt.Printf("Fences for %s: %d, %d = %d\n", v.letter, len(v.area), fences, len(v.area)*fences)

		sum += (len(v.area) * fences)
	}

	return sum
}

func floodFill(letter string, queue []coordinate, visited []coordinate, grid [][]string) []coordinate {

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if contains(visited, n) {
			continue
		}

		if n.x < 0 || n.x > len(grid[0])-1 || n.y < 0 || n.y > len(grid)-1 {
			continue
		}

		if letter != grid[n.y][n.x] {
			continue
		}

		visited = append(visited, n)

		//top
		queue = append(queue, coordinate{x: n.x, y: n.y + 1})
		// bottom
		queue = append(queue, coordinate{x: n.x, y: n.y - 1})
		// right
		queue = append(queue, coordinate{x: n.x + 1, y: n.y})
		// left
		queue = append(queue, coordinate{x: n.x - 1, y: n.y})
	}

	return visited
}

func contains(visited []coordinate, n coordinate) bool {
	for _, v := range visited {
		if v.x == n.x && v.y == n.y {
			return true
		}
	}
	return false
}
