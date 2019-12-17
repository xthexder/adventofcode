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

func read(view [][]byte, x, y int) byte {
	if x < 0 || y < 0 || y >= len(view) || x >= len(view[y]) {
		return '.'
	}
	return view[y][x]
}

func splitPath(path string) (string, string, string, string) {
	for i := 1; i < 20; i++ {
		if i >= len(path) || path[i] != ',' {
			continue
		}
		a := path[:i]
		for offsetJ := i + 1; offsetJ < len(path); offsetJ++ {
			if path[offsetJ-1] != ',' {
				continue
			}
			for j := 1; j < 20; j++ {
				if offsetJ+j >= len(path) || path[offsetJ+j] != ',' {
					continue
				}
				b := path[offsetJ : offsetJ+j]
				for offsetK := offsetJ + j + 1; offsetK < len(path); offsetK++ {
					if path[offsetK-1] != ',' {
						continue
					}
					for k := 1; k < 20; k++ {
						if offsetK+k >= len(path) || path[offsetK+k] != ',' {
							continue
						}
						c := path[offsetK : offsetK+k]

						main := strings.ReplaceAll(path, a, "A")
						main = strings.ReplaceAll(main, b, "B")
						main = strings.ReplaceAll(main, c, "C")

						if !strings.ContainsAny(main, "LR0123456789") {
							return main, a, b, c
						}
					}
				}
			}
		}
	}
	return "", "", "", ""
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

	posX, posY := 0, 0
	dirX, dirY := 0, 0
	total := 0
	for y, line := range view {
		if y > 0 && y < len(view)-1 {
			for x, ch := range line {
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
				if ch == '^' {
					posX, posY = x, y
					dirY = -1
				} else if ch == '>' {
					posX, posY = x, y
					dirX = 1
				} else if ch == 'v' {
					posX, posY = x, y
					dirY = 1
				} else if ch == '<' {
					posX, posY = x, y
					dirX = -1
				}
			}
		}
	}
	fmt.Println("Part 1:", total)

	path := []string{}
	for {
		count := 0
		for read(view, posX+dirX, posY+dirY) == '#' {
			count++
			posX += dirX
			posY += dirY
		}
		if count > 0 {
			path = append(path, strconv.Itoa(count))
		}

		left := complex(float32(dirX), float32(dirY)) * -1i
		right := complex(float32(dirX), float32(dirY)) * 1i
		if read(view, posX+int(real(left)), posY+int(imag(left))) == '#' {
			dirX = int(real(left))
			dirY = int(imag(left))
			path = append(path, "L")
		} else if read(view, posX+int(real(right)), posY+int(imag(right))) == '#' {
			dirX = int(real(right))
			dirY = int(imag(right))
			path = append(path, "R")
		} else {
			break
		}
	}
	main, a, b, c := splitPath(strings.Join(path, ","))
	// fmt.Println("Path:", strings.Join(path, ","))
	// fmt.Println("Main:", main)
	// fmt.Println("A:", a)
	// fmt.Println("B:", b)
	// fmt.Println("C:", c)

	program[0] = 2
	done = make(chan struct{})
	go func() {
		run(program, &io)
		close(done)
	}()

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
	}()

	for {
		select {
		case <-done:
			return
		case ch := <-io.output:
			if ch > 127 {
				fmt.Println("Part 2:", ch)
				return
			}
		}
	}
}
