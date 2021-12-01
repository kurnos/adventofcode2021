package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func day01a() (int, error) {
	file, err := os.Open("data/day01.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	previous := math.MaxInt
	growing_count := 0
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		if n > previous {
			growing_count += 1
		}
		previous = n
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return growing_count, nil
}

func day01b() (int, error) {
	file, err := os.Open("data/day01.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	a, b, c, d := 0, 0, 0, 0
	growing_count := 0
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		a, b, c, d = b, c, d, n
		if a != 0 && d != 0 && (a < d) {
			growing_count += 1
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return growing_count, nil
}

func main() {
	day01a, err := day01a()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("day01a:", day01a)
	day01b, err := day01b()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("day01b:", day01b)
}
