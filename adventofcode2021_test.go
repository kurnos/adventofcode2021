package main

import (
	"testing"
)

func TestDay01a(t *testing.T) {
	got := day01a()
	if got != 1288 {
		t.Fail()
	}
}

func TestDay01b(t *testing.T) {
	got := day01b()
	if got != 1311 {
		t.Fail()
	}
}

func TestDay02a(t *testing.T) {
	got := day02a()
	if got != 1924923 {
		t.Fail()
	}
}

func TestDay02b(t *testing.T) {
	got := day02b()
	if got != 1982495697 {
		t.Fail()
	}
}

func TestDay03a(t *testing.T) {
	got := day03a()
	if got != 3969000 {
		t.Fail()
	}
}

func TestDay03b(t *testing.T) {
	got := day03b()
	if got != 4267809 {
		t.Fail()
	}
}

func TestDay04a(t *testing.T) {
	got := day04a()
	if got != 11536 {
		t.Fail()
	}
}

func TestDay04b(t *testing.T) {
	got := day04b()
	if got != 1284 {
		t.Fail()
	}
}

func TestDay05a(t *testing.T) {
	got := day05a()
	if got != 6267 {
		t.Fail()
	}
}

func TestDay05b(t *testing.T) {
	got := day05b()
	if got != 20196 {
		t.Fail()
	}
}

func TestDay06a(t *testing.T) {
	got := day06a()
	if got != 386755 {
		t.Fail()
	}
}

func TestDay06b(t *testing.T) {
	got := day06b()
	if got != 1732731810807 {
		t.Fail()
	}
}

func TestDay07a(t *testing.T) {
	got := day07a()
	if got != 335330 {
		t.Fail()
	}
}

func TestDay07b(t *testing.T) {
	got := day07b()
	if got != 92439766 {
		t.Fail()
	}
}

func TestDay08a(t *testing.T) {
	got := day08a()
	if got != 367 {
		t.Fail()
	}
}

func TestDay08b(t *testing.T) {
	got := day08b()
	if got != 974512 {
		t.Fail()
	}
}

func TestDay09a(t *testing.T) {
	got := day09a()
	if got != 530 {
		t.Fail()
	}
}

func TestDay09b(t *testing.T) {
	got := day09b()
	if got != 1019494 {
		t.Fail()
	}
}

func TestDay10a(t *testing.T) {
	got := day10a()
	if got != 394647 {
		t.Fail()
	}
}

func TestDay10b(t *testing.T) {
	got := day10b()
	if got != 2380061249 {
		t.Fail()
	}
}

func TestDay11a(t *testing.T) {
	got := day11a()
	if got != 1747 {
		t.Fail()
	}
}

func TestDay11b(t *testing.T) {
	got := day11b()
	if got != 505 {
		t.Fail()
	}
}

func TestDay12a(t *testing.T) {
	got := day12a()
	if got != 4413 {
		t.Fail()
	}
}

func TestDay12b(t *testing.T) {
	got := day12b()
	if got != 118803 {
		t.Fail()
	}
}

func TestDay13a(t *testing.T) {
	got := day13a()
	if got != 610 {
		t.Fail()
	}
}

func TestDay13b(t *testing.T) {
	got := day13b()
	expected := `###..####.####...##.#..#.###..####.####
#..#....#.#.......#.#..#.#..#.#.......#
#..#...#..###.....#.####.#..#.###....#.
###...#...#.......#.#..#.###..#.....#..
#....#....#....#..#.#..#.#.#..#....#...
#....####.#.....##..#..#.#..#.#....####
`
	if got != expected {
		t.Fail()
	}
}
