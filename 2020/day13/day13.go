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
	var earliest int
	var data []int

	reader, err := os.Open("day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		earliest, _ = strconv.Atoi(scanner.Text())
	}
	if scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		for _, field := range fields {
			value, _ := strconv.Atoi(field)
			data = append(data, value)
		}
	}
	reader.Close()

	next := make([]int, len(data))
	for i, bus := range data {
		if bus > 0 {
			next[i] = bus - (earliest % bus)
		}
	}

	min := 0
	for i := range next {
		if data[i] > 0 && next[i] < next[min] {
			min = i
		}
	}
	fmt.Println("Part 1:", next[min]*data[min])

	var schedule [][2]int64
	for minute, bus := range data {
		if bus > 0 {
			schedule = append(schedule, [2]int64{int64(minute), int64(bus)})
		}
	}
	var t int64 = 0
	// fmt.Println(schedule)
	var multiplier int64 = schedule[0][1]
	nextIndex := 1
	var last int64 = -1
	for nextIndex < len(schedule) {
		if (t+schedule[nextIndex][0])%schedule[nextIndex][1] == 0 {
			if last < 0 {
				last = t
			} else {
				multiplier = t - last
				last = -1
				nextIndex++
			}
		}
		// fmt.Println(t, multiplier, nextIndex, last)
		t += multiplier
	}
	for t > multiplier {
		t -= multiplier
	}
	fmt.Println("Part 2:", t)

	// Alternate WolframAlpha Solution:
	// a * 17 == b * 41 - 7 == c * 643 - 17 == d * 23 - 25 == e * 13 - 30 == f * 29 - 46 == g * 433 - 48 == h * 37 - 54 == i * 19 - 67
	// Integer solution for a with n = 0
	//
	// fmt.Println("Part 2:", 17*(69583655027627*0+44715963560085))
}
