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

loop:
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == 2020 {
				fmt.Println("Part 1: ", data[i], "*", data[j], "=", data[i]*data[j])
				break loop
			}
		}
	}

loop2:
	for i := 0; i < len(data); i++ {
		if data[i] >= 2020 {
			continue
		}
		for j := i + 1; j < len(data); j++ {
			total := data[i] + data[j]
			if total >= 2020 {
				continue
			}
			for k := j + 1; k < len(data); k++ {
				if total+data[k] == 2020 {
					fmt.Println("Part 2: ", data[i], "*", data[j], "*", data[k], "=", data[i]*data[j]*data[k])
					break loop2
				}
			}
		}
	}
}
