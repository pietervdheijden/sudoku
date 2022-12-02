package services

import "github.com/pietervdheijden/sudoku/util"

func (sudoku *Sudoku) applyXWing() bool {
	// Use method X wing to eliminate options
	// Source: https://www.learn-sudoku.com/x-wing.html
	updated := false
	for c1 := 0; c1 < len(sudoku); c1++ {
		for c2 := c1+1; c2 < len(sudoku); c2++ {
			for c3 := c2+1; c3 < len(sudoku); c3++ {
				for c4 := c3+1; c4 < len(sudoku); c4++ {
					if sudoku[c1].Row != sudoku[c2].Row || sudoku[c3].Row != sudoku[c4].Row {
						continue
					}
					if sudoku[c1].Column != sudoku[c3].Column || sudoku[c2].Column != sudoku[c4].Column {
						continue
					}
					intersect := util.IntersectSlices(sudoku[c1].Options, sudoku[c2].Options, sudoku[c3].Options, sudoku[c4].Options)
					if len(intersect) == 0 {
						continue
					}

					// Check if any of the other cells in the same *row* contain the intersect option
					for _, option := range intersect {
						found := false
						for c5 := 0; c5 < len(sudoku); c5++ {
							if c5 == c1 || c5 == c2 || c5 == c3 || c5 == c4 {
								continue
							}
							if sudoku[c5].Row != sudoku[c1].Row && sudoku[c5].Row != sudoku[c3].Row {
								continue
							} 
							if util.Contains(sudoku[c5].Options, option) {
								found = true
							}
						}
						if found {
							continue
						}

						// None of the other cells in the same *rows* contain the intersect option
						// X wing found! Remove option from other cells in the same *columns*.
						for c5 := 0; c5 < len(sudoku); c5++ {
							if c5 == c1 || c5 == c2 || c5 == c3 || c5 == c4 {
								continue
							}
							if sudoku[c5].Column != sudoku[c1].Column && sudoku[c5].Column != sudoku[c2].Column {
								continue
							}
							removed := sudoku[c5].removeOption(option)
							if removed {
								updated = true
							}
						}
					}

					// Check if any of the other cells in the same *column* contain the intersect option
					for _, option := range intersect {
						found := false
						for c5 := 0; c5 < len(sudoku); c5++ {
							if c5 == c1 || c5 == c2 || c5 == c3 || c5 == c4 {
								continue
							}
							if sudoku[c5].Column != sudoku[c1].Column && sudoku[c5].Column != sudoku[c2].Column {
								continue
							} 
							if util.Contains(sudoku[c5].Options, option) {
								found = true
							}
						}
						if found {
							continue
						}

						// None of the other cells in the same *columns* contain the intersect option
						// X wing found! Remove option from other cells in the same *rows*.
						for c5 := 0; c5 < len(sudoku); c5++ {
							if c5 == c1 || c5 == c2 || c5 == c3 || c5 == c4 {
								continue
							}
							if sudoku[c5].Row != sudoku[c1].Row && sudoku[c5].Row != sudoku[c3].Row {
								continue
							}
							removed := sudoku[c5].removeOption(option)
							if removed {
								updated = true
							}
						}
					}
				}
			}
		}
	}
	return updated
}