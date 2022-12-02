package services

import "github.com/pietervdheijden/sudoku/util"

func (sudoku *Sudoku) applyEliminateByOverlapping9Cells() bool {
	// TODO: use standard terminology for this technique.
	// Eliminate options by "9cells" (9cells could be a row, column or square)
	// When number X is missing in "9cells" and the options which include X have exactly 1 row, 1 column or 1 square,
	// X can be eliminated for the other cells in the same row, column or square.
	// For example, consider row 1 where number 2 has to be in either cell 0,1,2, which all fall in square 1.
	// Then number 2 can be eliminated for the other cells in square 1.
	updated := false

	cellsList := []cell9Collection{}
	for i := 0; i < 9; i++ {
		row := sudoku.getRow(i)
		column := sudoku.getColumn(i)
		square := sudoku.getSquare(i)

		cellsList = append(cellsList, row, column, square)
	}

	for _, cells := range cellsList {  
		for number := 1; number <= 9; number++ {
			found := false
			rows := []int{}
			columns := []int{}
			squares := []int{}
			for c1 := 0; c1 < len(cells); c1++ {
				if cells[c1].Number == number {
					found = true
					break
				}
				if !util.Contains(cells[c1].Options, number) {
					continue
				}
				if !util.Contains(rows, cells[c1].Row) {
					rows = append(rows, cells[c1].Row)
				}
				if !util.Contains(columns, cells[c1].Column) {
					columns = append(columns, cells[c1].Column)
				}
				if !util.Contains(squares, cells[c1].Square) {
					squares = append(squares, cells[c1].Square)
				}
			}
			if found {
				continue
			}

			if len(rows) == 1 {
				row := sudoku.getRow(rows[0])
				for c2 := 0; c2 < len(row); c2++ {
					if cells.contains(row[c2].Id) {
						continue
					}
					removed := row[c2].removeOption(number)
					if removed {
						updated = true
					}
				}
			}
			if len(columns) == 1 {
				col := sudoku.getColumn(columns[0])
				for c2 := 0; c2 < len(col); c2++ {
					if cells.contains(col[c2].Id) {
						continue
					}
					removed := col[c2].removeOption(number)
					if removed {
						updated = true
					}
				}
			}
			if len(squares) == 1 {
				square := sudoku.getSquare(squares[0])
				for c2 := 0; c2 < len(square); c2++ {
					if cells.contains(square[c2].Id) {
						continue
					}
					removed := square[c2].removeOption(number)
					if removed {
						updated = true
					}
				}
			}
		}
	}

	return updated
}