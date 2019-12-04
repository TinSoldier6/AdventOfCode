package main

import (
	"fmt"
	"os"
)

const (
	add = 1
	mult = 2
	halt = 99
)

type machine struct {
	stopped bool
	ip, reg1, reg2, reg3 int
	mem []int
}

func (m *machine) next() {
	if m.stopped {
		return
	}

	ip := m.ip
	reg1, reg2, reg3 := m.reg1, m.reg2, m.reg3
	mem := m.mem
	opcode := mem[ip]

	switch opcode {
	case 1:
		reg1, reg2, reg3 = mem[ip+1], mem[ip+2], mem[ip+3]
		mem[reg3] = mem[reg1] + mem[reg2]
		ip += 4
	case 2:
		reg1, reg2, reg3 = mem[ip+1], mem[ip+2], mem[ip+3]
		mem[reg3] = mem[reg1] * mem[reg2]
		ip += 4
	case 99:
		ip++
		m.stopped = true
	default:
		fmt.Printf("illegal instruction: %d\n", opcode)
		m.stopped = true
		fmt.Printf("%#v\n", *m)
		os.Exit(1)
	}

	m.ip = ip
	m.reg1, m.reg2, m.reg3 = reg1, reg2, reg3
	return

}

func (m *machine) run() {
	for !m.stopped {
		m.next()
	}
}

func (m *machine) reset(mem []int) {
	m.stopped = false
	m.ip = 0
	m.mem = make([]int, len(mem))
	copy(m.mem, mem)
}