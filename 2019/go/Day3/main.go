package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

type point struct {
	x, y int
}

type path map[point]int

func walk(moves []string) path {
	path := make(path)

	x, y := 0, 0
	dx, dy := 0, 0
	total := 0

	for _, m := range moves {
		steps, err := strconv.Atoi(m[1:])
		check(err)
		switch m[0] {
		case 'R':
			dx, dy = 1, 0
		case 'L':
			dx, dy = -1, 0
		case 'U':
			dx, dy = 0, 1
		case 'D':
			dx, dy = 0, -1
		}

		for ; steps > 0; steps-- {
			x += dx
			y += dy
			total++
			path[point{x, y}] = total
		}
	}

	return path

}

func intersections(points1, points2 path) path {
	intersect := make(path)
	for point, length1 := range points1 {
		if length2, ok := points2[point]; ok {
			intersect[point] = length1 + length2
		}
	}

	return intersect

}

func part1(input [][]string) int {
	wire1 := walk(input[0])
	wire2 := walk(input[1])

	intersects := intersections(wire1, wire2)

	distances := make(path)
	for point := range intersects {
		distances[point] = intAbs(point.x) + intAbs(point.y)
	}

	min := 0
	for _, min = range distances {
		break
	}
	for _, d := range distances {
		if d < min {
			min = d
		}
	}
	
	return min

}

func part2(input [][]string) int {
	wire1 := walk(input[0])
	wire2 := walk(input[1])

	intersects := intersections(wire1, wire2)

	min := 0
	for _, min = range intersects {
		break
	}
	for _, d := range intersects {
		if d < min {
			min = d
		}
	}
	
	return min

}

func main() {
	var file io.Reader

	file = os.Stdin
	if len(os.Args) >= 2 {
		b, err := ioutil.ReadFile(os.Args[1])
		check(err)
		file = bytes.NewReader(b)
	}

	input, err := csv.NewReader(file).ReadAll()
	check(err)

	fmt.Println(part1(input))
	fmt.Println(part2(input))

}

func check(err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}
}

func intMin(a, b int) (c int) {
	if a <= b {
		return a
	}
	return b
}

func intAbs(a int) int {
	b := a >> (strconv.IntSize - 1)
	return (a ^ b) - b
}

func dist(p1, p2 point) int {
	return intAbs(p1.x-p2.x) + intAbs(p1.y-p2.y)
}
