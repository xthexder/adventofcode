package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(`
#include <stdio.h>

int main() {
`)

	fmt.Println("    int a = 1;")
	for i := 1; i < 8; i++ {
		fmt.Println("    int", string(i+'a'), "= 0;")
	}
	fmt.Println()

	reader, err := os.Open("day23.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	targets := make(map[int]struct{})

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		var line string
		if split[0] == "set" {
			line = "    " + split[1] + " = " + split[2] + ";"
		} else if split[0] == "sub" {
			line = "    " + split[1] + " -= " + split[2] + ";"
		} else if split[0] == "mul" {
			line = "    " + split[1] + " *= " + split[2] + ";"
		} else if split[0] == "jnz" {
			target, _ := strconv.Atoi(split[2])
			target += len(lines)
			if split[1] != "1" {
				line = "    if (" + split[1] + " != 0) goto line" + strconv.Itoa(target) + ";"
			} else {
				line = "    goto line" + strconv.Itoa(target) + ";"
			}
			targets[target] = struct{}{}
		} else {
			fmt.Println("Unknown command:", split)
		}
		lines = append(lines, line)

	}
	reader.Close()

	for i, line := range lines {
		if _, ok := targets[i]; ok {
			fmt.Println("line" + strconv.Itoa(i) + ":")
		}
		fmt.Println(line)
	}

	fmt.Println("line" + strconv.Itoa(len(lines)) + ":")
	fmt.Println("    printf(\"h=%d\\n\", h);")
	fmt.Println("}")
}
