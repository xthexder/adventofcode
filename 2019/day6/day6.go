package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	orbits := make(map[string]string)

	reader, err := os.Open("day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		orbit := strings.Split(scanner.Text(), ")")
		orbits[orbit[1]] = orbit[0]
	}
	reader.Close()

	total := 0
	for _, orbit := range orbits {
		ok := true
		for ok {
			total++
			orbit, ok = orbits[orbit]
		}
	}
	fmt.Println("Part 1:", total)

	santaOrbits := make(map[string]int)
	transfers := 0
	orbit, ok := orbits["SAN"]
	for ok {
		santaOrbits[orbit] = transfers
		orbit, ok = orbits[orbit]
		transfers++
	}

	transfers = 0
	orbit, ok = orbits["YOU"]
	for ok {
		if trans, ok := santaOrbits[orbit]; ok {
			fmt.Println("Part 2:", transfers + trans)
			break
		}
		orbit, ok = orbits[orbit]
		transfers++
	}
}
