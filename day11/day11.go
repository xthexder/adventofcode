package main

import "fmt"

const input = 5791
const gridSize = 300

// 1-indexed, not 0
func powerLevel(x, y int) int {
	rackId := x + 10
	powerLevel := rackId * y
	powerLevel += input
	powerLevel *= rackId
	powerLevel = powerLevel % 1000
	powerLevel /= 100
	powerLevel -= 5
	return powerLevel
}

func findMaxSquare(sumGrid [][]int, size int) (int, int, int) {
	maxX, maxY := 0, 0
	maxPowerLevel := -1000

	for x := 0; x <= gridSize-size; x++ {
		for y := 0; y <= gridSize-size; y++ {

			powerLevel := sumGrid[x+size-1][y+size-1]
			if x > 0 && y > 0 {
				powerLevel += sumGrid[x-1][y-1]
			}
			if x > 0 {
				powerLevel -= sumGrid[x-1][y+size-1]
			}
			if y > 0 {
				powerLevel -= sumGrid[x+size-1][y-1]
			}

			if powerLevel > maxPowerLevel {
				maxX = x + 1
				maxY = y + 1
				maxPowerLevel = powerLevel
			}
		}
	}
	return maxX, maxY, maxPowerLevel
}

func main() {
	grid := make([][]int, gridSize)
	for x := 1; x <= gridSize; x++ {
		grid[x-1] = make([]int, gridSize)
		for y := 1; y <= gridSize; y++ {
			grid[x-1][y-1] = powerLevel(x, y)
		}
	}

	sumGrid := make([][]int, gridSize)
	for x := 0; x < gridSize; x++ {
		sumGrid[x] = make([]int, gridSize)
		sumGrid[x][0] = grid[x][0]
		for y := 1; y < gridSize; y++ {
			sumGrid[x][y] = sumGrid[x][y-1] + grid[x][y]
		}
	}
	for x := 1; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			sumGrid[x][y] += sumGrid[x-1][y]
		}
	}

	maxX, maxY, maxPowerLevel := findMaxSquare(sumGrid, 3)

	fmt.Println("Part A:", maxX, maxY)

	maxSize := 3
	for size := 1; size <= gridSize; size++ {
		x, y, powerLevel := findMaxSquare(sumGrid, size)
		if powerLevel > maxPowerLevel {
			maxX = x
			maxY = y
			maxPowerLevel = powerLevel
			maxSize = size
		}
	}
	fmt.Println("Part B:", maxX, maxY, maxSize)
}
