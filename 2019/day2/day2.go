package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func run(program []int) []int {
	mem := make([]int, len(program))
	copy(mem, program)

	pc := 0
	for {
		if mem[pc] == 1 {
			if pc+3 >= len(mem) {
				panic("Opcode 1 overflows")
			}
			mem[mem[pc+3]] = mem[mem[pc+1]] + mem[mem[pc+2]]
		} else if mem[pc] == 2 {
			if pc+3 >= len(mem) {
				panic("Opcode 2 overflows")
			}
			mem[mem[pc+3]] = mem[mem[pc+1]] * mem[mem[pc+2]]
		} else if mem[pc] == 99 {
			return mem
		} else {
			panic("Unknown opcode " + strconv.Itoa(mem[pc]))
		}
		pc += 4
	}
}

func main() {
	var program []int

	reader, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, str := range line {
			i, err := strconv.Atoi(str)
			if err == nil {
				program = append(program, i)
			}
		}
	}
	reader.Close()

	program[1] = 12
	program[2] = 2

	mem := run(program)
	fmt.Println("Part 1:", mem[0])

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			program[1] = noun
			program[2] = verb
			mem := run(program)
			if mem[0] == 19690720 {
				fmt.Println("Part 2:", 100*noun+verb)
			}
		}
	}
}
