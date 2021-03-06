package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point [2]int

var board [2000][2000]byte
var distMap [1000][1000]int

var minx, maxx int = 1000, 1000
var miny, maxy int = 1000, 1000

func minmax(x, y int) {
	if x < minx {
		minx = x
	}
	if x > maxx {
		maxx = x
	}
	if y < miny {
		miny = y
	}
	if y > maxy {
		maxy = y
	}
}

func parseRegex(regex []byte) *Node {
	if len(regex) == 0 {
		// fmt.Println("Blank")
		return &Node{}
	}
	// fmt.Println("Start", string(regex))
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
				// fmt.Println("(: ", string(regex[:i]))
			}
		case '|':
			if nest == 1 {
				// fmt.Println("|: ", string(regex[sectionStart:i]))
				node.children = append(node.children, parseRegex(regex[sectionStart:i]))
				sectionStart = i + 1
			}
		case ')':
			if nest == 1 {
				// fmt.Println("): ", string(regex[sectionStart:i]))
				node.children = append(node.children, parseRegex(regex[sectionStart:i]))
				if i < len(regex)-1 {
					node.next = parseRegex(regex[i+1:])
				}
				// fmt.Println("End2", string(regex))
				return node
			}
			nest--
		}
	}
	node.value = regex
	// fmt.Println("End1", string(regex))
	return node
}

type Node struct {
	value    []byte
	children []*Node
	next     *Node
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}

	out := " [" + string(n.value)
	if len(n.children) > 1 {
		out += "("
	}
	for i, c := range n.children {
		if i > 0 {
			out += "|"
		}
		out += c.String()
	}
	if len(n.children) > 1 {
		out += ")"
	}
	out += n.next.String()
	out += "] "
	return out
}

func (n *Node) traverse(pos Point) []Point {
	if n == nil {
		return []Point{pos}
	}

	dist := distMap[pos[1]][pos[0]]
	for _, v := range n.value {
		switch v {
		case 'N':
			board[pos[1]*2-1][pos[0]*2] = '-'
			pos[1]--
		case 'S':
			board[pos[1]*2+1][pos[0]*2] = '-'
			pos[1]++
		case 'E':
			board[pos[1]*2][pos[0]*2+1] = '|'
			pos[0]++
		case 'W':
			board[pos[1]*2][pos[0]*2-1] = '|'
			pos[0]--
		default:
			panic("Unknown direction")
		}
		dist++
		board[pos[1]*2][pos[0]*2] = '.'
		minmax(pos[0]*2, pos[1]*2)
		if distMap[pos[1]][pos[0]] == 0 || distMap[pos[1]][pos[0]] > dist {
			distMap[pos[1]][pos[0]] = dist
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
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			board[y][x] = '#'
		}
	}
	board[1000][1000] = 'X'

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
	// fmt.Println(root)

	root.traverse(Point{500, 500})

	max := 0
	roomCount := 0
	for y := 0; y < len(distMap); y++ {
		for x := 0; x < len(distMap[y]); x++ {
			dist := distMap[y][x]
			if dist > max {
				max = dist
			}
			if dist >= 1000 {
				roomCount++
			}
		}
	}
	fmt.Println("Part A:", max)
	fmt.Println("Part A:", roomCount)

	for y := miny - 1; y <= maxy+1; y++ {
		fmt.Println(string(board[y][minx-1 : maxx+2]))
	}
}
