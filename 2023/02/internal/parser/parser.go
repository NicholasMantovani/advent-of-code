package parser

import (
	"strings"
	"fmt"
	"strconv"
	"unicode"
)

type GameSet struct {
	Blue int
	Red int
	Green int
}

func ParseLine(line string) map[int][]GameSet{
	output := map[int][]GameSet{}


	line = strings.ReplaceAll(line, "Game ", "")
	gameId := strings.Split(line, ":")
	id, _ := strconv.Atoi(gameId[0])
	gamesSet := strings.Split(gameId[1], ";")
	output[id] = []GameSet{}
	
	for _, set := range gamesSet {
		output[id] = append(output[id], parseGameSet(set))
	}	
	fmt.Println("Line", line)
	fmt.Println("output", output)
	return output
}


func parseGameSet(set string) GameSet{
	gameSet := GameSet{}

	splittedByComma := strings.Split(set, ",")

	for _, cube := range splittedByComma {
		if strings.Contains(cube, "blue") {
			gameSet.Blue = fromCubeNumColorToNum(cube)
		}
		if strings.Contains(cube, "red") {
			gameSet.Red = fromCubeNumColorToNum(cube)
		}
		if strings.Contains(cube, "green") {
			gameSet.Green = fromCubeNumColorToNum(cube)
		}
	}
	return gameSet;
}

func fromCubeNumColorToNum(cubeNumCol string) int {
	num := ""
	for _, digit := range cubeNumCol {
		if unicode.IsDigit(digit) {
			num += string(digit)
		}
	}

	out, _ := strconv.Atoi(num)
	return out
}