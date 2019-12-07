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

	state := &machine{}
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
