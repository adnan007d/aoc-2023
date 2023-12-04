package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Can do better but meh or I am doing row by row so cache will give free reads?
var neighbours = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
}

func isValid(inputs []string, row int, start int, end int) bool {
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

func extractNumber(input string, startIndex int) (int, error) {

	end := startIndex
	for ; end < len(input) && unicode.IsDigit(rune(input[end])); end++ {
	}

	start := startIndex
	for ; start >= 0 && unicode.IsDigit(rune(rune(input[start]))); start-- {
	}

	return strconv.Atoi(input[start+1 : end])
}

func day3Part2(inputs []string) int {
	var sum = 0

	// Find Star
	// Find coords of neighbouring numbers construct the number
	// Remove duplicate numbers (My method won't work if both the gear number are same as I am removing duplicates without verifying the position)
	// Do the math
	for row, input := range inputs {

		for col := 0; col < len(input); col++ {
			// Find Star
			if input[col] == '*' {
				coords := [][]int{}
				// Traversing neighbours
				for _, neighbour := range neighbours {
					r := row + neighbour[0]
					c := col + neighbour[1]
					// Finding neighbouring numbers
					if r >= 0 && r < len(inputs) && c >= 0 && c < len(input) && unicode.IsDigit(rune(inputs[r][c])) {
						coords = append(coords, []int{r, c})
					}

				}

				// Construct and remove duplicates
				var set = map[int]bool{}
				for _, coord := range coords {
					num, err := extractNumber(inputs[coord[0]], coord[1])
					if err != nil {
						panic(err)
					}
					set[num] = true
				}

				// Do the math
				if len(set) > 1 {
					var result = 1
					for num := range set {
						result *= num
					}
					sum += result
				}
			}
		}
	}
	return sum
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
