package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func dayFive() {
	file, err := os.Open("./inputs/input_day_5.txt")
	check(err)
	defer file.Close()

	lines := getFileLines(file)
	getCorrectlyOrderedSum(lines)
	getIncorrectlyOrderedSum(lines)
}

func getCorrectlyOrderedSum(lines []string) {
	var correct [][]int

	predecessors := make(map[string][]string)
	i := 0
	for ; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}
		pages := strings.Split(line, "|")
		predecessors[pages[1]] = append(predecessors[pages[1]], pages[0])
	}

	i++
	for ; i < len(lines); i++ {
		line := lines[i]
		pages := strings.Split(line, ",")

		isCorrect := true
		for j, page := range pages {
			preds := predecessors[page]
			for k := j; k < len(pages); k++ {
				for _, p := range preds {
					if p == pages[k] {
						isCorrect = false
						break
					}
				}
			}
		}
		if isCorrect {
			correct = append(correct, parseInts(strings.Join(pages, " ")))
		}
	}

	var sum int
	for i := 0; i < len(correct); i++ {
		mid := (len(correct[i]) - 1) / 2
		sum += correct[i][mid]
	}
	fmt.Printf("Correct Sum: %d\n", sum)
}

func getIncorrectlyOrderedSum(lines []string) {
	var incorrect [][]int

	predecessors := make(map[string][]string)
	i := 0
	for ; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}
		pages := strings.Split(line, "|")
		predecessors[pages[1]] = append(predecessors[pages[1]], pages[0])
	}

	i++
	for ; i < len(lines); i++ {
		line := lines[i]
		pages := strings.Split(line, ",")

		isCorrect := true
		for j, page := range pages {
			preds := predecessors[page]
			for k := j; k < len(pages); k++ {
				for _, p := range preds {
					if p == pages[k] {
						isCorrect = false
						break
					}
				}
			}
		}
		if !isCorrect {
			incorrect = append(incorrect, parseInts(strings.Join(pages, " ")))
		}
	}

	var correct [][]int
	for _, pagesList := range incorrect {
		corrected := make([]int, 0)
		for _, page := range pagesList {
			if len(corrected) == 0 {
				corrected = append(corrected, page)
				continue
			}

			preds := predecessors[strconv.Itoa(page)]
			if len(preds) == 0 {
				corrected = append([]int{page}, corrected...)
				continue
			}

			predCount := 0
			for i, item := range corrected {
				if slices.Contains(preds, strconv.Itoa(item)) {
					predCount++
				}
				pageStr := strconv.Itoa(page)
				if slices.Contains(predecessors[strconv.Itoa(item)], pageStr) {
					corrected = append(corrected[:i], append([]int{page}, corrected[i:]...)...)
					break
				}
				if predCount == len(preds) {
					corrected = append(corrected[:i+1], append([]int{page}, corrected[i+1:]...)...)
					break
				}
			}
			if !slices.Contains(corrected, page) {
				corrected = append(corrected, page)
			}
		}
		correct = append(correct, corrected)
	}

	var sum int
	for i := 0; i < len(correct); i++ {
		mid := (len(correct[i]) - 1) / 2
		sum += correct[i][mid]
	}
	fmt.Printf("Corrected Sum: %d\n", sum)
}
