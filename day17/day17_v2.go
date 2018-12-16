package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var data [2000][2000]byte

var miny, maxy int = 2000, 0

func open(x, y int) bool {
	return data[x][y] == 0 || data[x][y] == '|'
}

func fill(x, y int) {
	if y > maxy {
		return
	} else if !open(x, y) {
		return
	}

	if !open(x, y+1) {
		leftX := x
		for open(leftX, y) && !open(leftX, y+1) {
			data[leftX][y] = '|'
			leftX--
		}
		rightX := x + 1
		for open(rightX, y) && !open(rightX, y+1) {
			data[rightX][y] = '|'
			rightX++
		}
		if open(leftX, y+1) || open(rightX, y+1) {
			fill(leftX, y)
			fill(rightX, y)
		} else if data[leftX][y] == '#' && data[rightX][y] == '#' {
			for x2 := leftX + 1; x2 < rightX; x2++ {
				data[x2][y] = '~'
			}
		}
	} else if data[x][y] == 0 {
		data[x][y] = '|'
		fill(x, y+1)
		if data[x][y+1] == '~' {
			fill(x, y)
		}
	}
}

func main() {
	reader, err := os.Open("day17.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ", ")
		a, _ := strconv.Atoi(line[0][2:])
		bstr := strings.Split(line[1][2:], "..")
		bmin, _ := strconv.Atoi(bstr[0])
		bmax, _ := strconv.Atoi(bstr[1])
		if line[0][0] == 'x' {
			if bmin < miny {
				miny = bmin
			}
			if bmax > maxy {
				maxy = bmax
			}
			for y := bmin; y <= bmax; y++ {
				data[a][y] = '#'
			}
		} else {
			if a < miny {
				miny = a
			}
			if a > maxy {
				maxy = a
			}
			for x := bmin; x <= bmax; x++ {
				data[x][a] = '#'
			}
		}
	}
	reader.Close()

	fill(500, 0)

	water, touched := 0, 0
	for x := 0; x < 2000; x++ {
		for y := miny; y <= maxy; y++ {
			if data[x][y] == '|' {
				touched++
			} else if data[x][y] == '~' {
				water++
			}
		}
	}

	fmt.Println("Part A:", water+touched)
	fmt.Println("Part B:", water)
}
