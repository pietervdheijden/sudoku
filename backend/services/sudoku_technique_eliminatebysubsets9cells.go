package services

import "github.com/pietervdheijden/sudoku/util"

func (sudoku *Sudoku) applyEliminateBySubsets9Cells() bool {
	// TODO: use standard terminology for this technique.
	updated := false

	cellsList := []cell9Collection{}
	for i := 0; i < 9; i++ {
		row := sudoku.getRow(i)
		column := sudoku.getColumn(i)
		square := sudoku.getSquare(i)

		cellsList = append(cellsList, row, column, square)
	}

	// Eliminate options by subsets
	// When cell C1 has X options, and there exist X-1 other cells (c2) with a subset of X, then the X options can be removed from the remaining cells (c3).
	// Example 1: when cell 1 and 2 both have options (1,2), then options (1,2) can be removed from cells 3,...,9.
	// Example 2: when cell 1 has options (1,2,3), cell 2 has options (1,2) and cell 3 has options (2,3), then options (1,2,3) can be removed from cells 4,...,9.
	for _, cells := range cellsList {
		for c1 := 0; c1 < 9; c1++ {
			optionsC1 := (*cells[c1]).Options
			if len(optionsC1) == 0 {
				continue
			}
			subsetCells := []int{} // = cells with a subset of options (from c1) 
			for c2 := 0; c2 < 9; c2++ {
				if util.Subset(optionsC1, (*cells[c2]).Options) {
					subsetCells = append(subsetCells, c2)
				}
			}
			if len(optionsC1) != len(subsetCells) {
				// Skip if the number of subset cells is not exactly equal to the number of options from c1.
				continue
			}
			for c3 := 0; c3 < 9; c3++ {
				// Remove optionsC1 from all cells except the "subset cells"
				if util.Contains(subsetCells, c3) {
					continue
				}
				removed := cells[c3].removeOptions(optionsC1)
				if removed {
					updated = true
				}
			}
		}
	}

	return updated
}