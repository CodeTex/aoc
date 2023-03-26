// With the help of https://github.com/tsenart/advent/tree/master/2022/12
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

func timeTrack(start time.Time, action string) {
	elapsed := time.Since(start)
	fmt.Printf("\n%s took %s\n", action, elapsed)
}

func parseFile(file *os.File) ([]byte, int) {
	defer timeTrack(time.Now(), "Parsing the input")

	var lines []byte
	var width int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Bytes()
		if width == 0 {
			width = len(row)
		}
		lines = append(lines, row...)
	}

	return lines, width
}

func plot(heightmap []byte, width int, start, pos int, path []int) {
	trace := make([]bool, len(heightmap))
	for n := pos; n != start; n = path[n] {
		trace[n] = true
	}

	for i := range heightmap {
		if i%width == 0 {
			fmt.Println()
		}

		if i == pos || trace[i] {
			fmt.Printf("\x1b[31m%c\x1b[0m", heightmap[i])
		} else if path[i] != -1 {
			fmt.Printf("\x1b[30m%c\x1b[0m", heightmap[i])
		} else {
			fmt.Printf("%c", heightmap[i])
		}
	}

	fmt.Println()
}

func getNeighbours(row, col, width, totalLen int) []int {
	neighbours := make([]int, 0, 4)
	if row > 0 {
		neighbours = append(neighbours, (row-1)*width+col)
	}
	if row < (totalLen/width)-1 {
		neighbours = append(neighbours, (row+1)*width+col)
	}
	if col > 0 {
		neighbours = append(neighbours, row*width+(col-1))
	}
	if col < width-1 {
		neighbours = append(neighbours, row*width+(col+1))
	}
	return neighbours
}

func findShortestPath(heightmap *[]byte, start, end, width int, visualize bool) int {
	defer timeTrack(time.Now(), "Finding shortest path")

	var distance int
	mapSlice := make([]byte, len(*heightmap))
	copy(mapSlice, *heightmap)

	queue := [][2]int{{start, 0}}
	path := make([]int, len(mapSlice))

	for i := range path {
		path[i] = -1
	}

	mapSlice[start] = 'a'
	mapSlice[end] = 'z'

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		i, steps := current[0], current[1]
		row, col := i/width, i%width

		neighbours := getNeighbours(row, col, width, len(mapSlice))

		for _, n := range neighbours {
			if path[n] == -1 && mapSlice[n] <= mapSlice[i]+1 {
				if path[n] = i; n == end {
					distance = steps + 1
					if visualize {
						plot(mapSlice, width, start, n, path)
					}
					queue = nil
					break
				}
				queue = append(queue, [2]int{n, steps + 1})
			}
		}
	}

	return distance
}

func findOptimalPath(heightmap *[]byte, start, width int, visualize bool) int {
	defer timeTrack(time.Now(), "Finding shortest path")

	var distance int
	mapSlice := make([]byte, len(*heightmap))
	copy(mapSlice, *heightmap)

	queue := [][2]int{{start, 0}}
	path := make([]int, len(mapSlice))

	for i := range path {
		path[i] = -1
	}

	mapSlice[start] = 'z'
	mapSlice[bytes.IndexByte(mapSlice, 'S')] = 'a'

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		i, steps := current[0], current[1]
		row, col := i/width, i%width

		neighbours := getNeighbours(row, col, width, len(mapSlice))

		for _, n := range neighbours {
			if path[n] == -1 && mapSlice[n] >= mapSlice[i]-1 {
				if path[n] = i; mapSlice[n] == 'a' {
					distance = steps + 1
					if visualize {
						plot(mapSlice, width, start, n, path)
					}
					queue = nil
					break
				}
				queue = append(queue, [2]int{n, steps + 1})
			}
		}
	}

	return distance
}

func main() {
	mapData, width := parseFile(os.Stdin)
	start := bytes.IndexByte(mapData, 'S')
	end := bytes.IndexByte(mapData, 'E')

	shortestDistance := findShortestPath(&mapData, start, end, width, true)
	fmt.Printf(
		"\nShortest distance from (%v, %v) to (%v, %v): %d steps\n",
		start/width, start%width, end/width, end%width,
		shortestDistance,
	)

	bestHikeDistance := findOptimalPath(&mapData, end, width, true)
	fmt.Printf(
		"\nShortest distance from (%v, %v) to (%v, %v): %d steps\n",
		start/width, start%width, end/width, end%width,
		bestHikeDistance,
	)
}
