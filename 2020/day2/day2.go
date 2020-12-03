package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type input struct {
	Min, Max int
	Policy   byte
	Pass     string
}

func main() {
	var data []input

	reader, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return false
			case r >= 'a' && r <= 'z':
				return false
			default:
				return true
			}
		})
		min, _ := strconv.Atoi(fields[0])
		max, _ := strconv.Atoi(fields[1])
		pass := input{
			Min:    min,
			Max:    max,
			Policy: fields[2][0],
			Pass:   fields[3],
		}
		data = append(data, pass)
	}
	reader.Close()

	validPasswords := 0
	for i := 0; i < len(data); i++ {
		total := 0
		for _, ch := range data[i].Pass {
			if ch == rune(data[i].Policy) {
				total++
			}
		}
		if total >= data[i].Min && total <= data[i].Max {
			validPasswords++
		}
	}
	fmt.Println("Part 1:", validPasswords)

	validPasswords = 0
	for i := 0; i < len(data); i++ {
		if data[i].Min <= len(data[i].Pass) && data[i].Max <= len(data[i].Pass) {
			a := data[i].Pass[data[i].Min-1] == data[i].Policy
			b := data[i].Pass[data[i].Max-1] == data[i].Policy
			if (a || b) && !(a && b) {
				validPasswords++
			}
		}
	}
	fmt.Println("Part 2:", validPasswords)
}
