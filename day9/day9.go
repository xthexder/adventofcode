package main

import "fmt"

const players = 455
const lastMarble = 71223 // * 100

type Marble struct {
	number int
	prev   *Marble // Counter-clockwise
	next   *Marble // Clockwise
}

func insertAfter(marble *Marble, number int) *Marble {
	newMarble := &Marble{number, marble, marble.next}
	marble.next.prev = newMarble
	marble.next = newMarble
	return newMarble
}

func removeMarble(marble *Marble) *Marble {
	marble.prev.next = marble.next
	marble.next.prev = marble.prev
	return marble
}

func main() {
	current := &Marble{0, nil, nil}
	current.prev = current
	current.next = current

	scores := make([]int, players)

	marble := 1
	for marble <= lastMarble {
		for elf := 0; elf < players && marble <= lastMarble; elf++ {
			if marble%23 == 0 {
				scores[elf] += marble
				removed := removeMarble(current.prev.prev.prev.prev.prev.prev.prev)
				scores[elf] += removed.number
				current = removed.next
			} else {
				current = insertAfter(current.next, marble)
			}
			marble++
		}
	}

	maxScore := 0
	winner := -1
	for elf, score := range scores {
		if score > maxScore {
			winner = elf
			maxScore = score
		}
	}
	fmt.Println(winner, "wins with score", maxScore)
}
