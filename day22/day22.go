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

var tileChars = []string{".", "=", "|"}
var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func erosion(index int) int {
	return (index + depth) % 20183
}

type Point [2]int

var weights [depth][depth][3]int

type Entry [3]int // x, y, gear

func (e Entry) add(a [2]int) [2]int {
	return Point{e[0] + a[0], e[1] + a[1]}
}

func (e Entry) Weight() *int {
	return &weights[e[0]][e[1]][e[2]]
}

func main() {
	var geologicIndex [depth][depth]int
	var erosionLevel [depth][depth]int

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
	for y := 0; y <= targetY; y++ {
		for x := 0; x <= targetX; x++ {
			riskLevel += erosionLevel[x][y]
			// if x == 0 && y == 0 {
			// 	fmt.Print("M")
			// } else if x == targetX && y == targetX {
			// 	fmt.Print("T")
			// } else {
			// 	fmt.Print(tileChars[erosionLevel[x][y]])
			// }
		}
		// fmt.Println()
	}
	fmt.Println("Part A:", riskLevel)

	tiles := erosionLevel
	visited := make(map[Entry]struct{})
	unvisited := make(map[Entry]struct{})

	visited[Entry{0, 0, TORCH}] = struct{}{}
	unvisited[Entry{0, 0, TORCH}] = struct{}{}
	weights[0][0][TORCH] = 1

	for len(unvisited) > 0 {
		for entry, _ := range visited {
			for _, dir := range directions {
				pos := entry.add(dir)
				if pos[0] < 0 || pos[1] < 0 {
					continue
				}
				if entry[2] != tiles[pos[0]][pos[1]] {
					// Don't need to change gear
					posWeight := &weights[pos[0]][pos[1]][entry[2]]
					if *posWeight == 0 || *posWeight > *entry.Weight()+1 {
						*posWeight = *entry.Weight() + 1
						unvisited[Entry{pos[0], pos[1], entry[2]}] = struct{}{}
					}
				}
			}
			gear := 3 ^ (entry[2] ^ tiles[entry[0]][entry[1]])
			posWeight := &weights[entry[0]][entry[1]][gear]
			if *posWeight == 0 || *posWeight > *entry.Weight()+7 {
				*posWeight = *entry.Weight() + 7
				unvisited[Entry{entry[0], entry[1], gear}] = struct{}{}
			}
		}
		min := -1
		for entry, _ := range unvisited {
			if min < 0 || *entry.Weight() < min {
				min = *entry.Weight()
			}
		}
		for entry, _ := range unvisited {
			if *entry.Weight() == min {
				if entry[0] == targetX && entry[1] == targetY && entry[2] == TORCH {
					fmt.Println("Part B:", *entry.Weight()-1)
					return
				}
				visited[entry] = struct{}{}
				delete(unvisited, entry)
			}
		}
	}
}
