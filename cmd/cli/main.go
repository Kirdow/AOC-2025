package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc/internal/day1"
	"aoc/internal/day2"
	"aoc/internal/day3"
	"aoc/internal/day4"
	"aoc/internal/day5"
	"aoc/internal/day6"
	"aoc/internal/day7"
	"aoc/internal/day8"
	"aoc/internal/day9"
	"aoc/internal/day10"
	"aoc/internal/day11"
	"aoc/internal/day12"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var day int
	for {
		fmt.Print("Enter day (1-12): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err == nil && val >= 1 && val <= 12 {
			day = val
			break
		}
		fmt.Println("Invalid input, try again.")
	}

	inputPath := fmt.Sprintf("inputs/day%d.txt", day)
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Input not found")
		os.Exit(1)
	}
	input := string(data)

	var result string
	switch day {
	case 1: result, err = day1.Main(input)
	case 2: result, err = day2.Main(input)
	case 3: result, err = day3.Main(input)
	case 4: result, err = day4.Main(input)
	case 5: result, err = day5.Main(input)
	case 6: result, err = day6.Main(input)
	case 7: result, err = day7.Main(input)
	case 8: result, err = day8.Main(input)
	case 9: result, err = day9.Main(input)
	case 10: result, err = day10.Main(input)
	case 11: result, err = day11.Main(input)
	case 12: result, err = day12.Main(input)
	default:
		fmt.Printf("Failure to run AOC 2025. Day %d not found\n", day)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Failure to run AOC 2025 Day %d: %v\n", day, err)
		os.Exit(1)
	}

	fmt.Printf("AOC 2025 Day %d Result:\n%s\n", day, result)
}
