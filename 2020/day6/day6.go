package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
)

func main() {
	var data [][]int

	reader, err := os.Open("day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	var yes []int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			y := 0
			for _, ch := range line {
				y = y | (1 << (ch - 'a'))
			}
			yes = append(yes, y)
		} else {
			yes2 := make([]int, len(yes))
			copy(yes2, yes)
			data = append(data, yes2)
			yes = yes[:0]
		}
	}
	reader.Close()

	total := 0
	for _, yes := range data {
		group := 0
		for _, y := range yes {
			group = group | y
		}
		total += bits.OnesCount(uint(group))
	}

	fmt.Println("Part 1:", total)

	total = 0
	for _, yes := range data {
		group := 0xFFFFFFFF
		for _, y := range yes {
			group = group & y
		}
		total += bits.OnesCount(uint(group))
	}

	fmt.Println("Part 2:", total)
}
