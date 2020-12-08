package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type context struct {
	pc  int
	acc int
}

type instruction struct {
	op       string
	arg      int
	executed bool
}

func (instr *instruction) Exec(ctx *context) {
	if instr.executed {
		// fmt.Println("Loop detected")
		ctx.pc = -1
		return
	}
	switch instr.op {
	case "acc":
		ctx.acc += instr.arg
		ctx.pc++
	case "jmp":
		ctx.pc += instr.arg
	case "nop":
		ctx.pc++
	default:
		panic("Unknown instruction:" + instr.op)
	}
	instr.executed = true
}

func main() {
	var data []instruction

	reader, err := os.Open("day8.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 1 {
			arg, _ := strconv.Atoi(fields[1])
			data = append(data, instruction{fields[0], arg, false})
		}
	}
	reader.Close()

	ctx := context{0, 0}
	for ctx.pc >= 0 && ctx.pc < len(data) {
		data[ctx.pc].Exec(&ctx)
	}
	fmt.Println("Part 1:", ctx.acc)

	for change := 0; change < len(data); change++ {
		if data[change].op == "nop" {
			data[change].op = "jmp"
		} else if data[change].op == "jmp" {
			data[change].op = "nop"
		} else {
			continue
		}
		ctx = context{0, 0}
		for i := range data {
			data[i].executed = false
		}
		for ctx.pc >= 0 && ctx.pc < len(data) {
			data[ctx.pc].Exec(&ctx)
		}
		if ctx.pc == len(data) {
			fmt.Println("Part 2:", ctx.acc)
			return
		}

		if data[change].op == "nop" {
			data[change].op = "jmp"
		} else if data[change].op == "jmp" {
			data[change].op = "nop"
		}
	}
}
