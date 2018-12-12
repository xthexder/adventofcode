package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var width int = 0
var height int = 0
var data [][]byte

type Cart struct {
	pos           int
	dir           int
	intersections int
}

func (c *Cart) Tick(posLookup map[int]*Cart) *Cart {
	delete(posLookup, c.pos)
	switch c.dir {
	case 0: // Up
		c.pos -= width
	case 1: // Right
		c.pos++
	case 2: // Down
		c.pos += width
	case 3: // Left
		c.pos--
	}
	x, y := c.Pos()
	if crash, exists := posLookup[c.pos]; exists {
		fmt.Println("Crash at:", x, y)
		delete(posLookup, c.pos)
		return crash
	}
	posLookup[c.pos] = c
	if data[y][x] == '+' {
		switch c.intersections % 3 {
		case 0:
			// Turn left
			c.dir--
			if c.dir < 0 {
				c.dir += 4
			}
		case 1:
			// Go Straight
		case 2:
			// Go Right
			c.dir++
			if c.dir > 3 {
				c.dir -= 4
			}
		}
		c.intersections++
	} else if data[y][x] == '/' {
		switch c.dir {
		case 0:
			c.dir = 1
		case 1:
			c.dir = 0
		case 2:
			c.dir = 3
		case 3:
			c.dir = 2
		}
	} else if data[y][x] == '\\' {
		switch c.dir {
		case 0:
			c.dir = 3
		case 1:
			c.dir = 2
		case 2:
			c.dir = 1
		case 3:
			c.dir = 0
		}
	}
	return nil
}

func (c *Cart) Pos() (int, int) {
	return c.pos % width, c.pos / width
}

type CartSort []*Cart

func (c CartSort) Len() int      { return len(c) }
func (c CartSort) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CartSort) Less(i, j int) bool {
	if c[j] == nil {
		return true
	} else if c[i] == nil {
		return false
	}
	return c[i].pos < c[j].pos
}

func main() {
	reader, err := os.Open("day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		bytes := make([]byte, len(line))
		copy(bytes, line)
		data = append(data, bytes)
	}
	reader.Close()

	height = len(data)
	width = len(data[0])

	var carts []*Cart
	posLookup := make(map[int]*Cart)

	for y, line := range data {
		for x, c := range line {
			index := x + y*len(line)
			switch c {
			case '^':
				carts = append(carts, &Cart{index, 0, 0})
				posLookup[index] = carts[len(carts)-1]
				line[x] = '|'
			case '>':
				carts = append(carts, &Cart{index, 1, 0})
				posLookup[index] = carts[len(carts)-1]
				line[x] = '-'
			case 'v':
				carts = append(carts, &Cart{index, 2, 0})
				posLookup[index] = carts[len(carts)-1]
				line[x] = '|'
			case '<':
				carts = append(carts, &Cart{index, 3, 0})
				posLookup[index] = carts[len(carts)-1]
				line[x] = '-'
			}
		}
	}

	cartCount := len(carts)
	for tick := 0; cartCount > 1; tick++ {
		sort.Sort(CartSort(carts))

		for i, cart := range carts {
			if cart != nil {
				collision := cart.Tick(posLookup)
				if collision != nil {
					carts[i] = nil
					for i, c := range carts {
						if c == collision {
							carts[i] = nil
							break
						}
					}
					cartCount -= 2
				}
			}
		}
	}

	for _, cart := range carts {
		if cart != nil {
			x, y := cart.Pos()
			fmt.Println("Last cart:", x, y)
		}
	}
}
