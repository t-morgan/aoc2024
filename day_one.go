package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func dayOne() {
	file, err := os.Open("./inputs/input_day_1.txt")
	check(err)
	defer file.Close()

	left, right := getLrLists(file)
	getListDiff(left, right)
	getListSimilarity(left, right)
}

func getLrLists(file *os.File) (left, right []int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		n_left, err := strconv.Atoi(fields[0])
		check(err)
		left = append(left, n_left)
		n_right, err := strconv.Atoi(fields[1])
		check(err)
		right = append(right, n_right)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return left, right
}

func getListDiff(left, right []int) {
	slices.Sort(left)
	slices.Sort(right)

	diff := 0
	for i := 0; i < len(left); i++ {
		i_diff := left[i] - right[i]
		if i_diff < 0 {
			i_diff = -i_diff
		}
		diff += i_diff
	}

	fmt.Printf("Diff: %d\n", diff)
}

func getListSimilarity(left, right []int) {
	similarity_score := 0
	right_map := make(map[int]int)

	for i := 0; i < len(right); i++ {
		right_map[right[i]]++
	}

	for i := 0; i < len(left); i++ {
		similarity_score += left[i] * right_map[left[i]]
	}

	fmt.Printf("Similarity: %d\n", similarity_score)
}
