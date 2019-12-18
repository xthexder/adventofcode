package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var dirs [][2]int = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type branch struct {
	key1, key2       byte
	weight1, weight2 int
}

func main() {
	var board [][]byte
	var weight [][]int

	reader, err := os.Open("day18.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		board = append(board, line)
		weight = append(weight, make([]int, len(line)))
	}
	reader.Close()

	for y := range board {
		fmt.Println(string(board[y]))
	}
	fmt.Println()

	freeKeys := make(map[byte]bool)
	depends := make(map[byte][]byte)
	var branches []branch

	// Simpify board
	end := false
	for !end {
		end = true
		for y := range board {
			for x, ch := range board[y] {
				door := ch >= 'A' && ch <= 'Z'
				key := ch >= 'a' && ch <= 'z'
				if ch == '.' || key || door {
					count := 0
					for _, dir := range dirs {
						if board[y+dir[0]][x+dir[1]] == '#' {
							count++
						}
					}
					if count == 3 {
						if door {
							freeKeys[ch-'A'+'a'] = true
						}
						if key {
							for _, dir := range dirs {
								ch2 := board[y+dir[0]][x+dir[1]]
								if ch2 >= 'a' && ch2 <= 'z' {
									branches = append(branches, branch{
										ch, ch2,
										weight[y][x] + 1, weight[y+dir[0]][x+dir[1]],
									})
								}
								_, free := freeKeys[ch2]
								requires := (ch2 >= 'A' && ch2 <= 'Z') || (ch2 >= 'a' && ch2 <= 'z')
								if requires {
									depends[ch] = append(depends[ch], ch2)
								}
								if ch2 == '.' || free || requires {
									board[y+dir[0]][x+dir[1]] = ch
									weight[y+dir[0]][x+dir[1]] = weight[y][x] + 1
									board[y][x] = '#'
									end = false
									break
								}
							}
						} else {
							board[y][x] = '#'
							end = false
						}
					}
				}
			}
		}
	}

	// Print board
	for y := range board {
		fmt.Println(string(board[y]))
	}
	fmt.Println()

	fmt.Println("Weights:")
	for y := range board {
		for x, ch := range board[y] {
			if ch >= 'a' && ch <= 'z' {
				fmt.Println(string(ch), weight[y][x])
			}
		}
	}
	fmt.Println()

	fmt.Println("Free keys:")
	for k := range freeKeys {
		fmt.Print(string(k), " ")
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Dependencies:")
	for k, dep := range depends {
		fmt.Print(string(k), " =>")
		for _, v := range dep {
			fmt.Print(" ", string(v))
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Branches:")
	for _, branch := range branches {
		fmt.Println(string(branch.key1), strconv.Itoa(branch.weight1), "==", string(branch.key2), strconv.Itoa(branch.weight2))
	}
}
