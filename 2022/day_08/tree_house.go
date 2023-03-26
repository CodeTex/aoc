package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readSquareGrid(fp string) [][]int {
	var grid [][]int

	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowIndex := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		if rowIndex == 0 {
			grid = make([][]int, len(line))
		}
		for _, v := range line {
			vInt, _ := strconv.Atoi(v)
			grid[rowIndex] = append(grid[rowIndex], vInt)
		}
		rowIndex++
	}

	return grid
}

func square(x int) int {
	return x * x
}

func isVisible(a int, s []int) bool {
	for _, v := range s {
		if v >= a {
			return false
		}
	}
	return true
}

func getColumn(grid [][]int, colIndex int) []int {
	col := make([]int, 0)
	for _, row := range grid {
		col = append(col, row[colIndex])
	}
	return col
}

func countVisibleTrees(grid [][]int) int {
	var sum int
	gridSize := len(grid)

	// count trees on the edge
	sum += square(gridSize) - square(gridSize-2)

	// TODO: test performance improvement if col based matrix is used instead of getColumn

	for i := 1; i < gridSize-1; i++ {
		for j := 1; j < gridSize-1; j++ {
			v := grid[i][j]
			row := grid[i]
			col := getColumn(grid, j)
			if isVisible(v, row[:j]) || isVisible(v, row[j+1:]) || isVisible(v, col[:i]) || isVisible(v, col[i+1:]) {
				sum++
			}
		}
	}

	return sum
}

func reverse(s []int) []int {
	t := make([]int, len(s))
	for i, v := range s {
		t[len(t)-1-i] = v
	}
	return t
}

func getViewingDistance(a int, s []int) int {
	var distance int
	for _, v := range s {
		distance++
		if v >= a {
			break
		}
	}
	return distance
}

func findHighestScenicScore(grid [][]int) ([]int, int) {
	var score int
	var location []int
	gridSize := len(grid)

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			currentScore := 1
			v := grid[i][j]
			row := grid[i]
			col := getColumn(grid, j)
			currentScore *= getViewingDistance(v, reverse(row[:j]))
			currentScore *= getViewingDistance(v, row[j+1:])
			currentScore *= getViewingDistance(v, reverse(col[:i]))
			currentScore *= getViewingDistance(v, col[i+1:])
			// fmt.Println(i, j, v)
			// fmt.Println(row[:j], getViewingDistance(v, row[:j]))
			// fmt.Println(row[j+1:], getViewingDistance(v, row[j+1:]))
			// fmt.Println(col[:i], getViewingDistance(v, col[:i]))
			// fmt.Println(col[i+1:], getViewingDistance(v, col[i+1:]))
			// fmt.Println(currentScore)
			if currentScore > score {
				score = currentScore
				location = []int{i, j}
			}
		}
	}
	return location, score
}

func main() {
	input_fp := "day_08/input.txt"
	grid := readSquareGrid(input_fp)

	val1 := countVisibleTrees(grid)
	fmt.Println("\nVisible trees:\t\t", val1)

	loc, val2 := findHighestScenicScore(grid)
	fmt.Println("Best location:\t\t", loc)
	fmt.Println("Highest scenic score:\t", val2)
}
