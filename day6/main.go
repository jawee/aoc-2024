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
	// file, err := os.Open(pwd + "/day6/atest.txt")
	// file, err := os.Open(pwd + "/day6/btest.txt")
	file, err := os.Open(pwd + "/day6/input.txt")

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
	i := 0
	guardPosX := 0
	guardPosY := 0

	theMap := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		theMap = append(theMap, []string{})
		for x, v := range split {
			if v == "^" {
				guardPosY = i
				guardPosX = x
			}
			theMap[i] = append(theMap[i], v)
		}
		i++
	}

	// fmt.Printf("%+v\n", theMap)
	// fmt.Printf("Guard position: %d,%d\n", guardPosX, guardPosY)

	sum = guardWalkStuck(guardPosX, guardPosY, theMap)

	return sum
}

func a(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	i := 0
	guardPosX := 0
	guardPosY := 0

	theMap := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		theMap = append(theMap, []string{})
		for x, v := range split {
			if v == "^" {
				guardPosY = i
				guardPosX = x
			}
			theMap[i] = append(theMap[i], v)
		}
		i++
	}

	// fmt.Printf("%+v\n", theMap)
	// fmt.Printf("Guard position: %d,%d\n", guardPosX, guardPosY)

	sum = guardWalk(guardPosX, guardPosY, theMap)

	return sum
}

func createVisited(theMap [][]string) [][]bool {
	visited := make([][]bool, len(theMap))
	for i := range theMap {
		visited[i] = make([]bool, len(theMap[0]))
	}

	return visited
}

func deepCopy(theMap [][]string) [][]string {
	new := [][]string{}
	for _, row := range theMap {
		newRow := make([]string, len(row))
		copy(newRow, row)
		new = append(new, newRow)
	}

	return new
}

type obstacle struct {
	x, y int
}

func guardWalkStuck(guardX, guardY int, theMap [][]string) int {
	count := 0
	for i := range theMap {
		for j := range theMap[i] {
			isLoop := isLoop(guardX, guardY, i, j, theMap)
			if isLoop {
				count++
			}
		}
	}

	return count
}

func isLoop(guardX, guardY, x, y int, theMap [][]string) bool {
	if guardX == x && guardY == y {
		return false
	}
	newTheMap := deepCopy(theMap)
	newTheMap[y][x] = "#" // add new obstacle
	obstacles := []obstacle{}

	isOnBoard := true
	dx := 0
	dy := -1

	path := []obstacle{}

	for isOnBoard {
		if (guardX+dx) < 0 || (guardX+dx) >= len(newTheMap[0]) || (guardY+dy) < 0 || (guardY+dy) >= len(newTheMap) {
			// fmt.Printf("%d,%d, fell off the board\n", i, j)
			isOnBoard = false
			break
		}

		// obstacle
		if newTheMap[guardY+dy][guardX+dx] == "#" {
			if isCycle(obstacles, guardX+dx, guardY+dy) {
				// fmt.Printf("Obstacles: %+v\n", obstacles)
				// fmt.Printf("%d,%d, Hit cycle at %d,%d\n", x, y, guardX+dx, guardY+dy)
				// fmt.Printf("Path: %+v\n", path)
				return true
			}
			obstacles = append(obstacles, obstacle{guardX + dx, guardY + dy})
			if dy == -1 {
				dy = 0
				dx = 1
				continue
			}
			if dx == 1 {
				dy = 1
				dx = 0
				continue
			}
			if dy == 1 {
				dy = 0
				dx = -1
				continue
			}
			if dx == -1 {
				dy = -1
				dx = 0
				continue
			}
		}

		path = append(path, obstacle{x: guardX + dx, y: guardY + dy})

		guardX += dx
		guardY += dy
	}

	return false
}

func isCycle(obstacles []obstacle, x, y int) bool {
	foundCount := 0
	for _, o := range obstacles {
		if o.x == x && o.y == y {
			foundCount++
		}
	}
	return foundCount > 2
}

func guardWalk(guardX, guardY int, theMap [][]string) int {
	visited := make([][]bool, len(theMap))
	for i := range theMap {
		visited[i] = make([]bool, len(theMap[0]))
	}

	count := 0
	isOnBoard := true
	dx := 0
	dy := -1

	for isOnBoard {
		if (guardX+dx) < 0 || (guardX+dx) >= len(theMap[0]) || (guardY+dy) < 0 || (guardY+dy) >= len(theMap) {
			isOnBoard = false
			break
		}

		// obstacle
		if theMap[guardY+dy][guardX+dx] == "#" {
			// fmt.Printf("Hit obstacle at %d,%d\n", guardX+dx, guardY+dy)
			if dy == -1 {
				dy = 0
				dx = 1
				continue
			}
			if dx == 1 {
				dy = 1
				dx = 0
				continue
			}
			if dy == 1 {
				dy = 0
				dx = -1
				continue
			}
			if dx == -1 {
				dy = -1
				dx = 0
				continue
			}
		}

		guardX += dx
		guardY += dy
		if !visited[guardY][guardX] {
			visited[guardY][guardX] = true
			count++
		}
	}

	return count
}
