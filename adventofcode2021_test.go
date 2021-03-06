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

func TestDay14a(t *testing.T) {
	got := day14a()
	if got != 3259 {
		t.Fail()
	}
}

func TestDay14b(t *testing.T) {
	got := day14b()
	if got != 3459174981021 {
		t.Fail()
	}
}

func TestDay15a(t *testing.T) {
	got := day15a()
	if got != 527 {
		t.Fail()
	}
}

func TestDay15b(t *testing.T) {
	got := day15b()
	if got != 2887 {
		t.Fail()
	}
}

func TestDay16a(t *testing.T) {
	got := day16a()
	if got != 929 {
		t.Fail()
	}
}

func TestDay16b(t *testing.T) {
	got := day16b()
	if got != 911945136934 {
		t.Fail()
	}
}

func TestDay17a(t *testing.T) {
	got := day17a()
	if got != 17766 {
		t.Fail()
	}
}

func TestDay17b(t *testing.T) {
	got := day17b()
	if got != 1733 {
		t.Fail()
	}
}

func TestDay18a(t *testing.T) {
	got := day18a()
	if got != 4433 {
		t.Fail()
	}
}

func TestDay18b(t *testing.T) {
	got := day18b()
	if got != 4559 {
		t.Fail()
	}
}

func TestDay19(t *testing.T) {
	a, b := day19()
	if a != 472 {
		t.Fail()
	}
	if b != 12092 {
		t.Fail()
	}
}

func TestDay20a(t *testing.T) {
	got := day20a()
	if got != 5229 {
		t.Fail()
	}
}

func TestDay20b(t *testing.T) {
	got := day20b()
	if got != 17009 {
		t.Fail()
	}
}

func TestDay21a(t *testing.T) {
	got := day21a()
	if got != 989352 {
		t.Fail()
	}
}

func TestDay21b(t *testing.T) {
	got := day21b()
	if got != 430229563871565 {
		t.Fail()
	}
}

func TestDay22a(t *testing.T) {
	got := day22a()
	if got != 601104 {
		t.Fail()
	}
}

func TestDay22b(t *testing.T) {
	got := day22b()
	if got != 1262883317822267 {
		t.Fail()
	}
}

func TestDay23a(t *testing.T) {
	got := day23a()
	if got != 15385 {
		t.Fail()
	}
}

func TestDay23b(t *testing.T) {
	got := day23b()
	if got != 49803 {
		t.Fail()
	}
}

func TestDay24a(t *testing.T) {
	got := day24a()
	if got != 45989929946199 {
		t.Fail()
	}
}

func TestDay24b(t *testing.T) {
	got := day24b()
	if got != 11912814611156 {
		t.Fail()
	}
}

func TestDay25a(t *testing.T) {
	got := day25a()
	if got != 386 {
		t.Fail()
	}
}
