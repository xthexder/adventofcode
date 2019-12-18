package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var reactions map[string]reaction

type element struct {
	name     string
	quantity int64
}

func (e element) String() string {
	return strconv.FormatInt(e.quantity, 10) + " " + e.name
}

func newElement(input []string) element {
	quantity, _ := strconv.Atoi(input[0])
	return element{
		input[1],
		int64(quantity),
	}
}

type reaction struct {
	input  []element
	result element
}

func (r reaction) String() string {
	str := ""
	for _, e := range r.input {
		if e.quantity != 0 {
			if len(str) > 0 {
				str += ", "
			}
			str += e.String()
		}
	}
	str += " => " + r.result.String()
	return str
}

func (r reaction) Copy() reaction {
	newR := reaction{
		make([]element, len(r.input)),
		r.result,
	}
	copy(newR.input, r.input)
	return newR
}

func (r reaction) Simplify() reaction {
	allOre := false
	for !allOre {
		allOre = true
		for i, e := range r.input {
			if e.name == "ORE" || e.quantity <= 0 {
				continue
			}
			allOre = false
			simplified := reactions[e.name].Copy()
			multiple := int64(math.Ceil(float64(e.quantity) / float64(simplified.result.quantity)))
			for j := range simplified.input {
				simplified.input[j].quantity *= multiple
			}
			simplified.result.quantity *= multiple
			r.input[i].quantity -= simplified.result.quantity
			for j := range simplified.input {
				found := false
				for k := range r.input {
					if r.input[k].name == simplified.input[j].name {
						r.input[k].quantity += simplified.input[j].quantity
						found = true
						break
					}
				}
				if !found {
					r.input = append(r.input, simplified.input[j])
				}
			}
		}
	}

	return r
}

func newReaction(input []string) reaction {
	r := reaction{
		nil,
		newElement(input[len(input)-2:]),
	}
	for i := 0; i < len(input)-2; i += 2 {
		r.input = append(r.input, newElement(input[i:i+2]))
	}
	return r
}

func main() {
	reactions = make(map[string]reaction)

	reader, err := os.Open("day14.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return false
			case r >= 'A' && r <= 'Z':
				return false
			default:
				return true
			}
		})
		r := newReaction(fields)
		reactions[r.result.name] = r
	}
	reader.Close()

	simple := reactions["FUEL"].Copy().Simplify()
	var total int64 = 0
	for _, e := range simple.input {
		if e.quantity > 0 && e.name == "ORE" {
			total += e.quantity
		}
	}
	fmt.Println(simple)
	fmt.Println("Part 1:", total)

	max := reactions["FUEL"].Copy()
	max.input = append(max.input, element{"ORE", -1000000000000})
	max.Simplify()
	total = 0
	for _, e := range max.input {
		if e.quantity > 0 && e.name == "ORE" {
			total += e.quantity
		} else {
			// max.input[i].quantity = 0
		}
	}
	fmt.Println("Part 2:")
	fmt.Println(max)
}
