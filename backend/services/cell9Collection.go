package services

import "github.com/pietervdheijden/sudoku/util"

// cell9Collection can contain a row, a column or a square.
type cell9Collection [9]*Cell

func (cells cell9Collection) valid() bool {
	numbers := []int{}
	for _, cell := range cells {
		number := (*cell).Number
		if util.Contains(numbers, number) {
			return false
		}
		if number != 0 {
			numbers = append(numbers, number)
		}
		if cell.Number == 0 && len(cell.Options) == 0 {
			return false
		}
	}
	return true
}

func (cells cell9Collection) contains(id int) bool {
	for _, cell := range cells {
		if cell.Id == id {
			return true
		}
	}
	return false
}

func (cells cell9Collection) updateNumbers() bool {
	updated := false

	// Set number when there's only a single cell which options contains the number
	for number := 1; number <= 9; number++ {
		numberFound := false
		potentialCellCount := 0
		lastCellIndex := -1
		for i := 0; i < 9; i++ {
			if (*cells[i]).Number == number {
				numberFound = true
				break
			}
			if util.Contains((*cells[i]).Options, number) {
				potentialCellCount++
				lastCellIndex = i
			}
		}
		if !numberFound && potentialCellCount == 1 {
			(*cells[lastCellIndex]).Number = number
			(*cells[lastCellIndex]).Options = []int{}
			updated = true
		}
	}
	return updated
}