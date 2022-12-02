package services

import (
	"fmt"
	"strings"
)

type Sudoku [81]Cell

func (sudoku *Sudoku) Solve() bool {
	i := 0
	for {
		i++
		fmt.Printf("######### Iteration %d ##########\n", i)
		sudoku.print()
		if !sudoku.valid() {
			fmt.Println("Sudoku is not valid!")
			return false
		}
        for {
            // Keep updating options until all "updates" are found
            // For example, if there are 3 algorithms and the 3rd algorithm removes an option,
            // run all 3 algorithms again to check if the other algorithms can now also remove an option.
            updated := sudoku.updateOptions()
            if !updated {
                break
            }
        }
		updated := sudoku.updateNumbers()
		if updated {
			fmt.Println("updated")
		}
		if !updated {
			// Guess / Brute force / Forcing pair
			// TODO: refactor
			for i := 0; i < len(sudoku); i++ {
				if sudoku[i].Number == 0 {
					for _, option := range sudoku[i].Options {
						sudokuCopy := sudoku.deepCopy()
						sudokuCopy[i].Number = option
						solved := sudokuCopy.Solve()
						if solved {
							sudoku[i].Number = option
							break
						}
					}
					break
				}
			}
		}
		if sudoku.done() {
			break
		}
	}

	fmt.Println("######################")
	fmt.Println("Done! Sudoku is solved!")
	sudoku.print()
	return true
}

func (sudoku Sudoku) deepCopy() Sudoku {
	// Copy all properties
	sudokuCopy := sudoku

	// Create a deep copy for *slice* options
	// Source: https://www.delftstack.com/howto/go/deep-copy-in-golang/
	for i := 0; i < len(sudoku); i++ {
		sudokuCopy[i].Options = make([]int, len(sudoku[i].Options))
		copy(sudokuCopy[i].Options, sudoku[i].Options)
	}

	return sudokuCopy
}

func (sudoku Sudoku) print() {
	fmt.Println("Sudoku:")

	for row := 0; row < 9; row++ {
		for _, cell := range sudoku.getRow(row) {
			fmt.Printf("%d", (*cell).Number)
		}
		fmt.Printf("\n")
	}
}

func (sudoku Sudoku) printOptions() {
    fmt.Println("Options:")
    for i := 0; i < 9; i++ {
        if i == 0 || i == 3 || i == 6 {
            fmt.Println("__________________________________________________________________________________________")
        }
        cells := sudoku.getRow(i)
        j := 0
        for _, cells := range cells {
            if j == 3 || j == 6 {
                fmt.Printf("|\t")
            }
            j++

            delim := ""
            optionsString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(cells.Options)), delim), "[]")
            fmt.Printf("%v\t", optionsString)
        }
        fmt.Println()
        if (i == 8) {
            fmt.Println("__________________________________________________________________________________________")
        }
    }
}

func (sudoku *Sudoku) getColumn(column int) cell9Collection {
	cells := cell9Collection{}
	i := 0
	for j := 0; j < len(sudoku); j++ {
		if (sudoku[j].Column == column) {
			cells[i] = &sudoku[j]
			i++
		}
	}
	return cells
}

func (sudoku *Sudoku) getRow(row int) cell9Collection{
	cells := cell9Collection{}
	i := 0
	for j := 0; j < len(sudoku); j++ {
		if (sudoku[j].Row == row) {
			cells[i] = &sudoku[j]
			i++
		}
	}
	return cells
}

func (sudoku *Sudoku) getSquare(square int) cell9Collection{
	cells := cell9Collection{}
	i := 0
	for j := 0; j < len(sudoku); j++ {
		if (sudoku[j].Square == square) {
			cells[i] = &sudoku[j]
			i++
		}
	}
	return cells
}

func (sudoku Sudoku) valid() bool {
	for i := 0; i < 9; i++ {
		row := sudoku.getRow(i)
		column := sudoku.getColumn(i)
		square := sudoku.getSquare(i)

		if !row.valid() || !column.valid() || !square.valid() {
			fmt.Printf("i=%d, row.valid()=%v, column.valid()=%v, square.valid()=%v\n", i, row.valid(), column.valid(), square.valid())
			return false
		}
	}
	return true
}

func (sudoku Sudoku) done() bool {
	for i := 0; i < len(sudoku); i++ {
		if sudoku[i].Number == 0 {
			fmt.Printf("Not done: sudoku[row:%d][column:%d] == 0\n", sudoku[i].Row, sudoku[i].Column)
			return false
		}
	}
	return true
}

func (sudoku *Sudoku) updateOptions() bool {
    fmt.Println("Update options")
    
	updated := false
	updated = sudoku.applyEliminateByNumbers9Cells() || updated 
	updated = sudoku.applyEliminateBySubsets9Cells() || updated 
	updated = sudoku.applyEliminateByOverlapping9Cells() || updated
	updated = sudoku.applyUniqueRectangle() || updated
	updated = sudoku.applyXWing() || updated
	updated = sudoku.applyXYWing() || updated
	updated = sudoku.applyXYZWing() || updated
	// TODO: implement swordfish and jellyfish
	// - Swordfish: https://www.sudokuwiki.org/Sword_Fish_Strategy
	// - Jellyfish: https://www.sudokuwiki.org/Jelly_Fish_Strategy 	

    return updated
}

func (sudoku *Sudoku) updateNumbers() bool {
    fmt.Println("updateNumbers()")
	updated := false

	// Set number when the cell only has 1 option
	for i := 0; i < len(sudoku); i++ {
		if len(sudoku[i].Options) == 1 {
			sudoku[i].Number = sudoku[i].Options[0]
			sudoku[i].Options = []int{}
			updated = true
		}
	}

	// Update numbers for each row, column and square
	for i := 0; i < 9; i++ {
		row := sudoku.getRow(i)
		column := sudoku.getColumn(i)
		square := sudoku.getSquare(i)

		updatedRow := row.updateNumbers()
		updatedColumn := column.updateNumbers()
		updatedSquare := square.updateNumbers()
		if updatedRow || updatedColumn || updatedSquare {
			updated = true
		}
	}

	return updated
}

// returns if cell c1 and c2 affect each other
func (sudoku *Sudoku) affect(c1 int, c2 int) bool {
	return sudoku[c1].Row == sudoku[c2].Row || sudoku[c1].Column == sudoku[c2].Column || sudoku[c1].Square == sudoku[c2].Square
}
