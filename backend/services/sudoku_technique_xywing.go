package services

import "github.com/pietervdheijden/sudoku/util"

func (sudoku *Sudoku) applyXYWing() bool {
	// Use method XY wing to eliminate options
	// Source: https://www.learn-sudoku.com/xy-wing.html
	updated := false
	for c1 := 0; c1 < len(sudoku); c1++ {
		if len(sudoku[c1].Options) != 2 {
			// skip if c1 does not exactly have 2 options
			continue
		}
		for c2 := 0; c2 < len(sudoku); c2++ {
			if len(sudoku[c2].Options) != 2 {
				// skip if c2 does not exactly have 2 options
				continue
			}
			if util.Equals(sudoku[c1].Options, sudoku[c2].Options) {
				// skip if c2 has exactly the same options as c1
				continue
			}
			if !sudoku.affect(c1, c2) {
				// skip if c1 and c2 do not affect each other
				continue
			} 
			intersectC2 := util.Intersect(sudoku[c2].Options, sudoku[c1].Options)
			exceptC2 := util.Except(sudoku[c2].Options, sudoku[c1].Options)
			if len(intersectC2) != 1 || len(exceptC2) != 1 {
				continue
			}
			for c3 := 0; c3 < len(sudoku); c3++ {
				if len(sudoku[c3].Options) != 2 {
					// skip if c3 does not exactly have 2 options
					continue
				}
				if util.Equals(sudoku[c1].Options, sudoku[c3].Options) {
					// skip if c3 has exactly the same options as c1
					continue
				}
				if !sudoku.affect(c1, c3) {
					// skip if c1 and c3 do not affect eachother
					continue
				} 
				intersectC3 := util.Intersect(sudoku[c3].Options, sudoku[c1].Options)
				exceptC3 := util.Except(sudoku[c3].Options, sudoku[c1].Options)
				if len(intersectC3) != 1 || len(exceptC3) != 1 {
					continue
				}
				if intersectC2[0] == intersectC3[0] || exceptC2[0] != exceptC3[0] {
					continue
				}

				// XY wing found! Remove exceptC2 (=exceptC3) from cells affected by c2 and c3.
				for c4 := 0; c4 < len(sudoku); c4++ {
					if c4 == c1 || c4 == c2 || c4 == c3 {
						continue
					}
					if sudoku.affect(c2, c4) && sudoku.affect(c3, c4) {
						removed := sudoku[c4].removeOption(exceptC2[0])
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