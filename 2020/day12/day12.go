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
	value float32
}

func (instr *instruction) Exec(pos, dir *complex64) {
	switch instr.op {
	case 'N':
		*pos -= complex(0, instr.value)
	case 'S':
		*pos += complex(0, instr.value)
	case 'E':
		*pos += complex(instr.value, 0)
	case 'W':
		*pos -= complex(instr.value, 0)
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
		*pos += *dir * complex(instr.value, 0)
	}
	fmt.Println(string(instr.op), instr.value, *pos, *dir)
}

func (instr *instruction) Exec2(pos, waypoint *complex64) {
	switch instr.op {
	case 'N':
		*waypoint -= complex(0, instr.value)
	case 'S':
		*waypoint += complex(0, instr.value)
	case 'E':
		*waypoint += complex(instr.value, 0)
	case 'W':
		*waypoint -= complex(instr.value, 0)
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
		*pos += *waypoint * complex(instr.value, 0)
	}
	fmt.Println(string(instr.op), instr.value, *pos, *waypoint)
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
			data = append(data, instruction{line[0], float32(value)})
		}
	}
	reader.Close()

	var pos complex64
	var dir complex64 = 1

	for _, instr := range data {
		instr.Exec(&pos, &dir)
	}

	fmt.Println("Part 1:", math.Abs(float64(real(pos)))+math.Abs(float64(imag(pos))))

	pos = 0
	dir = 10 - 1i

	for _, instr := range data {
		instr.Exec2(&pos, &dir)
	}

	fmt.Println("Part 2:", math.Abs(float64(real(pos)))+math.Abs(float64(imag(pos))))
}
