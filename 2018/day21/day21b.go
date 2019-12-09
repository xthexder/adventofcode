package main

import "fmt"

func main() {
	var inputs [0xFFFFFF + 1]bool
	count := 0
	last := -1
	a := 0
	for {
		inputs[a] = true
		b := a | 65536 // [2]
		a = 6663054    // [1]
		for {
			a += b & 0xFF
			a &= 0xFFFFFF
			a *= 65899
			a &= 0xFFFFFF
			if b < 256 {
				break
			}
			b = b >> 8
		}
		if count == 0 {
			fmt.Println("Part A:", a)
		}
		if inputs[a] {
			fmt.Println("Part B:", last)
			break
		}
		last = a
		count++
	}
}
