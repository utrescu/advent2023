package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func StringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func Abs(number int) int {
	if number > 0 {
		return number
	}
	return number * -1
}

func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func ReadLines(path string) ([]string, error) {
	result := []string{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		result = append(result, line)
	}
	return result, scanner.Err()
}

func ReadLinesCharacters(path string) ([][]string, error) {
	result := [][]string{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		linechars := strings.Split(line, "")
		result = append(result, linechars)
	}
	return result, scanner.Err()
}
