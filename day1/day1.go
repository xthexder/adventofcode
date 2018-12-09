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

	sum := make([]int, len(data))

	total := 0
	for i := 0; i < len(data); i++ {
		total += data[i]
	}
	fmt.Println("Sum:", total)

	count := 0
	for {
		for i := 0; i < len(data); i++ {
			sum[i] += data[(i+count)%len(data)]
			if sum[i] == 0 {
				total = 0
				for j := i - 1; j >= 0; j-- {
					total += data[j]
				}
				fmt.Println("First repeat:", total)
				return
			}
		}
		count++
	}
}
