package main

import (
	"bufio"
	"fmt"
	"os"
)

type rucksack = []byte

func read_input(path string) []rucksack {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var arr []rucksack
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, []byte(scanner.Text()))
	}
	return arr
}

func find_priority(r rucksack) byte {
	r_mid := len(r) / 2
	comp_a := r[:r_mid]
	comp_b := r[r_mid:]
	for _, v := range comp_a {
		for _, w := range comp_b {
			if w == v {
				return v
			}
		}
	}
	panic("No priority found")
}

func convert_priority(p byte) int {
	if p < 91 {
		return int(p - 38)
	} else {
		return int(p - 96)
	}
}

func sum_priorities(data []rucksack) int {
	sum := 0
	for _, v := range data {
		sum += convert_priority(find_priority(v))
	}
	return sum
}

func chunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func unique(s []byte) []byte {
	keys := make(map[byte]bool)
	list := []byte{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func find_common_item(data []rucksack) byte {
	for _, v := range data[0] {
		for _, w := range data[1] {
			if w == v {
				for _, x := range data[2] {
					if x == w {
						return x
					}
				}
			}
		}
	}
	panic("No common item found")
}

func sum_badge_priorities(data []rucksack) int {
	sum := 0
	chunk_size := 3
	for i := 0; i < len(data); i += chunk_size {
		end := i + chunk_size
		if end > len(data) {
			end = len(data)
		}
		sum += convert_priority(find_common_item(data[i:end]))
	}
	return sum
}

func main() {
	input_fp := "day_03/input.txt"
	input := read_input(input_fp)

	val1 := sum_priorities(input)
	fmt.Println("\nPriority sum:\t\t", val1)

	val2 := sum_badge_priorities(input)
	fmt.Println("Badge priority sum:\t", val2)
}
