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

func (p1 point) right() point {
	return point{p1.x + 1, p1.y}
}

func (p1 point) left() point {
	return point{p1.x - 1, p1.y}
}

func (p1 point) up() point {
	return point{p1.x, p1.y + 1}
}

func (p1 point) down() point {
	return point{p1.x, p1.y - 1}
}

const bufsize = 1 << 16

func newPath(moves []string) []point {
	next := point{0, 0}
	path := make([]point, 0, bufsize)

	for _, move := range moves {
		dir := move[0]
		mag, err := strconv.Atoi(move[1:])
		check(err)
		fmt.Println(rune(dir), mag)
		switch dir {
		case 'R':
			for ; mag > 0; mag-- {
				next = next.right()
				path = append(path, next)
				fmt.Println(next)
			}
		case 'L':
			for ; mag > 0; mag-- {
				next = next.left()
				path = append(path, next)
			}
		case 'U':
			for ; mag > 0; mag-- {
				next = next.up()
				path = append(path, next)
			}
		case 'D':
			for ; mag > 0; mag-- {
				next = next.down()
				path = append(path, next)
			}
		default:
		}
	}

	return path
}

func part1(input [][]string) int {
	wire1 := newPath(input[0])
	wire2 := newPath(input[1])

	mark := make(map[point]int)

	for _, p := range wire1 {
		mark[p] |= 1
	}

	for _, p := range wire2 {
		mark[p] |= 2
	}

	origin := point{0, 0}
	cross := make([]int, 0, len(mark))
	for k, v := range mark {
		if v == 3 {
			cross = append(cross, dist(origin, k))
		}
	}

	min := cross[0]
	for _, d := range cross {
		min = intMin(min, d)
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
