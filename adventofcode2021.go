package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"log"
	"math"
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

func ParseDay11(fname string) (result [10][10]int8) {
	lines := readlines(fname)
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			result[y][x] = int8(lines[y][x] - '0')
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

func OctoStep(octopi *[10][10]int8) int {
	blinked := 0
	var pending []Point

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if octopi[y][x] <= 0 {
				octopi[y][x] = 1
			} else if octopi[y][x] < 9 {
				octopi[y][x] += 1
			} else {
				blinked += 1
				pending = append(pending, Point{x, y})
				octopi[y][x] = math.MinInt8
			}
		}
	}
	var b Point
	for len(pending) > 0 {
		b, pending = pending[len(pending)-1], pending[:len(pending)-1]
		VisitNeighbours(b, 10, func(p Point) {
			octopi[p.y][p.x] += 1
			if octopi[p.y][p.x] > 9 {
				blinked += 1
				pending = append(pending, p)
				octopi[p.y][p.x] = math.MinInt8
			}
		})
	}

	return blinked
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

func ParseDay12(fname string) map[int][]int {
	room_map := map[string]int{"start": 0, "end": 1}
	smalls, bigs := 1, 0

	map_rooms := func(s string) int {
		if mapping, ok := room_map[s]; ok {
			return mapping
		} else if strings.ToLower(s) == s {
			smalls++
			room_map[s] = smalls
			return smalls
		} else {
			bigs--
			room_map[s] = bigs
			return bigs
		}
	}

	result := make(map[int][]int)
	for _, line := range readlines(fname) {
		asdf := strings.Split(line, "-")
		a, b := map_rooms(asdf[0]), map_rooms(asdf[1])

		result[a] = append(result[a], b)
		result[b] = append(result[b], a)
	}
	return result
}

func day12a() int {
	edges := ParseDay12("data/day12.txt")

	var pathCount func(int, uint16) int
	pathCount = func(current int, visited uint16) (result int) {
		if current == 1 {
			return 1
		}

		for _, next := range edges[current] {
			if next < 0 {
				result += pathCount(next, visited)
			} else if visited&(1<<next) == 0 {
				result += pathCount(next, visited|(1<<next))
			}
		}
		return
	}

	return pathCount(0, 1)
}

func day12b() int {
	edges := ParseDay12("data/day12.txt")

	var pathCount func(int, uint16) int
	pathCount = func(current int, visited uint16) (result int) {
		if current == 1 {
			return 1
		}

		for _, next := range edges[current] {
			if next < 0 {
				result += pathCount(next, visited)
			} else if visited&(1<<next) == 0 {
				result += pathCount(next, visited|(1<<next))
			} else if next != 0 && visited&(1<<15) == 0 {
				result += pathCount(next, visited|(1<<15))
			}
		}
		return
	}

	return pathCount(0, 1)
}

func ParseDay13(fname string) (points map[Point]bool, folds []int) {
	points = make(map[Point]bool)

	for _, line := range readlines(fname) {
		if strings.HasPrefix(line, "fold along ") {
			n, _ := strconv.Atoi(line[13:])
			if strings.Contains(line, "x") {
				folds = append(folds, -n)
			} else {
				folds = append(folds, n)
			}
		} else if len(line) > 0 {
			coords := strings.SplitN(line, ",", 2)
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			points[Point{x, y}] = true
		}
	}
	return
}

func FoldPoints(points map[Point]bool, foldpoint int) map[Point]bool {
	result := make(map[Point]bool)
	i, f := (signum(foldpoint)+1)/2, Abs(foldpoint)
	for p := range points {
		x := [2]int{p.x, p.y}
		if x[i] > f {
			x[i] = 2*f - x[i]
		}
		result[Point{x[0], x[1]}] = true
	}
	return result
}

func ShowOrigami(points map[Point]bool) string {
	max_x, max_y := math.MinInt, math.MinInt
	for p := range points {
		if p.x > max_x {
			max_x = p.x
		}
		if p.y > max_y {
			max_y = p.y
		}
	}
	var buf bytes.Buffer
	for y := 0; y <= max_y; y++ {
		for x := 0; x <= max_x; x++ {
			if points[Point{x, y}] {
				fmt.Fprint(&buf, "#")
			} else {
				fmt.Fprint(&buf, ".")
			}
		}
		fmt.Fprintln(&buf)
	}
	return buf.String()
}

