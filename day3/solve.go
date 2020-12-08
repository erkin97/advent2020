package day3

import (
	"bufio"
	"os"
)

func count(grid []string, slopeY, slopeX int) int {
	count := 0

	posRow := slopeY
	posCol := slopeX

	for posRow < len(grid) {
		if grid[posRow][posCol%len(grid[posRow])] == byte('#') {
			count++
		}
		posRow += slopeY
		posCol += slopeX
	}

	return count
}

func input(filename string) ([]string, error) {
	var grid []string
	inputFile, err := os.Open(filename)
	if err != nil {
		return grid, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := scanner.Text()
		grid = append(grid, row)
	}

	return grid, nil
}

type slope struct {
	down  int
	right int
}

//Solve day3
func Solve() (int, error) {
	grid, err := input("day3/input")
	if err != nil {
		return 0, err
	}

	slopes := []slope{
		{down: 1, right: 1},
		{down: 3, right: 1},
		{down: 5, right: 1},
		{down: 7, right: 1},
		{down: 1, right: 2},
	}

	ans := 1
	for _, it := range slopes {
		ans *= count(grid, it.right, it.down)
	}

	return ans, nil
}
