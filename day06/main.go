package main

import (
	"advent2023/utils"
	"fmt"
	"strings"
)

type Race struct {
	Milliseconds int
	Distance     int
}

func Step2(race Race) int {
	first := 0
	last := race.Milliseconds / 2
	newdistance := 0

	for last-first > 1 {

		candidate := first + (last - first) / 2
		newdistance = candidate * (race.Milliseconds - candidate)
		if newdistance > race.Distance {
			last = candidate
		} else {
			first = candidate
		}

	}

	correction := 0
	if race.Milliseconds % 2 == 0  {
		correction = 1
	}
	return ((race.Milliseconds / 2) - first) * 2 - correction
}

func Part1(races []Race) int {
	result := 1

	for _, race := range races {
		result *= Step2(race)
	}

	return result
}

func ProcessInput(lines []string) ([]Race, Race) {
	races := []Race{}

	// Elimina prepend
	timesString, _ := strings.CutPrefix(lines[0], "Time:    ")
	distaString, _ := strings.CutPrefix(lines[1], "Distance:")

	times := utils.SplitLineNumbers(timesString, ' ')
	distances := utils.SplitLineNumbers(distaString, ' ')

	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			Milliseconds: times[i],
			Distance:     distances[i],
		})
	}

	timeString2 := strings.ReplaceAll(timesString, " ", "")
	distaString2 := strings.ReplaceAll(distaString, " ", "")

	time, _ := utils.StringToInt(timeString2)
	distance, _ := utils.StringToInt(distaString2)

	return races, Race{
		Milliseconds: time,
		Distance:     distance,
	}
}

func main() {
	lines, _ := utils.ReadLines("input")
	races, race2 := ProcessInput(lines)

	fmt.Println(races)

	result := Part1(races)
	fmt.Println("Part 1:", result)

	result2 := Step2(race2)
	fmt.Println("Part 2:", result2)

}
