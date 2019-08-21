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

	fmt.Println(validSolution(testBoard))

	//aTester := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	/*
		tia1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		tia2 := []int{1, 1, 3, 4, 5, 6, 7, 8, 9}
		tia3 := []int{7, 2, 3, 4, 5, 6, 7, 8, 9}

		fmt.Println(testIntArray(tia1))
		fmt.Println(testIntArray(tia2))
		fmt.Println(testIntArray(tia3))
	*/

}

func validSolution(board [][]int) bool {
	for i := 0; i < 9; i++ {
		//_ = getSubgroup(board, i)
		//fmt.Printf("  SubGroup %d = %v\n", i, getSubgroup(board, i))
		if testIntArray(getSubgroup(board, i)) == false {
			return false
		}
		fmt.Printf("       Row %d = %v\n", i, getBoardRow(board, i))
		if testIntArray(getSubgroup(board, i)) == false {
			return false
		}
		fmt.Printf("       Col %d = %v\n", i, getBoardCol(board, i))
		if testIntArray(getSubgroup(board, i)) == false {
			return false
		}

	}
	return true
}

// testIntArray will make sure there are the numbers 1-9 exactly once in an array
func testIntArray(ia []int) bool {
	aTester := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < len(ia); i++ {
		aTester[ia[i]-1]++
	}

	for i := 0; i < len(aTester); i++ {

		if aTester[i] != 1 {
			return false
		}
	}
	return true
}

func getSubgroup(board [][]int, sgIdx int) []int {
	ret := make([]int, 9)

	top := (sgIdx / 3) * 3
	left := (sgIdx % 3) * 3

	//fmt.Printf("sgIdx:%d  top:%d  left:%d\n", sgIdx, top, left)

	for x := 0; x < 3; x++ {
		r := top + x
		for y := 0; y < 3; y++ {
			c := y + left
			//fmt.Printf("r:%d  c:%d\n", r, c)
			ret[3*x+y] = board[r][c]
		}
	}
	fmt.Println()

	return ret
}

func getBoardRow(board [][]int, rowIdx int) []int {
	ret := make([]int, 9)
	for x := 0; x < 9; x++ {
		ret[x] = board[rowIdx][x]
	}

	return ret
}

func getBoardCol(board [][]int, colIdx int) []int {
	ret := make([]int, 9)
	for x := 0; x < 9; x++ {
		ret[x] = board[x][colIdx]
	}

	return ret
}
