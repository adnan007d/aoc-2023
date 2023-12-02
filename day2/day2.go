package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func day2Part1() int {
	f, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	var sum = 0
	cubesLimit := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input := scanner.Text()
		colonIndex := strings.Index(input, ":")

		gameId, err := strconv.Atoi(input[5:colonIndex])
		if err != nil {
			log.Fatal("Error while extracting gameID")
		}

		cubeMap := map[string]int{}

		// +1 to consume the space
		start := colonIndex + 1

		var cubeCount = -1
		var cube = ""
		var isValid = true
		// Trying to do it in o(n)

		for i := colonIndex + 1; i < len(input); i++ {

			// These are the delimeters space, comma, semicolon and end of string
			if input[i] == ' ' || input[i] == ',' || input[i] == ';' || i == len(input)-1 {

				// If the previous character before the delimeter is digit then its a whole digit
				if unicode.IsDigit(rune(input[i-1])) {
					_cubeCount, err := strconv.Atoi(input[start:i])
					if err != nil {
						panic(err)
					}
					cubeCount = _cubeCount
				} else {
					cube = input[start:i]
				}

				// To consume the last draw
				if i == len(input)-1 {
					cube = input[start:]
				}

				// Changing the start value to current delimeter + 1
				start = i + 1

				// As we are doing it in one rotation the cube count will be calculated first
				// and then the cube itself so just a so that both exists
				// reset the values to default once done
				if cubeCount != -1 && cube != "" {
					cubeMap[cube] += cubeCount
					cubeCount = -1
					cube = ""
				}
			}

			// iF one set is complete check whether the draws are in limit or not
			// if not in limit break out of the look
			if input[i] == ';' || i == len(input)-1 {
				for cube := range cubeMap {
					if cubeMap[cube] > cubesLimit[cube] {
						isValid = false
						break
					}
				}
				// Clearing for next set
				clear(cubeMap)
			}

			// If its not valid for a single set the whole game gets invalid
			if !isValid {
				break
			}

		}

		if isValid {
			sum += gameId
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}

	return sum
}

func day2Part2() int {
	f, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	var sum = 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return sum
}

func Day2() {
	fmt.Println("Day 2 Part 1", day2Part1())
	fmt.Println("Day 2 Part 2", day2Part2())
}
