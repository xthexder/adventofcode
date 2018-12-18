package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op   string
	a, b int
	outr int
}

func (i instruction) exec(registers []int) {
	switch i.op {
	case "addr":
		registers[i.outr] = registers[i.a] + registers[i.b]
	case "addi":
		registers[i.outr] = registers[i.a] + i.b

	case "mulr":
		registers[i.outr] = registers[i.a] * registers[i.b]
	case "muli":
		registers[i.outr] = registers[i.a] * i.b

	case "banr":
		registers[i.outr] = registers[i.a] & registers[i.b]
	case "bani":
		registers[i.outr] = registers[i.a] & i.b

	case "borr":
		registers[i.outr] = registers[i.a] | registers[i.b]
	case "bori":
		registers[i.outr] = registers[i.a] | i.b

	case "setr":
		registers[i.outr] = registers[i.a]
	case "seti":
		registers[i.outr] = i.a

	case "gtir":
		if i.a > registers[i.b] {
			registers[i.outr] = 1
		} else {
			registers[i.outr] = 0
		}
	case "gtri":
		if registers[i.a] > i.b {
			registers[i.outr] = 1
		} else {
			registers[i.outr] = 0
		}
	case "gtrr":
		if registers[i.a] > registers[i.b] {
			registers[i.outr] = 1
		} else {
			registers[i.outr] = 0
		}

	case "eqir":
		if i.a == registers[i.b] {
			registers[i.outr] = 1
		} else {
			registers[i.outr] = 0
		}
	case "eqri":
		if registers[i.a] == i.b {
			registers[i.outr] = 1
		} else {
			registers[i.outr] = 0
		}
	case "eqrr":
		if registers[i.a] == registers[i.b] {
			registers[i.outr] = 1
		} else {
			registers[i.outr] = 0
		}
	default:
		fmt.Println("Unknown opcode", i.op)
		panic("Unknown upcode")
	}
}

func main() {
	ipRegister := 0
	var registers [6]int
	var program []instruction

	reader, err := os.Open("day19.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if fields[0] == "#ip" {
			ipRegister, _ = strconv.Atoi(fields[1])
		} else {
			a, _ := strconv.Atoi(fields[1])
			b, _ := strconv.Atoi(fields[2])
			outr, _ := strconv.Atoi(fields[3])
			program = append(program, instruction{
				fields[0], a, b, outr,
			})
		}
	}
	reader.Close()

	ip := 0
	for ip >= 0 && ip < len(program) {
		registers[ipRegister] = ip
		program[ip].exec(registers[:])
		ip = registers[ipRegister]
		ip++
	}
	fmt.Println("Part A:", registers[0])
}
