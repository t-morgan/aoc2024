package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileLines(file *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func parseInts(line string) []int {
	var nums []int
	strNums := strings.Fields(line)
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			continue // Skip non-integer values
		}
		nums = append(nums, num)
	}
	return nums
}
