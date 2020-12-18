package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type expr interface {
	Eval() int
	Eval2() int
	String() string
}

type constant struct {
	value int
}

func (c constant) Eval() int {
	return c.value
}

func (c constant) Eval2() int {
	return c.value
}

func (c constant) String() string {
	return strconv.Itoa(c.value)
}

type operations struct {
	operators   []byte
	expressions []expr
}

func (o operations) Eval() int {
	lhs := o.expressions[0].Eval()
	for i := 1; i < len(o.expressions); i++ {
		rhs := o.expressions[i].Eval()
		if o.operators[i-1] == '+' {
			lhs += rhs
		} else if o.operators[i-1] == '*' {
			lhs *= rhs
		}
	}
	return lhs
}

func (o operations) Eval2() int {
	values := make([]int, len(o.expressions))
	for i, e := range o.expressions {
		values[i] = e.Eval2()
	}
	sums := []int{values[0]}
	for i := 0; i < len(o.operators); i++ {
		if o.operators[i] == '+' {
			sums[len(sums)-1] += values[i+1]
		} else {
			sums = append(sums, values[i+1])
		}
	}
	result := sums[0]
	for i := 1; i < len(sums); i++ {
		result *= sums[i]
	}
	return result
}

func (o operations) String() string {
	str := o.expressions[0].String()
	for i := 1; i < len(o.expressions); i++ {
		str += " " + string(o.operators[i-1]) + " " + o.expressions[i].String()
	}
	return str
}

type brackets struct {
	contents expr
}

func (b brackets) Eval() int {
	return b.contents.Eval()
}

func (b brackets) Eval2() int {
	return b.contents.Eval2()
}

func (b brackets) String() string {
	return "(" + b.contents.String() + ")"
}

func parse(line string) expr {
	// fmt.Println("in", line)
	i := len(line) - 1
	var rhs expr
	var index int = -1
	if line[i] >= '0' && line[i] <= '9' {
		index = strings.LastIndexAny(line, "+*")
		var value int
		if index >= 0 {
			value, _ = strconv.Atoi(line[index+1:])
		} else {
			value, _ = strconv.Atoi(line)
		}
		rhs = constant{value}
	} else if line[i] == ')' {
		index = i - 1
		depth := 1
		for depth > 0 && index >= 0 {
			if line[index] == '(' {
				depth--
			} else if line[index] == ')' {
				depth++
			}
			if depth > 0 {
				index--
			}
		}
		if index < 0 {
			rhs = brackets{parse(line[:i])}
		} else {
			rhs = brackets{parse(line[index+1 : i])}
			index--
		}
	} else {
		panic("Invalid state: " + string(line[i]))
	}
	if index < 0 {
		// fmt.Println("out 0 +", rhs)
		return rhs
	}
	// fmt.Println("out", line[:index], string(line[index]), rhs)
	return operations{[]byte{line[index]}, []expr{parse(line[:index]), rhs}}
}

func parse2(line string) expr {
	var operators []byte
	var expressions []expr
	for len(line) > 0 {
		if line[0] >= '0' && line[0] <= '9' {
			index := strings.IndexAny(line, "+*")
			var value int
			if index >= 0 {
				value, _ = strconv.Atoi(line[:index])
				line = line[index:]
			} else {
				value, _ = strconv.Atoi(line)
				line = ""
			}
			expressions = append(expressions, constant{value})
		} else if line[0] == '(' {
			depth := 1
			i := 1
			for depth > 0 && i < len(line) {
				if line[i] == '(' {
					depth++
				} else if line[i] == ')' {
					depth--
				}
				if depth > 0 {
					i++
				}
			}
			expressions = append(expressions, brackets{parse2(line[1:i])})
			line = line[i+1:]
		} else if line[0] == '+' || line[0] == '*' {
			operators = append(operators, line[0])
			line = line[1:]
		} else {
			panic("Ah")
		}
	}
	if len(expressions) == 1 {
		return expressions[0]
	}
	return operations{operators, expressions}
}

func main() {
	var data []expr
	var data2 []expr

	reader, err := os.Open("day18.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println("Input:", line)
		if len(line) > 0 {
			minified := strings.Map(func(r rune) rune {
				switch {
				case r >= '0' && r <= '9':
					return r
				case r == '+' || r == '*' || r == '(' || r == ')':
					return r
				default:
					return -1
				}
			}, line)
			data = append(data, parse(minified))
			data2 = append(data2, parse2(minified))
		}
	}
	reader.Close()

	// fmt.Println()
	sum := 0
	for _, e := range data {
		eval := e.Eval()
		// fmt.Println(e, "=", eval)
		sum += eval
	}
	fmt.Println("Part 1:", sum)

	sum = 0
	for _, e := range data2 {
		eval := e.Eval2()
		// fmt.Println(e, "=", eval)
		sum += eval
	}
	fmt.Println("Part 2:", sum)
}
