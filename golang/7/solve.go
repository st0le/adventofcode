package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner := bufio.NewReader(file)
	line, _ := scanner.ReadString('\n')
	w, h := 25, 6
	// w, h := 2, 2
	image := make([]byte, w*h)
	for i := 0; i < w*h; i++ {
		image[i] = ' '
	}

	minZeroCount, minZeroIndex := w*h, 0
	for i := 0; i < len(line); i += w * h {
		zeros := 0
		for j := 0; j < w*h; j++ {
			if line[i+j] == '0' {
				zeros++
			}
			if image[j] == ' ' && line[i+j] != '2' {
				image[j] = line[i+j]
			}
		}

		if zeros < minZeroCount {
			minZeroCount, minZeroIndex = zeros, i
		}
	}

	// fmt.Println(minZeroCount, minZeroIndex)
	ones, twos := 0, 0
	for i := 0; i < w*h; i++ {
		switch line[i+minZeroIndex] {
		case '1':
			ones++
		case '2':
			twos++
		}
	}
	fmt.Println(ones * twos)

	finalImage := string(image)
	for i := 0; i < w*h; i += w {
		fmt.Println(strings.ReplaceAll(strings.ReplaceAll(finalImage[i:i+w], "0", "."), "1", "#"))
	}

}
