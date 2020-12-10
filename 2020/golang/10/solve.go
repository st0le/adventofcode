package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	arr := make([]int, 0)
	for scan.Scan() {
		n, _ := strconv.Atoi(scan.Text())
		arr = append(arr, n)
	}

	sort.Ints(arr)
	arr = append(arr, arr[len(arr)-1]+3)
	target := arr[len(arr)-1]
	// fmt.Println(arr)

	// part 1
	freq := make(map[int]int)
	freq[1], freq[2], freq[3] = 0, 0, 0
	prev := 0
	for _, v := range arr {
		freq[v-prev]++
		prev = v
	}
	fmt.Println(freq[1] * freq[3])

	// part 2
	ways := make(map[int]int64)
	ways[0] = 1
	for _, v := range arr {
		ways[v] = ways[v-1] + ways[v-2] + ways[v-3]
	}
	fmt.Println(ways[target])
}
