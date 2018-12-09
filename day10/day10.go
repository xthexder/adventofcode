package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func distance(a, b []int, second int) int {
	dx := (a[0] + a[2]*second) - (b[0] + b[2]*second)
	dy := (a[1] + a[3]*second) - (b[1] + b[3]*second)
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}
	return dx + dy
}

func main() {
	var data [][]int

	reader, err := os.Open("day10.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return false
			case r == '-':
				return false
			default:
				return true
			}
		})
		line := make([]int, 4)
		for i, field := range fields {
			line[i], _ = strconv.Atoi(field)
		}
		data = append(data, line)
	}
	reader.Close()

	second := 2
	lastDistance := distance(data[0], data[1], 0)
	nextDistance := distance(data[0], data[1], 1)
	for nextDistance < lastDistance {
		lastDistance = nextDistance
		nextDistance = distance(data[0], data[1], second)
		second++
	}
	fmt.Println("Closest second:", second)

	// Output in CSV format to graph, velocity can be used to offset points (-2 in this case)
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i][0] + data[i][2]*second)
		fmt.Print(",", data[i][1]+data[i][3]*second)
		fmt.Print(",", data[i][2])
		fmt.Println(",", data[i][3])
	}
}
