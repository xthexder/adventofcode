package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type Opcode int

const (
	addr Opcode = iota
	addi

	mulr
	muli

	banr
	bani

	borr
	bori

	setr
	seti

	gtir
	gtri
	gtrr

	eqir
	eqri
	eqrr
)
const NumOpcodes = 16

func run(i Opcode, a, b int, registers [4]int) int {
	switch i {
	case addr:
		return registers[a] + registers[b]
	case addi:
		return registers[a] + b

	case mulr:
		return registers[a] * registers[b]
	case muli:
		return registers[a] * b

	case banr:
		return registers[a] & registers[b]
	case bani:
		return registers[a] & b

	case borr:
		return registers[a] | registers[b]
	case bori:
		return registers[a] | b

	case setr:
		return registers[a]
	case seti:
		return a

	case gtir:
		if a > registers[b] {
			return 1
		} else {
			return 0
		}
	case gtri:
		if registers[a] > b {
			return 1
		} else {
			return 0
		}
	case gtrr:
		if registers[a] > registers[b] {
			return 1
		} else {
			return 0
		}

	case eqir:
		if a == registers[b] {
			return 1
		} else {
			return 0
		}
	case eqri:
		if registers[a] == b {
			return 1
		} else {
			return 0
		}
	case eqrr:
		if registers[a] == registers[b] {
			return 1
		} else {
			return 0
		}
	}
	panic("Unknown opcode")
}

func main() {
	var data [][4]int
	var instructionMap [NumOpcodes]uint16

	reader, err := os.Open("day16_hint.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return false
			default:
				return true
			}
		})
		ints := [4]int{}
		for i := 0; i < len(fields); i++ {
			ints[i], _ = strconv.Atoi(fields[i])
		}
		data = append(data, ints)
	}
	reader.Close()

	for i := 0; i < NumOpcodes; i++ {
		instructionMap[i] = 0xFFFF
	}

	partA := 0
	for i := 0; i < len(data); i += 4 {
		for op := Opcode(0); op < NumOpcodes; op++ {
			outr := data[i+1][3]
			out := run(op, data[i+1][1], data[i+1][2], data[i])
			if data[i+2][outr] != out {
				instructionMap[data[i+1][0]] &= ^(1 << uint(op))
			}
		}
		if bits.OnesCount16(instructionMap[data[i+1][0]]) >= 3 {
			partA++
		}
	}

	fmt.Println("Part A:", partA)

	unknown := uint16(0xFFFF)
	for unknown > 0 {
		for op := 0; op < NumOpcodes; op++ {
			if bits.OnesCount16(instructionMap[op]) == 1 {
				unknown &= ^instructionMap[op]
				for i := 0; i < NumOpcodes; i++ {
					if i != op {
						instructionMap[i] &= ^instructionMap[op]
					}
				}
			}
		}
	}
	for op := 0; op < NumOpcodes; op++ {
		instructionMap[op] = uint16(bits.Len16(instructionMap[op]) - 1)
		fmt.Printf("%2d -> %2d\n", op, instructionMap[op])
	}

	data = data[:0]
	reader, err = os.Open("day16_program.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		ints := [4]int{}
		for i := 0; i < len(fields); i++ {
			ints[i], _ = strconv.Atoi(fields[i])
		}
		data = append(data, ints)
	}
	reader.Close()

	registers := [4]int{0, 0, 0, 0}
	for i := range data {
		outr := data[i][3]
		out := run(Opcode(instructionMap[data[i][0]]), data[i][1], data[i][2], registers)
		registers[outr] = out
	}
	fmt.Println("Part B:", registers[0])
}
