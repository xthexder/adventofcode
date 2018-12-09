package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var data [][]int
	var suit []int = make([]int, 1000*1000)

	reader, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return false
			default:
				return true
			}
		})
		line := make([]int, 5)
		for i, field := range fields {
			line[i], _ = strconv.Atoi(field)
		}
		data = append(data, line)
	}
	reader.Close()

	total := 0
	for _, square := range data {
		for x := 0; x < square[3]; x++ {
			for y := 0; y < square[4]; y++ {
				index := square[1] + x + (square[2]+y)*1000
				if suit[index] == 1 {
					total++
				}
				suit[index]++
			}
		}
	}
	fmt.Println("Overlapping:", total)

	for _, square := range data {
		func() {
			for x := 0; x < square[3]; x++ {
				for y := 0; y < square[4]; y++ {
					index := square[1] + x + (square[2]+y)*1000
					if suit[index] != 1 {
						return
					}
				}
			}
			fmt.Println("Not overlapping:", square[0])
		}()
	}
}
