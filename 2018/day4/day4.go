package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func findMaxMinute(times []time.Time) (int, int) {
	minutes := make([]int, 60)
	maxMinute := -1
	for i := 0; i < len(times); i += 2 {
		for m := times[i].Minute(); m != times[i+1].Minute(); m = (m + 1) % 60 {
			minutes[m]++
			if maxMinute < 0 || minutes[maxMinute] < minutes[m] {
				maxMinute = m
			}
		}
	}
	return maxMinute, minutes[maxMinute]
}

func main() {
	var guards = make(map[int][]time.Time)

	reader, err := os.Open("day4_sorted.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lastGuard int
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		t, err := time.Parse("[2006-01-02 15:04]", line[:18])
		line = line[19:]
		if err == nil {
			if strings.Contains(line, "#") {
				lastGuard, _ = strconv.Atoi(strings.Fields(line)[1][1:])
			} else if strings.HasPrefix(line, "falls") {
				guards[lastGuard] = append(guards[lastGuard], t)
			} else if strings.HasPrefix(line, "wakes") {
				guards[lastGuard] = append(guards[lastGuard], t)
			}
		}
	}
	reader.Close()

	maxGuard := -1
	var maxDuration time.Duration
	for guard, times := range guards {
		var total time.Duration
		for i := 0; i < len(times); i += 2 {
			total += times[i+1].Sub(times[i])
		}
		if maxGuard < 0 || total > maxDuration {
			maxGuard = guard
			maxDuration = total
		}
	}
	fmt.Println(maxGuard, "slept the longest at", maxDuration)

	guardMinutes := make(map[int]int)
	guardMinutesCount := make(map[int]int)
	for guard, times := range guards {
		guardMinutes[guard], guardMinutesCount[guard] = findMaxMinute(times)
	}

	fmt.Println("Part A:", maxGuard, "*", guardMinutes[maxGuard], "=", maxGuard*guardMinutes[maxGuard])

	maxGuard = -1
	for guard, minutes := range guardMinutesCount {
		if maxGuard < 0 || minutes > guardMinutesCount[maxGuard] {
			maxGuard = guard
		}
	}
	fmt.Println("Part B:", maxGuard, "*", guardMinutes[maxGuard], "=", maxGuard*guardMinutes[maxGuard])
}
