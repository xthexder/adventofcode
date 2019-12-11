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
	mem := make([]int, len(program)+1000)
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

type robot struct {
	input  chan int
	output chan int
}

func (a *robot) Input() int {
	return <-a.input
}

func (a *robot) Output(out int) {
	a.output <- out
}

func main() {
	var program []int

	reader, err := os.Open("day11.txt")
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

	for part := 0; part < 2; part++ {
		fmt.Println("Part", part+1)
		panel := make([][]byte, 1000)
		for y := 0; y < len(panel); y++ {
			panel[y] = make([]byte, 1000)
		}
		robotPos := 500 + 500i
		robotDir := -1i
		if part > 0 {
			panel[500][500] = '#'
		}

		minX := 1000
		minY := 1000
		maxX := 0
		maxY := 0

		var r robot
		r.input = make(chan int)
		r.output = make(chan int)
		done := make(chan struct{})
		go func() {
			run(program, &r)
			close(done)
		}()
		func() {
			for {
				x := int(real(robotPos))
				y := int(imag(robotPos))
				if panel[y][x] == '#' {
					select {
					case <-done:
						return
					case r.input <- 1:
					}
				} else {
					select {
					case <-done:
						return
					case r.input <- 0:
					}
				}
				paint := <-r.output
				if paint == 0 {
					if part == 0 {
						panel[y][x] = '.'
					} else {
						panel[y][x] = ' ' // Space makes the text easier to read
					}
				} else {
					panel[y][x] = '#'
				}
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
				dir := <-r.output
				if dir == 0 { // Left
					robotDir *= 1i
				} else { // Right
					robotDir *= -1i
				}
				robotPos += robotDir
			}
		}()
		count := 0
		for y, line := range panel {
			if y >= minY && y <= maxY {
				for x := minX; x <= maxX; x++ {
					if line[x] == 0 {
						line[x] = ' '
					} else {
						count++
					}
				}
				for x := maxX; x >= minX; x-- {
					fmt.Print(string(line[x]))
				}
				fmt.Println()
			}
		}
		fmt.Println("Count:", count)
	}
}
