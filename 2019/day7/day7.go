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
	mem := make([]int, len(program))
	copy(mem, program)

	pc := 0
	for {
		op := mem[pc] % 100
		mode1 := (mem[pc] / 100) % 10
		mode2 := (mem[pc] / 1000) % 10
		param1 := 0
		param2 := 0
		if op >= 1 && op <= 8 {
			if mode1 == 0 {
				param1 = mem[mem[pc+1]]
			} else if mode1 == 1 {
				param1 = mem[pc+1]
			} else {
				panic("Unknown param1 mode")
			}
		}
		if (op >= 1 && op <= 2) || (op >= 5 && op <= 8) {
			if mode2 == 0 {
				param2 = mem[mem[pc+2]]
			} else if mode2 == 1 {
				param2 = mem[pc+2]
			} else {
				panic("Unknown param2 mode")
			}
		}
		if op == 1 { // Add
			mem[mem[pc+3]] = param1 + param2
			pc += 4
		} else if op == 2 { // Multiply
			mem[mem[pc+3]] = param1 * param2
			pc += 4
		} else if op == 3 { // Input
			mem[mem[pc+1]] = io.Input()
			pc += 2
		} else if op == 4 { // Output
			io.Output(param1)
			pc += 2
		} else if op == 5 { // Jump-if-true
			if param1 != 0 {
				pc = param2
			} else {
				pc += 3
			}
		} else if op == 6 { // Jump-if-false
			if param1 == 0 {
				pc = param2
			} else {
				pc += 3
			}
		} else if op == 7 { // Less than
			if param1 < param2 {
				mem[mem[pc+3]] = 1
			} else {
				mem[mem[pc+3]] = 0
			}
			pc += 4
		} else if op == 8 { // Equals
			if param1 == param2 {
				mem[mem[pc+3]] = 1
			} else {
				mem[mem[pc+3]] = 0
			}
			pc += 4
		} else if op == 99 {
			return mem
		} else {
			panic("Unknown opcode " + strconv.Itoa(mem[pc]))
		}
	}
}

type amplifier struct {
	phaseDone bool
	phase     int
	input     int
	output    int
}

func (a *amplifier) Input() int {
	if !a.phaseDone {
		a.phaseDone = true
		return a.phase
	}
	return a.input
}

func (a *amplifier) Output(out int) {
	a.output = out
}

type loopAmplifier struct {
	input  chan int
	output chan int
}

func (a *loopAmplifier) Input() int {
	return <-a.input
}

func (a *loopAmplifier) Output(out int) {
	a.output <- out
}

func permute(phases []int, i int, callback func([]int)) {
	if i == len(phases) {
		callback(phases)
		return
	}
	for j := i; j < len(phases); j++ {
		phases[i], phases[j] = phases[j], phases[i]
		permute(phases, i+1, callback)
		phases[i], phases[j] = phases[j], phases[i]
	}
}

func main() {
	var program []int

	reader, err := os.Open("day7.txt")
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

	phases := []int{0, 1, 2, 3, 4}
	best := 0

	permute(phases, 0, func(ph []int) {
		var amp amplifier
		for i := 0; i < 5; i++ {
			amp.phase = ph[i]
			amp.phaseDone = false
			run(program, &amp)
			amp.input = amp.output
		}
		if amp.output > best {
			best = amp.output
		}
	})
	fmt.Println("Part 1:", best)

	phases = []int{5, 6, 7, 8, 9}
	best = 0

	amps := make([]loopAmplifier, len(phases))

	permute(phases, 0, func(ph []int) {
		for i := 0; i < 5; i++ {
			amps[i].output = make(chan int)
			amps[(i+1)%5].input = amps[i].output
		}
		go run(program, &amps[0])
		go run(program, &amps[1])
		go run(program, &amps[2])
		done := make(chan struct{})
		go func() {
			run(program, &amps[3])
			close(done)
		}()
		go run(program, &amps[4])
		for i := 0; i < 5; i++ {
			amps[i].input <- ph[i]
		}
		amps[0].input <- 0
		<-done
		output := <-amps[4].output
		if output > best {
			best = output
		}
	})
	fmt.Println("Part 2:", best)
}
