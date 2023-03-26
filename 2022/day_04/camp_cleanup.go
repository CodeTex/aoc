package main

import (
	"bufio"
	"fmt"
	"os"
)

type sRange = [2]int
type pAssignment = [2]sRange

func read_input(path string) []pAssignment {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var arr []pAssignment
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var a, b, c, d int
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		if err != nil {
			panic("Format parsing error on input read")
		}
		arr = append(arr, [2][2]int{{a, b}, {c, d}})
	}
	return arr
}

func contains(s1 sRange, s2 sRange) bool {
	if s1[0] <= s2[0] && s1[1] >= s2[1] {
		return true
	}
	return false
}

func intersect(s1 sRange, s2 sRange) bool {
	if s1[0] <= s2[1] && s2[0] <= s1[1] {
		return true
	}
	return false
}

func find_overlapping_assignments(data []pAssignment) int {
	sum := 0
	for _, v := range data {
		if contains(v[0], v[1]) || contains(v[1], v[0]) {
			sum++
		}
	}
	return sum
}

func find_intersecting_assignments(data []pAssignment) int {
	sum := 0
	for _, v := range data {
		if intersect(v[0], v[1]) {
			sum++
		}
	}
	return sum
}

func main() {
	input_fp := "day_04/input.txt"
	input := read_input(input_fp)

	val1 := find_overlapping_assignments(input)
	fmt.Println("No. of overlaps:\t", val1)

	val2 := find_intersecting_assignments(input)
	fmt.Println("No. of intersects:\t", val2)
}
