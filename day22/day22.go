package main

import "fmt"

const depth = 5913
const targetX, targetY = 8, 701

const (
	ROCKY int = iota
	WET
	NARROW
)

const (
	NOTHING int = iota
	TORCH
	CLIMBING_GEAR
)

var tileChars = []string{".", "=", "|", "#"}
var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func erosion(index int) int {
	return (index + depth) % 20183
}

type Point [2]int

func (p Point) add(a Point) Point {
	return Point{p[0] + a[0], p[1] + a[1]}
}

var weights [depth][depth][3]int
var geologicIndex [depth][depth]int
var erosionLevel [depth][depth]int

type Entry struct {
	pos      Point
	gear     int
	previous *Entry
}

func (e *Entry) Weight() *int {
	return &weights[e.pos[0]][e.pos[1]][e.gear]
}

func printBoard() {
	for y := 0; y <= targetY+10; y++ {
		for x := 0; x <= targetX+50; x++ {
			if x == 0 && y == 0 {
				fmt.Print("M")
			} else if x == targetX && y == targetY {
				fmt.Print("T")
			} else {
				fmt.Print(tileChars[erosionLevel[x][y]])
			}
		}
		fmt.Println()
	}
}

func main() {
	for x := 0; x < depth; x++ {
		geologicIndex[x][0] = x * 16807
		erosionLevel[x][0] = erosion(geologicIndex[x][0])
		for y := 1; y < depth; y++ {
			if x == 0 {
				geologicIndex[0][y] = y * 48271
				erosionLevel[0][y] = erosion(geologicIndex[0][y])
			} else if x == targetX && y == targetY {
				erosionLevel[x][y] = erosion(0)
			} else {
				geologicIndex[x][y] = erosionLevel[x-1][y] * erosionLevel[x][y-1]
				erosionLevel[x][y] = erosion(geologicIndex[x][y])
			}
		}
	}

	var riskLevel int = 0
	for x := 0; x < depth; x++ {
		for y := 0; y < depth; y++ {
			erosionLevel[x][y] %= 3
		}
	}
	for x := 0; x <= targetX; x++ {
		for y := 0; y <= targetY; y++ {
			riskLevel += erosionLevel[x][y]
		}
	}
	fmt.Println("Part A:", riskLevel)

	weights[0][0][TORCH] = 1
	tiles := erosionLevel
	queue := make([]*Entry, 1, 10000)
	queue[0] = &Entry{Point{0, 0}, TORCH, nil}

	for len(queue) > 0 {
		min := -1
		for _, entry := range queue {
			if min < 0 || *entry.Weight() < min {
				min = *entry.Weight()
			}
		}
		for {
			startLen := len(queue)
			for _, entry := range queue {
				if *entry.Weight() == min {
					if entry.pos[0] == targetX && entry.pos[1] == targetY && entry.gear == TORCH {
						for e := entry; e != nil; e = e.previous {
							erosionLevel[e.pos[0]][e.pos[1]] = 3
						}
						printBoard()
						fmt.Println("Part B:", *entry.Weight()-1)
						return
					}

					for _, dir := range directions {
						pos := entry.pos.add(dir)
						if pos[0] < 0 || pos[1] < 0 {
							continue
						}
						if entry.gear != tiles[pos[0]][pos[1]] {
							// Don't need to change gear
							posWeight := &weights[pos[0]][pos[1]][entry.gear]
							if *posWeight == 0 || *posWeight > *entry.Weight()+1 {
								*posWeight = *entry.Weight() + 1
								queue = append(queue, &Entry{pos, entry.gear, entry})
							}
						}
					}
					gear := 3 ^ (entry.gear ^ tiles[entry.pos[0]][entry.pos[1]])
					posWeight := &weights[entry.pos[0]][entry.pos[1]][gear]
					if *posWeight == 0 || *posWeight > *entry.Weight()+7 {
						*posWeight = *entry.Weight() + 7
						queue = append(queue, &Entry{entry.pos, gear, entry})
					}
				}
			}
			if len(queue) == startLen {
				j := 0
				for i, entry := range queue[:] {
					if *entry.Weight() == min {
						queue[j], queue[i] = queue[i], queue[j]
						j++
					}
				}
				queue = queue[j:]
				break
			}
		}
	}
}
