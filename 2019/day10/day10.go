package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type asteroid [2]int

func (a asteroid) Angle(laser asteroid) float64 {
	angle := math.Atan2(float64(a[0]-laser[0]), float64(laser[1]-a[1]))
	if angle < 0 {
		angle += math.Pi * 2
	}
	return angle
}

type asteroidSort struct {
	list  []asteroid
	laser asteroid
}

func (c asteroidSort) Len() int      { return len(c.list) }
func (c asteroidSort) Swap(i, j int) { c.list[i], c.list[j] = c.list[j], c.list[i] }
func (c asteroidSort) Less(i, j int) bool {
	return c.list[i].Angle(c.laser) < c.list[j].Angle(c.laser)
}

func main() {
	var field [][]byte
	var width, height int

	reader, err := os.Open("day10.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		field = append(field, line)
		if width == 0 {
			width = len(line)
		}
	}
	height = len(field)
	reader.Close()

	var asteroids []asteroid
	maxAsteroids := 0
	maxX := 0
	maxY := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if field[y][x] != '#' {
				continue
			}
			count := 0
			var tmpAsteroids []asteroid
			for x2 := 0; x2 < width; x2++ {
				for y2 := 0; y2 < height; y2++ {
					if (x != x2 || y != y2) && field[y2][x2] == '#' {
						blocked := false
						if x2 == x {
							dy := 1
							if y2 < y {
								dy = -1
							}
							for y3 := y + dy; y3 != y2; y3 += dy {
								if field[y3][x] == '#' {
									blocked = true
									break
								}
							}
						} else if y2 == y {
							dx := 1
							if x2 < x {
								dx = -1
							}
							for x3 := x + dx; x3 != x2; x3 += dx {
								if field[y][x3] == '#' {
									blocked = true
									break
								}
							}
						} else {
							dx := 1
							if x2 < x {
								dx = -1
							}
							dy := float64(y2-y) / math.Abs(float64(x2-x))
							y3 := float64(y) + dy
							for x3 := x + dx; x3 != x2; x3 += dx {
								y3round := math.Round(y3*1000) / 1000
								if math.Floor(y3round) == y3round && field[int(y3round)][x3] == '#' {
									blocked = true
									break
								}
								y3 += dy
							}
						}
						if !blocked {
							count++
							tmpAsteroids = append(tmpAsteroids, asteroid{x2, y2})
						}
					}
				}
			}
			if count > maxAsteroids {
				maxAsteroids = count
				asteroids = tmpAsteroids
				maxX = x
				maxY = y
			}
		}
	}
	laser := asteroid{maxX, maxY}
	fmt.Println("Part 1:", maxAsteroids, laser)

	sort.Sort(asteroidSort{asteroids, laser})
	fmt.Println("Part 2:", asteroids[199], "=", asteroids[199][0]*100+asteroids[199][1])
}
