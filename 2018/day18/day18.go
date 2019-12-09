package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countAdjacent(x, y int, grid [][]byte, search byte) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if y+dy < 0 || y+dy >= len(grid) {
				continue
			}
			if x+dx < 0 || x+dx >= len(grid[y]) {
				continue
			}
			if grid[y+dy][x+dx] == search {
				count++
			}
		}
	}
	return count
}

func main() {
	var data [2][][]byte

	reader, err := os.Open("day18.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) > 0 {
			bufA := make([]byte, len(line))
			bufB := make([]byte, len(line))
			copy(bufA, line)
			copy(bufB, line)
			data[0] = append(data[0], bufA)
			data[1] = append(data[1], bufB)
		}
	}
	reader.Close()

	minute := 0
	var counts []int
	for minute < 1000 {
		grid := data[minute%2]
		gridOut := data[(minute+1)%2]
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				switch grid[y][x] {
				case '.':
					if countAdjacent(x, y, grid, '|') >= 3 {
						gridOut[y][x] = '|'
					} else {
						gridOut[y][x] = '.'
					}
				case '|':
					if countAdjacent(x, y, grid, '#') >= 3 {
						gridOut[y][x] = '#'
					} else {
						gridOut[y][x] = '|'
					}
				case '#':
					lumber := countAdjacent(x, y, grid, '#')
					trees := countAdjacent(x, y, grid, '|')
					if lumber < 1 || trees < 1 {
						gridOut[y][x] = '.'
					} else {
						gridOut[y][x] = '#'
					}
				}
			}
		}
		minute++

		treeCount, lumberCount := 0, 0
		for y := 0; y < len(gridOut); y++ {
			// fmt.Println(string(gridOut[y]))
			for x := 0; x < len(gridOut[y]); x++ {
				switch gridOut[y][x] {
				case '|':
					treeCount++
				case '#':
					lumberCount++
				}
			}
		}
		if minute == 10 {
			fmt.Println("Part A:", treeCount*lumberCount)
		}
		counts = append(counts, treeCount*lumberCount)
	}

	for i := len(counts) - 2; i >= 0; i-- {
		if counts[i] == counts[len(counts)-1] {
			matches := true
			for j := 0; j < len(counts)-i; j++ {
				if counts[i-j] != counts[len(counts)-j-1] {
					matches = false
					break
				}
			}
			if matches {
				fmt.Println("Part B:", counts[i+(1000000000-minute)%(len(counts)-i-1)])
				return
			}
		}
	}
}
