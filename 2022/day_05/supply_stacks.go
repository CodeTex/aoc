package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readStacks() [][]string {
	initPos := [9]string{
		"QWPSZRHD",
		"VBRWQHF",
		"CVSH",
		"HFG",
		"PGJBZ",
		"QTJHWFL",
		"ZTWDLVJN",
		"DTZCJGHF",
		"WPVMBH",
	}
	var stacks [][]string
	for _, v := range initPos {
		crates := strings.Split(v, "")
		stacks = append(stacks, crates)
	}
	return stacks
}

func readCommands(path string) [][3]int {
	var arr [][3]int

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowIndex := 1
	for scanner.Scan() {
		if rowIndex <= 10 {
			rowIndex++
			continue
		}
		var a, b, c int
		_, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &a, &b, &c)
		if err != nil {
			panic("Format parsing error on input read")
		}
		arr = append(arr, [3]int{a, b, c})
	}
	return arr
}

func moveCrates(stacks [][]string, commands [][3]int, reverse bool) string {
	var message string
	for _, command := range commands {
		quant := command[0]
		from := command[1] - 1
		to := command[2] - 1
		sliceIndex := len(stacks[from]) - quant
		cargo := stacks[from][sliceIndex:]
		if reverse {
			for i := len(cargo) - 1; i >= 0; i-- {
				stacks[to] = append(stacks[to], cargo[i])
			}
		} else {
			for i := 0; i < len(cargo); i++ {
				stacks[to] = append(stacks[to], cargo[i])
			}
		}
		stacks[from] = stacks[from][:sliceIndex]
	}
	for _, v := range stacks {
		message += v[len(v)-1]
	}
	return message
}

func main() {
	input_fp := "day_05/input.txt"
	commands := readCommands(input_fp)

	val1 := moveCrates(readStacks(), commands, true)
	fmt.Println("CrateMover 9000 message:\t", val1)

	val2 := moveCrates(readStacks(), commands, false)
	fmt.Println("CrateMover 9001 message:\t", val2)
}
