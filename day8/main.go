package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func input() []string {
	dat, _ := ioutil.ReadFile("input")
	str := strings.TrimSuffix(string(dat), "\n")
	inputs := strings.Split(str, "\n")
	return inputs
}

var accum int

func main() {
	accum = 0
	soln1()
	soln2()
}

func soln1() {
	i := input()
	cpu := CPU{}
	interrupt := cpu.LoadAndExec(i)
	fmt.Println("Soln1: Interrupt found: ", interrupt, "accum val", cpu.accum)
}

func soln2() {
	i := input()

	for n, cmdRow := range i {
		cmd, op := parseCmd(cmdRow)
		var newCmdRow string
		if cmd == "nop" {
			newCmdRow = fmt.Sprintf("%s %d", "jmp", op)
		} else if cmd == "jmp" {
			newCmdRow = fmt.Sprintf("%s %d", "nop", op)
		} else {
			continue
		}

		j := append([]string(nil), i...)
		j[n] = newCmdRow

		cpu := CPU{}
		interrupt := cpu.LoadAndExec(j)
		if interrupt == ExecutionEnd {
			fmt.Println("Soln2: Accum value", cpu.accum)
			break
		}
	}
}

func parseCmd(cmdRow string) (string, int) {
	splits := strings.Split(cmdRow, " ")
	op, _ := strconv.Atoi(splits[1])
	return splits[0], op
}

type CPU struct {
	accum int
	pos   int
}

type Interrupt int

const (
	InfiniteLoop Interrupt = iota
	ExecutionEnd
)

func (cpu *CPU) LoadAndExec(input []string) Interrupt {
	totalInstructions := len(input)
	previousExecutionPos := make(map[int]int)

	for {
		if cpu.pos >= totalInstructions {
			return ExecutionEnd
		}

		cmdRow := input[cpu.pos]
		if _, ok := previousExecutionPos[cpu.pos]; ok {
			return InfiniteLoop
		}

		previousExecutionPos[cpu.pos] = 1

		cmd, op := parseCmd(cmdRow)

		switch cmd {
		case "nop":
			cpu.pos++
		case "acc":
			cpu.accum = cpu.accum + op
			cpu.pos++
		case "jmp":
			cpu.pos = cpu.pos + op
		default:
			fmt.Println("Unknown command found:", cmd)
			break
		}
	}
}
