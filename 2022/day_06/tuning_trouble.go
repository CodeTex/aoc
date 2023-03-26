package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

func fifoQueue(q []rune, l int, e rune) []rune {
	q = append(q, e)
	if len(q) > l {
		return q[1:]
	}
	return q
}

func isUniqueSlice[T rune | string](s []T) bool {
	m := make(map[T]bool, len(s))

	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = true
			continue
		}
		return false
	}
	return true
}

func findMarkerPosition(path string, mLen int) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	var queue []rune
	pos := 1
	for {
		if c, _, err := scanner.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		} else {
			if unicode.IsLetter(c) {
				queue = fifoQueue(queue, mLen, c)
				if len(queue) >= mLen && isUniqueSlice(queue) {
					return pos, nil
				}
			}
		}
		pos++
	}
	return pos, errors.New("Marker not found")
}

func main() {
	input_fp := "day_06/input.txt"

	val1, err := findMarkerPosition(input_fp, 4)
	if err != nil {
		panic("No start-of-packet marker found")
	}
	fmt.Println("Marker position:\t", val1)

	val2, err := findMarkerPosition(input_fp, 14)
	if err != nil {
		panic("No start-of-message marker found")
	}
	fmt.Println("Message position:\t", val2)
}
