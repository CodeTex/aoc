package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const (
	ITMSL = len("  Starting items: ")
	OPRSL = len("  Operation: new = ")
	TESTL = len("  Test: divisible by ")
	TRUEL = len("    If true: throw to monkey ")
	FALSL = len("    If false: throw to monkey ")
)

const (
	ADD int = iota
	SUBTRACT
	MULTIPLY
	DIVIDE
	SQUARE
)

type Monkey struct {
	items        []int
	operator     int
	operationNum int
	testDivisor  int
	next         [2]int // 0 -> true, 1 -> false on throw condition
	inspected    int
}

var ErrOverflow = errors.New("integer overflow")

func parseFile(fp string) []string {
	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	data := strings.TrimRightFunc(string(bytes), unicode.IsSpace)
	return strings.Split(data, "\n\n")
}

func parseInput(data []string) []Monkey {
	monkeys := make([]Monkey, len(data))

	for i, v := range data {
		monkey := Monkey{}

		lines := strings.Split(v, "\n")
		for j, line := range lines[1:] {

			switch j {
			case 0:
				items := strings.Split(line[ITMSL:], ", ")
				for _, v := range items {
					num, err := strconv.Atoi(v)
					if err != nil {
						panic(err)
					}
					monkey.items = append(monkey.items, num)
				}
			case 1:
				op := line[OPRSL:]
				switch op {
				case "old * old":
					monkey.operator = SQUARE
					monkey.operationNum = 2 // unused
				default:
					fields := strings.Fields(op)
					switch fields[1] {
					case "+":
						monkey.operator = ADD
					case "-":
						monkey.operator = SUBTRACT
					case "*":
						monkey.operator = MULTIPLY
					case "/":
						monkey.operator = DIVIDE
					}
					num, err := strconv.Atoi(fields[2])
					if err != nil {
						panic(err)
					}
					monkey.operationNum = num
				}
			case 2:
				num, err := strconv.Atoi(line[TESTL:])
				if err != nil {
					panic(err)
				}
				monkey.testDivisor = num
			case 3:
				num, err := strconv.Atoi(line[TRUEL:])
				if err != nil {
					panic(err)
				}
				monkey.next[0] = num
			case 4:
				num, err := strconv.Atoi(line[FALSL:])
				if err != nil {
					panic(err)
				}
				monkey.next[1] = num
			}

		}
		monkeys[i] = monkey
	}

	return monkeys
}

func changeWorryLevelSafely(op int, a int, b int) (int, error) {
	if a > math.MaxInt64-b || a > math.MaxInt64/b {
		return 0, ErrOverflow
	}

	switch op {
	case ADD:
		a += b
	case SUBTRACT:
		a -= b
	case MULTIPLY:
		a *= b
	case DIVIDE:
		a /= b
	case SQUARE:
		a *= a
	}

	return a, nil
}

func playMonkeyInTheMiddle(monkeys []Monkey, rounds int, divisor int, modDiv bool) []int {
	counter := make([]int, len(monkeys))
	fmt.Println()

	for i := 1; i <= rounds; i++ {

		if rounds <= 20 {
			fmt.Println("Round:", i)
		} else if rounds <= 200 {
			if i%10 == 0 {
				fmt.Println("Round:", i)
			}
		} else {
			if i%500 == 0 {
				fmt.Println("Round:", i)
			}
		}

		// play round
		for idx, monkey := range monkeys {
			// monkey plays/inspects
			for _, worryLevel := range monkey.items {
				monkeys[idx].inspected++

				worryLevel, err := changeWorryLevelSafely(
					monkey.operator, worryLevel, monkey.operationNum,
				)
				if err != nil {
					panic(err)
				}

				// monkey gets disinterested, rounds down automatically
				if modDiv {
					worryLevel %= divisor
				} else {
					worryLevel /= divisor
				}

				// test where to throw
				isDivisible := worryLevel%monkey.testDivisor == 0

				// throw item
				next := monkey.next[0]
				if !isDivisible {
					next = monkey.next[1]
				}
				monkeys[next].items = append(monkeys[next].items, worryLevel)
			}

			// clean item slate
			monkeys[idx].items = []int{}
		}
	}

	// add inspected items per monkey to counter
	for i, monkey := range monkeys {
		counter[i] = monkey.inspected
	}

	return counter
}

func calculateMonkeyBusinessLevel(counter []int) int {
	businessLevel := 1
	sort.Ints(counter)

	for _, v := range counter[len(counter)-2:] {
		businessLevel *= v
	}
	return businessLevel
}

func main() {
	input_fp := "day_11/input.txt"
	fileData := parseFile(input_fp)

	monkeys := parseInput(fileData)

	fmt.Println("\nStarting state:")
	for i, m := range monkeys {
		fmt.Printf("Monkey %d |%8d|%v\n", i, m.inspected, m.items)
	}

	counter := playMonkeyInTheMiddle(monkeys, 20, 3, false)

	val1 := calculateMonkeyBusinessLevel(counter)
	fmt.Println("\n[1] Sum of signal strenghts:\t", val1)

	// TODO: pass by value for calculations to omit having to read in again
	monkeys = parseInput(fileData)

	divisor := 1
	for _, monkey := range monkeys {
		divisor *= monkey.testDivisor
	}
	counter = playMonkeyInTheMiddle(monkeys, 10000, divisor, true)

	val2 := calculateMonkeyBusinessLevel(counter)
	fmt.Println("\n[2] Sum of signal strenghts:\t", val2)

	fmt.Println("\nEnd state:")
	for i, m := range monkeys {
		fmt.Printf("Monkey %d |%8d|%v\n", i, m.inspected, m.items)
	}
}
