package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Nanobot struct {
	pos   [4]int
	count int
	scale int
}

func (n *Nanobot) distance(a *Nanobot) int {
	sum := 0
	for i := range n.pos[:3] {
		delta := a.pos[i] - n.pos[i]
		if delta < 0 {
			sum -= delta
		} else {
			sum += delta
		}
	}
	return sum
}

func (n *Nanobot) inArea(area *Nanobot) bool {
	sum := 0
	for i, p := range n.pos[:3] {
		if p < area.pos[i] {
			sum += area.pos[i] - p
		} else if p >= area.pos[i]+area.scale {
			sum += p - (area.pos[i] + area.scale - 1)
		}
	}
	return sum <= n.pos[3]
}

func (n *Nanobot) String() string {
	return fmt.Sprintf("[%v count:%d scale:%d]", n.pos, n.count, n.scale)
}

func findMax(queue []*Nanobot) ([]*Nanobot, []*Nanobot) {
	max := 0
	for _, bot := range queue {
		if bot.count > max {
			max = bot.count
		}
	}

	count := 0
	for i, bot := range queue {
		if bot.count == max {
			// Move the processed entries to the beginning so they can be quickly removed
			queue[count], queue[i] = queue[i], queue[count]
			count++
		}
	}
	return queue[:count], queue[count:]
}

func main() {
	var bots []*Nanobot
	var min, max int
	minmaxSet := false

	reader, err := os.Open("day23.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return false
			case r == '-':
				return false
			default:
				return true
			}
		})
		bot := &Nanobot{}
		for i, field := range fields {
			bot.pos[i], _ = strconv.Atoi(field)
			if !minmaxSet {
				min, max = bot.pos[i], bot.pos[i]
				minmaxSet = true
			} else if i < 3 {
				if bot.pos[i] < min {
					min = bot.pos[i]
				}
				if bot.pos[i] > max {
					max = bot.pos[i]
				}
			}
		}
		bots = append(bots, bot)
	}
	reader.Close()

	maxRange := 0
	var maxBot *Nanobot
	for i, bot := range bots {
		if bot.pos[3] > maxRange {
			maxRange = bot.pos[3]
			maxBot = bots[i]
		}
	}

	count := 0
	for _, bot := range bots {
		if bot.distance(maxBot) <= maxRange {
			count++
		}
	}
	fmt.Println("Part A:", count)

	var list []*Nanobot
	queue := []*Nanobot{
		&Nanobot{
			pos:   [4]int{min, min, min, 0},
			scale: max - min + 1,
		},
	}
	for len(queue) > 0 {
		list, queue = findMax(queue)
		// fmt.Println("List:", list)
		// fmt.Println("Queue:", len(queue))
		done := true
		for _, bot := range list {
			if bot.scale > 1 {
				done = false
				break
			}
		}
		if done {
			closest := list[0]
			for _, test := range list[1:] {
				if test.distance(&Nanobot{}) < closest.distance(&Nanobot{}) {
					closest = test
				}
			}
			fmt.Println("Part B point:", closest)
			fmt.Println("Part B:", closest.distance(&Nanobot{}))
			return
		}
		for _, bot := range list {
			if bot.scale <= 1 {
				queue = append(queue, bot)
				continue
			} else {
				bot.scale = (bot.scale + 1) / 2
			}
			// fmt.Println("For:", bot)
			for x := 0; x <= bot.scale; x += bot.scale {
				for y := 0; y <= bot.scale; y += bot.scale {
					for z := 0; z <= bot.scale; z += bot.scale {
						test := &Nanobot{
							pos:   [4]int{bot.pos[0] + x, bot.pos[1] + y, bot.pos[2] + z, 0},
							scale: bot.scale,
						}
						for _, bot := range bots {
							if bot.inArea(test) {
								test.count++
							}
						}
						// fmt.Println(test)
						// Optimization assuming best count includes most of the bots
						if test.count >= len(bots)/2 {
							queue = append(queue, test)
						}
					}
				}
			}
		}
	}
}
