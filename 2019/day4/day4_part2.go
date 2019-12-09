package main

import (
	"fmt"
)

// 123257-647015
var min []byte = []byte{1, 2, 3, 2, 5, 7}
var max []byte = []byte{6, 4, 7, 0, 1, 5}

func count(prefix []byte, double bool) int {
	if len(prefix) == len(min) {
		if double || (prefix[len(min)-1] == prefix[len(min)-2] && prefix[len(min)-1] != prefix[len(min)-3]) {
			return 1
		} else {
			return 0
		}
	}
	start := min[len(prefix)]
	end := max[len(prefix)]
	for i, num := range prefix {
		if min[i] != num {
			start = prefix[len(prefix)-1]
		}
		if max[i] != num {
			end = 9
		}
	}
	if start < prefix[len(prefix)-1] {
		start = prefix[len(prefix)-1]
	}

	total := 0
	for i := start; i <= end; i++ {
		newPrefix := make([]byte, len(prefix), len(prefix)+1)
		copy(newPrefix, prefix)
		newPrefix = append(newPrefix, i)
		newDouble := double
		if len(prefix) > 1 && prefix[len(prefix)-1] == prefix[len(prefix)-2] && i != prefix[len(prefix)-1] && (len(prefix) < 3 || prefix[len(prefix)-2] != prefix[len(prefix)-3]) {
			newDouble = true
		}
		total += count(newPrefix, newDouble)
	}
	return total
}

func main() {
	total := 0
	for i := min[0]; i <= max[0]; i++ {
		total += count([]byte{i}, false)
	}
	fmt.Println("Part 2:", total)
}
