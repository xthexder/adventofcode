package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkPlant(plants []byte, plantMap [][]byte) byte {
	for i := 0; i < 5; i++ {
		if plants[i] == 0 {
			plants[i] = '.'
		}
	}
	for _, m := range plantMap {
		if bytes.Equal(plants[:5], m[:5]) {
			return m[5]
		}
	}
	return 0
}

func printPlants(plants []byte, gen, center int) {
	first, last := -1, -1
	for i, c := range plants {
		if c == '#' {
			if first < 0 {
				first = i
			}
			last = i
		}
	}
	if center < first {
		first = center
	}
	fmt.Println(gen, string(plants[first:last]))
	fmt.Println(strings.Repeat(" ", center-first+len(strconv.Itoa(gen))), "^")
}

func main() {
	reader, err := os.Open("day12.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	initialState := scanner.Bytes()[15:]

	plants := make([]byte, (len(initialState)+200)*2)
	center := len(initialState) + 200
	for i, c := range initialState {
		plants[center+i] = c
	}
	printPlants(plants, 0, center)

	var plantMap [][]byte
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 10 {
			m := line[:6]
			m[5] = line[9]
			plantMap = append(plantMap, m)
		}
	}
	reader.Close()

	sum98 := 0
	sum99 := 0
	for i := 0; i < 100; i++ {
		nextGen := make([]byte, len(plants))
		copy(nextGen, plants)
		for j := 2; j < len(plants)-2; j++ {
			out := checkPlant(plants[j-2:], plantMap)
			if out != 0 {
				nextGen[j] = out
			}
		}
		printPlants(plants, i+1, center)
		plants = nextGen

		if i == 19 {
			sum := 0
			for i, c := range plants {
				if c == '#' {
					sum += i - center
				}
			}
			fmt.Println("Part A:", sum)
		} else if i == 98 {
			for i, c := range plants {
				if c == '#' {
					sum98 += i - center
				}
			}
		} else if i == 99 {
			for i, c := range plants {
				if c == '#' {
					sum99 += i - center
				}
			}
			fmt.Println("Part B:", sum99+(sum99-sum98)*(50000000000-100))
		}
	}
}
