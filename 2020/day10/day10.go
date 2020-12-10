package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var variantCache []int

func getVariants(data []int, i int) int {
	if len(data[i:]) <= 1 {
		return 1
	} else if variantCache[i] != 0 {
		return variantCache[i]
	}
	// fmt.Println("Base:", data[1:])
	variants := 0
	diff := data[i+1] - data[i]
	if diff < 3 && len(data) > 2 {
		diff2 := data[i+2] - data[i]
		if diff2 <= 3 {
			// fmt.Println("Variant2:", data[2:])
			variants += getVariants(data, i+2)
		}
		if diff2 < 3 && len(data) > 3 {
			diff3 := data[i+3] - data[i]
			if diff3 <= 3 {
				// fmt.Println("Variant3:", data[3:])
				variants += getVariants(data, i+3)
			}
		}
	}
	variants += getVariants(data, i+1)
	variantCache[i] = variants
	return variants
}

func main() {
	var data []int

	reader, err := os.Open("day10.txt")
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

	data = append(data, 0)
	sort.Ints(data)
	data = append(data, data[len(data)-1]+3)

	variantCache = make([]int, len(data))

	oneCount := 0
	threeCount := 0
	for i := 1; i < len(data); i++ {
		diff := data[i] - data[i-1]
		if diff == 1 {
			oneCount++
		} else if diff == 3 {
			threeCount++
		}
	}
	fmt.Println("Part 1:", oneCount*threeCount)

	fmt.Println("Part 2:", getVariants(data, 0))
}
