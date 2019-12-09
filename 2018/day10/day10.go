package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func distance(a, b []int, second int) int {
	dx := (a[0] + a[2]*second) - (b[0] + b[2]*second)
	dy := (a[1] + a[3]*second) - (b[1] + b[3]*second)
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}
	return dx + dy
}

func main() {
	var data [][]int

	reader, err := os.Open("day10.txt")
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
		line := make([]int, 4)
		for i, field := range fields {
			line[i], _ = strconv.Atoi(field)
		}
		data = append(data, line)
	}
	reader.Close()

	second := 2
	lastDistance := distance(data[0], data[1], 0)
	nextDistance := distance(data[0], data[1], 1)
	for nextDistance < lastDistance {
		lastDistance = nextDistance
		nextDistance = distance(data[0], data[1], second)
		second++
	}
	second -= 2
	fmt.Println("Closest second:", second)

	min := [2]int{
		data[0][0] + data[0][2]*second,
		data[0][1] + data[0][3]*second,
	}
	max := [2]int{
		min[0],
		min[1],
	}
	for i := 1; i < len(data); i++ {
		x := data[i][0] + data[i][2]*second
		y := data[i][1] + data[i][3]*second
		if x < min[0] {
			min[0] = x
		}
		if x > max[0] {
			max[0] = x
		}
		if y < min[1] {
			min[1] = y
		}
		if y > max[1] {
			max[1] = y
		}
	}
	width := max[0] - min[0] + 1
	height := max[1] - min[1] + 1

	grid := make([][]byte, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]byte, width)
		for x := 0; x < width; x++ {
			grid[y][x] = ' '
		}
	}

	for i := 0; i < len(data); i++ {
		x := data[i][0] + data[i][2]*second - min[0]
		y := data[i][1] + data[i][3]*second - min[1]
		grid[y][x] = '#'
	}
	for y := 0; y < height; y++ {
		fmt.Println(string(grid[y]))
	}
}
