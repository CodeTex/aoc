package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op  string
	val int
}

func readInput(fp string) []Instruction {
	var instructions []Instruction

	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := 0
		line := strings.Fields(scanner.Text())
		if len(line) == 2 {
			val, _ = strconv.Atoi(line[1])
		}
		instructions = append(instructions, Instruction{line[0], val})
	}

	return instructions
}

func run(instructions []Instruction) []int {
	register := 1
	registerSlice := []int{register}
	for _, instruction := range instructions {
		switch instruction.op {
		case "addx":
			registerSlice = append(registerSlice, register)
			register += instruction.val
		}
		registerSlice = append(registerSlice, register)
	}
	return registerSlice
}

func sumSignalStrenghts(registerOps []int, cycles []int) int {
	var sum int
	for _, readCycle := range cycles {
		sum += readCycle * registerOps[readCycle-1]
	}
	return sum
}

func draw(registerOps []int) {
	var pixel string
	for i, v := range registerOps {
		col := i % 40
		if col >= v-1 && col <= v+1 {
			pixel = "#"
		} else {
			pixel = "."
		}
		if col == 39 {
			fmt.Println(pixel)
		} else {
			fmt.Print(pixel)
		}
	}
}

func main() {
	input_fp := "day_10/input.txt"
	instructions := readInput(input_fp)

	registerOps := run(instructions)
	cycles := []int{20, 60, 100, 140, 180, 220}
	val1 := sumSignalStrenghts(registerOps, cycles)
	fmt.Println("\nSum of signal strenghts:\t", val1)

	fmt.Println("CRT message:")
	draw(registerOps)
}
