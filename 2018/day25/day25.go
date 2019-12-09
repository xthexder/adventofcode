package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point [4]int

func (p Point) distance(a Point) int {
	sum := 0
	for i := range p {
		delta := a[i] - p[i]
		if delta < 0 {
			sum -= delta
		} else {
			sum += delta
		}
	}
	return sum
}

func main() {
	var data []Point

	reader, err := os.Open("day25.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		p := Point{}
		for i := range p {
			p[i], _ = strconv.Atoi(line[i])
		}
		data = append(data, p)
	}
	reader.Close()

	var list [][]int
	mapping := make([]*int, len(data))
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			if data[i].distance(data[j]) <= 3 {
				if mapping[i] == nil && mapping[j] == nil {
					tmp := len(list)
					mapping[i] = &tmp
					mapping[j] = &tmp
					list = append(list, []int{i, j})
				} else if mapping[i] == nil {
					mapping[i] = mapping[j]
					list[*mapping[j]] = append(list[*mapping[j]], i)
				} else if mapping[j] == nil {
					mapping[j] = mapping[i]
					list[*mapping[i]] = append(list[*mapping[i]], j)
				} else if *mapping[i] != *mapping[j] {
					fromList := *mapping[i]
					toList := *mapping[j]
					if len(list[*mapping[j]]) < len(list[fromList]) {
						fromList = *mapping[j]
						toList = *mapping[i]
					}
					for k := range list[fromList] {
						*mapping[list[fromList][k]] = toList
						list[toList] = append(list[toList], list[fromList][k])
					}
					list[fromList] = nil
				}
			}
		}
	}
	unique := make(map[int]struct{})
	for i := range data {
		if mapping[i] == nil {
			tmp := len(list)
			mapping[i] = &tmp
			list = append(list, []int{i})
		}
		unique[*mapping[i]] = struct{}{}
	}

	fmt.Println("Part A:", len(unique))
}
