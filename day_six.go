package main

import (
	"fmt"
	"os"
	"strings"
)

var simpleInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func daySix() {
	file, err := os.Open("./inputs/input_day_6.txt")
	check(err)
	defer file.Close()

	lines := getFileLines(file)
	// lines := strings.Split(simpleInput, "\n")
	var guardMap [][]string
	for _, line := range lines {
		guardMap = append(guardMap, strings.Split(line, ""))
	}

	getMapPositionCount(copy2DStringSlice(guardMap))
	getPotentialLoopCount(copy2DStringSlice(guardMap))
}

type Coordinates struct {
	X, Y int
}

type Guard struct {
	Location    Coordinates
	Orientation string
}

func getMapPositionCount(guardMap [][]string) {
	x, y, orientation := getInitialGuardPosition(guardMap)
	guard := Guard{Location: Coordinates{X: x, Y: y}, Orientation: orientation}

	next := getNextPosition(guard)
	for isXYInBounds(next.X, next.Y, guardMap) {
		if guardMap[next.X][next.Y] == "#" {
			guard.Orientation = getNextOrientation(guard)
			next = getNextPosition(guard)
		}
		guardMap[guard.Location.X][guard.Location.Y] = "X"
		guard.Location = next
		next = getNextPosition(guard)
	}
	guardMap[guard.Location.X][guard.Location.Y] = "X"

	countDistinctPositions := getGuardPositionCount(guardMap)
	fmt.Printf("Distinct Position Count: %d\n", countDistinctPositions)
}

func getInitialGuardPosition(guardMap [][]string) (x, y int, orientation string) {
	for i, row := range guardMap {
		for j, col := range row {
			if col == "^" || col == "v" || col == "<" || col == ">" {
				return i, j, col
			}
		}
	}
	return 0, 0, ""
}

func getNextPosition(guard Guard) Coordinates {
	switch guard.Orientation {
	case "^":
		return Coordinates{X: guard.Location.X - 1, Y: guard.Location.Y}
	case "v":
		return Coordinates{X: guard.Location.X + 1, Y: guard.Location.Y}
	case "<":
		return Coordinates{X: guard.Location.X, Y: guard.Location.Y - 1}
	case ">":
		return Coordinates{X: guard.Location.X, Y: guard.Location.Y + 1}
	default:
		return Coordinates{X: -1, Y: -1}
	}
}

func getNextOrientation(guard Guard) string {
	switch guard.Orientation {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
	default:
		return ""
	}
}

func getGuardPositionCount(guardMap [][]string) int {
	count := 0
	for _, row := range guardMap {
		for _, col := range row {
			if col == "X" {
				count++
			}
		}
	}
	return count
}

func isXYInBounds(x, y int, guardMap [][]string) bool {
	return x >= 0 && x < len(guardMap) && y >= 0 && y < len(guardMap[0])
}

func getPotentialLoopCount(guardMap [][]string) {
	loopCount := 0
	x, y, orientation := getInitialGuardPosition(guardMap)

	obstacles := getObstacles(guardMap)

	for _, obstacle := range obstacles {
		guardMapTmp := copy2DStringSlice(guardMap)
		guardMapTmp[obstacle.X][obstacle.Y] = "#"
		guard := Guard{Location: Coordinates{X: x, Y: y}, Orientation: orientation}
		visited := map[string]bool{fmt.Sprintf("%d:%d:%s", guard.Location.X, guard.Location.Y, guard.Orientation): true}
		next := getNextPosition(guard)
		for isXYInBounds(next.X, next.Y, guardMapTmp) {
			if visited[fmt.Sprintf("%d:%d:%s", next.X, next.Y, guard.Orientation)] {
				loopCount++
				break
			}
			visited[fmt.Sprintf("%d:%d:%s", next.X, next.Y, guard.Orientation)] = true

			if guardMapTmp[next.X][next.Y] == "#" {
				guard.Orientation = getNextOrientation(guard)
				next = getNextPosition(guard)
				continue
			}
			guard.Location = next
			next = getNextPosition(guard)
		}
	}

	fmt.Printf("Potential Loop Count: %d\n", loopCount)
}

func getObstacles(guardMap [][]string) []Coordinates {
	guardMap = copy2DStringSlice(guardMap)
	var obstacles []Coordinates
	x, y, orientation := getInitialGuardPosition(guardMap)
	guard := Guard{Location: Coordinates{X: x, Y: y}, Orientation: orientation}

	obstaclesFound := map[Coordinates]bool{{X: x, Y: y}: true}
	next := getNextPosition(guard)
	for isXYInBounds(next.X, next.Y, guardMap) {
		if guardMap[next.X][next.Y] == "#" {
			guard.Orientation = getNextOrientation(guard)
			next = getNextPosition(guard)
		}
		coord := Coordinates{next.X, next.Y}
		if !obstaclesFound[coord] {
			obstacles = append(obstacles, Coordinates{coord.X, coord.Y})
			obstaclesFound[coord] = true
		}
		guard.Location = next
		next = getNextPosition(guard)
	}
	return obstacles
}
