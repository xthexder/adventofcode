package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var boost = 0

type Group struct {
	units, hp, attack, initiative int
	defense                       map[string]int
	attackType, team              string
	target                        *Group
	selected                      bool
}

func NewGroup(line, team string) *Group {
	// fmt.Println(line, team)

	group := &Group{}
	split := strings.Split(line, " ")
	group.units, _ = strconv.Atoi(split[0])
	group.hp, _ = strconv.Atoi(split[4])
	group.attack, _ = strconv.Atoi(split[len(split)-6])
	group.initiative, _ = strconv.Atoi(split[len(split)-1])
	group.attackType = split[len(split)-5]
	group.team = team
	group.defense = make(map[string]int)

	if split[7][0] == '(' {
		i := 7
		for {
			if split[i][0] == '(' {
				split[i] = split[i][1:]
			}
			defense := 1
			if split[i] == "immune" {
				defense = 0
			} else if split[i] == "weak" {
				defense = 2
			} else {
				fmt.Println(split[i])
				panic("Unknown defense modifier")
			}
			i += 2
			for {
				defenseType := split[i]
				defenseType = defenseType[:len(defenseType)-1]
				group.defense[defenseType] = defense
				if split[i][len(split[i])-1] == ';' {
					break
				} else if split[i][len(split[i])-1] == ')' {
					break
				} else {
					i++
				}
			}

			if split[i][len(split[i])-1] == ')' {
				break
			} else {
				i++
			}
		}
	}
	// fmt.Println(group)
	return group
}

func (g *Group) EffectivePower() int {
	if g.units <= 0 {
		return 0
	} else if g.team == "immune" {
		return g.units * (g.attack + boost)
	}
	return g.units * g.attack
}

func (g *Group) Damage(target *Group) int {
	if g.units <= 0 || target.units <= 0 {
		return 0
	}
	if defense, ok := target.defense[g.attackType]; ok {
		return g.EffectivePower() * defense
	}
	return g.EffectivePower()
}

func (g *Group) Attack() int {
	if g.target == nil || g.units <= 0 {
		g.selected = false
		g.target = nil
		return 0
	}
	// fmt.Println(g.units, " attacking", g.target.units)
	dmg := g.Damage(g.target)
	lost := dmg / g.target.hp
	// fmt.Println(g.target.units, "loses", lost, "units")
	g.target.units -= lost
	g.selected = false
	g.target = nil
	return lost
}

type GroupSort struct {
	list        []*Group
	sortByPower bool
}

func (g GroupSort) Len() int      { return len(g.list) }
func (g GroupSort) Swap(i, j int) { g.list[i], g.list[j] = g.list[j], g.list[i] }
func (g GroupSort) Less(i, j int) bool {
	if g.sortByPower && g.list[i].EffectivePower() != g.list[j].EffectivePower() {
		return g.list[i].EffectivePower() > g.list[j].EffectivePower()
	} else {
		return g.list[i].initiative > g.list[j].initiative
	}
}

func simulate(groups []*Group) (string, int) {
	done := false
	for !done {
		done = true
		// fmt.Println("Selection phase:")
		sort.Sort(GroupSort{groups, true})
		for _, group := range groups {
			if group.units <= 0 {
				continue
			}
			maxDamage := 0
			var maxTarget *Group
			for _, target := range groups {
				if group == target || target.units <= 0 || group.team == target.team || target.selected {
					continue
				}
				dmg := group.Damage(target)
				// fmt.Println(group.units, "would deal", dmg, "to target", target.units)
				if dmg > maxDamage {
					maxDamage = dmg
					maxTarget = target
				}
			}
			if maxTarget != nil {
				// fmt.Println(group.units, "selected target", maxTarget.units)
				maxTarget.selected = true
			}
			group.target = maxTarget
		}
		// fmt.Println("Attack phase:")
		sort.Sort(GroupSort{groups, false})
		for _, group := range groups {
			if group.Attack() > 0 {
				done = false
			}
		}
		// fmt.Println()
	}

	// fmt.Println()
	count := 0
	team := ""
	for _, group := range groups {
		if group.units > 0 {
			if team != "" && group.team != team {
				return "tie", -1
			}
			// fmt.Println(group.team, group.units)
			count += group.units
			team = group.team
		}
	}
	return team, count
}

func main() {
	var original []*Group
	var groups []*Group

	reader, err := os.Open("day24.txt")
	if err != nil {
		log.Fatal(err)
	}

	team := ""
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "Immune System:" {
			team = "immune"
			continue
		} else if line == "Infection:" {
			team = "infection"
			continue
		} else if len(line) == 0 {
			continue
		}
		groups = append(groups, NewGroup(line, team))
		original = append(original, NewGroup(line, team))
	}
	reader.Close()

	team, count := simulate(groups)
	fmt.Println("Part A:", count)

	for team != "immune" {
		for i := range groups {
			*groups[i] = *original[i]
		}
		boost++
		// fmt.Println("Boost:", boost)
		team, count = simulate(groups)
	}
	fmt.Println("Part B:", count)
}
