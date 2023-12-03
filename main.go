package main

import (
	"log"
	"os"

	"github.com/adnan007d/aoc-2023/day1"
	"github.com/adnan007d/aoc-2023/day2"
	"github.com/adnan007d/aoc-2023/day3"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("./aoc-2023 <day>")
	}

	day := os.Args[1]

	switch day {
	case "1":
		day1.Day1()
	case "2":
		day2.Day2()
	case "3":
		day3.Day3()
	}
}
