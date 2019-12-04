package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	solution_1 := 0
	solution_2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		n = (n/3 - 2)
		solution_1 += n
		for n > 0 {
			solution_2 += n
			n = (n/3 - 2)
		}
	}
	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
