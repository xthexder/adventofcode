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

func (n Nanobot) add(dir [3]int, multiplier int) Nanobot {
	return Nanobot{n[0] + dir[0]*multiplier, n[1] + dir[1]*multiplier, n[2] + dir[2]*multiplier}
}

var directions = [][3]int{
	{0, 0, 1},
	{0, 0, -1},
	{0, 1, 0},
	{0, -1, 0},
	{1, 0, 0},
	{-1, 0, 0},
}

func totalDistance(bot Nanobot, bots []Nanobot) (int, int) {
	sum := 0
	count := 0
	for i := range bots {
		if bots[i].distance(bot) <= bots[i][3] {
			sum += bots[i].distance(bot) - bots[i][3]
			count++
		}
	}
	return sum, count
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
	dist, count := totalDistance(pos, bots)
	maxCount := count
	maxCountPos := pos
	moved := true
	multiplier := 100000
	for moved {
		moved = false
		for _, dir := range directions {
			newPos := pos.add(dir, multiplier)
			newDist, newCount := totalDistance(newPos, bots)
			if newDist < dist || (newDist == dist && newCount > count) {
				moved = true
				pos = newPos
				dist = newDist
				if newCount > maxCount {
					// fmt.Println(newCount)
					maxCount = newCount
					maxCountPos = newPos
				}
				count = newCount
				break
			}
		}
		if !moved && multiplier > 1 {
			multiplier /= 2
			moved = true
		}
	}
	fmt.Println("Count: ", maxCount, maxCountPos)

	moved = true
	multiplier = 10000
	pos = maxCountPos
	count = maxCount
	for moved {
		moved = false
		for _, dir := range directions {
			newPos := pos.add(dir, multiplier)
			if newPos.distance(Nanobot{0, 0, 0, 0}) >= pos.distance(Nanobot{0, 0, 0, 0}) {
				continue
			}
			_, newCount := totalDistance(newPos, bots)
			if newCount >= count {
				moved = true
				pos = newPos
				if newCount > count {
					// fmt.Println(count)
				}
				count = newCount
				break
			}
		}
		if !moved && multiplier > 1 {
			multiplier--
			moved = true
		}
	}
	// This is definitely wrong, but it happened to get the right answer anyway.
	fmt.Println("Part B:", pos.distance(Nanobot{0, 0, 0, 0}), pos, count)
}
