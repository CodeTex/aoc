package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	val_rock     = 1
	val_paper    = 2
	val_scissors = 3
	score_loss   = 0
	score_draw   = 3
	score_win    = 6
)

func read_input(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var arr [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, 2)
		str_row := strings.Split(scanner.Text(), " ")
		if len(str_row) != 2 {
			fmt.Println("To many elements read:", str_row)
			break
		}
		for i, v := range str_row {
			switch v {
			case "A", "X":
				row[i] = val_rock
			case "B", "Y":
				row[i] = val_paper
			case "C", "Z":
				row[i] = val_scissors
			}
		}
		arr = append(arr, row)
	}
	return arr
}

func play(a int, b int) int {
	if a == b {
		return score_draw + b
	}
	if (a < b && (b-a) == 1) || (a > b && (a-b) == 2) {
		return score_win + b
	}
	return score_loss + b
}

func choose_strategy(a int, b int) int {
	var pick int
	switch b {
	case val_rock:
		switch a {
		case val_rock:
			pick = val_scissors
		case val_paper:
			pick = val_rock
		case val_scissors:
			pick = val_paper
		}
	case val_paper:
		pick = a
	case val_scissors:
		switch a {
		case val_rock:
			pick = val_paper
		case val_paper:
			pick = val_scissors
		case val_scissors:
			pick = val_rock
		}
	}
	return pick
}

func play_strategy(a int, b int) int {
	pick := choose_strategy(a, b)
	// fmt.Println(a, b, "->", a, pick, ":", play(a, pick))
	return play(a, pick)
}

func eval(data [][]int, f func(int, int) int) int {
	score := 0
	for _, v := range data {
		score += f(v[0], v[1])
	}
	return score
}

func main() {
	input_fp := "day_02/input.txt"
	input := read_input(input_fp)

	val1 := eval(input, play)
	fmt.Println("\nMaximum score:\t", val1)

	val2 := eval(input, play_strategy)
	fmt.Println("Ideal score:\t", val2)
}
