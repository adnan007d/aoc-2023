package day4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func day4Part1(inputs []string) int {

	var sum = 0

	for _, input := range inputs {
		// can use split and shit but it will be too easy
		startOfPoints := strings.Index(input, ":")

		setIntersection := map[string]int{}

		start := startOfPoints + 1

		for i := start; input[i] != '|'; i++ {
			if unicode.IsDigit(rune(input[i])) {
				start = i
				for unicode.IsDigit(rune(input[i])) {
					i++
				}
				setIntersection[input[start:i]] += 1
				start = i + 1
			}
		}

		start += 2

		for i := start; i < len(input); i++ {
			if unicode.IsDigit(rune(input[i])) {
				start = i
				for i<len(input) && unicode.IsDigit(rune(input[i])) {
					i++
				}
				setIntersection[input[start:i]] += 1
				start = i + 1
			}
		}

		var n = 0
		// can put this in the above loop but meh
		for _, points := range setIntersection {
			if points == 2 {
				n++
			}
		}
		//
		sum += int(math.Pow(2, float64(n-1)))

	}
	return sum
}

func day4Part2(inputs []string) int {
	return 0
}

func Day4() {
	f, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	inputs := []string{}
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	fmt.Println("Day 4 Part 1", day4Part1(inputs))
	fmt.Println("Day 4 Part 2", day4Part2(inputs))
}
