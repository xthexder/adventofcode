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

	for len(nodeList) > 0 {
		var output []byte
		for node := range nodeList {
			depList, hasDep := deps[node]
			if !hasDep {
				output = append(output, node)
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
				if noDeps {
					output = append(output, node)
				}
			}
		}

		sort.Sort(ByteSort(output))

		if len(output) > 0 {
			delete(nodeList, output[0])
		}

		fmt.Print(string(output[0]))
	}
}
