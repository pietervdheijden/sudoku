package services

import "github.com/pietervdheijden/sudoku/util"

func (sudoku *Sudoku) applyUniqueRectangle() bool {
	updated := false
	// Use method unique rectangle to eliminate options
    // Docs: https://www.learn-sudoku.com/unique-rectangle.html
	for c1 := 0; c1 < len(sudoku); c1++  {
        if len(sudoku[c1].Options) != 2 {
            // skip if cell does not have exactly 2 options
            continue
        }
		for c2 := c1 + 1; c2 < len(sudoku); c2++ {
			if len(sudoku[c2].Options) != 2 {
                // skip if cell does not have exactly 2 options
                continue
            }
            for c3 := c2 + 1; c3 < len(sudoku); c3++ {
                if len(sudoku[c3].Options) != 2 {
                    // skip if cell does not have exactly 2 options
                    continue
                }
				if c1 == c2 || c2 == c3 || c1 == c3 {
					// skip if cells (partially) overlap
					continue
				}
				if sudoku[c1].Square != sudoku[c2].Square && sudoku[c2].Square != sudoku[c3].Square && sudoku[c1].Square != sudoku[c3].Square {
					// 2 cells have to be in the same square
					continue
				}
				if sudoku[c1].Row != sudoku[c2].Row && sudoku[c2].Row != sudoku[c3].Row && sudoku[c1].Row != sudoku[c3].Row {
					// 2 cells have to be in the same row
					continue
				}
				if sudoku[c1].Column != sudoku[c2].Column && sudoku[c2].Column != sudoku[c3].Column && sudoku[c1].Column != sudoku[c3].Column {
					// 2 cells have to be in the same column
					continue
				}
				if !util.Equals(sudoku[c1].Options, sudoku[c2].Options) || !util.Equals(sudoku[c2].Options, sudoku[c3].Options) || !util.Equals(sudoku[c1].Options, sudoku[c3].Options) {
					// skip if cells do not have the same options 
					continue
				}
				c4 := c1 + c3 - c2
				if !util.Subset(sudoku[c4].Options, sudoku[c1].Options) {
					// skip if c4 does not contain all options from c1 (note: c1 has the same options as c2 and c3, see above)
					continue
				}

				// Unique rectangle found! Remove c1.Options from c4.Options
				sudoku[c4].removeOptions(sudoku[c1].Options)
                updated = true		
			}
		}
	}

	return updated
}