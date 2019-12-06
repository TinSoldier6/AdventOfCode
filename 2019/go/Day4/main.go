package main

import (
	"fmt"
	"strconv"
)

func part1(start, end int) []int {
	guesses := []int{}

	for guess := start; guess <= end; guess++ {
		b := []byte(strconv.Itoa(guess))
		if isAscending(b) && hasDouble(b) {
			guesses = append(guesses, guess)
		}
	}
	return guesses
}

func part2(list []int) []int {
	guesses := []int{}

outer:
	for _, guess := range list {
		b := []byte(strconv.Itoa(guess))
		for start, count := 0, 0; len(b) > 0; b = b[start+count:] {
			start = findPair(b)
			count = runLength(b[start:])
			if count == 2 {
				guesses = append(guesses, guess)
				continue outer
			}
		}
	}
	return guesses
}

func isAscending(b []byte) bool {
	for n, c := range b {
		if n+1 < len(b) && c > b[n+1] {
			return false
		}
	}
	return true
}

func hasDouble(b []byte) bool {

	return findPair(b) < len(b)

}

func findPair(b []byte) int {
	for n, c := range b {
		if n+1 < len(b) && c == b[n+1] {
			return n
		}
	}
	return len(b)
}

func runLength(b []byte) int {
	if len(b) < 1 {
		return 0
	}

	start := b[0]
	for n, c := range b[1:] {
		if c != start {
			return n + 1
		}
	}
	return len(b)
}

func main() {
	guesses := part1(387638, 919123)
	fmt.Println(len(guesses))
	fmt.Println(len(part2(guesses)))
}
