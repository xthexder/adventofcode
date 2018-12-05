package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func distance(x1, y1, x2, y2 int) int {
	dx := x1 - x2
	dy := y1 - y2
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}
	return dx + dy
}

func main() {
	var data []int

	reader, err := os.Open("day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return false
			default:
				return true
			}
		})
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		data = append(data, x, y)
	}
	reader.Close()

	minx := data[0]
	miny := data[1]
	maxx := data[0]
	maxy := data[1]
	for i := 2; i < len(data)-1; i += 2 {
		if data[i] < minx {
			minx = data[i]
		}
		if data[i+1] < miny {
			miny = data[i+1]
		}
		if data[i] > maxx {
			maxx = data[i]
		}
		if data[i+1] > maxy {
			maxy = data[i+1]
		}
	}
	width := maxx - minx + 1
	height := maxy - miny + 1
	boardId := make([]int, width*height)
	boardDistance := make([]int, width*height)
	areaSizes := make([]int, len(data)/2)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			index := x + y*width
			boardDistance[index] = distance(data[0], data[1], x+minx, y+miny)
			for i := 2; i < len(data)-1; i += 2 {
				distance := distance(data[i], data[i+1], x+minx, y+miny)
				if distance < boardDistance[index] {
					boardDistance[index] = distance
					boardId[index] = i / 2
				} else if distance == boardDistance[index] {
					boardId[index] = -1
				}
			}
			if boardId[index] >= 0 {
				areaSizes[boardId[index]]++
			}
		}
	}
	for x := 0; x < width; x++ {
		if boardId[x] >= 0 {
			areaSizes[boardId[x]] = 0
		}
		if boardId[x+width*(height-1)] >= 0 {
			areaSizes[boardId[x+width*(height-1)]] = 0
		}
	}
	for y := 0; y < height; y++ {
		if boardId[y*width] >= 0 {
			areaSizes[boardId[y*width]] = 0
		}
		if boardId[width-1+y*width] >= 0 {
			areaSizes[boardId[width-1+y*width]] = 0
		}
	}
	// Part 1
	// for _, area := range areaSizes {
	// 	fmt.Println(area)
	// }

	// Part 2
	areaSize := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dist := distance(data[0], data[1], x+minx, y+miny)
			for i := 2; i < len(data)-1; i += 2 {
				dist += distance(data[i], data[i+1], x+minx, y+miny)
			}
			if dist < 10000 {
				areaSize++
			}
		}
	}
	fmt.Println("Part B:", areaSize)
}
