package main

import (
	"bufio"
	"fmt"
	"os"
)

func treeCount(lines []string, dx int, dy int) int {
	x := 0
	y := 0
	c := 0
	for x < len(lines) {
		if lines[x][y] == '#' {
			c++
		}
		x, y = x+dx, (y+dy)%len(lines[0])
	}
	return c
}

func main() {

	lines := make([]string, 0)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	fmt.Println(treeCount(lines, 1, 3))
	fmt.Println(treeCount(lines, 1, 1) * treeCount(lines, 1, 3) * treeCount(lines, 1, 5) * treeCount(lines, 1, 7) * treeCount(lines, 2, 1))
}
