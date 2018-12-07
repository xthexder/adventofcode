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
    int mulCount = 0;
`)

	for i := 0; i < 8; i++ {
		fmt.Println("    int", string(i+'a'), "= 0;")
	}
	fmt.Println()

	reader, err := os.Open("day23.txt")
	if err != nil {
		log.Fatal(err)
	}

	line := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		fmt.Print("    line" + strconv.Itoa(line) + ":\t")
		if split[0] == "set" {
			fmt.Println("   ", split[1], "=", split[2]+";")
		} else if split[0] == "sub" {
			fmt.Println("   ", split[1], "-=", split[2]+";")
		} else if split[0] == "mul" {
			fmt.Println("   ", split[1], "*=", split[2]+"; mulCount++;")
		} else if split[0] == "jnz" {
			target, _ := strconv.Atoi(split[2])
			target += line
			fmt.Println("    if ("+split[1], "!= 0) goto line"+strconv.Itoa(target)+";")
		} else {
			fmt.Println("Unknown command:", split)
		}
		line++
	}
	reader.Close()

	fmt.Println("    line" + strconv.Itoa(line) + ":")
	fmt.Print("    printf(\"Registers: a=%d")
	for i := 1; i < 8; i++ {
		fmt.Print(" " + string(i+'a') + "=%d")
	}
	fmt.Print("\\n\"")
	for i := 0; i < 8; i++ {
		fmt.Print(", " + string(i+'a'))
	}
	fmt.Println(");")

	fmt.Println("    printf(\"mulCount = %d\\n\", mulCount);")
	fmt.Println("}")
}
