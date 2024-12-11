package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var simpleInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func daySeven() {
	file, err := os.Open("./inputs/input_day_7.txt")
	check(err)
	defer file.Close()

	lines := getFileLines(file)
	// lines := strings.Split(simpleInput, "\n")
	result := 0
	for _, line := range lines {
		answerInputs := strings.Split(line, ": ")
		answer, err := strconv.Atoi(answerInputs[0])
		check(err)
		inputStrings := strings.Split(answerInputs[1], " ")
		var inputs []int
		for _, input := range inputStrings {
			inputInt, err := strconv.Atoi(input)
			check(err)
			inputs = append(inputs, inputInt)
		}

		if slices.Contains(getCalibrations(inputs), answer) {
			result += answer
		}
	}

	fmt.Printf("Calibration Results: %d\n", result)
}

func getCalibrations(inputs []int) []int {
	if len(inputs) == 1 {
		return inputs
	}
	addition := getCalibrations(append([]int{inputs[0] + inputs[1]}, inputs[2:]...))
	multiplication := getCalibrations(append([]int{inputs[0] * inputs[1]}, inputs[2:]...))
	concatted, err := strconv.Atoi(strconv.Itoa(inputs[0]) + strconv.Itoa(inputs[1]))
	check(err)
	concat := getCalibrations(append([]int{concatted}, inputs[2:]...))
	return append(addition, append(multiplication, concat...)...)
}
