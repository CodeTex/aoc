package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_input(path string) ([]int16, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int16
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return lines, err
		}
		lines = append(lines, int16(val))
	}

	return lines, scanner.Err()
}

func arr_sum(v []int16) int {
	arrSum := 0
	for _, e := range v {
		arrSum = arrSum + int(e)
	}
	return arrSum
}

func calc_step_one(data []int16) int {
	prev := data[0]
	inc := 0
	for _, e := range data[1:] {
		if e > prev {
			inc++
		}
		prev = e
	}
	return inc
}

func calc_step_two(data []int16) int {
	prev := arr_sum(data[:3])
	inc := 0
	for i := range data[1:] {
		if i >= len(data)-2 {
			break
		}
		new := arr_sum((data[i : i+3]))
		if new > prev {
			inc++
		}
		prev = new
	}
	return inc
}

func main() {
	input_fp := "day_01/input.txt"
	data, err := read_input(input_fp)
	if err != nil {
		panic(err)
	}

	sol1 := calc_step_one(data)
	fmt.Println("\nSolution:\t", sol1)
	sol2 := calc_step_two(data)
	fmt.Println("\nSolution:\t", sol2)
}
