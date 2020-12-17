package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func printBoard(board [][][]byte) {
	found := false
	var minx, miny, minz int = -1, -1, -1
	var maxx, maxy, maxz int = -1, -1, -1
	for z := range board {
		for y := range board[z] {
			for x, value := range board[z][y] {
				if value == '#' {
					if minx < 0 || minx > x {
						minx = x
					}
					if maxx < x {
						maxx = x
					}

					if miny < 0 || miny > y {
						miny = y
					}
					if maxy < y {
						maxy = y
					}

					if minz < 0 || minz > z {
						minz = z
					}
					if maxz < z {
						maxz = z
					}
					found = true
				}
			}
		}
	}
	if found {
		for z := minz; z <= maxz; z++ {
			fmt.Println("z =", z-iterations)
			for y := miny; y <= maxy; y++ {
				fmt.Println(string(board[z][y][minx : maxx+1]))
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func advance(cur, prev [][][]byte) {
	for z := 1; z < len(prev)-1; z++ {
		for y := 1; y < len(prev[z])-1; y++ {
			for x := 1; x < len(prev[z][y])-1; x++ {
				neighbors := 0
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						for dz := -1; dz <= 1; dz++ {
							if dx != 0 || dy != 0 || dz != 0 {
								if prev[z+dz][y+dy][x+dx] == '#' {
									neighbors++
								}
							}
						}
					}
				}
				if prev[z][y][x] == '#' {
					if neighbors == 2 || neighbors == 3 {
						cur[z][y][x] = '#'
					} else {
						cur[z][y][x] = '.'
					}
				} else {
					if neighbors == 3 {
						cur[z][y][x] = '#'
					} else {
						cur[z][y][x] = '.'
					}
				}
			}
		}
	}
}

func advance2(cur, prev [][][][]byte) {
	for w := 1; w < len(prev)-1; w++ {
		for z := 1; z < len(prev[w])-1; z++ {
			for y := 1; y < len(prev[w][z])-1; y++ {
				for x := 1; x < len(prev[w][z][y])-1; x++ {
					neighbors := 0
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							for dz := -1; dz <= 1; dz++ {
								for dw := -1; dw <= 1; dw++ {
									if dx != 0 || dy != 0 || dz != 0 || dw != 0 {
										if prev[w+dw][z+dz][y+dy][x+dx] == '#' {
											neighbors++
										}
									}
								}
							}
						}
					}
					if prev[w][z][y][x] == '#' {
						if neighbors == 2 || neighbors == 3 {
							cur[w][z][y][x] = '#'
						} else {
							cur[w][z][y][x] = '.'
						}
					} else {
						if neighbors == 3 {
							cur[w][z][y][x] = '#'
						} else {
							cur[w][z][y][x] = '.'
						}
					}
				}
			}
		}
	}
}

const iterations = 6

func main() {
	var init [][]byte

	reader, err := os.Open("day17.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			init = append(init, []byte(line))
		}
	}
	reader.Close()

	var boardA, boardB [][][]byte
	boardA = make([][][]byte, 1+iterations*4)
	boardB = make([][][]byte, 1+iterations*4)
	for z := range boardA {
		boardA[z] = make([][]byte, len(init)+iterations*4)
		boardB[z] = make([][]byte, len(init)+iterations*4)
		for y := range boardA[z] {
			boardA[z][y] = make([]byte, len(init[0])+iterations*4)
			boardB[z][y] = make([]byte, len(init[0])+iterations*4)
			for x := range boardA[z][y] {
				boardA[z][y][x] = '.'
				boardB[z][y][x] = '.'
			}
		}
	}
	for y := range init {
		for x, value := range init[y] {
			boardA[iterations*2][y+iterations*2][x+iterations*2] = value
		}
	}

	// fmt.Println("Init")
	// printBoard(boardA)
	for cycle := 0; cycle < iterations; cycle++ {
		// fmt.Println("Cycle", cycle)
		if cycle&1 == 0 {
			advance(boardB, boardA)
			// printBoard(boardB)
		} else {
			advance(boardA, boardB)
			// printBoard(boardA)
		}
	}
	lastBoard := boardB
	if iterations&1 == 0 {
		lastBoard = boardA
	}

	total := 0
	for z := range lastBoard {
		for y := range lastBoard[z] {
			for _, value := range lastBoard[z][y] {
				if value == '#' {
					total++
				}
			}
		}
	}
	fmt.Println("Part 1:", total)

	var boardA2, boardB2 [][][][]byte
	boardA2 = make([][][][]byte, 1+iterations*4)
	boardB2 = make([][][][]byte, 1+iterations*4)
	for w := range boardA2 {
		boardA2[w] = make([][][]byte, 1+iterations*4)
		boardB2[w] = make([][][]byte, 1+iterations*4)
		for z := range boardA2[w] {
			boardA2[w][z] = make([][]byte, len(init)+iterations*4)
			boardB2[w][z] = make([][]byte, len(init)+iterations*4)
			for y := range boardA2[w][z] {
				boardA2[w][z][y] = make([]byte, len(init[0])+iterations*4)
				boardB2[w][z][y] = make([]byte, len(init[0])+iterations*4)
				for x := range boardA2[w][z][y] {
					boardA2[w][z][y][x] = '.'
					boardB2[w][z][y][x] = '.'
				}
			}
		}
	}
	for y := range init {
		for x, value := range init[y] {
			boardA2[iterations*2][iterations*2][y+iterations*2][x+iterations*2] = value
		}
	}

	for cycle := 0; cycle < iterations; cycle++ {
		if cycle&1 == 0 {
			advance2(boardB2, boardA2)
		} else {
			advance2(boardA2, boardB2)
		}
	}
	lastBoard2 := boardB2
	if iterations&1 == 0 {
		lastBoard2 = boardA2
	}

	total = 0
	for w := range lastBoard2 {
		for z := range lastBoard2[w] {
			for y := range lastBoard2[w][z] {
				for _, value := range lastBoard2[w][z][y] {
					if value == '#' {
						total++
					}
				}
			}
		}
	}
	fmt.Println("Part 2:", total)
}
