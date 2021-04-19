package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type puzzle [][]int8

func printPuzzle(puzzle puzzle) {
	for _, row := range puzzle {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Print("\n")
	}
}

func parsePuzzles() []puzzle {
	re := regexp.MustCompile("\\s+")
	reader := bufio.NewReader(os.Stdin)

	numPuzzlesStr, _ := reader.ReadString('\n')
	numPuzzles, _ := strconv.ParseInt(strings.TrimRight(numPuzzlesStr, "\n"), 0, 10)

	puzzles := make([]puzzle, numPuzzles)

	for i := range puzzles {
		puzzle := make([][]int8, 9)

		for j := 0; j < 9; j++ {
			puzzleRow := make([]int8, 9)

			puzzleRowStr, _ := reader.ReadString('\n')
			puzzleRowStr = strings.TrimRight(puzzleRowStr, "\n")
			puzzleRowStr = re.ReplaceAllString(puzzleRowStr, "")

			for idx, num := range puzzleRowStr {
				parsedInt, _ := strconv.ParseInt(string(num), 0, 10)
				puzzleRow[idx] = int8(parsedInt)
			}

			puzzle[j] = puzzleRow
		}

		puzzles[i] = puzzle
	}

	return puzzles
}

func findEmptyPosition(puzzle puzzle) [2]int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if puzzle[i][j] == 0 {
				return [2]int{i, j}
			}
		}
	}

	return [2]int{-1, -1}
}

func canPlaceNumber(number int8, row int, col int, puzzle puzzle) bool {
	for i := 0; i < 9; i++ {
		if puzzle[row][i] == number {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if puzzle[i][col] == number {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if puzzle[i+(row-row%3)][j+(col-col%3)] == number {
				return false
			}
		}
	}

	return true
}

func solvePuzzle(puzzle puzzle) bool {
	position := findEmptyPosition(puzzle)

	if position[0] == -1 {
		printPuzzle(puzzle)
		return true
	}

	for i := 1; i < 10; i++ {
		number := int8(i)
		if canPlaceNumber(number, position[0], position[1], puzzle) {
			puzzle[position[0]][position[1]] = number

			if solvePuzzle(puzzle) {
				return true
			}

			puzzle[position[0]][position[1]] = 0
		}
	}

	return false
}

func main() {
	sudokuPuzzles := parsePuzzles()

	for _, puzzle := range sudokuPuzzles {
		solvePuzzle(puzzle)
	}
}
