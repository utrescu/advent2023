package main

import (
	"advent2023/utils"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type Card struct {
	instances uint64
	winning []string
	have    []string
}

func (c Card) Matches() []string {
	correctNumbers := []string{}
	for _, card := range c.winning {
		if slices.Contains(c.have, card) {
			correctNumbers = append(correctNumbers, card)
		}
	}
	return correctNumbers
}

func FillCards(lines []string) map[int]Card {
	cards := map[int]Card{}

	for num, line := range lines {
		var re = regexp.MustCompile(`(?m)(\d+)`)

		numbers := strings.Split(line, ":")[1]
		parts := strings.Split(numbers, " |")

		winning := re.FindAllString(parts[0], -1)
		have := re.FindAllString(parts[1], -1)

		cards[num] = Card{instances: 1, winning: winning, have: have}
	}

	return cards
}

func main() {

	var lines, _ = utils.ReadLines("input")
	cards := FillCards(lines)

	points := 0
	for _, card := range cards {
		result := card.Matches()
		if len(result) != 0 {
			roundPoints := 1 
			for i:=0; i < (len(result) - 1); i++ {
				roundPoints *= 2
			}
			points += roundPoints
		}
	}
	fmt.Println("Part 1:", points)

	totalCards := len(cards)
	var part2 uint64 
	for i:=0; i< totalCards; i++ {
		result := len(cards[i].Matches())
		for j:=i+1; j<= i+result; j++ {
			if j < totalCards {
				oldCard := cards[j]
				oldCard.instances += cards[i].instances
				cards[j] = oldCard
			}
		}
		part2 += cards[i].instances
	}

	fmt.Println("Part 2:", part2)
}
