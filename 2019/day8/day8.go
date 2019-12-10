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

	reader, err := os.Open("day8.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			data = append(data, i)
		}
	}
	reader.Close()

	i := 0
	minZeros := 0x7FFFFFFF
	layerOnes := 0
	layerTwos := 0
	for layer := 0; i < len(data); layer++ {
		countZeros := 0
		countOnes := 0
		countTwos := 0
		for x := 0; x < 25 && i < len(data); x++ {
			for y := 0; y < 6 && i < len(data); y++ {
				if data[i] == 0 {
					countZeros++
				} else if data[i] == 1 {
					countOnes++
				} else if data[i] == 2 {
					countTwos++
				}
				i++
			}
		}
		if countZeros < minZeros {
			minZeros = countZeros
			layerOnes = countOnes
			layerTwos = countTwos
		}
	}
	fmt.Println("Part 1:", layerOnes*layerTwos)

	display := make([][]byte, 6)
	for y := 0; y < len(display); y++ {
		display[y] = make([]byte, 25)
	}
	i = 0
	for layer := 0; i < len(data); layer++ {
		for y := 0; y < 6 && i < len(data); y++ {
			for x := 0; x < 25 && i < len(data); x++ {
				if display[y][x] == 0 {
					if data[i] == 0 {
						display[y][x] = ' ' // Black
					} else if data[i] == 1 {
						display[y][x] = '#' // White
					} else if data[i] == 2 {
						display[y][x] = 0 // Transparent
					}
				}
				i++
			}
		}
	}
	fmt.Println("Part 2:")
	for y := 0; y < len(display); y++ {
		fmt.Println(string(display[y]))
	}
}
