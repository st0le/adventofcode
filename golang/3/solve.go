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

// Line (A,B)
type Line struct {
	A Point
	B Point
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func intersect(u, v Line) (bool, Point) {

	isUHorizontal := u.A.X == u.B.X
	isVHorizontal := v.A.X == v.B.X

	// one is horizontal and other is vertical
	if (isUHorizontal && !isVHorizontal) || (!isUHorizontal && isVHorizontal) {
		horizontal, vertical := u, v
		if isVHorizontal {
			horizontal, vertical = v, u
		}

		if min(horizontal.A.X, horizontal.B.X) <= vertical.A.X && vertical.A.X <= max(horizontal.A.X, horizontal.B.X) &&
			min(vertical.A.Y, vertical.B.Y) <= horizontal.A.Y && horizontal.A.Y <= max(vertical.A.Y, vertical.B.Y) {
			return true, Point{vertical.A.X, horizontal.A.Y}
		}

	}

	return false, Point{}

}

func parse(path string) []Line {

	directions := strings.Split(strings.TrimSpace(path), ",")
	var points []Point = nil
	var lines []Line = nil
	points = append(points, Point{0, 0})
	for i, v := range directions {
		offset, _ := strconv.Atoi(v[1:])
		// fmt.Println(v, offset)

		switch v[0] {
		case 'R':
			points = append(points, Point{points[i].X + offset, points[i].Y})
		case 'L':
			points = append(points, Point{points[i].X - offset, points[i].Y})
		case 'U':
			points = append(points, Point{points[i].X, points[i].Y - offset})
		case 'D':
			points = append(points, Point{points[i].X, points[i].Y + offset})
		}
	}

	for i, point := range points {
		if i > 0 {
			lines = append(lines, Line{points[i-1], point})
		}
	}

	return lines
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewReader(file)
	path1, _ := scanner.ReadString('\n')
	path2, _ := scanner.ReadString('\n')
	lines1, lines2 := parse(path1), parse(path2)
	fmt.Println(lines1)
	fmt.Println(lines2)

	for i, l1 := range lines1 {
		for j, l2 := range lines2 {
			if i > 0 && j > 0 {
				hasPoint, point := intersect(l1, l2)
				if hasPoint {
					fmt.Println(point, Abs(point.X)+Abs(point.Y))
				}
			}
		}
	}
}
