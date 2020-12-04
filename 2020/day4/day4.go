package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type passport struct {
	byr string // Birth Year
	iyr string // Issue Year
	eyr string // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid string // Passport ID
	cid string // Country ID
}

func (p passport) ValidPart1() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func (p passport) ValidPart2() bool {
	if v, err := strconv.Atoi(p.byr); err != nil || v < 1920 || v > 2002 {
		return false
	}
	if v, err := strconv.Atoi(p.iyr); err != nil || v < 2010 || v > 2020 {
		return false
	}
	if v, err := strconv.Atoi(p.eyr); err != nil || v < 2020 || v > 2030 {
		return false
	}
	if strings.HasSuffix(p.hgt, "cm") {
		if v, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "cm")); err != nil || v < 150 || v > 193 {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		if v, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "in")); err != nil || v < 59 || v > 76 {
			return false
		}
	} else {
		return false
	}
	if strings.HasPrefix(p.hcl, "#") && len(p.hcl) == 7 {
		trimmed := strings.TrimRightFunc(p.hcl, func(r rune) bool {
			switch {
			case r >= '0' && r <= '9':
				return true
			case r >= 'a' && r <= 'f':
				return true
			default:
				return false
			}
		})
		if len(trimmed) != 1 {
			return false
		}
	} else {
		return false
	}
	if p.ecl != "amb" && p.ecl != "blu" && p.ecl != "brn" && p.ecl != "gry" && p.ecl != "grn" && p.ecl != "hzl" && p.ecl != "oth" {
		return false
	}
	if len(p.pid) != 9 {
		return false
	}
	trimmed := strings.TrimRightFunc(p.pid, func(r rune) bool {
		switch {
		case r >= '0' && r <= '9':
			return true
		default:
			return false
		}
	})
	return len(trimmed) == 0
}

func main() {
	var data []passport

	reader, err := os.Open("day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	var pass passport
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			data = append(data, pass)
			pass = passport{}
		}
		parts := strings.Split(line, " ")
		if len(parts) > 0 {
			for _, part := range parts {
				kv := strings.Split(part, ":")
				if kv[0] == "byr" {
					pass.byr = kv[1]
				} else if kv[0] == "iyr" {
					pass.iyr = kv[1]
				} else if kv[0] == "eyr" {
					pass.eyr = kv[1]
				} else if kv[0] == "hgt" {
					pass.hgt = kv[1]
				} else if kv[0] == "hcl" {
					pass.hcl = kv[1]
				} else if kv[0] == "ecl" {
					pass.ecl = kv[1]
				} else if kv[0] == "pid" {
					pass.pid = kv[1]
				} else if kv[0] == "cid" {
					pass.cid = kv[1]
				}
			}
		}
	}
	reader.Close()

	validCount := 0
	for _, pass := range data {
		if pass.ValidPart1() {
			validCount++
		}
	}
	fmt.Println("Part 1:", validCount)

	validCount = 0
	for _, pass := range data {
		if pass.ValidPart2() {
			validCount++
		}
	}
	fmt.Println("Part 2:", validCount)
}
