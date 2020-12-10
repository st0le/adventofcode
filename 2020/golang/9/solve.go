package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func canAdd(arr []int, lo, hi, target int) bool {
	for j := lo; j < hi; j++ {
		for k := lo + 1; k < hi; k++ {
			if arr[j]+arr[k] == target {
				return true
			}
		}
	}
	return false
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max2(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findSum(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		min, max, s := arr[i], arr[i], arr[i]
		for j := i + 1; j < len(arr) && s < target; j++ {
			min = min2(min, arr[j])
			max = max2(max, arr[j])
			s += arr[j]

			if s == target {
				return min + max
			}
		}
	}
	return -1
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	arr := make([]int, 0)
	for scan.Scan() {
		n, _ := strconv.Atoi(scan.Text())
		arr = append(arr, n)
	}

	N := 25
	if len(arr) < N {
		N = 5
	}

	first := 0
	for i, v := range arr {
		if i >= N {
			if !canAdd(arr, i-N, i, v) {
				first = v
				break
			}
		}
	}

	fmt.Println(first)
	fmt.Println(findSum(arr, first))

}
