package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type moon struct {
	pos      []int
	velocity []int
}

func (m moon) Copy() moon {
	pos := make([]int, len(m.pos))
	vel := make([]int, len(m.velocity))
	copy(pos, m.pos)
	copy(vel, m.velocity)
	return moon{pos, vel}
}

func (m moon) ApplyGravity(moons []moon) {
	for _, m2 := range moons {
		for i := range m.pos {
			if m.pos[i] != m2.pos[i] {
				if m.pos[i] > m2.pos[i] {
					m.velocity[i]--
				} else {
					m.velocity[i]++
				}
			}
		}
	}
}

func (m moon) ApplyVelocity() {
	for i := range m.pos {
		m.pos[i] += m.velocity[i]
	}
}

func (m moon) Energy() int {
	potential := 0
	kinetic := 0
	for i := range m.pos {
		if m.pos[i] >= 0 {
			potential += m.pos[i]
		} else {
			potential -= m.pos[i]
		}
		if m.velocity[i] >= 0 {
			kinetic += m.velocity[i]
		} else {
			kinetic -= m.velocity[i]
		}
	}
	return potential * kinetic
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(x, y, z int64) int64 {
	result := x * y / GCD(x, y)
	result = result * z / GCD(result, z)
	return result
}

func main() {
	var input []moon

	reader, err := os.Open("day12.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return false
			case r == '-':
				return false
			default:
				return true
			}
		})
		line := make([]int, 3)
		for i, val := range fields {
			line[i], _ = strconv.Atoi(val)
		}
		input = append(input, moon{line, make([]int, 3)})
	}
	reader.Close()

	moons := make([]moon, len(input))
	for i := range input {
		moons[i] = input[i].Copy()
	}
	for i := 0; i < 1000; i++ {
		for _, moon := range moons {
			moon.ApplyGravity(moons)
		}
		for _, moon := range moons {
			moon.ApplyVelocity()
		}
	}
	total := 0
	for _, moon := range moons {
		total += moon.Energy()
	}
	fmt.Println("Part 1:", total)

	fmt.Println("Part 2:")
	for i := range input {
		moons[i] = input[i].Copy()
	}
	for _, moon := range moons {
		fmt.Println(moon)
	}
	loop := make([]int64, 3)
	for step := 0; ; step++ {
		for _, moon := range moons {
			moon.ApplyGravity(moons)
		}
		for _, moon := range moons {
			moon.ApplyVelocity()
		}
		done := true
		for i := 0; i < len(loop); i++ {
			if loop[i] != 0 {
				continue
			}

			loop[i] = int64(step + 1)
			for j, moon := range moons {
				if moon.pos[i] != input[j].pos[i] || moon.velocity[i] != 0 {
					loop[i] = 0
					break
				}
			}
			if loop[i] == 0 {
				done = false
			} else {
				fmt.Println("Axis:", i, "Steps:", (step + 1))
				for _, moon := range moons {
					fmt.Println(moon)
				}
			}
		}
		if done {
			break
		}
	}
	fmt.Println("Least common multiple:", LCM(loop[0], loop[1], loop[2]))
}
