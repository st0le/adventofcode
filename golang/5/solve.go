package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(operations []string) []int {
	output := make([]int, 0)
	opcode := make([]int, len(operations))
	for i, v := range operations {
		opcode[i], _ = strconv.Atoi(v)
	}
	i := 0
	for opcode[i] != 99 {

		hold := opcode[i]
		op := hold % 100
		hold /= 100
		pc := hold % 10
		hold /= 10
		pb := hold % 10
		hold /= 10
		pa := hold % 10

		switch op {
		case 1:

			if 


			opcode[opcode[i+3]] = opcode[opcode[i+1]] + opcode[opcode[i+2]]
			i += 4
		case 2:
			opcode[opcode[i+3]] = opcode[opcode[i+1]] * opcode[opcode[i+2]]
			i += 4
		case 3:
			opcode[opcode[i+1]] = opcode[i+1]
			i += 2
		case 4:
			output = append(output, opcode[i+1])
			i += 2
		default:
			println("something went wrong")
		}
	}
	return output
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	operations := strings.Split(scanner.Text(), ",")
	println(solve(operations, 12, 2))

}
