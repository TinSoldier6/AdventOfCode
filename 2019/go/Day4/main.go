package main

import (
	"fmt"
	"strconv"
)

func part1(start, end int) int {
	guesses := []int{}

	for guess := start; guess <= end; guess++ {
		if isAscending(guess) && hasDouble(guess) {
			guesses = append(guesses, guess)
		}
	}
	return len(guesses)
}

func isAscending(i int) bool {
	a := []byte(strconv.Itoa(i))

	for n, c := range a {
		if n+1 < len(a) {
			if c > a[n+1] {
				return false
			}
		}
	}
	return true
}

func hasDouble(i int) bool {
	a := []byte(strconv.Itoa(i))

	for n, c := range a {
		if n+1 < len(a) {
			if c == a[n+1] {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(part1(387638, 919123))
}
