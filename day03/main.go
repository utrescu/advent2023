package main

import (
	"advent2023/utils"
	"errors"
	"fmt"
)

type Pos struct {
	x int
	y int
}

type SegmentNumber struct {
	from   Pos
	to     Pos
	number int
}

type SegmentSimbol struct {
	pos   Pos
	value string
}

type EngineParts struct {
	PartNumbers []SegmentNumber
	PartSimbols []SegmentSimbol
}

func ScanEngineParts(lines [][]string) EngineParts {
	e := EngineParts{
		PartNumbers: []SegmentNumber{},
		PartSimbols: []SegmentSimbol{},
	}

	for row, line := range lines {

		number := ""
		skip := 0
		for i, cell := range line {
			if utils.IsInt(cell) {
				if number == "" {
					start := i
					for start < len(line) && utils.IsInt(line[start]) {
						number = fmt.Sprintf("%s%s", number, line[start])
						start++
						skip++
					}
					num, _ := utils.StringToInt(number)
					e.PartNumbers = append(e.PartNumbers, SegmentNumber{from: Pos{x: i, y: row}, to: Pos{x: start - 1, y: row}, number: num})
				}
				skip--
				if skip == 0 {
					number = ""
				}

			} else if cell != "." {
				e.PartSimbols = append(e.PartSimbols, SegmentSimbol{pos: Pos{x: i, y: row}, value: cell})
			}
		}
	}
	return e
}

func (e EngineParts) IsValid(partnumber SegmentNumber) (int, error) {
	for _, partsimbol := range e.PartSimbols {

		if utils.Abs(partnumber.from.y-partsimbol.pos.y) <= 1 {
			// Candidate!
			for x := partnumber.from.x; x <= partnumber.to.x; x++ {
				if utils.Abs(x-partsimbol.pos.x) <= 1 {
					return partnumber.number, nil
				}

			}
		}
	}
	return -1, errors.New("not Valid")
}

func (e EngineParts) IsGear(partsimbol SegmentSimbol) (int, error) {
	numbers := []int{}
	for _, partnumber := range e.PartNumbers {
		if utils.Abs(partnumber.from.y-partsimbol.pos.y) <= 1 {
			for x := partnumber.from.x; x <= partnumber.to.x; x++ {
				if utils.Abs(x-partsimbol.pos.x) <= 1 {
					numbers = append(numbers, partnumber.number)
					break
				}
			}
		}
	}

	if len(numbers) != 2 {
		return -1, errors.New("not gear")
	}
	return numbers[0] * numbers[1], nil
}

func (e EngineParts) GetEngineNumbers() []int {

	result := []int{}

	for _, partnumber := range e.PartNumbers {
		number, err := e.IsValid(partnumber)
		if err == nil {
			result = append(result, number)
		}
	}
	return result
}

func (e EngineParts) GetGears() []int {
	result := []int{}

	for _, partsimbol := range e.PartSimbols {
		if partsimbol.value == "*" {
			number, err := e.IsGear(partsimbol)
			if err == nil {
				result = append(result, number)
			}
		}
	}

	return result
}

func main() {

	lines, _ := utils.ReadLinesCharacters("input")

	// mapa := Mapa{}
	// mapa.Construct(lines)

	// for _, linia := range mapa.Cells {
	// 	fmt.Println(linia)
	// }

	engine := ScanEngineParts(lines)

	part1 := 0
	resultats := engine.GetEngineNumbers()
	for _, resultat := range resultats {
		part1 += resultat
	}
	fmt.Println("Part 1:", part1)

	part2 := 0
	resultats2 := engine.GetGears()
	for _, resultat2 := range resultats2 {
		part2 += resultat2
	}
	fmt.Println("Part 2:", part2)

}
