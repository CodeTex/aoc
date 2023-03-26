package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func arr_sum(v []int) int {
	arrSum := 0
	for _, e := range v {
		arrSum = arrSum + e
	}
	return arrSum
}

func buffest_inventory(path string) (int, error) {
	buffest_inventory := 0
	var cal_count int

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if cal_count > buffest_inventory {
				buffest_inventory = cal_count
			}
			cal_count = 0
		}
		cal_count += val
	}

	return buffest_inventory, scanner.Err()
}

func replace_min(arr []int, val int) {
	idx := 0
	min := arr[idx]
	for i, v := range arr[idx+1:] {
		if v < min {
			min = v
			idx = i + 1
		}
	}
	if val > min {
		arr[idx] = val
	}
}

func buffest_inventories(path string, n int) (int, error) {
	buffest_inventories := make([]int, n)
	var cal_count int

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			replace_min(buffest_inventories, cal_count)
			cal_count = 0
		}
		cal_count += val
	}

	return arr_sum(buffest_inventories), scanner.Err()
}

func main() {
	input_fp := "day_01/input.txt"

	val, err := buffest_inventory(input_fp)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nBiggest inventory:\t", val)

	top_n := 3
	val, err = buffest_inventories(input_fp, top_n)
	if err != nil {
		panic(err)
	}
	fmt.Println("Top", top_n, "inventories:\t", val)
}
