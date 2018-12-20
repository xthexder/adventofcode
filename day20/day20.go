package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

// 0: Max continuable length, 1: Max length if stopping
func (n *Node) maxLength() [2]int {
	if n == nil {
		return [2]int{0, 0}
	}

	mirror := len(n.value)%2 == 0
	for i := 0; mirror && i < len(n.value)/2; i++ {
		opposite := n.value[len(n.value)-i-1]
		switch n.value[i] {
		case 'N':
			mirror = opposite == 'S'
		case 'S':
			mirror = opposite == 'N'
		case 'E':
			mirror = opposite == 'W'
		case 'W':
			mirror = opposite == 'E'
		default:
			panic("Unknown direction")
		}
	}

	max := [2]int{0, 0}
	for _, c := range n.children {
		tmp := c.maxLength()
		if tmp[0] > max[0] {
			max[0] = tmp[0]
		}
		if tmp[1] > max[1] {
			max[1] = tmp[1]
		}
	}
	nextLen := n.next.maxLength()
	if mirror {
		if max[1] < len(n.value)/2 {
			max[1] = len(n.value) / 2
		}
		if max[1] < max[0]+nextLen[1] {
			max[1] = max[0] + nextLen[1]
		}
	} else {
		max[0] += len(n.value)
		max[1] += len(n.value)

		if max[1] < max[0]+nextLen[1] {
			max[1] = max[0] + nextLen[1]
		}
	}

	if max[0] > max[1] {
		max[1] = max[0]
	}
	return max
}

func main() {
	var regex []byte

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
	fmt.Println("Part A:", root.maxLength()[1])
}
