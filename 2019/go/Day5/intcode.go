package main

import (
	"fmt"
)

// Opcodes
const (
	NOP = iota
	ADD
	MUL
	INP
	OUT
	JNZ
	JZ
	LT
	EQ
	HLT = 99
)

// Addressing modes
const (
	REL = iota
	IMM
)

var MODEBITS = []int{100, 1000, 10000}

type machine struct {
	ticks  int
	halted bool
	ip, op int
	regs   [3]int
	mem    []int
}

func (m *machine) fetch() int {
	op := m.mem[m.ip]
	m.ip++
	return op
}

func (m *machine) load(r int) {
	addr := m.ip
	op := m.fetch()
	switch (m.op / MODEBITS[r]) % 10 {
	case REL:
		addr = op
	case IMM:
	}
	m.regs[r] = m.mem[addr]
}

func (m *machine) store(r int) {
	addr := m.ip
	op := m.fetch()
	switch (m.op / MODEBITS[r]) % 10 {
	case REL:
		addr = op
	case IMM:
	}
	m.mem[addr] = m.regs[r]
}

func (m *machine) add() {
	m.load(0)
	m.load(1)
	m.regs[2] = m.regs[0] + m.regs[1]
	m.store(2)
}

func (m *machine) mul() {
	m.load(0)
	m.load(1)
	m.regs[2] = m.regs[0] * m.regs[1]
	m.store(2)
}

func (m *machine) inp() {
	fmt.Printf(" :")
	fmt.Scanf("%v", &m.regs[0])
	fmt.Println()
	m.store(0)
}

func (m *machine) out() {
	m.load(0)
	fmt.Printf("%v ", m.regs[0])
}

func (m *machine) jnz() {
	m.load(0)
	m.load(1)
	if m.regs[0] != 0 {
		m.ip = m.regs[1]
	}
}

func (m *machine) jz() {
	m.load(0)
	m.load(1)
	if m.regs[0] == 0 {
		m.ip = m.regs[1]
	}
}

func (m *machine) lt() {
	m.load(0)
	m.load(1)
	m.regs[2] = 0
	if m.regs[0] < m.regs[1] {
		m.regs[2] = 1
	}
	m.store(2)
}

func (m *machine) eq() {
	m.load(0)
	m.load(1)
	m.regs[2] = 0
	if m.regs[0] == m.regs[1] {
		m.regs[2] = 1
	}
	m.store(2)
}

func (m *machine) hlt() {
	m.halted = true
}

func (m *machine) next() bool {
	if m.halted {
		return false
	}

	m.ticks++
	m.op = m.fetch()
	switch m.op % 100 {
	case NOP:
	case ADD:
		m.add()
	case MUL:
		m.mul()
	case INP:
		m.inp()
	case OUT:
		m.out()
	case JNZ:
		m.jnz()
	case JZ:
		m.jz()
	case LT:
		m.lt()
	case EQ:
		m.eq()
	case HLT:
		m.hlt()
		return false
	}
	return true
}

func (m *machine) run() {
	m.halted = false
	for m.next() {
	}
}

func (m *machine) reset() {
	*m = machine{mem: m.mem}
}

func (m *machine) memload(mem []int) {
	m.mem = make([]int, len(mem))
	copy(m.mem, mem)
}
