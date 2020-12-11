package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PrintBoard(data [][]byte) {
	for y := 0; y < len(data); y++ {
		fmt.Println(string(data[y]))
	}
}

func main() {
	var dataInit [][]byte
	var dataPrev [][]byte
	var data [][]byte

	reader, err := os.Open("day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			dataInit = append(dataInit, []byte(line))
		}
	}
	reader.Close()

	for y := 0; y < len(dataInit); y++ {
		data = append(data, make([]byte, len(dataInit[y])))
		dataPrev = append(dataPrev, make([]byte, len(dataInit[y])))
		copy(data[y], dataInit[y])
		copy(dataPrev[y], dataInit[y])
	}

	changed := true
	for changed {
		changed = false
		for y := 0; y < len(data); y++ {
			for x := 0; x < len(data[y]); x++ {
				if dataPrev[y][x] == 'L' || dataPrev[y][x] == '#' {
					count := 0
					for dy := -1; dy <= 1; dy++ {
						for dx := -1; dx <= 1; dx++ {
							if dx == 0 && dy == 0 {
								continue
							}
							if x+dx >= 0 && x+dx < len(data[y]) && y+dy >= 0 && y+dy < len(data) {
								if dataPrev[y+dy][x+dx] == '#' {
									count++
								}
							}
						}
					}
					if dataPrev[y][x] == 'L' && count == 0 {
						data[y][x] = '#'
						changed = true
					} else if dataPrev[y][x] == '#' && count >= 4 {
						data[y][x] = 'L'
						changed = true
					}
				}
			}
		}

		// fmt.Println()
		// PrintBoard(data)

		for y := 0; y < len(data); y++ {
			copy(dataPrev[y], data[y])
		}
	}
	count := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == '#' {
				count++
			}
		}
	}
	fmt.Println("Part 1:", count)

	for y := 0; y < len(dataInit); y++ {
		copy(data[y], dataInit[y])
		copy(dataPrev[y], dataInit[y])
	}

	changed = true
	for changed {
		changed = false
		for y := 0; y < len(data); y++ {
			for x := 0; x < len(data[y]); x++ {
				if dataPrev[y][x] == 'L' || dataPrev[y][x] == '#' {
					count := 0
					for dy := -1; dy <= 1; dy++ {
						for dx := -1; dx <= 1; dx++ {
							if dx == 0 && dy == 0 {
								continue
							}
							x2 := x + dx
							y2 := y + dy
							for x2 >= 0 && x2 < len(data[y]) && y2 >= 0 && y2 < len(data) {
								if dataPrev[y2][x2] == '#' {
									count++
								}
								if dataPrev[y2][x2] != '.' {
									break
								}
								x2 += dx
								y2 += dy
							}
						}
					}
					if dataPrev[y][x] == 'L' && count == 0 {
						data[y][x] = '#'
						changed = true
					} else if dataPrev[y][x] == '#' && count >= 5 {
						data[y][x] = 'L'
						changed = true
					}
				}
			}
		}

		// fmt.Println()
		// PrintBoard(data)

		for y := 0; y < len(data); y++ {
			copy(dataPrev[y], data[y])
		}
	}
	count = 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == '#' {
				count++
			}
		}
	}
	fmt.Println("Part 2:", count)
}