func day13a() int {
	points, folds := ParseDay13("data/day13.txt")
	return len(FoldPoints(points, folds[0]))
}
func day13b() string {
	points, folds := ParseDay13("data/day13.txt")
	for _, f := range folds {
		points = FoldPoints(points, f)
	}
	return ShowOrigami(points)
}

func ParseDay14(fname string) (map[byte]int, map[[2]byte]int, map[[2]byte]byte) {
	lines := readlines(fname)
	template := lines[0]
	mappings := lines[2:]

	state := make(map[[2]byte]int)
	for i := 0; i < len(template)-1; i++ {
		state[[2]byte{template[i], template[i+1]}] += 1
	}
	counts := make(map[byte]int)
	for i := 0; i < len(template); i++ {
		counts[template[i]] += 1
	}
	transformations := make(map[[2]byte]byte)
	for _, t := range mappings {
		x := strings.SplitN(t, " -> ", 2)
		transformations[[2]byte{x[0][0], x[0][1]}] = x[1][0]
	}

	return counts, state, transformations
}

func Polymerize(transformations map[[2]byte]byte, counts map[byte]int, state map[[2]byte]int) (map[byte]int, map[[2]byte]int) {
	new_counts, result := counts, make(map[[2]byte]int)
	for pair, n := range state {
		x := transformations[pair]
		new_counts[x] += n
		result[[2]byte{pair[0], x}] += n
		result[[2]byte{x, pair[1]}] += n
	}
	return new_counts, result
}

func minmaxdiff(c map[byte]int) (result int) {
	min := math.MaxInt
	for _, n := range c {
		if n < min {
			min = n
		}
		if n > result {
			result = n
		}
	}
	result -= min
	return
}

func day14a() int {
	a, b, c := ParseDay14("data/day14.txt")

	for i := 1; i <= 10; i++ {
		a, b = Polymerize(c, a, b)
	}
	return minmaxdiff(a)
}

func day14b() int {
	a, b, c := ParseDay14("data/day14.txt")

	for i := 1; i <= 40; i++ {
		a, b = Polymerize(c, a, b)
	}
	return minmaxdiff(a)
}

func ParseDay15(fname string) (result [][]int) {
	for _, line := range readlines(fname) {
		var xs []int
		for _, x := range line {
			xs = append(xs, int(x-'0'))
		}
		result = append(result, xs)
	}
	return
}

type Vertex struct {
	pos         Point
	cost, index int
}

type PriorityQueue []*Vertex

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Vertex)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Update(vertex *Vertex, cost int) {
	vertex.cost = cost
	heap.Fix(pq, vertex.index)
}

func ShortestPath(grid [][]int, start, goal Point) int {
	var verticies [][](*Vertex)
	var queue PriorityQueue
	for y, line := range grid {
		var vs [](*Vertex)
		for x := range line {
			vs = append(vs, &Vertex{Point{x, y}, math.MaxInt, -1})
			queue.Push(vs[len(vs)-1])
		}
		verticies = append(verticies, vs)
	}

	heap.Init(&queue)
	queue.Update(verticies[0][0], 0)

	for {
		v := heap.Pop(&queue).(*Vertex)
		if v.pos == goal {
			return v.cost
		}
		p := v.pos
		neighbours := []Point{
			{p.x, p.y + 1},
			{p.x, p.y - 1},
			{p.x + 1, p.y},
			{p.x - 1, p.y},
		}
		for _, t := range neighbours {
			if t.y >= 0 && t.y < len(grid) && t.x >= 0 && t.x < len(grid[0]) {
				cost := v.cost + grid[t.y][t.x]
				if verticies[t.y][t.x].index != -1 && cost < verticies[t.y][t.x].cost {
					queue.Update(verticies[t.y][t.x], cost)
				}
			}
		}
	}
}

func MultiGrid(grid [][]int) (result [][]int) {
	for my := 0; my < 5; my++ {
		for _, line := range grid {
			var cs []int
			for mx := 0; mx < 5; mx++ {
				for _, c := range line {
					c += mx + my
					cs = append(cs, ((c-1)%9)+1)
				}
			}
			result = append(result, cs)
		}
	}
	return
}

