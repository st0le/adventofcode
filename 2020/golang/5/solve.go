package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decode(pass string) int64 {
	fb := strings.ReplaceAll(strings.ReplaceAll(pass[0:7], "F", "0"), "B", "1")
	lr := strings.ReplaceAll(strings.ReplaceAll(pass[7:], "L", "0"), "R", "1")
	row, _ := strconv.ParseInt(fb, 2, 64)
	col, _ := strconv.ParseInt(lr, 2, 64)
	return row*8 + col

}

func main() {
	min := int64(10000)
	max := int64(0)
	sum := int64(0)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		pass := scan.Text()
		id := decode(pass)
		if min > id {
			min = id
		}
		if max < id {
			max = id
		}
		sum += id
	}
	fmt.Println(max)
	fmt.Println((((min + max) * (max - min + 1)) / 2) - sum)
}
