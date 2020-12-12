package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type instruction struct {
	op    byte
	value int
}

func (instr *instruction) Exec(pos *[2]int, dir *complex64) {
	switch instr.op {
	case 'N':
		pos[1] -= instr.value
		// *dir = -1i
	case 'S':
		pos[1] += instr.value
		// *dir = 1i
	case 'E':
		pos[0] += instr.value
		// *dir = 1
	case 'W':
		pos[0] -= instr.value
		// *dir = -1
	case 'L', 'R':
		if instr.value == 180 {
			*dir *= -1
		} else {
			if instr.value == 270 {
				instr.op ^= 'L' ^ 'R'
			}
			if instr.op == 'L' {
				*dir *= -1i
			} else {
				*dir *= 1i
			}
		}
	case 'F':
		pos[0] += int(real(*dir)) * instr.value
		pos[1] += int(imag(*dir)) * instr.value
	}
	// fmt.Println(string(instr.op), instr.value, *pos, *dir)
}

func (instr *instruction) Exec2(pos *[2]int, waypoint *complex64) {
	switch instr.op {
	case 'N':
		*waypoint -= complex(float32(0), float32(instr.value))
	case 'S':
		*waypoint += complex(float32(0), float32(instr.value))
	case 'E':
		*waypoint += complex(float32(instr.value), float32(0))
	case 'W':
		*waypoint -= complex(float32(instr.value), float32(0))
	case 'L', 'R':
		if instr.value == 180 {
			*waypoint *= -1
		} else {
			if instr.value == 270 {
				instr.op ^= 'L' ^ 'R'
			}
			if instr.op == 'L' {
				*waypoint *= -1i
			} else {
				*waypoint *= 1i
			}
		}
	case 'F':
		pos[0] += int(real(*waypoint)) * instr.value
		pos[1] += int(imag(*waypoint)) * instr.value
	}
	// fmt.Println(string(instr.op), instr.value, *pos, *waypoint)
}

func main() {
	var data []instruction

	reader, err := os.Open("day12.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			value, _ := strconv.Atoi(line[1:])
			data = append(data, instruction{line[0], value})
		}
	}
	reader.Close()

	var pos [2]int
	var dir complex64 = 1

	for _, instr := range data {
		instr.Exec(&pos, &dir)
	}

	fmt.Println("Part 1:", math.Abs(float64(pos[0]))+math.Abs(float64(pos[1])))

	pos = [2]int{0, 0}
	dir = complex(float32(10), float32(-1))

	for _, instr := range data {
		instr.Exec2(&pos, &dir)
	}

	fmt.Println("Part 2:", math.Abs(float64(pos[0]))+math.Abs(float64(pos[1])))
}
