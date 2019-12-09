package main

import "fmt"

const depth = 5913
const targetX, targetY = 8, 701
const buffer = 150

const (
	NOTHING int = iota
	TORCH
	CLIMBING_GEAR
)

var tileChars = []string{".", "=", "|", "#"}
var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

var weights [targetX + buffer][targetY + buffer][3]int

type Entry struct {
	x, y     int
	gear     int
	previous *Entry
}

func (e *Entry) Weight() int {
	return weights[e.x][e.y][e.gear]
}

func main() {
	var geologicIndex [targetX + buffer][targetY + buffer]int
	var erosionLevel [targetX + buffer][targetY + buffer]int

	for x := 0; x < len(geologicIndex); x++ {
		geologicIndex[x][0] = x * 16807
		erosionLevel[x][0] = (geologicIndex[x][0] + depth) % 20183
		for y := 1; y < len(geologicIndex[x]); y++ {
			if x == 0 {
				geologicIndex[0][y] = y * 48271
				erosionLevel[0][y] = (geologicIndex[0][y] + depth) % 20183
			} else if x == targetX && y == targetY {
				erosionLevel[x][y] = depth % 20183
			} else {
				geologicIndex[x][y] = erosionLevel[x-1][y] * erosionLevel[x][y-1]
				erosionLevel[x][y] = (geologicIndex[x][y] + depth) % 20183
			}
		}
	}

	var riskLevel int = 0
	for x := 0; x < len(erosionLevel); x++ {
		for y := 0; y < len(erosionLevel[x]); y++ {
			erosionLevel[x][y] %= 3
			if x <= targetX && y <= targetY {
				riskLevel += erosionLevel[x][y]
			}
		}
	}
	fmt.Println("Part A:", riskLevel)

	tiles := erosionLevel
	queue := []*Entry{
		&Entry{0, 0, TORCH, nil},
	}
	weights[0][0][TORCH] = 1

	for len(queue) > 0 {
		// Find the minimum weight
		min := -1
		for _, entry := range queue {
			if min < 0 || entry.Weight() < min {
				min = entry.Weight()
			}
		}

		minCount := 0
		for i, entry := range queue {
			if entry.Weight() == min {
				if entry.x == targetX && entry.y == targetY && entry.gear == TORCH {
					// Print the path taken
					// for e := entry; e != nil; e = e.previous {
					// 	tiles[e.x][e.y] = 3
					// }
					// for y := 0; y < len(tiles[0]); y++ {
					// 	for x := 0; x < len(tiles); x++ {
					// 		if x == 0 && y == 0 {
					// 			fmt.Print("M")
					// 		} else if x == targetX && y == targetY {
					// 			fmt.Print("T")
					// 		} else {
					// 			fmt.Print(tileChars[tiles[x][y]])
					// 		}
					// 	}
					// 	fmt.Println()
					// }
					fmt.Println("Part B:", entry.Weight()-1)
					return
				}

				for _, dir := range directions {
					x := entry.x + dir[0]
					y := entry.y + dir[1]
					if x < 0 || y < 0 {
						continue
					} else if x >= len(tiles) || y >= len(tiles[x]) {
						continue
					}

					if entry.gear != tiles[x][y] {
						// Current gear can be used on this tile
						posWeight := &weights[x][y][entry.gear]
						if *posWeight == 0 || *posWeight > entry.Weight()+1 {
							*posWeight = entry.Weight() + 1
							queue = append(queue, &Entry{x, y, entry.gear, entry})
						}
					}
				}
				// Try switching gear
				gear := 3 ^ (entry.gear ^ tiles[entry.x][entry.y])
				posWeight := &weights[entry.x][entry.y][gear]
				if *posWeight == 0 || *posWeight > entry.Weight()+7 {
					*posWeight = entry.Weight() + 7
					queue = append(queue, &Entry{entry.x, entry.y, gear, entry})
				}

				// Move the processed entries to the beginning so they can be quickly removed
				queue[minCount], queue[i] = queue[i], queue[minCount]
				minCount++
			}
		}
		queue = queue[minCount:]
	}
}