func day15a() int {
	grid := ParseDay15("data/day15.txt")
	return ShortestPath(grid, Point{0, 0}, Point{len(grid) - 1, len(grid[0]) - 1})
}

func day15b() int {
	grid := MultiGrid(ParseDay15("data/day15.txt"))
	return ShortestPath(grid, Point{0, 0}, Point{len(grid) - 1, len(grid[0]) - 1})
}

func ParseDay16(fname string) (result []byte) {
	for _, c := range readlines(fname)[0] {
		var n byte
		if c <= '9' {
			n = byte(c - '0')
		} else {
			n = 10 + byte(c-'A')
		}
		result = append(append(append(append(result, (n>>3)&1), (n>>2)&1), (n>>1)&1), n&1)
	}
	return result
}

func ReadBit(bits []byte) (byte, []byte) {
	return bits[0], bits[1:]
}

func ReadInt(bits []byte, n int) (int, []byte) {
	result := 0
	var b byte
	for ; n > 0; n-- {
		b, bits = ReadBit(bits)
		result = result<<1 + int(b)
	}
	return result, bits
}

type Packet struct {
	version       int
	type_id       int
	literal_value int
	sub_packets   []Packet
}

func ReadPacket(bits []byte) (Packet, []byte) {
	version, type_id, bits := ReadHeader(bits)
	if type_id == 4 {
		value, bits := ReadLiteralValue(bits)
		return Packet{version, type_id, value, nil}, bits
	} else {
		sub_packets, bits := ReadSubPackets(bits)
		return Packet{version, type_id, -1, sub_packets}, bits
	}
}

func ReadHeader(bits []byte) (int, int, []byte) {
	version, bits := ReadInt(bits, 3)
	type_id, bits := ReadInt(bits, 3)
	return version, type_id, bits
}

func ReadLiteralValue(bits []byte) (int, []byte) {
	var marker byte
	var d int
	result := 0
	for {
		marker, bits = ReadBit(bits)
		d, bits = ReadInt(bits, 4)
		result = result<<4 + d
		if marker == 0 {
			break
		}
	}
	return result, bits
}

func ReadSubPackets(bits []byte) ([]Packet, []byte) {
	length_type, bits := ReadBit(bits)

	var packet Packet
	var packets []Packet

	if length_type == 0 {
		var bit_length int
		bit_length, bits = ReadInt(bits, 15)
		sub_bits, bits := bits[:bit_length], bits[bit_length:]
		for len(sub_bits) > 0 {
			packet, sub_bits = ReadPacket(sub_bits)
			packets = append(packets, packet)
		}
		return packets, bits
	} else {
		msg_length, bits := ReadInt(bits, 11)
		for i := 0; i < msg_length; i++ {
			packet, bits = ReadPacket(bits)
			packets = append(packets, packet)
		}
		return packets, bits
	}
}

func day16a() int {
	p, _ := ReadPacket(ParseDay16("data/day16.txt"))

	var version_sum func(Packet) int
	version_sum = func(p Packet) int {
		result := p.version
		for _, p := range p.sub_packets {
			result += version_sum(p)
		}
		return result
	}
	return version_sum(p)
}

func day16b() int {
	p, _ := ReadPacket(ParseDay16("data/day16.txt"))

	var eval func(Packet) int
	eval = func(p Packet) int {
		switch p.type_id {
		case 0:
			result := 0
			for _, p := range p.sub_packets {
				result += eval(p)
			}
			return result
		case 1:
			result := 1
			for _, p := range p.sub_packets {
				result *= eval(p)
			}
			return result
		case 2:
			result := math.MaxInt
			for _, p := range p.sub_packets {
				if v := eval(p); v < result {
					result = v
				}
			}
			return result
		case 3:
			result := math.MinInt
			for _, p := range p.sub_packets {
				if v := eval(p); v > result {
					result = v
				}
			}
			return result
		case 4:
			return p.literal_value
		case 5:
			v1, v2 := eval(p.sub_packets[0]), eval(p.sub_packets[1])
			if v1 > v2 {
				return 1
			} else {
				return 0
			}
		case 6:
			v1, v2 := eval(p.sub_packets[0]), eval(p.sub_packets[1])
			if v1 < v2 {
				return 1
			} else {
				return 0
			}
		case 7:
			v1, v2 := eval(p.sub_packets[0]), eval(p.sub_packets[1])
			if v1 == v2 {
				return 1
			} else {
				return 0
			}
		default:
			log.Fatal()
			return -1
		}
	}
	return eval(p)
}

