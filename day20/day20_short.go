package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point [2]int

var distMap map[Point]int

func parseRegex(regex []byte) *Node {
	if len(regex) == 0 {
		return &Node{}
	}
	node := &Node{}
	nest := 0
	sectionStart := 0
	for i, b := range regex {
		switch b {
		case '(':
			nest++
			if nest == 1 {
				node.value = regex[:i]
				sectionStart = i + 1
			}
		case '|':
			if nest == 1 {
				node.children = append(node.children, parseRegex(regex[sectionStart:i]))
				sectionStart = i + 1
			}
		case ')':
			if nest == 1 {
				node.children = append(node.children, parseRegex(regex[sectionStart:i]))
				if i < len(regex)-1 {
					node.next = parseRegex(regex[i+1:])
				}
				return node
			}
			nest--
		}
	}
	node.value = regex
	return node
}

type Node struct {
	value    []byte
	children []*Node
	next     *Node
}

func (n *Node) traverse(pos Point) []Point {
	if n == nil {
		return []Point{pos}
	}

	dist := distMap[pos]
	for _, v := range n.value {
		switch v {
		case 'N':
			pos[1]--
		case 'S':
			pos[1]++
		case 'E':
			pos[0]++
		case 'W':
			pos[0]--
		default:
			panic("Unknown direction")
		}
		dist++
		if oldDist, ok := distMap[pos]; !ok || oldDist > dist {
			distMap[pos] = dist
		}
	}
	if len(n.children) == 0 {
		return n.next.traverse(pos)
	}

	var ends []Point
	endSet := make(map[Point]struct{})
	for _, child := range n.children {
		childEnds := child.traverse(pos)
		for _, end := range childEnds {
			endSet[end] = struct{}{}
		}
	}
	for end, _ := range endSet {
		ends = append(ends, n.next.traverse(end)...)
	}
	return ends
}

func main() {
	var regex []byte
	distMap = make(map[Point]int)

	reader, err := os.Open("day20.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		regex = scanner.Bytes()
		regex = regex[1 : len(regex)-1]
	}
	reader.Close()

	root := parseRegex(regex)

	root.traverse(Point{0, 0})

	max := 0
	roomCount := 0
	for _, dist := range distMap {
		if dist > max {
			max = dist
		}
		if dist >= 1000 {
			roomCount++
		}
	}
	fmt.Println("Part A:", max)
	fmt.Println("Part A:", roomCount)
}
