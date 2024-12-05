package main

import (
	"fmt"
	"os"
)

func dayFour() {
	file, err := os.Open("./inputs/input_day_4.txt")
	check(err)
	defer file.Close()

	lines := getFileLines(file)
	wordsearch := makeWordsearch(lines)
	getXmasCount(wordsearch)
	getXMASCount(wordsearch)
}

func makeWordsearch(lines []string) [][]rune {
	var wordsearch [][]rune
	for _, line := range lines {
		chars := []rune(line)
		wordsearch = append(wordsearch, chars)
	}
	return wordsearch
}

func getXmasCount(wordsearch [][]rune) {
	count := 0
	// rows
	for i := 0; i < len(wordsearch); i++ {
		for j := 0; j < len(wordsearch[i])-3; j++ {
			word := string(wordsearch[i][j : j+4])
			if wordIsXMAS(word) {
				count++
			}
		}
	}

	// columns
	for i := 0; i < len(wordsearch[0]); i++ {
		for j := 0; j < len(wordsearch)-3; j++ {
			word := string([]rune{wordsearch[j][i], wordsearch[j+1][i], wordsearch[j+2][i], wordsearch[j+3][i]})
			if wordIsXMAS(word) {
				count++
			}
		}
	}

	// diagonal starting at top left
	for i := 0; i < 2*len(wordsearch)-2; i++ {
		var diag []rune
		for j := 0; j <= i; j++ {
			k := i - j
			if j < len(wordsearch) && k < len(wordsearch) {
				diag = append(diag, wordsearch[k][j])
			}
		}
		diagString := string(diag)
		for i := 0; i < len(diagString)-3; i++ {
			word := diagString[i : i+4]
			if wordIsXMAS(word) {
				count++
			}
		}
	}

	// diagonal starting at bottom left
	for i := 0; i < 2*len(wordsearch)-2; i++ {
		var diag []rune
		for j := 0; j <= i; j++ {
			k := len(wordsearch) - (i - j)
			if k >= 0 && j < len(wordsearch) && k < len(wordsearch) {
				diag = append(diag, wordsearch[k][j])
			}
		}
		diagString := string(diag)
		for i := 0; i < len(diagString)-3; i++ {
			word := diagString[i : i+4]
			if wordIsXMAS(word) {
				count++
			}
		}
	}

	fmt.Printf("XMAS Count: %d\n", count)
}

func wordIsXMAS(word string) bool {
	return word == "XMAS" || word == "SAMX"
}

func getXMASCount(wordsearch [][]rune) {
	count := 0

	squares := getSquares(wordsearch)
	for _, square := range squares {
		if string(square[1][1]) != "A" {
			continue
		}
		mCount := 0
		sCount := 0
		if string(square[0][0]) == "M" {
			mCount++
		} else if string(square[0][0]) == "S" {
			sCount++
		}
		if string(square[0][2]) == "M" {
			mCount++
		} else if string(square[0][2]) == "S" {
			sCount++
		}
		if string(square[2][0]) == "M" {
			mCount++
		} else if string(square[2][0]) == "S" {
			sCount++
		}
		if string(square[2][2]) == "M" {
			mCount++
		} else if string(square[2][2]) == "S" {
			sCount++
		}

		if mCount != 2 && sCount != 2 {
			continue
		}

		if (string(square[0][0]) == "M" && string(square[2][2]) != "S") ||
			(string(square[0][0]) == "S" && string(square[2][2]) != "M") {
			continue
		}
		if (string(square[0][2]) == "M" && string(square[2][0]) != "S") ||
			(string(square[0][2]) == "S" && string(square[2][0]) != "M") {
			continue
		}
		if (string(square[2][0]) == "M" && string(square[0][2]) != "S") ||
			(string(square[2][0]) == "S" && string(square[0][2]) != "M") {
			continue
		}
		if (string(square[2][2]) == "M" && string(square[0][0]) != "S") ||
			(string(square[2][2]) == "S" && string(square[0][0]) != "M") {
			continue
		}
		count++
	}

	fmt.Printf("X-MAS Count: %d\n", count)
}

func getSquares(wordsearch [][]rune) [][][]rune {
	var squares [][][]rune

	for i := 0; i < len(wordsearch)-2; i++ {
		for j := 0; j < len(wordsearch[0])-2; j++ {
			var square [][]rune
			for _, row := range wordsearch[i : i+3] {
				square = append(square, row[j:j+3])
			}
			squares = append(squares, square)
		}
	}

	return squares
}
