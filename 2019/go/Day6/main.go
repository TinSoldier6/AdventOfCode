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

type orbitalSystem map[string]string

func newOrbitalSystem(orbits [][]string) (system orbitalSystem) {
	system = make(orbitalSystem)

	for _, planet := range orbits {
		system[planet[1]] = planet[0]
	}

	return system
}

func (o orbitalSystem) getPath(to, from string) (path []string) {

	if to == from {
		return path
	}

	for nextPlanet := o[from]; nextPlanet != ""; nextPlanet = o[nextPlanet] {
		path = append(path, nextPlanet)
	}

	if len(path) > 0 && path[len(path)-1] != to { // No path to destination
		path = []string{}
	}

	return path
}

func reverse(in []string) {
	for left, right := 0, len(in)-1; left < right; left, right = left+1, right-1 {
		in[left], in[right] = in[right], in[left]
	}
}

func countAllJumps(input []string) int {
	orbits := parseOrbits(input)
	system := newOrbitalSystem(orbits)

	jumps := 0
	for planet := range system {
		path := system.getPath("COM", planet)
		jumps += len(path)
	}

	return jumps
}

func countOrbitalTransfers(input []string) int {
	orbits := parseOrbits(input)
	system := newOrbitalSystem(orbits)

	myPath := system.getPath("COM", "YOU")
	santaPath := system.getPath("COM", "SAN")

	reverse(myPath)
	reverse(santaPath)

	for len(myPath) > 0 && len(santaPath) > 0 && myPath[0] == santaPath[0] {
		myPath, santaPath = myPath[1:], santaPath[1:]
	}

	return len(myPath) + len(santaPath)
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

	fmt.Println(countAllJumps(input))
	fmt.Println(countOrbitalTransfers(input))
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
