package main

// The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pepperonirollz/advent-of-code-2023/pkg/utils"
)

func main() {
	fmt.Println(SolvePart1("../inputs/day2-2.txt"))
	fmt.Println(SolvePart2("../inputs/day2-2.txt"))
}

func SolvePart1(filename string) int {
	lines := utils.Parse(filename)
	sum := 0
	for _, line := range lines {
		game := parseGame(line)
		if validateGame(game) {
			sum += game.id
		}
	}

	return sum
}

func SolvePart2(filename string) int {
	lines := utils.Parse(filename)
	sum := 0
	for _, line := range lines {
		game := parseGame(line)
		sum += power(game)
	}
	return sum
}

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	color string
	count int
}

func parseGame(line string) Game {
	round := strings.Split(line, ":")
	id, _ := strconv.Atoi(strings.Fields(round[0])[1])

	setString := strings.Split(round[1], ";")
	var sets []Set
	for _, set := range setString {
		colorFields := strings.Split(set, ",")
		for _, colorField := range colorFields {
			color := strings.Fields(colorField)[1]
			count, _ := strconv.Atoi(strings.Fields(colorField)[0])
			sets = append(sets, Set{color, count})
		}
	}

	return Game{id, sets}
}

func validateGame(g Game) bool {
	for _, set := range g.sets {
		if set.color == "red" && set.count > 12 {
			return false
		}
		if set.color == "green" && set.count > 13 {
			return false
		}
		if set.color == "blue" && set.count > 14 {
			return false
		}
	}
	return true
}

func power(g Game) int {
	m := make(map[string]int)

	for _, set := range g.sets {
		count, ok := m[set.color]
		if ok {
			if count < set.count {
				m[set.color] = set.count
			}
		} else {
			m[set.color] = set.count
		}
	}
	power := 1
	for _, value := range m {
		power *= value
	}
	return power
}
