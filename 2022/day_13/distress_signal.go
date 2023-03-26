package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"unicode"
)

type Packet []interface{}

func timeTrack(start time.Time, action string) {
	elapsed := time.Since(start)
	fmt.Printf("\n%s took %s\n", action, elapsed)
}

func parseRawString(raw string) Packet {
	ans := Packet{}
	json.Unmarshal([]byte(raw), &ans)
	return ans
}

func parseFile(file *os.File) (ans [][2]Packet) {
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	data := strings.TrimRightFunc(string(bytes), unicode.IsSpace)
	for _, packetPairs := range strings.Split(data, "\n\n") {
		pairs := strings.Split(packetPairs, "\n")
		ans = append(ans, [2]Packet{
			parseRawString(pairs[0]),
			parseRawString(pairs[1]),
		})
	}
	return ans
}

func isOrderedVerbose(left, right Packet, level int) (bool, error) {
	for j := 0; j < level; j++ {
		fmt.Print("  ")
	}
	fmt.Println("- Compare", left, "vs", right)
	for i := 0; i < len(left); i++ {

		// check if amount of list items is the same, left should have less or equal
		if i > len(right)-1 {
			for j := 0; j <= level; j++ {
				fmt.Print("  ")
			}
			fmt.Println("- Right side ran out of items, so inputs are NOT in the right order")
			return false, nil
		}

		// float
		leftNum, isLeftNum := left[i].(float64)
		rightNum, isRightNum := right[i].(float64)

		if isLeftNum && isRightNum {
			for j := 0; j <= level; j++ {
				fmt.Print("  ")
			}
			fmt.Println("- Compare", leftNum, "vs", rightNum)
			if leftNum == rightNum {
				if i == len(left)-1 {
					if len(right) > len(left) {
						for j := 0; j <= level; j++ {
							fmt.Print("  ")
						}
						fmt.Println("- Left side ran out of items, so inputs are in the right order")
						return true, nil
					}
					// last item is equal, continue one level up
					return false, errors.New("List is equal")
				}
				// on equality test next item
				continue
			} else {
				for j := 0; j <= level+1; j++ {
					fmt.Print("  ")
				}
				if leftNum < rightNum {
					fmt.Println("- Left side is smaller, so inputs are in the right order")
				} else {
					fmt.Println("- Right side is smaller, so inputs are NOT in the right order")
				}
				return leftNum < rightNum, nil
			}
		} else if isLeftNum || isRightNum {
			for j := 0; j < level+1; j++ {
				fmt.Print("  ")
			}
			if isLeftNum {
				left[i] = []interface{}{leftNum}
				fmt.Println("- Mixed types; convert left to", left[i], "and retry comparison")
			} else if isRightNum {
				right[i] = []interface{}{rightNum}
				fmt.Println("- Mixed types; convert right to", right[i], "and retry comparison")
			} else {
				panic(fmt.Sprintf("expected one num %T:%v, %T:%v", left[i], left[i], right[i], right[i]))
			}
		}

		// list
		leftList, isLeftList := left[i].([]interface{})
		rightList, isRightList := right[i].([]interface{})

		if isLeftList && isRightList {
			if len(leftList) == 0 && len(rightList) == 0 {
				continue
			}
			if ans, err := isOrderedVerbose(leftList, rightList, level+1); err != nil {
				// lists are identical, continue with next item
				continue
			} else {
				return ans, nil
			}
		} else {
			panic(fmt.Sprintf("expected two lists %T:%v, %T:%v", left[i], left[i], right[i], right[i]))
		}

	}
	for j := 0; j <= level; j++ {
		fmt.Print("  ")
	}
	fmt.Println("- Left side ran out of items default, so inputs are in the right order")
	return true, nil
}

func isOrdered(left, right Packet) (bool, error) {
	for i := 0; i < len(left); i++ {

		// check if amount of list items is the same, left should have less or equal
		if i > len(right)-1 {
			return false, nil
		}

		// float
		leftNum, isLeftNum := left[i].(float64)
		rightNum, isRightNum := right[i].(float64)

		if isLeftNum && isRightNum {
			if leftNum == rightNum {
				if i == len(left)-1 {
					if len(right) > len(left) {
						return true, nil
					}
					// last item is equal, continue one level up
					return false, errors.New("List is equal")
				}
				// on equality test next item
				continue
			} else {
				return leftNum < rightNum, nil
			}
		} else if isLeftNum || isRightNum {
			if isLeftNum {
				left[i] = []interface{}{leftNum}
			} else if isRightNum {
				right[i] = []interface{}{rightNum}
			} else {
				panic(fmt.Sprintf("expected one num %T:%v, %T:%v", left[i], left[i], right[i], right[i]))
			}
		}

		// list
		leftList, isLeftList := left[i].([]interface{})
		rightList, isRightList := right[i].([]interface{})

		if isLeftList && isRightList {
			if len(leftList) == 0 && len(rightList) == 0 {
				continue
			}
			if ans, err := isOrdered(leftList, rightList); err != nil {
				// lists are identical, continue with next item
				continue
			} else {
				return ans, nil
			}
		} else {
			panic(fmt.Sprintf("expected two lists %T:%v, %T:%v", left[i], left[i], right[i], right[i]))
		}

	}
	return true, nil
}

func sumSlice(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}

func checkOrder(pairs [][2]Packet, verbose bool) int {
	defer timeTrack(time.Now(), "Verifying packets")
	var orderedPairs []int
	for idx, pair := range pairs {
		if !verbose {
			if res, err := isOrdered(pair[0], pair[1]); err == nil && res {
				orderedPairs = append(orderedPairs, idx+1)
			}
		} else {
			fmt.Println("=== Pair", idx+1, "===")
			if res, err := isOrderedVerbose(pair[0], pair[1], 0); err == nil && res {
				orderedPairs = append(orderedPairs, idx+1)
			}
			fmt.Println()
		}
	}
	if verbose {
		fmt.Println(orderedPairs)
	}
	return sumSlice(orderedPairs)
}

func findKey(pairs [][2]Packet) uint16 {
	defer timeTrack(time.Now(), "Searching decoder key")
	// packets start at index 1, add +1 for [[6]] to account for [[2]]
	dividerLoc := []uint16{1, 2}
	for _, pair := range pairs {
		for _, line := range pair {
			// put divider as left, since [[6]] has a problem with [[[[6]],2,[[],2],4],[4,9],[5,4,0,[4],[10]],[],[[8,7,3,10],9,10]]
			if res, err := isOrdered([]interface{}{float64(2)}, line); err != nil {
				panic(fmt.Sprintf("Packets are equal: %v and %v", line, "[[2]]"))
			} else if !res {
				dividerLoc[0]++
			}
			if res, err := isOrdered([]interface{}{float64(6)}, line); err != nil {
				panic(fmt.Sprintf("Packets are equal: %v and %v", line, "[[6]]"))
			} else if !res {
				dividerLoc[1]++
			}
		}
	}
	return dividerLoc[0] * dividerLoc[1]
}

func main() {
	arg := os.Stdin
	data := parseFile(arg)

	nOrderedPairs := checkOrder(data, false)
	fmt.Println("Packets in correct order:\t\t", nOrderedPairs)

	arg.Seek(0, 0)
	data = parseFile(arg)

	decoderKey := findKey(data)
	fmt.Println("Decoder key for distress signal:\t", decoderKey)
}