func AtoI(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return i
}

func ParseDay17(fname string) (int, int, int, int) {
	lines := readlines(fname)
	s := lines[0][13:]
	a := strings.SplitN(s, ", ", 2)
	xs := strings.SplitN(a[0][2:], "..", 2)
	ys := strings.SplitN(a[1][2:], "..", 2)
	return AtoI(xs[0]), AtoI(xs[1]), AtoI(ys[0]), AtoI(ys[1])
}

func day17a() int {
	// Assuming we can lob the projectile so it can fall vertically into the target...
	_, _, ymin, ymax := ParseDay17("data/day17.txt")

	best, max := 0, 0
	for dy0 := 0; dy0 <= -ymin; dy0++ {
		max += dy0
		for dy, y := -dy0, 0; y > ymin; dy, y = dy-1, y+dy {
			if y <= ymax && max > best {
				best = max
				break
			}
		}
	}
	return best
}

func day17b() int {
	xmin, xmax, ymin, ymax := ParseDay17("data/day17.txt")

	var dxts, dxToMinT [][2]int
	for dx0 := xmax; dx0 > 0; dx0-- {
		for x, dx, t := 0, dx0, 0; x <= xmax; x, dx, t = x+dx, dx-1, t+1 {
			if dx == 0 {
				if x >= xmin && x <= xmax {
					dxToMinT = append(dxToMinT, [2]int{dx0, t})
				}
				break
			}
			if x >= xmin && x <= xmax {
				dxts = append(dxts, [2]int{dx0, t})
			}
		}
	}

	var dyts [][2]int
	for dy0 := ymin; dy0 < -ymin; dy0++ {
		for y, dy, t := 0, dy0, 0; y >= ymin; y, dy, t = y+dy, dy-1, t+1 {
			if y >= ymin && y <= ymax {
				dyts = append(dyts, [2]int{dy0, t})
			}
		}
	}

	cartesian_product := make(map[[2]int]bool)
	for _, dyt := range dyts {
		dy, t0 := dyt[0], dyt[1]
		for _, dxt := range dxts {
			dx, t1 := dxt[0], dxt[1]
			if t0 == t1 {
				cartesian_product[[2]int{dx, dy}] = true
			}
		}
		for _, dxMinT := range dxToMinT {
			if t0 >= dxMinT[1] {
				cartesian_product[[2]int{dxMinT[0], dy}] = true
			}
		}
	}
	return len(cartesian_product)
}

type Snail struct {
	n, m       byte
	prev, next *Snail
}

func NewSnail(n, m byte) *Snail {
	return &Snail{n, m, nil, nil}
}

func ReadSnails(s string) (head *Snail) {
	current, m := new(Snail), byte(1)
	head = current
	for _, c := range []byte(s) {
		switch c {
		case '[':
			m *= 3
		case ',':
			m = 2 * (m / 3)
		case ']':
			m = m / 2
		default:
			current = current.insertAfter(NewSnail(byte(c-'0'), m))
		}
	}
	head = head.next
	head.prev.remove()
	return head
}

func ParseDay18(fname string) (result []*Snail) {
	for _, line := range readlines(fname) {
		result = append(result, ReadSnails(line))
	}
	return
}

func (a *Snail) insertAfter(b *Snail) *Snail {
	if a.next != nil {
		a.next, b.prev, b.next, a.next.prev = b, a, a.next, b
	} else {
		a.next, b.prev, b.next = b, a, nil
	}
	return a.next
}

func (b *Snail) remove() {
	if b.prev != nil {
		b.prev.next = b.next
	}
	if b.next != nil {
		b.next.prev = b.prev
	}
}

func SnailMagnitude(snail *Snail) (result int) {
	for ; snail != nil; snail = snail.next {
		result += int(snail.n) * int(snail.m)
	}
	return
}

