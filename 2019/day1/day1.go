package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var data []int

	reader, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			data = append(data, i)
		}
	}
	reader.Close()

	total := 0
	for i := 0; i < len(data); i++ {
		total += (data[i] / 3) - 2
	}
	fmt.Println("Part 1:", total)

	total = 0
	for i := 0; i < len(data); i++ {
		fuel := (data[i] / 3) - 2
		for fuel > 0 {
			total += fuel
			fuel = (fuel / 3) - 2
		}
	}
	fmt.Println("Part 2:", total)
}
