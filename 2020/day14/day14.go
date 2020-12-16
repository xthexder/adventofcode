package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type mask [2]uint64

func newMask(text string) mask {
	m := mask{0xFFFFFFFFF, 0}
	for i, bit := range text {
		if bit == '1' || bit == '0' {
			if bit == '1' {
				m[1] |= 1 << (len(text) - i - 1)
			}
			m[0] ^= 1 << (len(text) - i - 1)
		}
	}
	return m
}

func (m mask) Apply(value uint64) uint64 {
	return (value & m[0]) | m[1]
}

func (m mask) ApplyAddr(addr uint64) [2]uint64 {
	return [2]uint64{m[0], addr | m[0] | m[1]}
}

func main() {
	var input [][]string

	reader, err := os.Open("day14.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return false
			case r >= 'a' && r <= 'z':
				return false
			case r == 'X':
				return false
			default:
				return true
			}
		})
		if len(fields) > 0 {
			input = append(input, fields)
		}
	}
	reader.Close()

	activeMask := mask{^uint64(0), 0}
	mem := make(map[uint64]uint64)
	for _, fields := range input {
		if fields[0] == "mask" {
			activeMask = newMask(fields[1])
		} else if fields[0] == "mem" {
			addr, _ := strconv.ParseUint(fields[1], 10, 64)
			value, _ := strconv.ParseUint(fields[2], 10, 64)
			mem[addr] = activeMask.Apply(value)
		}
	}

	var sum uint64 = 0
	for _, value := range mem {
		sum += value
	}
	fmt.Println("Part 1:", sum)

	activeMask = mask{^uint64(0), 0}
	var data [][3]uint64
	for _, fields := range input {
		if fields[0] == "mask" {
			activeMask = newMask(fields[1])
		} else if fields[0] == "mem" {
			addr, _ := strconv.ParseUint(fields[1], 10, 64)
			value, _ := strconv.ParseUint(fields[2], 10, 64)
			fixedAddr := activeMask.ApplyAddr(addr)
			data = append(data, [3]uint64{fixedAddr[0], fixedAddr[1], value})
		}
	}

	mem = make(map[uint64]uint64)
	for _, instr := range data {
		maxCount := uint64(1 << bits.OnesCount64(instr[0]))
		var count uint64
		for count = 0; count < maxCount; count++ {
			bitmap := count
			for i := 0; i < 36; i++ {
				if instr[0]&(1<<i) != 0 {
					if bitmap&1 == 0 {
						instr[1] &= ^(1 << i)
					} else {
						instr[1] |= (1 << i)
					}
					bitmap = bitmap >> 1
				}
			}
			mem[instr[1]] = instr[2]
		}
	}

	sum = 0
	for _, value := range mem {
		sum += value
	}
	fmt.Println("Part 2:", sum)
}
