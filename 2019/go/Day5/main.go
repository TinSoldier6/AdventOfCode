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

func main() {
	var file io.Reader

	file = os.Stdin
	if len(os.Args) >= 2 {
		b, err := ioutil.ReadFile(os.Args[1])
		check(err)
		file = bytes.NewReader(b)
	}

	input, err := readIntcode(file)
	check(err)

	// test := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
	// 	1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
	// 	999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	state := &machine{}
	// state.memload(test)
	state.memload(input)
	state.run()

}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}
}

func readIntcode(r io.Reader) (mem []int, err error) {
	c := csv.NewReader(r)
	in, err := c.Read()
	if err != nil {
		return
	}

	mem = make([]int, len(in))
	for i, v := range in {
		mem[i], err = strconv.Atoi(v)
		if err != nil {
			return
		}
	}
	return mem, err
}
