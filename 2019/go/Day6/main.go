package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func parseOrbits(input []string) [][]string {
	list := make([][]string, 0, len(input))

	for _, s := range input {
		split := strings.Split(s, ")")
		list = append(list, split)
	}

	return list
}

func getPrimaryOrbits(input [][]string) map[string]string {
	orbits := make(map[string]string)

	for _, planet := range input {
		orbits[planet[1]] = planet[0]
	}

	orbits["COM"] = ""

	return orbits
}

func countOrbits(orbits map[string]string) int {
	count := 0

	for planet := range orbits {
		for planet != "COM" {
			satellite := planet
			planet = orbits[satellite]
			count++
		}
	}

	return count
}

func part1(input []string) {
	orbits := parseOrbits(input)
	primary := getPrimaryOrbits(orbits)
	count := countOrbits(primary)
	fmt.Println(count)
}

func main() {
	var file io.Reader

	file = os.Stdin
	if len(os.Args) >= 2 {
		b, err := ioutil.ReadFile(os.Args[1])
		check(err)
		file = bytes.NewReader(b)
	}

	input, err := readLines(file)
	check(err)

	part1(input)
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
		fmt.Fprintf(os.Stderr, "%s: %s", os.Args[0], err)
		os.Exit(1)
	}
}
