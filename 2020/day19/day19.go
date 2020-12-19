package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func compileInput(input map[int][]string) map[int]string {
	rules := make(map[int]string)
	for len(input) > 0 {
	nextRule:
		for rule, fields := range input {
			regex := ""
			for _, field := range fields {
				if strings.HasPrefix(field, "\"") {
					regex += strings.Trim(field, "\"")
				} else if value, err := strconv.Atoi(field); err == nil {
					if str, ok := rules[value]; ok {
						if len(str) > 1 {
							regex += "(" + str + ")"
						} else {
							regex += str
						}
					} else {
						continue nextRule
					}
				} else if field == "|" {
					regex += "|"
				} else {
					panic("Unknown field: " + field)
				}
			}
			rules[rule] = regex
			delete(input, rule)
			break
		}
	}
	return rules
}

func main() {
	input := make(map[int][]string)
	var data []string

	reader, err := os.Open("day19.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			fields := strings.FieldsFunc(line, func(r rune) bool {
				switch {
				case r >= '0' && r <= '9':
					return false
				case r >= 'a' && r <= 'z':
					return false
				case r == '|' || r == '"':
					return false
				default:
					return true
				}
			})
			rule, _ := strconv.Atoi(fields[0])

			input[rule] = fields[1:]
		} else {
			break
		}
	}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			data = append(data, line)
		}
	}
	reader.Close()

	rules := compileInput(input)

	// fmt.Println(rules[0])
	regex, err := regexp.Compile("^" + rules[0] + "$")
	if err != nil {
		panic(err)
	}

	count := 0
	for _, str := range data {
		if regex.Match([]byte(str)) {
			count++
		}
	}
	fmt.Println("Part 1:", count)

	regex8, err := regexp.Compile("^(" + rules[8] + ")+$")
	if err != nil {
		panic(err)
	}

	regex11L, err := regexp.Compile("(" + rules[42] + ")$")
	if err != nil {
		panic(err)
	}

	regex11R, err := regexp.Compile("(" + rules[31] + ")$")
	if err != nil {
		panic(err)
	}

	// 0: 8 11
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	count = 0
	for _, str := range data {
		input := []byte(str)
		matches := 0
		match := regex11R.Find(input)
		for match != nil {
			input = input[:len(input)-len(match)]
			matches++
			match = regex11R.Find(input)
		}
		if matches == 0 {
			continue
		}
		match = regex11L.Find(input)
		for matches > 0 && match != nil {
			input = input[:len(input)-len(match)]
			matches--
			match = regex11L.Find(input)
		}
		if matches != 0 {
			continue
		}
		if regex8.Match(input) {
			count++
		}
	}
	fmt.Println("Part 2:", count)
}
