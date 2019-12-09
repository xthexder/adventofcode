package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// This implementation was slower and longer than expected... Still an interesting exercise

var width, carts, iteration int
var tiles []func(*Cart)

func get(pos complex64) *func(*Cart) {
	return &tiles[int(real(pos))+int(imag(pos))*width]
}

func Null(c *Cart) {}

type Cart struct {
	pos                      complex64
	dir                      complex64
	iteration, intersections int
	tile                     func(*Cart)
}

func NewCart(pos, dir complex64) func(*Cart) {
	c := &Cart{pos, dir, 0, 0, Null}
	carts++
	return c.Tick
}

func (c *Cart) Tick(cart *Cart) {
	if cart != nil {
		fmt.Println("Crash at:", real(c.pos), imag(c.pos))
		carts -= 2
		*get(c.pos) = c.tile
		return
	} else if carts == 1 {
		fmt.Println("Last cart:", real(c.pos), imag(c.pos))
		carts--
		return
	} else if iteration == c.iteration {
		return
	}
	c.iteration = iteration

	*get(c.pos) = c.tile
	c.pos += c.dir
	c.tile = *get(c.pos)
	*get(c.pos) = c.Tick
	c.tile(c)
}

func main() {
	reader, err := os.Open("day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for y := 0; scanner.Scan(); y++ {
		bytes := scanner.Bytes()
		if width == 0 {
			width = len(bytes)
		}
		for i, b := range bytes {
			pos := complex(float32(i), float32(y))
			switch b {
			case '^':
				tiles = append(tiles, NewCart(pos, -1i))
			case '>':
				tiles = append(tiles, NewCart(pos, 1))
			case 'v':
				tiles = append(tiles, NewCart(pos, 1i))
			case '<':
				tiles = append(tiles, NewCart(pos, -1))
			case '+':
				tiles = append(tiles, func(c *Cart) {
					if c != nil {
						switch c.intersections % 3 {
						case 0:
							c.dir *= -1i
						case 1:
							c.dir *= 1
						case 2:
							c.dir *= 1i
						}
						c.intersections++
					}
				})
			case '/':
				tiles = append(tiles, func(c *Cart) {
					if c != nil {
						c.dir = complex(-imag(c.dir), -real(c.dir))
					}
				})
			case '\\':
				tiles = append(tiles, func(c *Cart) {
					if c != nil {
						c.dir = complex(imag(c.dir), real(c.dir))
					}
				})
			default:
				tiles = append(tiles, Null)
			}
		}
	}
	reader.Close()

	for carts > 0 {
		iteration++
		for _, tile := range tiles {
			tile(nil)
		}
	}
}
