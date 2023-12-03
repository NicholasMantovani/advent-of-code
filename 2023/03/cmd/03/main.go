package main

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/NicholasMantovani/2023/03/internal/input_reader"
)

var validSymbols = []rune{'=', '-', '/', '+', '*', '%', '&', '@', '#', '$'}

var numberRegex = regexp.MustCompile("[0-9]+")

func main() {
	input := getInputs("input.txt")

	resultPartOne := executePartOne(input)
	fmt.Println("Part one result: ", resultPartOne)
}

func getInputs(inputFile string) []string {
	fileRead := input_reader.ReadInput("/workspaces/advent-of-code/2023/03/" + inputFile)

	lines := []string{}

	fileScanner := bufio.NewScanner(fileRead)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		lines = append(lines, text)
	}

	fileRead.Close()

	return lines
}

func executePartOne(input []string) int {
	sum := 0
	validNumbers := findValidNumbers(input)
	for _, num := range validNumbers {
		sum += num
	}
	return sum
}

func findValidNumbers(input []string) []int {
	validNumbers := []int{}
	
	for i, line := range input {
		numbersInLine := findAllNumbersInLine(line)

		var indexOfSymbolsInLineBefore = []int{}
		var indexOfSymbolsInLineAfter = []int{}
		var first, last = 0, 0

		if i != 0 {
			// i don't need to check the line before this
			indexOfSymbolsInLineBefore = findAllIndexOfValidSymbolsInLine(input[i-1])

		}
		if i != len(input)-1 {
			// i don't need to check the line after this
			indexOfSymbolsInLineAfter = findAllIndexOfValidSymbolsInLine(input[i+1])
		}

		for _, number := range numbersInLine {
			first, last = findFirstAndLastIndexOfString(number, line)
			numberInt, _ := strconv.Atoi(number)

			// Finds the value in the current line
			if first != 0 {
				if isRuneAValidSymbols(rune(line[first-1])) {
					validNumbers = append(validNumbers, numberInt)
					line = removeCheckedIndexesFromString(first, last, line)

					continue
				}
			}
			if last != len(line)-1 {
				if isRuneAValidSymbols(rune(line[last+1])) {
					validNumbers = append(validNumbers, numberInt)
					line = removeCheckedIndexesFromString(first, last, line)

					continue
				}
			}


			// Find value in the before and after line
			if isNumberIndexInArrayOfIndexes(first - 1, last + 1, indexOfSymbolsInLineBefore) {
				validNumbers = append(validNumbers, numberInt)
				line = removeCheckedIndexesFromString(first, last, line)

				continue
			}

			if isNumberIndexInArrayOfIndexes(first - 1, last + 1, indexOfSymbolsInLineAfter) {
				validNumbers = append(validNumbers, numberInt)
				line = removeCheckedIndexesFromString(first, last, line)

				continue
			}
			line = removeCheckedIndexesFromString(first, last, line)

		}
	}

	return validNumbers
}

func findAllNumbersInLine(line string) []string {
	return numberRegex.FindAllString(line, -1)
}

func findAllIndexOfValidSymbolsInLine(line string) []int {
	output := []int{}
	for index, run := range line {
		if isRuneAValidSymbols(run) {
			output = append(output, index)
		}
	}
	return output
}

func findFirstAndLastIndexOfString(num, line string) (firstIndex int, lastIndex int) {
	firstIndex = strings.Index(line, num)
	lastIndex = firstIndex + len(num) - 1
	return firstIndex, lastIndex
}

func isRuneAValidSymbols(val rune) bool {
	return slices.Contains(validSymbols, val)
}

func isNumberIndexInArrayOfIndexes(first, last int, symbolsIndexes []int) bool {

	if len(symbolsIndexes) == 0 {
		return false
	}

	fromFirstToLastIndexes := []int{first}

	count := first

	for count < last {
		count++
		fromFirstToLastIndexes = append(fromFirstToLastIndexes, count)
	}

	for _, index := range fromFirstToLastIndexes {
		if slices.Contains(symbolsIndexes, index) {
			return true
		}
	}
	return false
}

func removeCheckedIndexesFromString(first, last int, line string) string {
	count := first
	for count <= last {
		line = replaceAt(line, count, '.')
		count++
	}
	return line
}


func replaceAt(s string, i int, c rune) string {
    r := []rune(s)
    r[i] = c
    return string(r)
}