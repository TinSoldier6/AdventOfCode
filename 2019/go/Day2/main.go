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

func part1(input []int) (result int) {

	state := &machine{}
	state.reset(input)

	// Preconditions regardless of input file
	state.mem[1] = 12
	state.mem[2] = 2
	state.run()
	return state.mem[0]

	// for i := 0; i < len(input); i += 4 {
	// 	opcode := input[i]
	// 	op1 := input[i+1]
	// 	op2 := input[i+2]
	// 	op3 := input[i+3]

	// 	switch opcode {
	// 	case 1:
	// 		input[op3] = input[op1] + input[op2]
	// 	case 2:
	// 		input[op3] = input[op1] * input[op2]
	// 	case 99:
	// 		return input[0]
	// 	default:
	// 		fmt.Printf("invalid opcode: %d\n", opcode)
	// 	}
	// }
	// fmt.Println("no end code found")
	// return input[0]
}

func part2(input []int) (result int) {

	expect := 19690720
	state := &machine{}
		
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			state.reset(input)
			state.mem[1] = noun
			state.mem[2] = verb
			state.run()
			if state.mem[0] == expect {
				return 100 * noun + verb
			}
		}
	}
		
	return 0
}

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
	fmt.Println(part1(input))
	fmt.Println(part2(input))

}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}
}

func readIntcode(r io.Reader) (intcode []int, err error) {
	c := csv.NewReader(r)
	in, err := c.Read()
	if err != nil {
		return
	}

	intcode = make([]int, len(in))
	for i, v := range in {
		intcode[i], err = strconv.Atoi(v)
		if err != nil {
			return
		}
	}
	return intcode, err
}
