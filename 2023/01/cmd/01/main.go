package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"github.com/NicholasMantovani/2023/01/internal/input_reader"
)

func main() {

	lines := getInputs("input")
	resultPartOne := executePartOne(lines)

	fmt.Println("\nThe result of part one is:", resultPartOne)

	resultPartTwo := executePartTwo(lines)

	fmt.Println("\nThe result of part two is:", resultPartTwo)

	resultPartTwoEasy := executePartTwoEasy(lines)
	fmt.Println("\nThe result of part two easy is:", resultPartTwoEasy)

	resultPartTwoEasier := executePartTwoEasier(lines)
	fmt.Println("\nThe result of part two easier is:", resultPartTwoEasier)


}

func getInputs(inputFile string) []string {
	fileRead := input_reader.ReadInput("/workspaces/advent-of-code/2023/01/" + inputFile)

	lines := []string{}

	fileScanner := bufio.NewScanner(fileRead)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		lines = append(lines, text)
	}

	fileRead.Close()

	fmt.Println("Input", lines)

	fmt.Println("\nInput length", len(lines))
	return lines
}

func executePartOne(lines []string) int {
	calibrationValues := []int{}
	sum := 0

	for _, line := range lines {
		firstNumber := ""
		lastNumber := ""

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				firstNumber = string(line[i])
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lastNumber = string(line[i])
				break
			}
		}

		res, _ := strconv.Atoi(firstNumber + lastNumber)

		calibrationValues = append(calibrationValues, res)

		// Just to debug
		// fmt.Printf("\nline %v | firstNum %v | lastNum %v | concatVal %v", line, firstNumber, lastNumber, res)

	}

	for _, calibrationVal := range calibrationValues {
		sum += calibrationVal
	}
	return sum
}

func executePartTwo(lines []string) int {
	calibrationValues := []int{}
	sum := 0

	for _, line := range lines {
		firstNumIndex := len(line)
		lastNumIndex := 0
		firstNumber := ""
		lastNumber := ""

		for key, value := range getNumMap() {

			resFirst := strings.Index(line, key)
			if resFirst != -1 {
				if resFirst <= firstNumIndex {
					firstNumIndex = resFirst
					firstNumber = fmt.Sprint(value)
				}
			}
			resLast := strings.LastIndex(line, key)
			if resLast >= lastNumIndex {
				lastNumIndex = resLast + len(key) - 1
				lastNumber = fmt.Sprint(value)
			}
		}

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				if i <= firstNumIndex {
					firstNumber = string(line[i])
					break
				}
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				
				if i >= lastNumIndex {
					lastNumber = string(line[i])
					break
				}
				
			}
		}

		res, _ := strconv.Atoi(firstNumber + lastNumber)

		calibrationValues = append(calibrationValues, res)

		// Just to debug
		//fmt.Printf("\nline %v | firstNum %v | lastNum %v | concatVal %v", line, firstNumber, lastNumber, res)
		

	}
	for _, calibrationVal := range calibrationValues {
		sum += calibrationVal
	}
	return sum
}

type NumIndexKey struct {
	index int
	key string
}

// This is to exercise myself 
func executePartTwoEasy(lines []string) int {
	calibrationValues := []int{}
	sum := 0
	var numMap map[string]string = getNumMapString()

	for _, line := range lines {
		firstIndexNum := NumIndexKey{index: len(line), key: ""}
		lastIndexNum := NumIndexKey{index: 0, key: ""}
		firstIndexString := NumIndexKey{index: len(line), key: ""}
		lastIndexString := NumIndexKey{index: 0, key: ""}

		for key, value := range numMap {

			// NUM
			resFirstNum := strings.Index(line, fmt.Sprint(value));

			setFirstNumIndexKey(resFirstNum, key, &firstIndexNum)

			resLastNum := strings.LastIndex(line, fmt.Sprint(value));

			setLastNumIndexKey(resLastNum, key, &lastIndexNum, nil)

			// STRING
			resFirstString := strings.Index(line, key);

			setFirstNumIndexKey(resFirstString, key, &firstIndexString)

			resLastString := strings.LastIndex(line, key);

			setLastNumIndexKey(resLastString, key, &lastIndexString, func (x int, k string) int {
				return x + len(k) -1
			})


		}
		
		values := []NumIndexKey{firstIndexNum, lastIndexNum, firstIndexString, lastIndexString}

		values = filterNumIndexKey(values, func(x NumIndexKey) bool {
			return x.key != ""
		})

		sort.Slice(values, func(i, j int) bool {

			return values[i].key != "" && values[j].key != "" && values[i].index < values[j].index
		})
		minVal := values[0]
		maxVal := values[len(values) - 1]

		// Just to debug
		// fmt.Printf("\nMinVal %v | MaxVal %v | Values %v, | Line %v", minVal, maxVal, values, line)

		res, _ := strconv.Atoi(numMap[minVal.key] + numMap[maxVal.key])

		calibrationValues = append(calibrationValues,  res)

	}

	for _, calibrationVal := range calibrationValues {
		sum += calibrationVal
	}
	return sum
}

// This is to exercise myself
func executePartTwoEasier(lines []string) int {
	calibrationValues := []int{}
	sum := 0
	var numMap map[string]string = getNumMapString()

	for _, line := range lines {
		firstIndex := NumIndexKey{index: len(line), key: ""}
		lastIndex :=  NumIndexKey{index: 0, key: ""}


		for key, value := range numMap {

			// NUM
			resFirstNum := strings.Index(line, fmt.Sprint(value));

			setFirstNumIndexKey(resFirstNum, key, &firstIndex)

			resLastNum := strings.LastIndex(line, fmt.Sprint(value));

			setLastNumIndexKey(resLastNum, key, &lastIndex, nil)

			// STRING
			resFirstString := strings.Index(line, key);

			setFirstNumIndexKey(resFirstString, key, &firstIndex)

			resLastString := strings.LastIndex(line, key);

			setLastNumIndexKey(resLastString, key, &lastIndex, func (x int, k string) int {
				return x + len(k) -1
			})


		}
		
		res, _ := strconv.Atoi(numMap[firstIndex.key] + numMap[lastIndex.key])

		calibrationValues = append(calibrationValues,  res)

	}

	for _, calibrationVal := range calibrationValues {
		sum += calibrationVal
	}
	return sum
}

func getNumMap() map[string]int {
	return map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9}
}

func getNumMapString() map[string]string {
	return map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9"}
}

func filterNumIndexKey(numIndexKeys []NumIndexKey, condition func(NumIndexKey) bool) []NumIndexKey {
    var filteredPersons []NumIndexKey

    for _, p := range numIndexKeys {
        if condition(p) {
            filteredPersons = append(filteredPersons, p)
        }
    }

    return filteredPersons
}

func setFirstNumIndexKey(res int, key string, num *NumIndexKey) { 
	if res != -1 {
		if res <= num.index {
			num.index = res
			num.key = key
		}
	}
}

func setLastNumIndexKey(res int, key string, num *NumIndexKey, calcolateResut func(x int, k string) int) { 
	if res >= num.index {
		if(calcolateResut != nil) {
			res = calcolateResut(res, key)
		}
		num.index = res
		num.key = key
	}
}

