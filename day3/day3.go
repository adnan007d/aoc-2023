package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func isValid(inputs []string, row int, start int, end int) bool {
	// Can do better but meh or I am doing row by row so cache will give free reads?
	neighbours := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
	}

	colLen := len(inputs[row])
	rowLen := len(inputs)
	for ; start < end; start++ {
		for _, neighbour := range neighbours {
			r := row + neighbour[0]
			c := start + neighbour[1]
			if r >= 0 && r < rowLen && c >= 0 && c < colLen {
				if !unicode.IsDigit(rune(inputs[r][c])) && inputs[r][c] != '.' {
					return true
				}
			}
		}
	}

	return false

}

func day3Part1(inputs []string) int {
	var sum = 0
	for row, input := range inputs {

		for col := 0; col < len(input); col++ {

			// parsing digit
			start := col
			for col < len(input) && unicode.IsDigit(rune(input[col])) {
				col++
			}

			// Digit found
			if start != col {
				if isValid(inputs, row, start, col) {
          num, err := strconv.Atoi(input[start:col])
          if err != nil {
            panic(err)
          }
          sum += num
				}

			}

		}
	}
	return sum
}

func day3Part2(inputs []string) int {
	return 0
}

func Day3() {
	f, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	inputs := []string{}
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	fmt.Println("Day 3 Part 1", day3Part1(inputs))
	fmt.Println("Day 3 Part 2", day3Part2(inputs))
}
