package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readints_radix(fname string, radix int) (result []int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.ParseInt(scanner.Text(), radix, 32)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, int(n))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func readints(fname string) []int {
	return readints_radix(fname, 10)
}

func readlines(fname string) (result []string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func day01a() (result int) {
	lines := readints("data/day01.txt")
	for i := 0; i < len(lines)-1; i += 1 {
		if lines[i] < lines[i+1] {
			result += 1
		}
	}
	return
}

func day01b() (result int) {
	lines := readints("data/day01.txt")
	for i := 0; i < len(lines)-3; i += 1 {
		if lines[i] < lines[i+3] {
			result += 1
		}
	}
	return
}

func day02a() int {
	x, y := 0, 0
	for _, line := range readlines("data/day02.txt") {
		up, down, forward := -1, -1, -1
		if fmt.Sscanf(line, "up %d", &up); up >= 0 {
			y -= up
		} else if fmt.Sscanf(line, "down %d", &down); down >= 0 {
			y += down
		} else if fmt.Sscanf(line, "forward %d", &forward); forward >= 0 {
			x += forward
		} else {
			log.Fatalf("can't match command %s", line)
		}
	}
	return x * y
}

func day02b() int {
	x, y, aim := 0, 0, 0
	for _, line := range readlines("data/day02.txt") {
		up, down, forward := -1, -1, -1
		if fmt.Sscanf(line, "up %d", &up); up >= 0 {
			aim -= up
		} else if fmt.Sscanf(line, "down %d", &down); down >= 0 {
			aim += down
		} else if fmt.Sscanf(line, "forward %d", &forward); forward >= 0 {
			x += forward
			y += aim * forward

		} else {
			log.Fatalf("can't match command %s", line)
		}
	}
	return x * y
}

func day03a() int {
	lines := readlines("data/day03.txt")
	var counts [12]int
	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				counts[i] += 1
			}
		}
	}
	var gamma, epsilon int
	for _, c := range counts {
		gamma, epsilon = 2*gamma, 2*epsilon
		if c > (len(lines) - c) {
			gamma += 1
		} else {
			epsilon += 1
		}
	}
	return gamma * epsilon
}

func bit_criteria(numbers []string, bitpos int, oxygen bool) (result []string) {
	var ones, zeroes []string
	for _, line := range numbers {
		if line[bitpos] == '1' {
			ones = append(ones, line)
		} else {
			zeroes = append(zeroes, line)
		}
	}
	if len(zeroes) <= len(ones) != oxygen {
		return ones
	} else {
		return zeroes
	}
}

func day03b() uint64 {
	lines := readlines("data/day03.txt")

	oxy_lines := lines
	for i := 0; len(oxy_lines) > 1; i = i + 1 {
		oxy_lines = bit_criteria(oxy_lines, i, true)
	}
	oxy, _ := strconv.ParseUint(oxy_lines[0], 2, 64)

	co2_lines := lines
	for i := 0; len(co2_lines) > 1; i = i + 1 {
		co2_lines = bit_criteria(co2_lines, i, false)
	}
	co2, _ := strconv.ParseUint(co2_lines[0], 2, 64)

	return oxy * co2
}

func ParseDay4(lines []string) (numbers []int, boards [][5][5]int) {
	for _, ns := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(ns)
		numbers = append(numbers, n)
	}
	for i := 2; i < len(lines); i += 6 {
		var board [5][5]int
		for j := 0; j < 5; j++ {
			for c, ns := range strings.Fields(lines[i+j]) {
				n, _ := strconv.Atoi(ns)
				board[j][c] = n
			}
		}
		boards = append(boards, board)
	}
	return
}

func day04a() int {
	lines := readlines("data/day04.txt")

	numbers, boards := ParseDay4(lines)

	bingo := [5]int{-1, -1, -1, -1, -1}

	for _, n := range numbers {
		for a, b := range boards {
			for x, r := range b {
				for i, cell := range r {
					if cell == n {
						boards[a][x][i] = -1
					}
				}
			}
		}
		for _, b := range boards {
			for i := 0; i < 5; i += 1 {
				if b[i] == bingo || [5]int{b[0][i], b[1][i], b[2][i], b[3][i], b[4][i]} == bingo {
					sum := 0
					for _, row := range b {
						for _, cell := range row {
							if cell != -1 {
								sum += cell
							}
						}
					}
					return sum * n
				}
			}
		}
	}
	return 0
}

func day04b() int {
	lines := readlines("data/day04.txt")
	numbers, boards := ParseDay4(lines)

	marked := [5]int{-1, -1, -1, -1, -1}

	completed := make(map[int]bool)

	for _, n := range numbers {
		for board_index, board := range boards {
			for r, row := range board {
				for c, cell := range row {
					if cell == n {
						boards[board_index][r][c] = -1
					}
				}
			}
		}
		for board_index, b := range boards {
			if completed[board_index] {
				continue
			}
			for i := 0; i < 5; i += 1 {
				if b[i] == marked || [5]int{b[0][i], b[1][i], b[2][i], b[3][i], b[4][i]} == marked {
					completed[board_index] = true
					if len(completed) == len(boards) {
						sum := 0
						for _, row := range b {
							for _, cell := range row {
								if cell != -1 {
									sum += cell
								}
							}
						}
						return sum * n
					}
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println("day01a:", day01a())
	fmt.Println("day01b:", day01b())
	fmt.Println("day02a:", day02a())
	fmt.Println("day02b:", day02b())
	fmt.Println("day03a:", day03a())
	fmt.Println("day03b:", day03b())
	fmt.Println("day04a:", day04a())
	fmt.Println("day04b:", day04b())
}
