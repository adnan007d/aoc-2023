package day1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func day1Part1() int {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sum := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		temp := scanner.Text()
		var result int
		for _, c := range temp {
			if unicode.IsDigit(c) {
				result = 10 * int(c-'0')
				break
			}
		}

		for i := len(temp) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(temp[i])) {
				result += int(temp[i] - '0')
				break
			}
		}

		sum += result
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

  return sum
}

func day1Part2() int {

	lookup := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	lookupNumbersPosition := make(map[string]int, 9)

	f, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var result int
		temp := scanner.Text()

		minI := len(temp)
		minDigit := ""
		maxI := -1
		maxDigit := ""

		for strDigit := range lookup {
			i := strings.Index(temp, strDigit)
			j := strings.LastIndex(temp, strDigit)

			lookupNumbersPosition[strDigit] = i

			if i <= minI && i != -1 {
				minDigit = strDigit
				minI = i
			}

			if j >= maxI && j != -1 {
				maxDigit = strDigit
				maxI = j
			}
		}

		for i, c := range temp {
			if unicode.IsDigit(c) {
				result = 10 * int(c-'0')
				break
			} else if i >= minI {
				result = 10 * lookup[minDigit]
				break
			}
		}

		for i := len(temp) - 1; i >= 0; i-- {
			if maxI >= i {
				result += lookup[maxDigit]
				break
			} else if unicode.IsDigit(rune(temp[i])) {
				result += int(temp[i] - '0')
				break
			}

		}

		sum += result
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

  return sum
}

func Day1() {
  fmt.Println("Day 1 Part 1", day1Part1())
  fmt.Println("Day 1 Part 2", day1Part2())
}
