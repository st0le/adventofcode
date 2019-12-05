package main

// incomplete

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Point (X,Y)
type Point struct {
	X int
	Y int
}

func parse(path string) map[Point]int {

	directions := strings.Split(strings.TrimSpace(path), ",")
	grid := make(map[Point]int)
	current := Point{0, 0}

	step := 0
	grid[current] = step
	for _, v := range directions {
		offset, _ := strconv.Atoi(v[1:])
		// fmt.Println(v, offset)
		for j := 0; j < offset; j++ {
			switch v[0] {
			case 'R':
				current.X++
			case 'L':
				current.X--
			case 'U':
				current.Y--
			case 'D':
				current.Y++
			}
			step++
			_, ok := grid[current]
			if !ok {
				grid[current] = step
			}
		}
	}

	return grid
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewReader(file)
	path1, _ := scanner.ReadString('\n')
	path2, _ := scanner.ReadString('\n')
	grid1, grid2 := parse(path1), parse(path2)

	solution1 := 10000000
	solution2 := 10000000
	for point, step1 := range grid1 {
		step2, ok := grid2[point]
		if ok && Abs(point.X)+Abs(point.Y) > 0 {
			solution1 = min(solution1, Abs(point.X)+Abs(point.Y))
			solution2 = min(solution2, step1+step2)
			// fmt.Println(p1)
		}
	}

	fmt.Println(solution1)
	fmt.Println(solution2)
}
