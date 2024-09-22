package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/pepperonirollz/advent-of-code-2023/pkg/utils"
)

func main() {
	fmt.Println(SolvePart1("../inputs/day4-1.txt"))
	fmt.Println(SolvePart2("../inputs/day4-1.txt"))
	// fmt.Println(SolvePart2Graph("../inputs/day4-1.txt"))
}

func SolvePart1(filename string) int {
	lines := utils.Parse(filename)
	sum := 0
	for i, line := range lines {
		card := parseLine(line, i)
		sum += int(math.Pow(2, float64(count(card)-1)))
	}
	return sum
}

func SolvePart2(filename string) int {
	lines := utils.Parse(filename)
	var q []Card
	originals := 0
	copies := 0
	for i, line := range lines {
		originals++
		card := parseLine(line, i)

		q = append(q, card)
		for len(q) > 0 {
			c := q[0]
			q = q[1:]
			i = c.cardNum
			numCopies := int(math.Min(float64(count(c)), float64(len(lines)-1-c.cardNum)))
			copies += numCopies
			for j := 1; j < numCopies+1; j++ {
				nextCard := parseLine(lines[i+j], i+j)
				q = append(q, nextCard)
			}
		}
	}
	return originals + copies
}

func SolvePart2Graph(filename string) int {
	lines := utils.Parse(filename)
	var q []*Card
	visited := make(map[int]bool)
	var cards []*Card
	for i, line := range lines {
		card := parseLine(line, i)
		numCopies := int(math.Min(float64(count(card)), float64(len(lines)-1-card.cardNum)))
		card.count = numCopies
		cards = append(cards, &card)
		fmt.Println(card.count)
	}
	for i, card := range cards {
		q = append(q, card)
		for len(q) > 0 {
			currentCard := q[0]
			fmt.Println(currentCard)
			q = q[1:]
			i = currentCard.cardNum
			sum := 0
			for j := 0; j < card.count; j++ {
				nextCard := cards[i+j+1]
				sum += nextCard.count
				if !visited[nextCard.cardNum] {
					q = append(q, nextCard)
				}
			}
			currentCard.sumCopies = sum
		}
	}
	s := 0
	for _, card := range cards {
		s += card.sumCopies
	}
	return s
}
func parseLine(line string, cardNum int) Card {
	split := strings.Split(line, ":")
	split = strings.Split(split[1], "|")
	winners := split[0]
	winnersArr := strings.Fields(winners)
	picks := split[1]
	picksArr := strings.Fields(picks)
	var intWinners []int
	var intPicks []int
	for _, winner := range winnersArr {
		num, _ := strconv.Atoi(winner)
		intWinners = append(intWinners, num)
	}
	for _, pick := range picksArr {
		num, _ := strconv.Atoi(pick)
		intPicks = append(intPicks, num)
	}
	sort.Ints(intWinners)

	card := Card{cardNum: cardNum, winners: intWinners, picks: intPicks}
	return card
}

type Card struct {
	cardNum   int
	winners   []int
	picks     []int
	count     int
	sumCopies int
}

func count(c Card) int {

	count := 0
	for _, pick := range c.picks {
		index := sort.Search(len(c.winners), func(i int) bool {
			return c.winners[i] >= pick
		})
		if index < len(c.winners) && c.winners[index] == pick {
			count++
		}
	}
	c.count = count
	return count
}
