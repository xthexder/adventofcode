package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type ByteSort []byte

func (a ByteSort) Len() int           { return len(a) }
func (a ByteSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByteSort) Less(i, j int) bool { return a[i] < a[j] }

func contains(data []byte, search byte) bool {
	for _, d := range data {
		if d == search {
			return true
		}
	}
	return false
}

func main() {
	deps := make(map[byte][]byte)
	nodeList := make(map[byte]struct{})

	reader, err := os.Open("day7.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		deps[line[36]] = append(deps[line[36]], line[5])
		nodeList[line[5]] = struct{}{}
		nodeList[line[36]] = struct{}{}
	}
	reader.Close()

	totalSeconds := 0
	workers := make([]int, 5)
	workersTarget := make([]byte, 5)
	for {
		for i, workFinished := range workers {
			if workFinished <= totalSeconds && workersTarget[i] > 0 {
				fmt.Println("worker", i, "finished", workersTarget[i])
				delete(nodeList, workersTarget[i])
				workersTarget[i] = 0
			}
		}

		var output []byte
		for node := range nodeList {
			depList, hasDep := deps[node]
			if !hasDep {
				if !contains(workersTarget, node) {
					output = append(output, node)
				}
			} else {
				noDeps := true
				for i, dep := range depList {
					if _, exists := nodeList[dep]; !exists {
						depList[i] = 0
					} else {
						noDeps = false
						break
					}
				}
				if noDeps && !contains(workersTarget, node) {
					output = append(output, node)
				}
			}
		}

		sort.Sort(ByteSort(output))

		freeWorkers := 0
		for i, target := range workersTarget {
			if target == 0 {
				if len(output) > 0 {
					fmt.Println(string(output[0]), "by worker", i)
					workers[i] = totalSeconds + 61 + int(output[0]-'A')
					workersTarget[i] = output[0]
					output = output[1:]
				} else {
					freeWorkers++
				}
			}
		}
		if freeWorkers >= 5 {
			fmt.Println(totalSeconds)
			return
		}

		totalSeconds++
	}
}
