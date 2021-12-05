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

func ParseDay4(lines []string) (numbers []int, boards [][25]int) {
	for _, ns := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(ns)
		numbers = append(numbers, n)
	}
	for i := 2; i < len(lines); i += 6 {
		var board [25]int
		for j := 0; j < 5; j++ {
			for c, ns := range strings.Fields(lines[i+j]) {
				n, _ := strconv.Atoi(ns)
				board[5*j+c] = n
			}
		}
		boards = append(boards, board)
	}
	return
}

func MarkNumberOnBoard(b *[25]int, n int) {
	for i := range b {
		if b[i] == n {
			b[i] = -1
		}
	}
}

func CheckBingo(b *[25]int) bool {
	for i := 0; i < 5; i++ {
		row, col := -1, -1
		for d := 0; d < 5; d++ {
			row &= b[5*i+d]
			col &= b[5*d+i]
		}
		if row < 0 || col < 0 {
			return true
		}
	}
	return false
}

func BoardScore(b *[25]int, n int) (result int) {
	for _, cell := range b {
		if cell != -1 {
			result += cell
		}
	}
	result *= n
	return
}

func day04a() int {
	lines := readlines("data/day04.txt")
	numbers, boards := ParseDay4(lines)

	for _, n := range numbers {
		for i := range boards {
			MarkNumberOnBoard(&boards[i], n)
		}
		for _, b := range boards {
			if CheckBingo(&b) {
				return BoardScore(&b, n)
			}
		}
	}
	return 0
}

func day04b() int {
	lines := readlines("data/day04.txt")
	numbers, boards := ParseDay4(lines)

	completed := make(map[int]bool)

	for _, n := range numbers {
		for i := range boards {
			MarkNumberOnBoard(&boards[i], n)
		}
		for board_index, b := range boards {
			if CheckBingo(&b) {
				completed[board_index] = true
				if len(completed) == len(boards) {
					return BoardScore(&b, n)
				}
			}
		}
	}
	return 0
}

type Point struct {
	x, y int
}

func ParseDay5(lines []string) (result [][2]Point) {
	for _, line := range lines {
		var p0, p1 Point
		fmt.Sscanf(line, "%d,%d -> %d,%d", &p0.x, &p0.y, &p1.x, &p1.y)
		result = append(result, [2]Point{p0, p1})
	}
	return
}

func signum(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	} else {
		return 0
	}
}

func MarkSegment(s [2]Point, vents map[Point]int) {
	dx, dy := signum(s[1].x-s[0].x), signum(s[1].y-s[0].y)
	for x, y := s[0].x, s[0].y; x != s[1].x || y != s[1].y; x, y = x+dx, y+dy {
		vents[Point{x, y}] += 1
	}
	vents[s[1]] += 1
}

func ScoreVents(vents map[Point]int) (result int) {
	for _, b := range vents {
		if b > 1 {
			result += 1
		}
	}
	return
}

func day05a() int {
	lines := readlines("data/day05.txt")
	segments := ParseDay5(lines)
	vents := make(map[Point]int)
	for _, segment := range segments {
		if segment[0].x == segment[1].x || segment[0].y == segment[1].y {
			MarkSegment(segment, vents)
		}
	}
	return ScoreVents(vents)
}

func day05b() int {
	lines := readlines("data/day05.txt")
	segments := ParseDay5(lines)
	vents := make(map[Point]int)
	for _, segment := range segments {
		MarkSegment(segment, vents)
	}
	return ScoreVents(vents)
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
	fmt.Println("day05a:", day05a())
	fmt.Println("day05b:", day05b())
}
