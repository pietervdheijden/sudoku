package services

import "github.com/pietervdheijden/sudoku/util"

func (sudoku *Sudoku) applyXYZWing() bool {
	// Use method XYZ wing to eliminate options
	// Source: https://www.sudokuwiki.org/XYZ_Wing
	updated := false
	for c1 := 0; c1 < len(sudoku); c1++ {
		if len(sudoku[c1].Options) != 3 {
			continue
		}
		for c2 := 0; c2 < len(sudoku); c2++ {
			if len(sudoku[c2].Options) != 2 {
				continue
			}
			if !sudoku.affect(c1, c2) {
				continue
			}
			if !util.Subset(sudoku[c1].Options, sudoku[c2].Options) {
				continue
			}
			intersectC2 := util.Intersect(sudoku[c2].Options, sudoku[c1].Options)
			if len(intersectC2) != 2 {
				continue
			}
			for c3 := 0; c3 < len(sudoku); c3++ {
				if len(sudoku[c3].Options) != 2 {
					continue
				}
				if util.Equals(sudoku[c1].Options, sudoku[c3].Options) {
					continue
				}
				if !sudoku.affect(c1, c3) {
					continue
				}
				if !util.Subset(sudoku[c1].Options, sudoku[c2].Options) {
					continue
				}
				intersectC3 := util.Intersect(sudoku[c3].Options, sudoku[c1].Options)
				if len(intersectC3) != 2 {
					continue
				}
				intersectC2C3 := util.Intersect(intersectC2, intersectC3)
				if len(intersectC2C3) != 1 {
					continue
				}

				// XYZ wing found! Remove intersectC2C3 from cells affected by c1, c2 and c3.
				for c4 := 0; c4 < len(sudoku); c4++ {
					if c4 == c1 || c4 == c2 || c4 == c3 {
						continue
					}
					if sudoku.affect(c1, c4) && sudoku.affect(c2, c4) && sudoku.affect(c3, c4) {
						removed := sudoku[c4].removeOption(intersectC2C3[0])
						if removed {
							updated = true
						}	
					}
				}
			}
		}
	}
	return updated
}