package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(opcode_s []string, noun int, verb int) int {
	opcode := make([]int, len(opcode_s))
	for i, v := range opcode_s {
		opcode[i], _ = strconv.Atoi(v)
	}
	opcode[1] = noun
	opcode[2] = verb

	i := 0
	for opcode[i] != 99 {

		switch opcode[i] {
		case 1:
			opcode[opcode[i+3]] = opcode[opcode[i+1]] + opcode[opcode[i+2]]
		case 2:
			opcode[opcode[i+3]] = opcode[opcode[i+1]] * opcode[opcode[i+2]]
		default:
			println("something went wrong")
		}
		i += 4
	}
	return opcode[0]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	opcode_s := strings.Split(scanner.Text(), ",")
	println(solve(opcode_s, 12, 2))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if solve(opcode_s, noun, verb) == 19690720 {
				println(noun*100 + verb)
			}
		}
	}
}
