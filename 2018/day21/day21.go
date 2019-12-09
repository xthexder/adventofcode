package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ipRegister int

type instruction struct {
	op   string
	a, b int
	outr int
}

func (i *instruction) String() string {
	out := ""
	switch i.op {
	case "addr":
		if i.a == i.outr {
			out = fmt.Sprintf("[%d] += [%d]", i.outr, i.b)
		} else if i.b == i.outr {
			out = fmt.Sprintf("[%d] += [%d]", i.outr, i.a)
		} else {
			out = fmt.Sprintf("[%d] = [%d] + [%d]", i.outr, i.a, i.b)
		}
	case "addi":
		if i.a == i.outr {
			out = fmt.Sprintf("[%d] += %d", i.outr, i.b)
		} else {
			out = fmt.Sprintf("[%d] = [%d] + %d", i.outr, i.a, i.b)
		}

	case "mulr":
		if i.a == i.outr {
			out = fmt.Sprintf("[%d] *= [%d]", i.outr, i.b)
		} else if i.b == i.outr {
			out = fmt.Sprintf("[%d] *= [%d]", i.outr, i.a)
		} else {
			out = fmt.Sprintf("[%d] = [%d] * [%d]", i.outr, i.a, i.b)
		}
	case "muli":
		if i.a == i.outr {
			out = fmt.Sprintf("[%d] *= %d", i.outr, i.b)
		} else {
			out = fmt.Sprintf("[%d] = [%d] * %d", i.outr, i.a, i.b)
		}

	case "banr":
		out = fmt.Sprintf("[%d] = [%d] & [%d]", i.outr, i.a, i.b)
	case "bani":
		out = fmt.Sprintf("[%d] = [%d] & %d (0b%b)", i.outr, i.a, i.b, i.b)

	case "borr":
		out = fmt.Sprintf("[%d] = [%d] | [%d]", i.outr, i.a, i.b)
	case "bori":
		out = fmt.Sprintf("[%d] = [%d] | %d (0b%b)", i.outr, i.a, i.b, i.b)

	case "setr":
		out = fmt.Sprintf("[%d] = [%d]", i.outr, i.a)
	case "seti":
		out = fmt.Sprintf("[%d] = %d (0b%b)", i.outr, i.a, i.a)

	case "gtir":
		out = fmt.Sprintf("[%d] = %d > [%d]", i.outr, i.a, i.b)
	case "gtri":
		out = fmt.Sprintf("[%d] = [%d] > %d", i.outr, i.a, i.b)
	case "gtrr":
		out = fmt.Sprintf("[%d] = [%d] > [%d]", i.outr, i.a, i.b)

	case "eqir":
		out = fmt.Sprintf("[%d] = %d == [%d]", i.outr, i.a, i.b)
	case "eqri":
		out = fmt.Sprintf("[%d] = [%d] == %d", i.outr, i.a, i.b)
	case "eqrr":
		out = fmt.Sprintf("[%d] = [%d] == [%d]", i.outr, i.a, i.b)
	default:
		fmt.Println("Unknown opcode", i.op)
		panic("Unknown upcode")
	}
	return strings.Replace(out, fmt.Sprintf("[%d]", ipRegister), "IP", -1)
}

func (i *instruction) exec(registers []int) {
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
	var registers [6]int
	registers[0] = 1
	var program []*instruction

	reader, err := os.Open("day21.txt")
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
			program = append(program, &instruction{
				fields[0], a, b, outr,
			})
		}
	}
	reader.Close()

	count := 0
	last := -1
	ip := 0
	var inputs [0xFFFFFF + 1]bool
	for ip >= 0 && ip < len(program) {
		if ip == 28 {
			if count == 0 {
				fmt.Println("Part A:", registers[1])
			}
			if inputs[registers[1]] {
				fmt.Println("Part B:", last)
				break
			}
			inputs[registers[1]] = true
			last = registers[1]
			count++
		}
		// fmt.Println(registers)
		// for a, i := range program {
		// 	if a == ip {
		// 		fmt.Printf("> %2d: %s\n", a, i)
		// 	} else {
		// 		fmt.Printf("  %2d: %s\n", a, i)
		// 	}
		// }
		// fmt.Println()

		registers[ipRegister] = ip
		program[ip].exec(registers[:])
		ip = registers[ipRegister]
		ip++
	}
}
