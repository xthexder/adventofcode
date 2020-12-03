package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	width := 0
	var data [][]byte

	reader, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			width = len(line)
			data = append(data, []byte(line))
		}
	}
	reader.Close()

	slope := 3
	count := 0
	x := 0
	for y := 0; y < len(data); y++ {
		if data[y][x%width] == '#' {
			count++
		}
		x += slope
	}

	fmt.Println("Part 1:", count)

	slopesX := []int{1, 3, 5, 7, 1}
	slopesY := []int{1, 1, 1, 1, 2}
	slopeCount := make([]int, len(slopesX))
	for i := 0; i < len(slopesX); i++ {
		x := 0
		y := 0
		for y < len(data) {
			if data[y][x%width] == '#' {
				slopeCount[i]++
			}
			x += slopesX[i]
			y += slopesY[i]
		}
	}
	fmt.Println("Part 2:", slopeCount[0]*slopeCount[1]*slopeCount[2]*slopeCount[3]*slopeCount[4])
}
