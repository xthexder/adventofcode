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

type camera struct {
	input  chan int
	output chan int
}

func (a *camera) Input() int {
	return <-a.input
}

func (a *camera) Output(out int) {
	a.output <- out
}

func main() {
	var program []int

	reader, err := os.Open("day17.txt")
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

	var io camera
	io.input = make(chan int)
	io.output = make(chan int)
	view := make([][]byte, 50)
	for i := range view {
		view[i] = make([]byte, 50)
	}
	done := make(chan struct{})
	go func() {
		run(program, &io)
		close(done)
	}()

	func() {
		x := 0
		y := 0
		for {
			select {
			case <-done:
				return
			case ch := <-io.output:
				if ch == '\n' {
					x = 0
					y++
				} else {
					view[y][x] = byte(ch)
					x++
				}
			}
		}
	}()

	for _, line := range view {
		fmt.Println(string(line))
	}

	// var pos complex64 = 0
	// var dir complex64 = 0
	total := 0
	for y, line := range view {
		if y > 0 && y < len(view)-1 {
			for x := range line {
				if x > 0 && x < len(line)-1 {
					intersection := true
					for i := -1; i < 2; i++ {
						if view[y+i][x] != '#' || view[y][x+i] != '#' {
							intersection = false
						}
					}
					if intersection {
						total += x * y
					}
				}
				// if ch == '^' {
				// 	pos = complex(float32(x), float32(y))
				// 	dir = -1i
				// } else if ch == '>' {
				// 	pos = complex(float32(x), float32(y))
				// 	dir = 1
				// } else if ch == 'v' {
				// 	pos = complex(float32(x), float32(y))
				// 	dir = 1i
				// } else if ch == '<' {
				// 	pos = complex(float32(x), float32(y))
				// 	dir = -1
				// }
			}
		}
	}
	fmt.Println("Part 1:", total)

	// path := ""
	// for {
	// 	frontX := int(real(pos+dir))
	// 	frontY := int(imag(pos+dir))
	// 	line, ok := view[frontY]
	// 	if !ok {
	// 		line = []byte{}
	// 	}
	// 	if front, ok := line[frontX]; ok && front == '#' {
	// 		pos +=
	// 	}

	// 	if view[frontY][frontX]
	// }

	program[0] = 2
	done = make(chan struct{})
	go func() {
		run(program, &io)
		close(done)
	}()

	// Counted manually, lol
	// R,8,L,10,L,12,R,4,R,8,L,12,R,4,R,4,R,8,L,10,L,12,R,4,R,8,L,10,R,8,R,8,L,10,L,12,R,4,R,8,L,12,R,4,R,4,R,8,L,10,R,8,R,8,L,12,R,4,R,4,R,8,L,10,R,8,R,8,L,12,R,4,R,4
	main := "B,A,B,C,B,A,C,A,C,A"
	a := "R,8,L,12,R,4,R,4"
	b := "R,8,L,10,L,12,R,4"
	c := "R,8,L,10,R,8"

	go func() {
		for _, ch := range main {
			io.input <- int(ch)
		}
		io.input <- '\n'

		for _, ch := range a {
			io.input <- int(ch)
		}
		io.input <- '\n'

		for _, ch := range b {
			io.input <- int(ch)
		}
		io.input <- '\n'

		for _, ch := range c {
			io.input <- int(ch)
		}
		io.input <- '\n'

		io.input <- int('n')
		io.input <- '\n'
		fmt.Println("Done input")
	}()

	// fmt.Println("Part 2:", <-io.output)
	for {
		select {
		case <-done:
			return
		case ch := <-io.output:
			if ch > 127 {
				fmt.Println(ch)
			}
		}
	}
}
