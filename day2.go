package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var data []string

	reader, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	reader.Close()

	two := 0
	three := 0
	for i := 0; i < len(data); i++ {
		count := make([]int, 256)
		twos := 0
		threes := 0
		for j := 0; j < len(data[i]); j++ {
			count[data[i][j]]++
			val := count[data[i][j]]
			if val == 2 {
				twos++
			} else if val == 3 {
				twos--
				threes++
			} else if val == 4 {
				threes--
			}
		}
		if twos > 0 {
			two++
		}
		if threes > 0 {
			three++
		}
	}
	fmt.Println("Checksum:", two, "*", three, "=", two*three)
}
