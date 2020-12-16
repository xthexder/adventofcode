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
	var data []int

	reader, err := os.Open("day15.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if len(fields) > 0 {
			for _, field := range fields {
				number, _ := strconv.Atoi(field)
				data = append(data, number)
			}
		}
	}
	reader.Close()

outerLoop:
	for i := len(data) - 1; i < 2020; i++ {
		for j := i - 1; j >= 0; j-- {
			if data[i] == data[j] {
				data = append(data, i-j)
				continue outerLoop
			}
		}
		data = append(data, 0)
	}
	fmt.Println("Part 1:", data[2019])

	lastTurn := make(map[int]int)
	for i := 0; i < len(data)-1; i++ {
		lastTurn[data[i]] = i
	}
	lastSpoken := data[len(data)-1]
	for i := len(data) - 1; i < 30000000-1; i++ {
		if last, found := lastTurn[lastSpoken]; found {
			lastTurn[lastSpoken] = i
			lastSpoken = i - last
		} else {
			lastTurn[lastSpoken] = i
			lastSpoken = 0
		}
	}
	fmt.Println("Part 2:", lastSpoken)
}
