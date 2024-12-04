package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var (
	mulRegexp        = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	mulEnabledRegexp = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)
)

func dayThree() {
	file, err := os.Open("./inputs/input_day_3.txt")
	check(err)
	defer file.Close()

	getMulSum(file)
	file.Seek(0, io.SeekStart)
	getMulSumEnabled(file)
}

func getMulSum(file *os.File) {
	scanner := bufio.NewScanner(file)
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
	fmt.Printf("Mul Sum\t%d\n", sum(muls))
}

func getMulSumEnabled(file *os.File) {
	scanner := bufio.NewScanner(file)
	var muls []int
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()
		matches := mulEnabledRegexp.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "don't()" {
				enabled = false
			} else if match[0] == "do()" {
				enabled = true
			} else if enabled {
				x, err := strconv.Atoi(match[1])
				check(err)
				y, err := strconv.Atoi(match[2])
				check(err)
				muls = append(muls, x*y)
			}
		}
	}
	fmt.Printf("Mul Sum Switch:\t%d\n", sum(muls))
}

func sum(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}
