package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

// Read input.txt
// Convert strings to ints (mass)
// For each mass, divide by 3, round down, subtract 2
// f = floor(m / 3) - 2
// total = sum(f)

func part1(r io.Reader) (result string) {
	input, err := readLines(r)
	check(err)

	var totalFuel int
	for _, line := range input {
		mass, err := strconv.Atoi(line)
		check(err)
		fuel := mass/3 - 2
		totalFuel += fuel
	}

	return strconv.Itoa(totalFuel)

}

func part2(r io.Reader) (result string) {
	return ""
}

func main() {
	var file io.Reader

	file = os.Stdin
	if len(os.Args) >= 2 {
		b, err := ioutil.ReadFile(os.Args[1])
		check(err)
		file = bytes.NewReader(b)
	}

	fmt.Println(part1(file))
	fmt.Println(part2(file))

}

func readLines(r io.Reader) (lines []string, err error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, s.Err()
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
