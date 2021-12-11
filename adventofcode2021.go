package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"sort"
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

func ParseDay6(line []string) map[int]int {
	result := make(map[int]int)
	for _, ns := range strings.Split(line[0], ",") {
		n, _ := strconv.Atoi(ns)
		result[n] += 1
	}
	return result
}

func EvolveLanternfish(counts map[int]int) map[int]int {
	new_counts := make(map[int]int)
	for age, count := range counts {
		if age > 0 {
			new_counts[age-1] += count
		} else {
			new_counts[6] += count
			new_counts[8] += count
		}
	}
	counts = new_counts
	return new_counts
}

func CountLanternfish(counts map[int]int) (result int) {
	for _, count := range counts {
		result += count
	}
	return
}

func day06a() int {
	counts := ParseDay6(readlines("data/day06.txt"))
	for i := 0; i < 80; i++ {
		counts = EvolveLanternfish(counts)
	}
	return CountLanternfish(counts)
}

func day06b() int {
	counts := ParseDay6(readlines("data/day06.txt"))
	for i := 0; i < 256; i++ {
		counts = EvolveLanternfish(counts)
	}
	return CountLanternfish(counts)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ParseDay7(fname string) (result []int) {
	lines := readlines("data/day07.txt")
	sum := 0
	for _, ns := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(ns)
		result = append(result, n)
		sum += n
	}
	return
}

func day07a() int {
	positions := ParseDay7("data/day07.txt")
	sort.Ints(positions)
	median := positions[len(positions)/2]
	diff := 0
	for _, p := range positions {
		diff += Abs(p - median)
	}
	return diff
}

func day07b() int {
	positions := ParseDay7("data/day07.txt")
	sum := 0
	for _, p := range positions {
		sum += p
	}
	mean := sum / len(positions)

	cost := 0
	for _, p := range positions {
		d := Abs(p - mean)
		cost += (d * (d + 1)) / 2
	}
	return cost
}

func ParseDay8(fname string) (result [][2][]byte) {
	to_bitmasks := func(ss []string) (result []byte) {
		for _, s := range ss {
			mask := byte(0)
			for _, r := range s {
				mask |= 1 << (int(r) - int('a'))
			}
			result = append(result, mask)
		}
		return
	}
	for _, line := range readlines(fname) {
		var sample [2][]byte
		asdf := strings.SplitN(line, "|", 2)
		sample[0] = to_bitmasks(strings.Fields(strings.TrimSpace(asdf[0])))
		sample[1] = to_bitmasks(strings.Fields(strings.TrimSpace(asdf[1])))
		result = append(result, sample)
	}
	return
}

func day08a() int {
	samples := ParseDay8("data/day08.txt")
	count := 0
	for _, ss := range samples {
		for _, s := range ss[1] {
			if x := bits.OnesCount8(s); x == 2 || x == 4 || x == 3 || x == 7 {
				count += 1
			}
		}
	}
	return count
}

func Decode(observations []byte) map[byte]int {
	var s1, s4, s7, s8 byte
	for _, ob := range observations {
		switch bits.OnesCount8(ob) {
		case 2:
			s1 = ob
		case 4:
			s4 = ob
		case 3:
			s7 = ob
		case 7:
			s8 = ob
		}
	}
	result := map[byte]int{s1: 1, s4: 4, s7: 7, s8: 8}
	for _, ob := range observations {
		switch bits.OnesCount8(ob ^ s1) {
		case 3: // 3^1=3
			result[ob] = 3
		case 6: // 6^1=6
			result[ob] = 6
		case 4: // 0^1=4, 9^1=4
			if bits.OnesCount8(ob^s4) == 4 {
				result[ob] = 0
			} else {
				result[ob] = 9
			}
		case 5: // 2^1=5 5^1=5 8^1=5
			switch bits.OnesCount8(ob & s4) { // 2&4=2 5&4=3 8&4=4
			case 2:
				result[ob] = 2
			case 3:
				result[ob] = 5
			}
		}
	}
	return result
}

func day08b() int {
	samples := ParseDay8("data/day08.txt")
	sum := 0
	for _, x := range samples {
		key := Decode(x[0])
		val := 0
		for _, d := range x[1] {
			val = 10*val + key[d]
		}
		sum += val
	}
	return sum
}

func ParseDay9(fname string) (result [][]int) {
	for _, line := range readlines(fname) {
		var asdf []int
		for _, c := range line {
			asdf = append(asdf, int(c)-int('0'))
		}
		result = append(result, asdf)
	}
	return
}

func Lowpoints(heights [][]int) (result []Point) {
	dx, dy := len(heights[0]), len(heights)
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			me := heights[y][x]
			if (x <= 0 || me < heights[y][x-1]) && (x == dx-1 || me < heights[y][x+1]) && (y == 0 || me < heights[y-1][x]) && (y == dy-1 || me < heights[y+1][x]) {
				result = append(result, Point{x, y})
			}
		}
	}
	return
}

