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
	color string
	num   int
}

var data map[string][]rule
var countMap map[string]int

func recursiveCount(container string) int {
	if num, ok := countMap[container]; ok {
		return num
	}

	count := 1
	for _, r := range data[container] {
		count += r.num * recursiveCount(r.color)
	}
	countMap[container] = count
	return count
}

func main() {
	data = make(map[string][]rule)
	countMap = make(map[string]int)

	reader, err := os.Open("day7.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) > 2 {
			var rules []rule
			for i := 2; i < len(words)-2; i++ {
				count, err := strconv.Atoi(words[i])
				if err == nil {
					color := words[i+1] + " " + words[i+2]
					rules = append(rules, rule{color, count})
				}
			}
			data[words[0]+" "+words[1]] = rules
		}
	}
	reader.Close()

	canContain := make(map[string]struct{})
	canContain["shiny gold"] = struct{}{}
	change := true
	for change {
		change = false
		for container, rules := range data {
			for _, r := range rules {
				if _, ok := canContain[r.color]; ok {
					if _, ok := canContain[container]; !ok {
						change = true
					}
					canContain[container] = struct{}{}
				}
			}
		}
	}
	count := 0
	for range canContain {
		count++
	}
	fmt.Println("Part 1:", count-1)
	fmt.Println("Part 2:", recursiveCount("shiny gold")-1)
}
