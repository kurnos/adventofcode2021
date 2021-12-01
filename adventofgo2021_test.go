package main

import "testing"

func TestDay01a(t *testing.T) {
	got, err := day01a()
	if err != nil || got != 1288 {
		t.Fail()
	}
}

func TestDay01b(t *testing.T) {
	got, err := day01b()
	if err != nil || got != 1311 {
		t.Fail()
	}
}
