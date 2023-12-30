package main

import (
	"advent2023/utils"
	"fmt"
	"regexp"
	"strings"
)

type DataInput struct {
	value int
	step  int
}

type ChangeFrom struct {
	Finsa   string
	Changes []Change
}

type Change struct {
	destination int
	source      int
	step        int
}

func splitLineNumbers(line string, separator string) []int {
	values := strings.Split(line, separator)
	result, _ := utils.SliceToInt(values, true)
	return result
}

func ProcessInput(lines []string) ([]int, map[string]ChangeFrom) {

	seeds := []int{}
	engine := map[string]ChangeFrom{}
	from := ""
	finsa := ""

	for i, line := range lines {
		if i == 0 {
			stringvalues, _ := strings.CutPrefix(line, "seeds: ")
			seeds = splitLineNumbers(stringvalues, " ")
		}
		if i > 1 {
			if line == "" {
				continue
			}
			var re = regexp.MustCompile(`(?ms)^(.+)-to-(.+) map:$`)
			found := re.FindStringSubmatch(line)
			if len(found) > 0 {
				from = found[1]
				finsa = found[2]
				engine[from] = ChangeFrom{Finsa: finsa, Changes: []Change{}}
				// Capsalera
			} else {
				data := splitLineNumbers(line, " ")

				change := Change{
					destination: data[0],
					source:      data[1],
					step:        data[2],
				}

				oldValues := engine[from]
				newvalues := append(oldValues.Changes, change)
				oldValues.Changes = newvalues
				engine[from] = oldValues
			}
		}
	}

	return seeds, engine
}

func Step(originals []DataInput, from ChangeFrom) []DataInput {
	newSeeds := []DataInput{}
	for _, original := range originals {
		currentDataInput := []DataInput{original}
		for len(currentDataInput) != 0 {
			firstSeed := currentDataInput[0].value
			stepSeed := currentDataInput[0].step
			lastSeed := firstSeed + stepSeed - 1
			valueDest := currentDataInput[0]

			for _, changefrom := range from.Changes {

				firstChangefrom := changefrom.source
				lastChangefrom := changefrom.source + changefrom.step - 1

				if firstSeed <= lastChangefrom && lastSeed >= firstChangefrom {

					//if lastChangefrom >= firstSeed && firstChangefrom <= lastChangefrom {

					start := max(firstChangefrom, firstSeed)
					end := min(lastChangefrom, lastSeed)
					newPos := changefrom.destination + utils.Abs(original.value-changefrom.source)
					newLen := end - start
					// Si està tot a dins ok i liquida'l

					valueDest = DataInput{
						value: newPos,
						step:  newLen,
					}

					if firstSeed < firstChangefrom {
						// Parteix-lo de nou
						currentDataInput = append(currentDataInput, DataInput{
							value: firstSeed,
							step:  firstChangefrom - firstSeed + 1,
						})
					}
					if lastSeed > lastChangefrom {
						// Parteix-lo
						currentDataInput = append(currentDataInput, DataInput{
							value: lastChangefrom + 1,
							step:  lastSeed - lastChangefrom,
						})
					}
					break
				}
			}
			currentDataInput = currentDataInput[1:]
			newSeeds = append(newSeeds, valueDest)

		}
	}
	return newSeeds
}

func Calculate(dades []DataInput, engine map[string]ChangeFrom) int {
	current := "seed"

	for current != "location" {
		changes := engine[current]
		newseeds := Step(dades, changes)
		current = changes.Finsa
		dades = newseeds

	}

	min := dades[0].value
	for _, seed := range dades {
		if seed.value < min {
			min = seed.value
		}
	}
	return min
}

func main() {

	lines, _ := utils.ReadLines("input")
	seeds, engine := ProcessInput(lines)

	part1Seeds := []DataInput{}
	for i := 0; i < len(seeds); i += 1 {
		part1Seeds = append(part1Seeds, DataInput{value: seeds[i], step: 1})
	}

	part1 := Calculate(part1Seeds, engine)
	fmt.Println("Part 1", part1) // 836040384

	part2Seeds := []DataInput{}
	for i := 0; i < len(seeds); i += 2 {
		part2Seeds = append(part2Seeds, DataInput{value: seeds[i], step: seeds[i+1]})
	}

	part2 := Calculate(part2Seeds, engine)
	fmt.Println("Part 2", part2) // 10834440

}
