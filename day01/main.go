package main

import (
	"advent2023/utils"
	"fmt"
	"strings"
)

func calculate(lines []string, numbers map[string]string) int {
	sum := 0
	for _, line := range lines {
		firstPos := len(line) + 1
		first := "No"

		lastPos := -1
		last := "No"

		for text, value := range numbers {
			pos0 := strings.Index(line, text)
			if pos0 != -1 {
				if pos0 < firstPos {
					first = value
					firstPos = pos0
				}
			}

			posn := strings.LastIndex(line, text)
			if posn != -1 {
				if posn > lastPos {
					last = value
					lastPos = posn
				}
			}
		}

		num, _ := utils.StringToInt(fmt.Sprintf("%s%s", first, last))
		fmt.Println("Num:", line, num)
		sum += num

	}
	return sum
}

func main() {
	lines, _ := utils.ReadLines("input")

	numbers := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
		"8": "8",
		"9": "9",
	}

	sum := calculate(lines, numbers)

	numbers["one"] = "1"
	numbers["two"] = "2"
	numbers["three"] = "3"
	numbers["four"] = "4"
	numbers["five"] = "5"
	numbers["six"] = "6"
	numbers["seven"] = "7"
	numbers["eight"] = "8"
	numbers["nine"] = "9"

	sum2 := calculate(lines, numbers)

	fmt.Println("Part 1", sum)
	fmt.Println("Part 2", sum2)

}
