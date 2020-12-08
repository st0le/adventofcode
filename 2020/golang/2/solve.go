package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	valid1 := 0
	valid2 := 0
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		var lo, hi int
		var ch byte
		var password string
		fmt.Sscanf(scan.Text(), "%d-%d %c: %s", &lo, &hi, &ch, &password)
		count := strings.Count(password, string([]byte{ch}))
		if lo <= count && count <= hi {
			valid1++
		}

		if (password[lo-1] == ch) != (password[hi-1] == ch) {
			valid2++
		}
	}
	fmt.Println(valid1)
	fmt.Println(valid2)
}
