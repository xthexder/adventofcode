package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var width int
var board [][2]int // [generation, distance]
var units []*int
var unitData = make(map[int]*Unit)
var elfs, goblins, round, generation int

type Unit struct {
	elf bool
	pos int
	hp  int
}

var directions = [][2]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

func move(pos, dx, dy int) (int, bool) {
	newX := (pos % width) + dx
	newY := pos + dy*width
	if newX < 0 || newX >= width || newY < 0 || newY >= len(board) {
		return pos, false
	}
	newPos := pos + dx + dy*width
	if board[newPos][0] < 0 {
		return pos, false
	}
	return newPos, true
}

func findStart(pos int) int {
	if board[pos][1] <= 0 {
		return -1
	} else if board[pos][1] == 1 {
		return pos
	}

	for i := 0; i < len(directions); i++ {
		if newPos, ok := move(pos, directions[i][0], directions[i][1]); ok {
			if board[newPos][0] == generation && board[newPos][1] < board[pos][1] {
				return findStart(newPos)
			}
		}
	}
	return -1
}

func findEnemy(elfIsEnemy bool, startPos int) int {
	generation++
	queue := make([]int, 1, 16)
	queue[0] = startPos
	board[startPos][0] = generation
	board[startPos][1] = 0

	for len(queue) > 0 {
		sort.Ints(queue)
		processing := queue[:len(queue)]
		queue = queue[len(queue):]
		for _, pos := range processing {
			for i := 0; i < len(directions); i++ {
				if newPos, ok := move(pos, directions[i][0], directions[i][1]); ok {
					if unit, ok := unitData[newPos]; ok {
						if unit.elf == elfIsEnemy {
							// fmt.Println("    Found enemy at (", newPos%width, newPos/width, ")")
							return findStart(pos)
						}
					} else if board[newPos][0] < generation || board[newPos][1] > (board[pos][1]+1) {
						board[newPos][0] = generation
						board[newPos][1] = board[pos][1] + 1
						queue = append(queue, newPos)
					}
				}
			}
		}
	}
	return -1
}

func printBoard() {
	for i := 0; i < len(board); i += width {
		for x := 0; x < width; x++ {
			if unit, ok := unitData[i+x]; ok {
				if unit.elf {
					fmt.Print("E")
				} else {
					fmt.Print("G")
				}
			} else if board[i+x][0] < 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type UnitSort []*int

func (c UnitSort) Len() int           { return len(c) }
func (c UnitSort) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c UnitSort) Less(i, j int) bool { return *c[i] < *c[j] }

func main() {
	reader, err := os.Open("day15.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		if width == 0 {
			width = len(line)
		}

		for _, c := range line {
			switch c {
			case '#':
				board = append(board, [2]int{-1, -1})
			case 'E', 'G':
				if c == 'E' {
					elfs++
				} else {
					goblins++
				}
				i := len(board)
				unitData[i] = &Unit{
					elf: c == 'E',
					pos: i,
					hp:  200,
				}
				units = append(units, &unitData[i].pos)
				fallthrough
			case '.':
				board = append(board, [2]int{0, 0})
			}
		}
	}
	reader.Close()

	printBoard()

	for elfs > 0 && goblins > 0 {
		// fmt.Println("Round", round)
		// printBoard()
		sort.Sort(UnitSort(units))

		for _, pos := range units {
			if *pos >= 0 {
				if unit, ok := unitData[*pos]; ok {
					// fmt.Println("Running unit (", unit.pos%width, unit.pos/width, "):")
					newPos := findEnemy(!unit.elf, unit.pos)
					if newPos >= 0 {
						// fmt.Println("    Moving to (", newPos%width, newPos/width, ")")
						delete(unitData, unit.pos)
						unit.pos = newPos
						unitData[unit.pos] = unit
					}
					var target *Unit
					for i := 0; i < len(directions); i++ {
						if newPos, ok := move(unit.pos, directions[i][0], directions[i][1]); ok {
							if enemy, ok := unitData[newPos]; ok && enemy.elf != unit.elf {
								if target == nil || target.hp > enemy.hp {
									target = enemy
								}
							}
						}
					}
					if target != nil {
						// fmt.Println("    Attacking enemy at (", target.pos%width, target.pos/width, "):", target)
						if unit.elf {
							target.hp -= 3 // 12
						} else {
							target.hp -= 3
						}
						if target.hp <= 0 {
							if target.elf {
								elfs--
								fmt.Println("Elf died at (", target.pos%width, target.pos/width, ")")
							} else {
								goblins--
								fmt.Println("Goblin died at (", target.pos%width, target.pos/width, ")")
							}
							delete(unitData, target.pos)
							target.pos = -1
						}
					}
				}
			}
		}
		if elfs > 0 && goblins > 0 {
			round++
		}
	}

	hpLeft := 0
	for _, unit := range unitData {
		hpLeft += unit.hp
	}

	if elfs > goblins {
		fmt.Print("Elfs")
	} else if goblins > elfs {
		fmt.Print("Goblins")
	} else {
		fmt.Print("Noone")
	}
	fmt.Println(" wins after", round, "rounds with", hpLeft, "hit points left.")
	fmt.Println("Part A:", round*hpLeft)
}
