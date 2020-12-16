package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	name   string
	index  []bool
	ranges [][2]int
}

func (r rule) Valid(value int) bool {
	for _, validRange := range r.ranges {
		if value >= validRange[0] && value <= validRange[1] {
			return true
		}
	}
	return false
}

func main() {
	var rules []rule
	var tickets [][]int

	reader, err := os.Open("day16.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsRune(line, ':') {
			fields := strings.Split(line, ":")
			if len(fields) > 1 {
				numbers := strings.FieldsFunc(fields[1], func(r rune) bool {
					switch {
					case r >= '0' && r <= '9':
						return false
					default:
						return true
					}
				})
				if len(numbers) > 0 {
					rule := rule{name: fields[0]}
					for i := 1; i < len(numbers); i += 2 {
						min, _ := strconv.Atoi(numbers[i-1])
						max, _ := strconv.Atoi(numbers[i])
						rule.ranges = append(rule.ranges, [2]int{min, max})
					}
					rules = append(rules, rule)
				}
			}
		} else if len(line) > 0 {
			fields := strings.Split(line, ",")
			numbers := make([]int, len(fields))
			for i, field := range fields {
				numbers[i], _ = strconv.Atoi(field)
			}
			tickets = append(tickets, numbers)
		}
	}
	reader.Close()

	totalError := 0
	for i := 1; i < len(tickets); i++ {
		for _, field := range tickets[i] {
			valid := false
			for _, rule := range rules {
				if rule.Valid(field) {
					valid = true
					break
				}
			}
			if !valid {
				totalError += field
				tickets[i] = nil
			}
		}
	}
	fmt.Println("Part 1:", totalError)

	fieldCount := len(tickets[0])
	for i := range rules {
		rules[i].index = make([]bool, fieldCount)
	}

	for _, rule := range rules {
		for i := range rule.index {
			valid := true
			for _, ticket := range tickets {
				if ticket == nil {
					continue
				}
				if !rule.Valid(ticket[i]) {
					valid = false
					break
				}
			}
			rule.index[i] = valid
		}
	}

	indexes := make([]*rule, fieldCount)
	done := false
	for !done {
		done = true
		for i, existing := range indexes {
			if existing != nil {
				continue
			}
			done = false
			var only *rule = nil
			for r, rule := range rules {
				if rule.index != nil && rule.index[i] {
					if only == nil {
						only = &rules[r]
					} else {
						only = nil
						break
					}
				}
			}
			if only != nil {
				indexes[i] = only
				only.index = nil
			}
		}
	}
	total := -1
	fmt.Println()
	fmt.Println("Your ticket:")
	for i, rule := range indexes {
		if strings.HasPrefix(rule.name, "departure") {
			if total < 0 {
				total = tickets[0][i]
			} else {
				total *= tickets[0][i]
			}
		}
		fmt.Println(rule.name, tickets[0][i])
	}
	fmt.Println()
	fmt.Println("Part 2:", total)
}
