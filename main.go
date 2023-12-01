package main

import (
	"log"
	"os"

	aoc "github.com/adnan007d/aoc-2023/day1"
)

func main() {
  if len(os.Args) !=  2 {
    log.Fatal("./aoc-2023 <day>")
  }

  if os.Args[1] == "1" {
      aoc.Day1() 
  }
}
