package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Nanobot [4]int

func (n Nanobot) distance(a Nanobot) int {
	sum := 0
	for i := range n[:3] {
		delta := a[i] - n[i]
		if delta < 0 {
			sum -= delta
		} else {
			sum += delta
		}
	}
	return sum
}

func (n Nanobot) add(dir [3]int) Nanobot {
	return Nanobot{n[0] + dir[0], n[1] + dir[1], n[2] + dir[2]}
}

var directions = [][3]int{
	{0, 0, 10},
	{0, 0, -10},
	{0, 10, 0},
	{0, -10, 0},
	{10, 0, 0},
	{-10, 0, 0},

	{0, 0, 1},
	{0, 0, -1},
	{0, 1, 0},
	{0, -1, 0},
	{1, 0, 0},
	{-1, 0, 0},
}

func totalDistance(bot Nanobot, bots []Nanobot) int {
	sum := 0
	for i := range bots {
		if bots[i].distance(bot) > bots[i][3] {
			sum += bots[i].distance(bot) - bots[i][3]
		}
	}
	return sum
}

func main() {
	var bots []Nanobot

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
		bot := Nanobot{}
		for i, field := range fields {
			bot[i], _ = strconv.Atoi(field)
		}
		bots = append(bots, bot)
	}
	reader.Close()

	maxRange := 0
	var maxBot *Nanobot
	for i, bot := range bots {
		if bot[3] > maxRange {
			maxRange = bot[3]
			maxBot = &bots[i]
		}
	}

	count := 0
	for _, bot := range bots {
		if bot.distance(*maxBot) <= maxRange {
			count++
		}
	}
	fmt.Println("Part A:", count)

	pos := Nanobot{0, 0, 0, 0}
	lastPos := pos
	dist := totalDistance(pos, bots)
	moved := true
	i := 0
	for moved {
		if i%10000 == 0 {
			fmt.Println(dist)
		}
		i++
		moved = false
		for _, dir := range directions {
			newPos := pos.add(dir)
			if newPos == lastPos {
				continue
			}
			newDist := totalDistance(newPos, bots)
			if newDist < dist {
				moved = true
				lastPos = pos
				pos = newPos
				dist = newDist
				break
			}
		}
	}
	fmt.Println("Part B:", pos.distance(Nanobot{0, 0, 0, 0}))
}
