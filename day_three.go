package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func dayThree() {
	file, err := os.Open("./inputs/input_day_3.txt")
	check(err)
	defer file.Close()

	getMulSum(file)
}

func getMulSum(file *os.File) {
	scanner := bufio.NewScanner(file)
	mulRegexp := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	var muls []int
	for scanner.Scan() {
		line := scanner.Text()
		matches := mulRegexp.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			x, err := strconv.Atoi(match[1])
			check(err)
			y, err := strconv.Atoi(match[2])
			check(err)
			muls = append(muls, x*y)
		}
	}
	fmt.Printf("Mul Sum: %d\n", sum(muls))
}

func sum(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}