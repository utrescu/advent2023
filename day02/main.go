package main

import (
	"advent2023/utils"
	"fmt"
	"strings"
)

type Bag struct {
	Cubes map[string]int
}

func InitBag(colors []string) Bag {
	bag := Bag{}
	bag.Cubes = map[string]int{}
	for _, color := range colors {
		bag.Cubes[color] = 0
	}
	return bag
}

func (c *Bag) AddCubes(name string, num int) {
	c.Cubes[name] = max(c.Cubes[name], num)
}

func (c Bag) IsCorrect(other Bag) bool {
	for k, v := range c.Cubes {
		if v > other.Cubes[k] {
			return false
		}
	}
	return true
}

func (c Bag) Power() int {
	result := 1
	for _, v := range c.Cubes {
		result = result * v
	}
	return result
}

func processGame(line string, colors []string, maxCubes Bag) (bool, int) {

	gameRounds := strings.Split(line, "; ")

	bag := InitBag(colors)

	for _, gameRound := range gameRounds {

		grabs := strings.Split(gameRound, ", ")
		for _, grab := range grabs {
			cube := strings.Split(grab, " ")
			cubename := cube[1]
			cubenum, _ := utils.StringToInt(cube[0])
			bag.AddCubes(cubename, cubenum)
		}
	}

	fmt.Println("Round 1", bag.Cubes, "---", maxCubes.Cubes, ":", bag.IsCorrect(maxCubes))
	return bag.IsCorrect(maxCubes), bag.Power()
}

func main() {
	result1 := 0

	colors := []string{"red", "green", "blue"}
	maxCubes := InitBag(colors)
	maxCubes.Cubes["red"] = 12
	maxCubes.Cubes["green"] = 13
	maxCubes.Cubes["blue"] = 14

	lines, _ := utils.ReadLines("input")
	result2 := 0

	for game, line := range lines {
		ok, power := processGame(strings.Split(line, ": ")[1], colors, maxCubes)
		if ok {
			result1 += game + 1
		}
		result2 = result2 + power
	}

	fmt.Println("Part1:", result1)
	fmt.Println("Part1:", result2)
}
