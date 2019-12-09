package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func react(polymer []byte, without rune) int {
	count := len(polymer)
	data := bytes.Map(func(r rune) rune {
		if (r^without)|32 == 32 {
			count--
			return 0
		}
		return r
	}, polymer)
	last := 0
	for i := 1; i < len(data); i++ {
		if data[i]^data[last] == 32 {
			data[i] = 0
			data[last] = 0
			count -= 2
			for j := last - 1; j >= 0; j-- {
				if data[j] != 0 {
					last = j
					break
				}
			}
		} else if data[i] != 0 {
			last = i
		}
	}
	return count
}

func main() {
	data, err := ioutil.ReadFile("day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = bytes.TrimSpace(data)

	minSize := react(data, 0)
	fmt.Println("Part A:", minSize)

	for a := 'a'; a <= 'z'; a++ {
		size := react(data, a)
		if size < minSize {
			minSize = size
		}
	}
	fmt.Println("Part B:", minSize)
}
