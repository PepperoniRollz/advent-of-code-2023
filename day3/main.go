package main

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"unicode"

	"github.com/pepperonirollz/advent-of-code-2023/pkg/utils"
)

func main() {
	fmt.Println(SolvePart1("../inputs/day3-1.txt"))
	fmt.Println(SolvePart2("../inputs/day3-1.txt"))
	fmt.Println(SolvePart2Concurrent("../inputs/day3-1.txt"))

}

func SolvePart1(filename string) int {
	lines := utils.Parse(filename)
	sum := 0
	re := regexp.MustCompile(`\d+`)
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // Up-Left, Up, Up-Right
		{0, -1}, {0, 1}, // Left,     Digit, Right
		{1, -1}, {1, 0}, {1, 1}, // Down-Left, Down, Down-Right
	}
	height := len(lines)
	for i, line := range lines {
		width := len(line)
		matches := re.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			startIndex := match[0]
			endIndex := match[1]
			numStr := line[startIndex:endIndex]
			num, _ := strconv.Atoi(numStr)
			done := false
			for j := startIndex; j < endIndex; j++ {
				if done {
					break
				}
				for _, dir := range directions {
					rowToCheck := i + dir[0]
					colToCheck := j + dir[1]
					if rowToCheck >= 0 && rowToCheck < height && colToCheck >= 0 && colToCheck < width {
						placeToCheck := lines[rowToCheck][colToCheck]
						if placeToCheck != '.' && !unicode.IsDigit(rune(placeToCheck)) {
							sum += num
							done = true
							break
						}
					}
				}
			}

		}
	}
	return sum
}

type Gear struct {
	row int
	col int
}

func SolvePart2(filename string) int {
	lines := utils.Parse(filename)
	sum := 0
	re := regexp.MustCompile(`\d+`)
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // Up-Left, Up, Up-Right
		{0, -1}, {0, 1}, // Left,     Digit, Right
		{1, -1}, {1, 0}, {1, 1}, // Down-Left, Down, Down-Right
	}
	//map gears to the part numbers around them
	m := make(map[Gear][]int)
	height := len(lines)
	for i, line := range lines {
		width := len(line)
		matches := re.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			startIndex := match[0]
			endIndex := match[1]
			numStr := line[startIndex:endIndex]
			num, _ := strconv.Atoi(numStr)
			done := false
			for j := startIndex; j < endIndex; j++ {
				if done {
					break
				}
				for _, dir := range directions {
					rowToCheck := i + dir[0]
					colToCheck := j + dir[1]
					if rowToCheck >= 0 && rowToCheck < height && colToCheck >= 0 && colToCheck < width {
						placeToCheck := lines[rowToCheck][colToCheck]
						if placeToCheck == '*' {
							gear := Gear{row: rowToCheck, col: colToCheck}
							m[gear] = append(m[gear], num)
							done = true
							break
						}
					}
				}
			}

		}
	}
	for _, val := range m {
		if len(val) == 2 {
			sum += val[0] * val[1]
		}
	}
	return sum
}

func SolvePart2Concurrent(filename string) int {
	lines := utils.Parse(filename)
	sum := 0
	re := regexp.MustCompile(`\d+`)
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // Up-Left, Up, Up-Right
		{0, -1}, {0, 1}, // Left,     Digit, Right
		{1, -1}, {1, 0}, {1, 1}, // Down-Left, Down, Down-Right
	}
	m := make(map[Gear][]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	height := len(lines)
	for i, line := range lines {
		wg.Add(1)

		go func(i int, line string) {
			defer wg.Done()
			width := len(line)
			matches := re.FindAllStringSubmatchIndex(line, -1)
			for _, match := range matches {
				startIndex := match[0]
				endIndex := match[1]
				numStr := line[startIndex:endIndex]
				num, _ := strconv.Atoi(numStr)
				done := false
				for j := startIndex; j < endIndex; j++ {
					if done {
						break
					}
					for _, dir := range directions {
						rowToCheck := i + dir[0]
						colToCheck := j + dir[1]
						if rowToCheck >= 0 && rowToCheck < height && colToCheck >= 0 && colToCheck < width {
							placeToCheck := lines[rowToCheck][colToCheck]
							if placeToCheck == '*' {
								mu.Lock()
								m[Gear{row: rowToCheck, col: colToCheck}] = append(m[Gear{row: rowToCheck, col: colToCheck}], num)
								mu.Unlock()
								done = true
								break
							}
						}
					}
				}

			}
		}(i, line)
	}
	wg.Wait()
	for _, val := range m {
		if len(val) == 2 {
			sum += val[0] * val[1]
		}
	}
	return sum
}