func SnailMagTooBig(m byte) bool {
	// [2**(5-n)*3**n for n in range(6)]
	return m == 32 || m == 48 || m == 72 || m == 108 || m == 162 || m == 243
}

func AddSnails(a, b *Snail) *Snail {
	head := new(Snail)
	current := head
	for s := a; s != nil; s = s.next {
		current = current.insertAfter(NewSnail(s.n, 3*s.m))
	}
	for s := b; s != nil; s = s.next {
		current = current.insertAfter(NewSnail(s.n, 2*s.m))
	}
	head = head.next
	head.prev.remove()
	for {
		var toSplit, toExplode *Snail
		for snail := head; snail != nil; snail = snail.next {
			if toSplit == nil && snail.n > 9 {
				toSplit = snail
			}
			if toExplode == nil && SnailMagTooBig(snail.m) {
				toExplode = snail
				break
			}
		}

		if toExplode != nil {
			n1, n2 := toExplode.n, toExplode.next.n
			toExplode.n, toExplode.m = 0, toExplode.m/3
			toExplode.next.remove()
			if toExplode.prev != nil {
				toExplode.prev.n += n1
			}
			if toExplode.next != nil {
				toExplode.next.n += n2
			}
		} else if toSplit != nil {
			half := toSplit.n / 2
			toSplit.insertAfter(&Snail{toSplit.n - half, 2 * toSplit.m, nil, nil})
			toSplit.n, toSplit.m = half, 3*toSplit.m
		} else {
			break
		}
	}
	return head
}

func day18a() int {
	snails2 := ParseDay18("data/day18.txt")
	a := snails2[0]
	for _, b := range snails2[1:] {
		a = AddSnails(a, b)
	}
	return SnailMagnitude(a)
}

func day18b() int {
	snails := ParseDay18("data/day18.txt")

	var results []int
	for i, a := range snails {
		for _, b := range snails[i+1:] {
			results = append(results, SnailMagnitude(AddSnails(a, b)))
			results = append(results, SnailMagnitude(AddSnails(b, a)))
		}
	}
	max := 0
	for _, m := range results {
		if m > max {
			max = m
		}
	}
	return max
}

func ParseDay19(fname string) (scanners []Scan) {
	var current Scan
	var n int
	for _, line := range readlines(fname) {
		if len(line) == 0 {
			current.fingerprint = fingerprint(current)
			scanners = append(scanners, current)
		} else if i, _ := fmt.Sscanf(line, "--- scanner %d ---", &n); i == 1 {
			current = Scan{n, nil, nil}
		} else {
			fields := strings.SplitN(line, ",", 3)
			current.points = append(current.points, Point3{AtoI(fields[0]), AtoI(fields[1]), AtoI(fields[2])})
		}
	}
	current.fingerprint = fingerprint(current)
	scanners = append(scanners, current)
	return
}

type Rot = [3][3]int
type Point3 = [3]int
type Scan struct {
	n           int
	fingerprint []int
	points      []Point3
}
type FixedScan struct {
	n           int
	center      Point3
	fingerprint []int
	points      map[Point3]bool
}

func NewFixedScan(s Scan, center Point3) FixedScan {
	return FixedScan{s.n, center, s.fingerprint, make(map[[3]int]bool)}
}

func rotations() [24]Rot {
	return [24]Rot{
		{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}},
		{{-1, 0, 0}, {0, -1, 0}, {0, 0, 1}},
		{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}},
		{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}},
		{{0, 0, -1}, {1, 0, 0}, {0, -1, 0}},
		{{0, 0, -1}, {0, -1, 0}, {-1, 0, 0}},
		{{0, 0, -1}, {-1, 0, 0}, {0, 1, 0}},
		{{-1, 0, 0}, {0, 1, 0}, {0, 0, -1}},
		{{0, 1, 0}, {1, 0, 0}, {0, 0, -1}},
		{{1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
		{{0, -1, 0}, {-1, 0, 0}, {0, 0, -1}},
		{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}},
		{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}},
		{{0, 0, 1}, {0, -1, 0}, {1, 0, 0}},
		{{0, 0, 1}, {-1, 0, 0}, {0, -1, 0}},
		{{0, -1, 0}, {0, 0, 1}, {-1, 0, 0}},
		{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}},
		{{-1, 0, 0}, {0, 0, 1}, {0, 1, 0}},
		{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
		{{-1, 0, 0}, {0, 0, -1}, {0, -1, 0}},
		{{0, 1, 0}, {0, 0, -1}, {-1, 0, 0}},
		{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}},
		{{0, -1, 0}, {0, 0, -1}, {1, 0, 0}},
	}
}