func Floodfill(heights [][]int, start Point) map[Point]bool {
	dx, dy := len(heights[0]), len(heights)

	seen := make(map[Point]bool)
	for frontier := []Point{start}; len(frontier) > 0; {
		p := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]
		if _, ok := seen[p]; ok {
			continue
		}
		seen[p] = true
		x, y := p.x, p.y
		if x > 0 && heights[y][x-1] != 9 {
			frontier = append(frontier, Point{x - 1, y})
		}
		if x < dx-1 && heights[y][x+1] != 9 {
			frontier = append(frontier, Point{x + 1, y})
		}
		if y > 0 && heights[y-1][x] != 9 {
			frontier = append(frontier, Point{x, y - 1})
		}
		if y < dy-1 && heights[y+1][x] != 9 {
			frontier = append(frontier, Point{x, y + 1})
		}
	}
	return seen
}

func day09a() int {
	heights := ParseDay9("data/day09.txt")
	result := 0
	for _, p := range Lowpoints(heights) {
		result += heights[p.y][p.x] + 1
	}
	return result
}

func day09b() int {
	heights := ParseDay9("data/day09.txt")
	var basin_sizes []int
	for _, p := range Lowpoints(heights) {
		basin_sizes = append(basin_sizes, len(Floodfill(heights, p)))
	}
	sort.Ints(basin_sizes)
	result := 1
	for _, p := range basin_sizes[len(basin_sizes)-3:] {
		result *= p
	}
	return result
}

func MatchParens(x string) (rune, []rune) {
	a := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	var stack []rune
	for _, c := range x {
		if _, ok := a[c]; ok {
			stack = append(stack, c)
		} else {
			if a[stack[len(stack)-1]] == c {
				stack = stack[:len(stack)-1]
			} else {
				return c, nil
			}
		}
	}
	return 'k', stack
}

func day10a() int {
	heights := readlines("data/day10.txt")
	scores := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137, 'k': 0}
	sum := 0
	for _, line := range heights {
		c, _ := MatchParens(line)
		sum += scores[c]
	}
	return sum
}

func day10b() int {
	heights := readlines("data/day10.txt")
	scoring := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	var scores []int
	for _, line := range heights {
		if c, s := MatchParens(line); c == 'k' {
			score := 0
			for i := len(s) - 1; i >= 0; i-- {
				score = score*5 + scoring[s[i]]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func ParseDay11(fname string) (result [10][10]byte) {
	lines := readlines(fname)
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			result[y][x] = lines[y][x] - '0'
		}
	}
	return
}

func VisitNeighbours(p Point, d int, visitor func(Point)) {
	for y := p.y - 1; y <= p.y+1; y += 1 {
		for x := p.x - 1; x <= p.x+1; x += 1 {
			if x >= 0 && y >= 0 && x < d && y < d && (x != p.x || y != p.y) {
				visitor(Point{x, y})
			}
		}
	}
}

func OctoStep(octopi *[10][10]byte) int {
	blinked := make(map[Point]bool)
	var pending []Point

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			octopi[y][x] += 1
			if octopi[y][x] > 9 {
				blinked[Point{x, y}] = true
				pending = append(pending, Point{x, y})
			}
		}
	}
	var b Point
	for len(pending) > 0 {
		b, pending = pending[len(pending)-1], pending[:len(pending)-1]
		VisitNeighbours(b, 10, func(p Point) {
			octopi[p.y][p.x] += 1
			if octopi[p.y][p.x] > 9 && !blinked[p] {
				blinked[p] = true
				pending = append(pending, p)
			}
		})
	}
	for b = range blinked {
		octopi[b.y][b.x] = 0
	}

	return len(blinked)
}

func day11a() int {
	octopi := ParseDay11("data/day11.txt")
	blinks := 0
	for i := 0; i < 100; i++ {
		blinks += OctoStep(&octopi)
	}
	return blinks
}
func day11b() int {
	octopi := ParseDay11("data/day11.txt")
	t := 1
	for ; OctoStep(&octopi) != 100; t++ {
	}
	return t
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
	fmt.Println("day06a:", day06a())
	fmt.Println("day06b:", day06b())
	fmt.Println("day07a:", day07a())
	fmt.Println("day07b:", day07b())
	fmt.Println("day08a:", day08a())
	fmt.Println("day08b:", day08b())
	fmt.Println("day09a:", day09a())
	fmt.Println("day09b:", day09b())
	fmt.Println("day10a:", day10a())
	fmt.Println("day10b:", day10b())
	fmt.Println("day11a:", day11a())
	fmt.Println("day11b:", day11b())
}
