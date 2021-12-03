package main

import "testing"

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