func rotate(p Point3, t Rot) Point3 {
	return Point3{
		p[0]*t[0][0] + p[1]*t[0][1] + p[2]*t[0][2],
		p[0]*t[1][0] + p[1]*t[1][1] + p[2]*t[1][2],
		p[0]*t[2][0] + p[1]*t[2][1] + p[2]*t[2][2],
	}
}

func add(a, b Point3) Point3 {
	return Point3{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func sub(a, b Point3) Point3 {
	return Point3{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func transform_scan(scan Scan, rotation Rot, new_center Point3) FixedScan {
	result := NewFixedScan(scan, new_center)
	for _, p := range scan.points {
		result.points[add(rotate(p, rotation), new_center)] = true
	}
	return result
}

func popat(slice []Scan, i int) ([]Scan, Scan) {
	slice[i], slice[len(slice)-1] = slice[len(slice)-1], slice[i]
	return slice[:len(slice)-1], slice[len(slice)-1]
}

func overlaps_with_transformation(s1 FixedScan, s2 Scan, r Rot, d Point3) bool {
	count := 0
	for _, p := range s2.points {
		tp := add(d, rotate(p, r))
		if s1.points[tp] {
			count++
			if count == 12 {
				return true
			}
		}
	}
	return false
}

func overlappingTransformation(s1 FixedScan, s2 Scan) (bool, Rot, Point3) {
	matches := 0
	for i, j := 0, 0; i < len(s1.fingerprint) && j < len(s2.fingerprint); {
		if s1.fingerprint[i] == s2.fingerprint[j] {
			i, j, matches = i+1, j+1, matches+1
		} else if s1.fingerprint[i] < s2.fingerprint[j] {
			i++
		} else {
			j++
		}
	}

	if matches < 66 {
		return false, Rot{}, Point3{}
	}

	rotations := rotations()
	for _, r := range rotations {
		for _, p2 := range s2.points {
			tp2 := rotate(p2, r)
			for p1 := range s1.points {
				d := sub(p1, tp2)
				if overlaps_with_transformation(s1, s2, r, d) {
					return true, r, d
				}
			}
		}
	}
	return false, Rot{}, Point3{}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func fingerprint(s Scan) (result []int) {
	for i, p1 := range s.points {
		for _, p2 := range s.points[i+1:] {
			d := sub(p1, p2)
			l1 := Abs(d[0]) + Abs(d[1]) + Abs(d[2])
			linf := Max(Max(Abs(d[0]), Abs(d[1])), Abs(d[2]))
			result = append(result, l1<<10|linf)
		}
	}
	sort.Ints(result)
	return
}

func day19() (int, int) {
	remaining := ParseDay19("data/day19.txt")

	remaining, s0 := popat(remaining, 0)
	fixed := []FixedScan{transform_scan(s0, rotations()[0], Point3{})}

	for len(remaining) > 0 {
		for _, s1 := range fixed {
			for i, s2 := range remaining {
				if ok, r, d := overlappingTransformation(s1, s2); ok {
					remaining, _ = popat(remaining, i)
					fixed = append(fixed, transform_scan(s2, r, d))
					break
				}
			}
		}
	}

	union := make(map[Point3]bool)
	for _, f := range fixed {
		for p := range f.points {
			union[p] = true
		}
	}

	max := 0
	for i, f1 := range fixed {
		c1 := f1.center
		for _, f2 := range fixed[i+1:] {
			c2 := f2.center
			d := Abs(c1[0]-c2[0]) + Abs(c1[1]-c2[1]) + Abs(c1[2]-c2[2])
			if d > max {
				max = d
			}
		}
	}

	return len(union), max
}

func ParseDay20(fname string) (key [512]byte, image [][]byte) {
	lines := readlines(fname)
	for i, c := range []byte(lines[0]) {
		key[i] = c & 1
	}
	for _, line := range lines[2:] {
		var is []byte
		for _, c := range []byte(line) {
			is = append(is, c&1)
		}
		image = append(image, is)
	}
	return
}

func EnhanceImage(image [][]byte, outer byte, key [512]byte) ([][]byte, byte) {
	size := len(image)
	new_image := make([][]byte, size+2)
	for y := 0; y < size+2; y++ {
		new_image[y] = make([]byte, size+2)
	}

	for y0 := -1; y0 < size+1; y0++ {
		for x0 := -1; x0 < size+1; x0++ {
			t := 0
			for y := y0 - 1; y <= y0+1; y++ {
				for x := x0 - 1; x <= x0+1; x++ {
					p := int(outer)
					if y >= 0 && x >= 0 && y < size && x < size {
						p = int(image[y][x])
					}
					t = t<<1 + p
				}
			}
			new_image[y0+1][x0+1] = key[t]
		}
	}
	outer = key[511*int(outer)]
	return new_image, outer
}

func CountPixels(image [][]byte) (count int) {
	for _, line := range image {
		for _, p := range line {
			count += int(p)
		}
	}
	return
}

func day20a() int {
	key, image := ParseDay20("data/day20.txt")
	outer := byte(0)
	image, outer = EnhanceImage(image, outer, key)
	image, _ = EnhanceImage(image, outer, key)
	return CountPixels(image)
}

func day20b() int {
	key, image := ParseDay20("data/day20.txt")
	outer := byte(0)
	for i := 0; i < 50; i++ {
		image, outer = EnhanceImage(image, outer, key)
	}
	return CountPixels(image)
}

func ParseDay21(fname string) (result [2]byte) {
	lines := readlines(fname)
	for i := 0; i < 2; i++ {
		result[i] = byte(AtoI(strings.SplitN(lines[i], ": ", 2)[1]) - 1)
	}
	return
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func day21a() int {
	pos := ParseDay21("data/day21.txt")
	scores := [2]int{}
	die, rolls := byte(100), 0
	roll := func() byte {
		die, rolls = (die%100)+1, rolls+1
		return die % 10
	}
	for i := 0; scores[0] < 1000 && scores[1] < 1000; i = 1 - i {
		pos[i] = (pos[i] + roll() + roll() + roll()) % 10
		scores[i] += int(pos[i]) + 1
	}
	return Min(scores[0], scores[1]) * rolls
}

type Asdf struct {
	ps [2]byte
	ss [2]int
}

func day21b() int {
	pos := ParseDay21("data/day21.txt")
	states := map[Asdf]int{{pos, [2]int{}}: 1}
	wins := [2]int{}
	rolls := [7]int{1, 3, 6, 7, 6, 3, 1}

	for i := 0; len(states) > 0; i = 1 - i {
		new := make(map[Asdf]int)
		for s, count := range states {
			for r, m := range rolls {
				next_s := s
				next_s.ps[i] = (next_s.ps[i] + byte(r) + 3) % 10
				next_s.ss[i] += int(next_s.ps[i] + 1)
				if next_s.ss[i] >= 21 {
					wins[i] += count * m
				} else {
					new[next_s] += count * m
				}
			}
		}
		states = new
	}

	return Max(wins[0], wins[1])
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
	fmt.Println("day12a:", day12a())
	fmt.Println("day12b:", day12b())
	fmt.Println("day13a:", day13a())
	fmt.Println("day13b:", day13b())
	fmt.Println("day14a:", day14a())
	fmt.Println("day14b:", day14b())
	fmt.Println("day15a:", day15a())
	fmt.Println("day15b:", day15b())
	fmt.Println("day16a:", day16a())
	fmt.Println("day16b:", day16b())
	fmt.Println("day17a:", day17a())
	fmt.Println("day17b:", day17b())
	fmt.Println("day18a:", day18a())
	fmt.Println("day18b:", day18b())
	a, b := day19()
	fmt.Print("day19:", a, b)
	fmt.Println("day20a:", day20a())
	fmt.Println("day20b:", day20b())
	fmt.Println("day21a:", day21a())
	fmt.Println("day21b:", day21b())
}
