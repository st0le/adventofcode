package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var arr map[int]bool = make(map[int]bool)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		var v int
		fmt.Sscanf(scan.Text(), "%d", &v)
		arr[v] = true
		_, ok := arr[2020-v]
		if ok {
			fmt.Println(v * (2020 - v))
		}
	}

	for a := range arr {
		for b := range arr {
			if a < b {
				c := 2020 - a - b
				if b < c {
					if _, ok := arr[c]; ok {
						fmt.Println(a * b * c)
					}
				}
			}
		}
	}
}
