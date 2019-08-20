package main

import "fmt"

func main() {
	testBoard := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	validSolution(testBoard)

}

func validSolution(board [][]int) bool {
	ret := false
	for i := 0; i < 9; i++ {
		fmt.Printf("  SubGroup %d = %v\n", i, getSubgroup(board, i))
		fmt.Printf("       Row %d = %v\n", i, getSubgroup(board, i))
		fmt.Printf("       Col %d = %v\n", i, getSubgroup(board, i))
	}
	return ret
}

func getSubgroup(board [][]int, sgIdx int) []int {
	ret := make([]int, 9)
	for i := 0; i < 3; i++ {
		row := i + (i * 3)
		for j := 0; j < 3; j++ {
			col := j * 3

			// Here we know te top left coordinates
			for x := 0; x < 3; x++ {
				row = row + x
				for y := 0; y < 3; y++ {
					col = col + 3
					ret[3*i+j] = board[row][col]
				}
			}

		}
	}

	return ret
}

func getRow(board [][]int, rowIdx int) []int {
	ret := make([]int, 9)
	return ret
}

func getColboard(board [][]int, colIdx int) []int {
	ret := make([]int, 9)
	return ret
}
