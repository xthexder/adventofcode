package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var width int
var data []byte
var carts []*complex64
var cartData map[complex64]*Cart = make(map[complex64]*Cart)

func read(pos complex64) byte {
	return data[int(real(pos))+int(imag(pos))*width]
}

type CartSort []*complex64

func (c CartSort) Len() int      { return len(c) }
func (c CartSort) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CartSort) Less(i, j int) bool {
	return imag(*c[i]) < imag(*c[j]) || (imag(*c[i]) == imag(*c[j]) && real(*c[i]) < real(*c[j]))
}

type Cart struct {
	pos   complex64
	dir   complex64
	state complex64
}

func (c *Cart) Tick() {
	delete(cartData, c.pos)
	c.pos += c.dir
	if crash, exists := cartData[c.pos]; exists {
		fmt.Println("Crash at:", real(c.pos), imag(c.pos))
		delete(cartData, c.pos)
		crash.pos, c.pos = 0, 0
		return
	}
	cartData[c.pos] = c
	if read(c.pos) == '+' {
		c.dir *= c.state
		switch c.state {
		case -1i:
			c.state = 1
		case 1:
			c.state = 1i
		case 1i:
			c.state = -1i
		}
	} else if read(c.pos) == '/' {
		c.dir = complex(-imag(c.dir), -real(c.dir))
	} else if read(c.pos) == '\\' {
		c.dir = complex(imag(c.dir), real(c.dir))
	}
}

func main() {
	reader, err := os.Open("day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		if width == 0 {
			width = len(line)
		}
		data = append(data, line...)
	}
	reader.Close()

	for i := 0; i < len(data); i++ {
		pos := complex(float32(i%width), float32(i/width))
		switch read(pos) {
		case '^':
			cartData[pos] = &Cart{pos, -1i, -1i}
		case '>':
			cartData[pos] = &Cart{pos, 1, -1i}
		case 'v':
			cartData[pos] = &Cart{pos, 1i, -1i}
		case '<':
			cartData[pos] = &Cart{pos, -1, -1i}
		default:
			continue
		}
		carts = append(carts, &cartData[pos].pos)
	}

	for len(cartData) > 1 {
		sort.Sort(CartSort(carts))

		for _, cart := range carts {
			if *cart != 0 {
				cartData[*cart].Tick()
			}
		}
	}

	for pos, _ := range cartData {
		fmt.Println("Last cart:", real(pos), imag(pos))
	}
}
