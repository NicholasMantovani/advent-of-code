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
		fmt.Println("Line:", line)
		numbersInLine := findAllNumbersInLine(line)
		var indexOfSymbolsInLineBefore = []int{}
		var indexOfSymbolsInLineAfter = []int{}

		if i != 0 {
			// i don't need to check the line before this
			indexOfSymbolsInLineBefore = findAllIndexOfValidSymbolsInLine(input[i-1])

		}
		if i != len(input)-1 {
			// i don't need to check the line after this
			indexOfSymbolsInLineAfter = findAllIndexOfValidSymbolsInLine(input[i+1])
		}

		for _, number := range numbersInLine {
			first, last := findFirstAndLastIndexOfString(number, line)
			numberInt, _ := strconv.Atoi(number)

			// Finds the value in the current line
			if first != 0 {
				if isRuneAValidSymbols(rune(line[first-1])) {
					validNumbers = append(validNumbers, numberInt)
					fmt.Println("First case", numberInt)
					continue
				}
			}
			if last != len(line)-1 {
				if isRuneAValidSymbols(rune(line[last+1])) {
					validNumbers = append(validNumbers, numberInt)
					fmt.Println("Second case", numberInt)
					continue
				}
			}

			// add tresholds to indexes (1 per side)
			first -= 1
			last += 1

			// Find value in the before and after line
			if isNumberIndexInArrayOfIndexes(first, last, indexOfSymbolsInLineBefore) {
				validNumbers = append(validNumbers, numberInt)
				fmt.Println("Third case", numberInt)
				continue
			}

			if isNumberIndexInArrayOfIndexes(first, last, indexOfSymbolsInLineAfter) {
				validNumbers = append(validNumbers, numberInt)
				fmt.Println("Fourth case", numberInt)
				continue
			}
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
