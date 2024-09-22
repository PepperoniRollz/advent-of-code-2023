package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/pepperonirollz/advent-of-code-2023/pkg/utils"
)

func main() {
	solveDayOne("../inputs/day1-1.txt")
	solvePart2("../inputs/day1-1.txt")
}

func solveDayOne(filename string) {
	var digits []rune
	sum := 0
	lines := utils.Parse(filename)
	for _, line := range lines {
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}
		num, err := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
		if err != nil {
			panic(err)
		}
		sum += num
		digits = digits[:0]
	}

	fmt.Println("Part 1 sum: ", sum)

}

func solvePart2(filename string) {
	digitMap := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
		"zero":  '0',
	}
	sum := 0
	lines := utils.Parse(filename)
	for _, line := range lines {

		indexMap := make(map[int]string)
		var values []string
		for word, _ := range digitMap {
			start := 0
			for {
				index := strings.Index(line[start:], word)
				if index == -1 {
					break
				}
				realIndex := start + index
				start = realIndex + len(word)
				indexMap[realIndex] = word
			}
		}
		for i, char := range line {
			if unicode.IsDigit(char) {
				indexMap[i] = string(char)
			}
		}
		for i, _ := range line {
			value, ok := indexMap[i]
			if ok {
				v, ok := digitMap[value]
				if ok {
					values = append(values, string(v))
				} else {
					values = append(values, value)
				}
			}

		}

		num, err := strconv.Atoi(string(values[0]) + string(values[len(values)-1]))
		if err != nil {
			panic(err)
		}
		sum += num

		values = values[:0]
	}
	fmt.Println("Part 2 sum: ", sum)
}
