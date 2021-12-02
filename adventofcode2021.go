package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readints(fname string) (result []int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
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

func main() {
	fmt.Println("day01a:", day01a())
	fmt.Println("day01b:", day01b())
	fmt.Println("day02a:", day02a())
	fmt.Println("day02b:", day02b())
}
