package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type inputOutput interface {
	Input() int
	Output(int)
}

func run(program []int, io inputOutput) []int {
	mem := make([]int, len(program)+10000)
	copy(mem, program)

	relativeBase := 0
	pc := 0
	for {
		op := mem[pc] % 100
		mode1 := (mem[pc] / 100) % 10
		mode2 := (mem[pc] / 1000) % 10
		mode3 := (mem[pc] / 10000) % 10
		param1 := 0
		param2 := 0
		param3 := 0
		if op >= 1 && op <= 9 {
			if mode1 == 0 {
				param1 = mem[pc+1]
			} else if mode1 == 1 {
				param1 = pc + 1
			} else if mode1 == 2 {
				param1 = mem[pc+1] + relativeBase
			} else {
				panic("Unknown param1 mode")
			}
		}
		if (op >= 1 && op <= 2) || (op >= 5 && op <= 8) {
			if mode2 == 0 {
				param2 = mem[pc+2]
			} else if mode2 == 1 {
				param2 = pc + 2
			} else if mode2 == 2 {
				param2 = mem[pc+2] + relativeBase
			} else {
				panic("Unknown param2 mode")
			}
		}
		if (op >= 1 && op <= 2) || (op >= 7 && op <= 8) {
			if mode3 == 0 {
				param3 = mem[pc+3]
			} else if mode3 == 1 {
				param3 = pc + 3
			} else if mode3 == 2 {
				param3 = mem[pc+3] + relativeBase
			} else {
				panic("Unknown param3 mode")
			}
		}
		if op == 1 { // Add
			mem[param3] = mem[param1] + mem[param2]
			pc += 4
		} else if op == 2 { // Multiply
			mem[param3] = mem[param1] * mem[param2]
			pc += 4
		} else if op == 3 { // Input
			mem[param1] = io.Input()
			pc += 2
		} else if op == 4 { // Output
			io.Output(mem[param1])
			pc += 2
		} else if op == 5 { // Jump-if-true
			if mem[param1] != 0 {
				pc = mem[param2]
			} else {
				pc += 3
			}
		} else if op == 6 { // Jump-if-false
			if mem[param1] == 0 {
				pc = mem[param2]
			} else {
				pc += 3
			}
		} else if op == 7 { // Less than
			if mem[param1] < mem[param2] {
				mem[param3] = 1
			} else {
				mem[param3] = 0
			}
			pc += 4
		} else if op == 8 { // Equals
			if mem[param1] == mem[param2] {
				mem[param3] = 1
			} else {
				mem[param3] = 0
			}
			pc += 4
		} else if op == 9 { // Adjust relative base
			relativeBase += mem[param1]
			pc += 2
		} else if op == 99 {
			return mem
		} else {
			panic("Unknown opcode " + strconv.Itoa(mem[pc]))
		}
	}
}

type springdroid struct {
	input  chan int
	output chan int
}

func (a *springdroid) Input() int {
	return <-a.input
}

func (a *springdroid) Output(out int) {
	a.output <- out
}

func (a *springdroid) Send(cmd string) {
	for _, ch := range cmd {
		a.input <- int(ch)
	}
	a.input <- '\n'
}

func main() {
	var program []int

	reader, err := os.Open("day21.txt")
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

	var io springdroid
	io.input = make(chan int, 100)
	io.output = make(chan int)

	done := make(chan struct{})
	go func() {
		run(program, &io)
		close(done)
	}()

	// J = (!A || !B || !C) && D
	io.Send("NOT A T")
	io.Send("NOT B J")
	io.Send("OR T J")
	io.Send("NOT C T")
	io.Send("OR T J")
	io.Send("AND D J")
	io.Send("WALK")

	func() {
		for {
			select {
			case <-done:
				return
			case n := <-io.output:
				if n > 127 {
					fmt.Println("Part 1:", n)
				}
			}
		}
	}()

	done = make(chan struct{})
	go func() {
		run(program, &io)
		close(done)
	}()

	// J = (!A || !B || !C) && D && (E || H)
	io.Send("NOT A T")
	io.Send("NOT B J")
	io.Send("OR T J")
	io.Send("NOT C T")
	io.Send("OR T J")
	io.Send("AND D J")

	io.Send("NOT J T")
	io.Send("OR E T")
	io.Send("OR H T")
	io.Send("AND T J")
	io.Send("RUN")

	for {
		select {
		case <-done:
			return
		case n := <-io.output:
			if n > 127 {
				fmt.Println("Part 2:", n)
			}
		}
	}
}
