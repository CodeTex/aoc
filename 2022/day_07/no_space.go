package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseSizeMap(fp string) map[string]int {

	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sizeMap := make(map[string]int)
	dirPath := []string{"/"}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch line := scanner.Text(); line[:4] {
		case "$ cd": // move into dir

			switch dir := line[5:]; dir {
			case "/":
				dirPath = []string{"/"}
			case "..":
				dirPath = dirPath[:len(dirPath)-1]
			default:
				dirPath = append(dirPath, dir)
			}

		case "$ ls", "dir ": // ignore
		default: // file
			fs := strings.Fields(line)
			fileSize, _ := strconv.Atoi(fs[0])
			for i := range dirPath {
				sizeMap[filepath.Join(dirPath[:(i+1)]...)] += fileSize
			}
		}
	}

	return sizeMap
}

func sumBelowCutoff(sizes map[string]int, cutoff int) int {
	var sum int
	for _, v := range sizes {
		if v <= cutoff {
			sum += v
		}
	}

	return sum
}

func findOptimalDir(sizeMap map[string]int, limit int) int {
	selectedSize := 0
	for _, v := range sizeMap {
		if selectedSize == 0 {
			selectedSize = v
		} else if v >= limit && v < selectedSize {
			selectedSize = v
		}
	}
	return selectedSize
}

func main() {
	input_fp := "day_07/input.txt"
	cutoff := 100000
	totalSpace := 70000000
	updateSpace := 30000000

	sizeMap := parseSizeMap(input_fp)
	val1 := sumBelowCutoff(sizeMap, cutoff)
	fmt.Println("\nAggregated dir sizes <100,000:\t", val1)

	requiredSpace := updateSpace - (totalSpace - sizeMap["/"])
	fmt.Println(requiredSpace)
	val2 := findOptimalDir(sizeMap, requiredSpace)
	fmt.Println("Size of directory to delete:\t", val2)
}
