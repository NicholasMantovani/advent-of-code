package main

import (
	"bufio"
	"fmt"
	"sort"
	"github.com/NicholasMantovani/2023/02/internal/input_reader"
	"github.com/NicholasMantovani/2023/02/internal/parser"
)

func main() {

	parsedInput := getInputs("input")

	resultPartOne := getPossibleGamesSum(parsedInput)
	fmt.Println("\n\nPart one result:", resultPartOne)
	
	resultPartTwo := getGamesPower(parsedInput)
	fmt.Println("\n\nPart two result:", resultPartTwo)
	
}

const maxRed int = 12
const maxGreen int = 13
const maxBlue int = 14


func getInputs(inputFile string) []map[int][]parser.GameSet {
	fileRead := input_reader.ReadInput("/workspaces/advent-of-code/2023/02/" + inputFile)

	parsedGames := []map[int][]parser.GameSet{}

	fileScanner := bufio.NewScanner(fileRead)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		parsedGames = append(parsedGames, parser.ParseLine(text))
	}

	defer fileRead.Close()

	return parsedGames;
}


func getPossibleGamesSum(games []map[int][]parser.GameSet) int {

	sumPossibleGamesId := 0
	for _, game := range games {
		for id, gamesSet := range game {
			if areGamesSetPossible(gamesSet) {
				sumPossibleGamesId += id
			}
			
		}
	}
	return sumPossibleGamesId;
}


func areGamesSetPossible(gamesSet []parser.GameSet) bool {
	possible := true

	for i := 0; i < len(gamesSet) && possible; i++ {
		set := gamesSet[i]
		if set.Blue > maxBlue || set.Green > maxGreen || set.Red > maxRed {
			possible = false
		}
	}

	return possible
}


func getGamesPower(games []map[int][]parser.GameSet) int {
	gamePower := 0

	for _, game := range games {
		for _, gamesSet := range game {
			gamePower += getGamePower(gamesSet)
			
		}
	}
	return gamePower;

}

func getGamePower(gamesSet []parser.GameSet) int {

	sort.SliceStable(gamesSet, func(i, j int) bool {
		return gamesSet[i].Blue > gamesSet[j].Blue

	})

	maxBlueVal := gamesSet[0].Blue

	sort.SliceStable(gamesSet, func(i, j int) bool {
		return gamesSet[i].Red > gamesSet[j].Red

	})

	maxRedVal := gamesSet[0].Red

	sort.SliceStable(gamesSet, func(i, j int) bool {
		return gamesSet[i].Green > gamesSet[j].Green

	})

	maxGreenVal := gamesSet[0].Green
	
	return multiplyValues(maxBlueVal, maxRedVal, maxGreenVal)
}


func multiplyValues(values ...int) int {
	var result int 
	for _, val := range values {
		if(val != 0) {
			if(result == 0) {
				result = val
			} else {
				result *= val
			}
			
		}
	}
	return result
}



