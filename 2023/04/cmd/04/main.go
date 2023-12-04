package main

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/NicholasMantovani/2023/04/internal/input_reader"
)

var numberRegex = regexp.MustCompile("[0-9]+")

type Card struct {
	cardId int
	winningNumbers []int
	actualNumbers []int
}

func main() {

	input := getInputs("input")
	fmt.Println("Input", input)
	cards := getCards(input)

	resultPartOne := executePartOne(cards)
	fmt.Println("\nPart one result: ", resultPartOne)


	resultPartTwo := executePartTwo(cards)
	fmt.Println("\nPart two result: ", resultPartTwo)
}


func getInputs(inputFile string) []string {
	fileRead := input_reader.ReadInput("/workspaces/advent-of-code/2023/04/" + inputFile)

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


func getCards(input []string) []Card {
	output := []Card{}

	for _, line := range input{
		card := Card{}
		cardIdGames := strings.Split(line, ":")
		res, _ := strconv.Atoi(findAllNumbersInString(cardIdGames[0])[0])
		card.cardId = res
		winningNumsActualNumbs := strings.Split(cardIdGames[1], "|")
		card.winningNumbers = convertStringsIntoNums(findAllNumbersInString(winningNumsActualNumbs[0]))
		card.actualNumbers = convertStringsIntoNums(findAllNumbersInString(winningNumsActualNumbs[1]))
		output = append(output, card)
	}
	return output
}

func findAllNumbersInString(line string) []string {
	return numberRegex.FindAllString(line, -1)
}

func convertStringsIntoNums( vals []string) []int {
	output := []int{}
	for _, val := range vals {
		res, _ := strconv.Atoi(val)
		output = append(output, res)
	}
	
	return output;
}


func executePartOne(cards []Card) int {
	winningPoints := 0
	for _, card := range cards {
		
		res := getWinningPointsForCard(card)
		winningPoints += res
		fmt.Printf("\nCard %#v | res %v | winningPoints %v", card, res, winningPoints)
	}
	return winningPoints
}

func getWinningPointsForCard(card Card) int {
	points := 0
	for _, actualNum := range card.actualNumbers {
		if slices.Contains(card.winningNumbers, actualNum) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func executePartTwo(cards []Card) int {
	sumTotalCards := 0

	totalCards := map[int]int{}
	
	for _, card := range cards {
		totalCards[card.cardId]++
	}

	for i, card := range cards {
		res := getWinningPointsForCardPartTwo(card)
		cardAmount :=  totalCards[card.cardId];
		
		for am := 0; am < cardAmount; am++ {
			for j := i + 1; j <= i + res && j < len(cards); j++ {
				totalCards[cards[j].cardId]++
			}
		}

		fmt.Printf("\nCard %#v | res %v | totalCards %v", card.cardId, res, totalCards)
	}

	for _, value := range totalCards {
		sumTotalCards += value
	}
	return sumTotalCards
}


func getWinningPointsForCardPartTwo(card Card) int {
	points := 0
	for _, actualNum := range card.actualNumbers {
		if slices.Contains(card.winningNumbers, actualNum) {
			points++
		}
	}
	return points
}