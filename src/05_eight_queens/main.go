package main

import "fmt"

type position struct {
	row int
	col int
}

type state struct {
	colsOccupied          [8]bool
	upDiagonalsOccupied   [15]bool
	downDiagonalsOccupied [15]bool
	queenPositions        [8]position
	numQueensPlaced       int
}

func placeQueen(s *state, row int) bool {
	for col := 0; col < 8; col++ {
		if s.colsOccupied[col] {
			continue
		}

		// formula is col = row + constant
		// let constant = diagNum - 7, so diagNum is 0-15.
		// So calculation is col - row + 7 = constant.
		upDiagonalIdx := col - row + 7
		if s.upDiagonalsOccupied[upDiagonalIdx] {
			continue
		}

		// formula is col = -row + diagNum
		// So calculation is col + row = diagNum
		downDiagonalIdx := col + row
		if s.downDiagonalsOccupied[downDiagonalIdx] {
			continue
		}

		s.queenPositions[s.numQueensPlaced] = position{row, col}
		s.colsOccupied[col] = true
		s.upDiagonalsOccupied[upDiagonalIdx] = true
		s.downDiagonalsOccupied[downDiagonalIdx] = true
		s.numQueensPlaced++

		if s.numQueensPlaced == 8 {
			return true
		}

		success := placeQueen(s, row+1)
		if success {
			return true
		}

		s.colsOccupied[col] = false
		s.upDiagonalsOccupied[upDiagonalIdx] = false
		s.downDiagonalsOccupied[downDiagonalIdx] = false
		s.numQueensPlaced--
	}

	return false
}

func main() {
	var s state
	result := placeQueen(&s, 0)
	if result {
		fmt.Printf("%#v\n", s.queenPositions)
	} else {
		fmt.Println("We have failed.")
	}
}
