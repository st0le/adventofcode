package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decode(pass string) int64 {
	pass = strings.ReplaceAll(strings.ReplaceAll(pass, "F", "0"), "B", "1")
	pass = strings.ReplaceAll(strings.ReplaceAll(pass, "L", "0"), "R", "1")
	id, _ := strconv.ParseInt(pass, 2, 64)
	return id
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
