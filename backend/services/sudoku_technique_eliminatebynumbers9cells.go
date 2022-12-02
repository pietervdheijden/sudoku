package services

func (sudoku *Sudoku) applyEliminateByNumbers9Cells() bool {
	// TODO: use standard terminology for this technique.
	updated := false

	cellsList := []cell9Collection{}
	for i := 0; i < 9; i++ {
		row := sudoku.getRow(i)
		column := sudoku.getColumn(i)
		square := sudoku.getSquare(i)

		cellsList = append(cellsList, row, column, square)
	}

	// When number != 0, remove the option from the other cells
	for _, cells := range cellsList {
		for c1 := 0; c1 < len(cells); c1++ {
			for c2 := 0; c2 < len(cells); c2++ {
				numberC2 := (*cells[c2]).Number 
				if c1 == c2 || numberC2 == 0 {
					continue
				}
				removed := cells[c1].removeOption(numberC2)
				if removed {
					updated = true
				}
			}
		}
	}
	
	return updated
}