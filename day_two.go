package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func dayTwo() {
	file, err := os.Open("./inputs/input_day_2.txt")
	check(err)
	defer file.Close()

	getSafeCount(file)
	file.Seek(0, io.SeekStart)
	getSafeCountWithDampener(file)
}

func getSafeCount(file *os.File) {
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := parseInts(line)

		isSafe := getIsSafe(levels)

		if isSafe {
			count += 1
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Safe Count: %d\n", count)
}

func getSafeCountWithDampener(file *os.File) {
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := parseInts(line)

		isSafe := getIsSafe(levels)

		if isSafe {
			count += 1
		} else {
			for i := 0; i < len(levels); i++ {
				subLevels := slices.Concat(levels[:i], levels[i+1:])

				isSafe = getIsSafe(subLevels)

				if isSafe {
					count += 1
					break
				}
			}
		}
	}

	fmt.Printf("Safe Count w/ Dampener: %d\n", count)
}

func getIsSafe(levels []int) bool {
	isSafe := true
	direction := "asc"
	for i, n := range levels {
		if i == 0 {
			if levels[1] < n {
				direction = "desc"
			}
		}
		if i == len(levels)-1 {
			break
		}
		if direction == "asc" && levels[i+1] < n {
			isSafe = false
			break
		}
		if direction == "desc" && levels[i+1] > n {
			isSafe = false
			break
		}
		diff := n - levels[i+1]
		if diff < 0 {
			diff = -diff
		}
		if !(1 <= diff && diff <= 3) {
			isSafe = false
			break
		}
	}
	return isSafe
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
