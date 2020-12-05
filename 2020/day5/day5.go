package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func decodeSeat(input string) (int, int) {
	row := 0
	col := 0
	for _, ch := range input {
		if ch == 'F' {
			row = row << 1
		} else if ch == 'B' {
			row = row << 1
			row |= 1
		} else if ch == 'L' {
			col = col << 1
		} else if ch == 'R' {
			col = col << 1
			col |= 1
		}
	}
	return row, col
}

func main() {
	var data []string

	reader, err := os.Open("day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			data = append(data, line)
		}
	}
	reader.Close()

	min := 1000000
	max := 0
	for _, seat := range data {
		row, col := decodeSeat(seat)
		id := row*8 + col
		if id > max {
			max = id
		}
		if id < min {
			min = id
		}
	}
	fmt.Println("Part 1:", max)

	fmt.Println("Total seats:", max-min+1)
	seats := make([]bool, max-min+1)
	for _, seat := range data {
		row, col := decodeSeat(seat)
		id := row*8 + col
		seats[id-min] = true
	}
	for i, filled := range seats {
		if !filled {
			fmt.Println("Part 2:", i+min)
			break
		}
	}
}
