package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := make([][]int, 100000)
	for x := 0; x < len(grid); x++ {
		grid[x] = make([]int, 100000)
	}
	centerX := 50000
	centerY := 50000

	reader, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		x := centerX
		y := centerY
		steps := 0
		for _, str := range line {
			dirX := 0
			dirY := 0
			if str[0] == 'U' {
				dirY = -1
			} else if str[0] == 'R' {
				dirX = 1
			} else if str[0] == 'D' {
				dirY = 1
			} else {
				dirX = -1
			}
			length, err := strconv.Atoi(str[1:])
			if err != nil {
				panic(err)
			}
			for i := 0; i < length; i++ {
				x += dirX
				y += dirY
				steps++
				if grid[x][y] == 0 {
					grid[x][y] = steps
				}
			}
		}
	}
	closest := 0x7FFFFFFF
	fewest := 0x7FFFFFFF
	if scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		x := centerX
		y := centerY
		steps := 0
		for _, str := range line {
			dirX := 0
			dirY := 0
			if str[0] == 'U' {
				dirY = -1
			} else if str[0] == 'R' {
				dirX = 1
			} else if str[0] == 'D' {
				dirY = 1
			} else {
				dirX = -1
			}
			length, err := strconv.Atoi(str[1:])
			if err != nil {
				panic(err)
			}
			for i := 0; i < length; i++ {
				x += dirX
				y += dirY
				steps++
				if grid[x][y] != 0 {
					distance := 0
					if x < centerX {
						distance += centerX - x
					} else {
						distance += x - centerX
					}
					if y < centerY {
						distance += centerY - y
					} else {
						distance += y - centerY
					}
					if distance != 0 && distance < closest {
						closest = distance
					}
					totalSteps := steps + grid[x][y]
					if totalSteps < fewest {
						fewest = totalSteps
					}
				}
			}
		}
	}
	reader.Close()

	fmt.Println("Part 1:", closest)
	fmt.Println("Part 2:", fewest)
}
