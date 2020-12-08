package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func validate2(passport string) bool {
	re := regexp.MustCompile("(\\w+):([#|\\w]+)")
	res := re.FindAllStringSubmatch(passport, -1)
	for i := range res {
		field := res[i][1]
		value := res[i][2]
		switch field {
		case "byr":
			if ok, _ := regexp.MatchString("^(\\d+)$", value); !ok {
				return false
			}
			byr := regexp.MustCompile("^(\\d+)$")
			parts := byr.FindStringSubmatch(value)
			year, err := strconv.Atoi(parts[1])
			if err != nil || !(1920 <= year && year <= 2002) {
				return false
			}

			break
		case "iyr":
			if ok, _ := regexp.MatchString("^(\\d+)$", value); !ok {
				return false
			}
			iyr := regexp.MustCompile("^(\\d+)$")
			parts := iyr.FindStringSubmatch(value)
			year, err := strconv.Atoi(parts[1])
			if err != nil || !(2010 <= year && year <= 2020) {
				return false
			}

			break
		case "eyr":
			if ok, _ := regexp.MatchString("^(\\d+)$", value); !ok {
				return false
			}
			eyr := regexp.MustCompile("^(\\d+)$")
			parts := eyr.FindStringSubmatch(value)
			year, err := strconv.Atoi(parts[1])
			if err != nil || !(2020 <= year && year <= 2030) {
				return false
			}

			break
		case "hgt":
			if ok, _ := regexp.MatchString("^(\\d+)(cm|in)$", value); !ok {
				return false
			}

			hgt := regexp.MustCompile("^(\\d+)(cm|in)$")
			parts := hgt.FindStringSubmatch(value)
			height, err := strconv.Atoi(parts[1])
			if err != nil {
				return false
			}

			if parts[2] == "cm" && !(150 <= height && height <= 193) {
				return false
			}

			if parts[2] == "in" && !(59 <= height && height <= 76) {
				return false
			}

			break
		case "hcl":
			if ok, _ := regexp.MatchString("^#[a-f0-9]{6}$", value); !ok {
				return false
			}
			break
		case "ecl":
			if ok, _ := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", value); !ok {
				return false
			}
			break
		case "pid":
			if ok, _ := regexp.MatchString("^\\d{9}$", value); !ok {
				return false
			}
			break

		}
	}
	return true
}

func validate1(passport string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	re := regexp.MustCompile("(\\w+):")
	res := re.FindAllStringSubmatch(passport, -1)
	fields := make(map[string]bool)
	for i := range res {
		field := res[i][1]
		fields[field] = true
	}

	for _, f := range required {
		if _, ok := fields[f]; !ok {
			return false
		}
	}
	return true
}

func main() {
	valid1 := 0
	valid2 := 0
	block := ""
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		if len(line) != 0 {
			block = block + line + " "
		} else {
			if validate1(block) {
				valid1++
				if validate2(block) {
					valid2++
				}
			}
			block = ""
		}
	}
	if validate1(block) {
		valid1++
		if validate2(block) {
			valid2++
		}
	}

	fmt.Println(valid1)
	fmt.Println(valid2)

}
