package main

import (
	"fmt"
	"math/rand"
)

const (
	rows  = 20
	cols  = 25
	mines = 99
	mine  = '*'
	empty = '.'
)

func GenerateRandomMinefield() [][]rune {
	board := make([][]rune, rows)
	for i := range board {
		board[i] = make([]rune, cols)
		for j := range board[i] {
			board[i][j] = empty
		}
	}

	mineCount := 0
	for mineCount < mines {
		r := rand.Intn(rows)
		c := rand.Intn(cols)
		if board[r][c] == empty {
			board[r][c] = mine
			mineCount++
		}
	}
	return board
}

func MarkBoard(board [][]rune) [][]rune {
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == empty {
				adjMineCount := 0
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < rows && nj >= 0 && nj < cols && board[ni][nj] == mine {
						adjMineCount++
					}
				}
				if adjMineCount > 0 {
					board[i][j] = rune('0' + adjMineCount)
				}
			}
		}
	}
	return board
}

func PrintBoard(board [][]rune) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(string(board[i][j]))
		}
		fmt.Println()
	}
}

func main() {
	minefield := GenerateRandomMinefield()

	fmt.Println("Initial Minefield:")
	PrintBoard(minefield)
	fmt.Println()

	markedBoard := MarkBoard(minefield)

	fmt.Println("Marked Minefield:")
	PrintBoard(markedBoard)
}
