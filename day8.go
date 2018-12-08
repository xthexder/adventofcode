package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var metadataSum int = 0

type Node struct {
	children []*Node
	metadata []int
}

func ReadNode(data []int) (*Node, int) {
	node := &Node{}
	children := data[0]
	metadata := data[1]
	size := 2
	for i := 0; i < children; i++ {
		child, s := ReadNode(data[size:])
		node.children = append(node.children, child)
		size += s
	}
	node.metadata = data[size : size+metadata]
	for _, meta := range node.metadata {
		metadataSum += meta
	}
	size += metadata
	return node, size
}

func GetValue(node *Node) int {
	value := 0
	if len(node.children) == 0 {
		for _, meta := range node.metadata {
			value += meta
		}
	} else {
		for _, meta := range node.metadata {
			if meta > 0 && meta <= len(node.children) {
				value += GetValue(node.children[meta-1])
			}
		}
	}
	return value
}

func main() {
	var data []int

	reader, err := os.Open("day8.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		data = append(data, num)
	}
	reader.Close()

	root, _ := ReadNode(data)

	fmt.Println("Sum:", metadataSum)
	fmt.Println("Root value:", GetValue(root))
}
