package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func cost(prev map[string]string, memo map[string]int, current string) int {
	c, ok := memo[current]
	if ok {
		return c
	}

	p, ok := prev[current]
	if ok {
		memo[current] = cost(prev, memo, p) + 1
		return memo[current]
	}
	memo[current] = 0
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	prev := make(map[string]string)
	memo := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moons := strings.Split(scanner.Text(), ")")
		prev[moons[1]] = moons[0]
	}

	sum := 0
	for k := range prev {
		sum += cost(prev, memo, k)
	}
	fmt.Println(sum)

	you := "YOU"
	santa := "SAN"
	youCost := cost(prev, memo, you)
	santaCost := cost(prev, memo, santa)

	diff := int(math.Abs(float64(youCost - santaCost)))
	for i := 0; i < diff; i++ {
		if youCost > santaCost {
			you = prev[you]
		} else {
			santa = prev[santa]
		}
	}

	for you != santa {
		diff += 2
		you, santa = prev[you], prev[santa]
		fmt.Println(you, santa)
	}

	// excluding you and santa
	fmt.Println(diff - 2)

}
