package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	people := 0
	part1 := 0
	part2 := 0
	answers := make(map[byte]int)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		if len(line) == 0 {

			part1 += len(answers)
			for _, v := range answers {
				if v == people {
					part2++
				}
			}
			answers = make(map[byte]int)
			people = 0

		} else {
			people++
			for _, c := range line {
				v, _ := answers[byte(c)]
				answers[byte(c)] = v + 1
			}
		}
	}
	part1 += len(answers)
	for _, v := range answers {
		if v == people {
			part2++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)

}
