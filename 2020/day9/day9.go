package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const preamble = 25

func main() {
	var data []int

	reader, err := os.Open("day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			arg, _ := strconv.Atoi(line)
			data = append(data, arg)
		}
	}
	reader.Close()

	invalid := 0
	numbers := make(map[int]struct{})
	for i := 0; i < preamble; i++ {
		numbers[data[i]] = struct{}{}
	}
	for i := preamble; i < len(data); i++ {
		found := false
		for num := range numbers {
			other := data[i] - num
			if _, ok := numbers[other]; ok {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Part 1:", data[i])
			invalid = data[i]
			break
		}
		delete(numbers, data[i-preamble])
		numbers[data[i]] = struct{}{}
	}

	start := 0
	for i := 0; i < len(data); i++ {
		if invalid == 0 {
			smallest := data[start]
			largest := data[start]
			for j := start + 1; j < i; j++ {
				if data[j] < smallest {
					smallest = data[j]
				} else if data[j] > largest {
					largest = data[j]
				}
			}
			fmt.Println("Part 2:", smallest+largest)
			break
		} else {
			invalid -= data[i]
			for invalid < 0 {
				invalid += data[start]
				start++
			}
		}
	}
}
