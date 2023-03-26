package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Motion struct {
	direction string
	distance  int
}

func readInput(fp string) []Motion {
	var motions []Motion

	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		distance, _ := strconv.Atoi(line[1])
		motion := Motion{line[0], distance}
		motions = append(motions, motion)
	}

	return motions
}

func countUnique(s []Position) int {
	map_ := make(map[Position]bool)
	for _, v := range s {
		if _, w := map_[v]; !w {
			map_[v] = true
		}
	}
	return len(map_)
}

func distance(p1, p2 Position) int {
	a := float64(p1.x - p2.x)
	b := float64(p1.y - p2.y)
	return int(math.Sqrt(a*a + b*b))
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func countVisitedPositions(motions []Motion, knotsN int) int {
	if knotsN < 2 {
		panic("Function needs a minimum of 2 knots")
	}

	knotPositions := make([]Position, knotsN)
	for i := 0; i < knotsN; i++ {
		knotPositions[i] = Position{0, 0}
	}

	tailPositions := []Position{knotPositions[knotsN-1]}

	for _, motion := range motions {
		for i := 0; i < motion.distance; i++ {
			switch motion.direction {
			case "U":
				knotPositions[0].y++
			case "R":
				knotPositions[0].x++
			case "D":
				knotPositions[0].y--
			case "L":
				knotPositions[0].x--
			default:
				panic("Unknown direction command given!")
			}

			for j := 1; j < knotsN; j++ {
				if distance(knotPositions[j-1], knotPositions[j]) >= 2 {
					knotPositions[j].x += sign(knotPositions[j-1].x - knotPositions[j].x)
					knotPositions[j].y += sign(knotPositions[j-1].y - knotPositions[j].y)
				}
			}

			tailPositions = append(tailPositions, knotPositions[knotsN-1])
		}
	}

	return countUnique(tailPositions)
}

func main() {
	input_fp := "day_09/input.txt"
	motions := readInput(input_fp)

	val1 := countVisitedPositions(motions, 2)
	fmt.Println("\nNo. of tail positions (2 knots):\t", val1)

	val2 := countVisitedPositions(motions, 10)
	fmt.Println("No. of tail positions (10 knots):\t", val2)
}
