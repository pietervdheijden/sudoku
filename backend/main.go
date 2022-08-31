package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)
type sudokuApiModel struct {
	Sudoku [81]int `json:"sudoku" binding:"required"`
}

func main() {
	router := gin.Default()
	// TODO: add difficulty to solve
	router.POST("/solve", solveSudoku)
	router.POST("/hint", hintSudoku)
	router.Run(":8080")
}

func solveSudoku(c *gin.Context) {
	var sudokuApiModel sudokuApiModel

	if err := c.BindJSON(&sudokuApiModel); err != nil {
		fmt.Println(err)
		return
	}

	sudoku := Sudoku{}
	sudoku.readFromArray(sudokuApiModel.Sudoku)
	solved := sudoku.solve()

	if solved {
		c.JSON(http.StatusOK, sudoku.toApiModel())
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Sudoku could not be solved!"})
	}
}

func hintSudoku(c *gin.Context) {
	// TODO: implement hint

}

func (sudoku *Sudoku) toApiModel() sudokuApiModel {
	sudokuApiModel := sudokuApiModel{}
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			id := row * 9 + column
			sudokuApiModel.Sudoku[id] = sudoku[id].number
		}
	}
	return sudokuApiModel
}

func test() {
	puzzles := []string{
		
		// TODO: add unit tests
		"puzzles/sudoku_level1.txt",
		"puzzles/sudoku_level6_1.txt", // edition 225, puzzle 7
		"puzzles/sudoku_level6_2.txt", // edition 225, puzzle 75
		"puzzles/sudoku_level6_3.txt", // edition 225, puzzle 77
		"puzzles/sudoku_level6_4.txt", // edition 225, puzzle 79
		"puzzles/sudoku_level6_5.txt", // edition 225, puzzle 83
		"puzzles/sudoku_level8_1.txt", // edition 134, puzzle 26
		"puzzles/sudoku_level8_2.txt", // edition 134, puzzle 28
		"puzzles/sudoku_level8_3.txt", // edition 134, puzzle 30
		"puzzles/sudoku_level9_1.txt", // edition 134, puzzle 17 (contradiction / forcing chain)
		"puzzles/sudoku_level9_2.txt", // edition 134, puzzle 25 (contradiction)
		"puzzles/sudoku_level9_3.txt", // edition 134, puzzle 27 (XY wing + XYZ wing)
		"puzzles/sudoku_level9_4.txt", // edition 134, puzzle 29 (contradiction)
		"puzzles/sudoku_level9_5.txt", // edition 134, puzzle 31 (X wing)
		"puzzles/sudoku_level9_6.txt", // edition 134, puzzle 33 (contradiction)
		"puzzles/sudoku_level9_7.txt", // edition 104, puzzle 9
		"puzzles/sudoku_level9_8.txt", // edition 104, puzzle 11
		"puzzles/sudoku_level9_9.txt", // edition 104, puzzle 13 (contradiction)
		"puzzles/sudoku_level9_10.txt", // edition 104, puzzle 15
		"puzzles/sudoku_level9_11.txt", // edition 104, puzzle 17
		"puzzles/sudoku_level9_12.txt", // edition 104, puzzle 19 (contradiction)		
	}

	for _, puzzle := range puzzles {
		sudoku := Sudoku{}
		sudoku.read(puzzle)
		success := sudoku.solve()
		if !success {
			fmt.Printf("Sudoku %v could not be solved...\n", puzzle)
			break
		}
		fmt.Printf("Sudoku %v: solved!\n", puzzle)
	}
}

func (sudoku *Sudoku) solve() bool {
	i := 0
	for {
		i++
		fmt.Printf("######### Iteration %d ##########\n", i)
		sudoku.print()
		if !sudoku.valid() {
			// panic("Sudoku is not valid!")
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

			// TODO: implement swordfish and jellyfish
        }
		updated := sudoku.updateNumbers()
		if updated {
			fmt.Println("updated")
		}
		if !updated {
			// Guess / Brute force / Forcing pair
			// TODO: refactor
			for i := 0; i < len(sudoku); i++ {
				if sudoku[i].number == 0 {
					for _, option := range sudoku[i].options {
						sudokuCopy := sudoku.deepCopy()
						sudokuCopy[i].number = option
						solved := sudokuCopy.solve()
						if solved {
							sudoku[i].number = option
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

// TODO: consider renaming options to candidates
// TODO: consider renaming Cell9Collection to 9Cells
// TODO: consider creating type for options
type Cell struct {
	id int
	row int
	column int
	square int
	number int
	options []int
}
type Sudoku [81]Cell
type Cell9Collection [9]*Cell

func (sudoku *Sudoku) read(fileLocation string) {
	fmt.Printf("Read Sudoku from fileLocation: %v\n", fileLocation)
	file, err := os.Open(fileLocation)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 0
	row := 0
    for scanner.Scan() {
		for column, v := range strings.Split(scanner.Text(), "") {
			number, err := strconv.Atoi(v)
			if err != nil {
				// ... handle error
				panic(err)
			}
			options := []int{1,2,3,4,5,6,7,8,9}
			if number != 0 {
				options = []int{}
			}
			square := (row/3)*3 + column/3 // (x/3)*3 != x, e.g. (4/3)*3=3
			sudoku[id] = Cell{
				id: id,
				row: row,
				column: column,
				square: square,
				number: number,
				options: options,	
			}
			id++
		}
		row++
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func (sudoku *Sudoku) readFromArray(input [81]int) {
	fmt.Printf("Read Sudoku from array: %v\n", input)
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			id := row * 9 + column
			number := input[id]
			options := []int{1,2,3,4,5,6,7,8,9}
			if number != 0 {
				options = []int{}
			}
			square := (row/3)*3 + column/3 // (x/3)*3 != x, e.g. (4/3)*3=3
			sudoku[id] = Cell{
				id: id,
				row: row,
				column: column,
				square: square,
				number: number,
				options: options,	
			}
		}
	}
}

func (sudoku Sudoku) deepCopy() Sudoku {
	// Copy all properties
	sudokuCopy := sudoku

	// Create a deep copy for *slice* options
	// Source: https://www.delftstack.com/howto/go/deep-copy-in-golang/
	for i := 0; i < len(sudoku); i++ {
		sudokuCopy[i].options = make([]int, len(sudoku[i].options))
		copy(sudokuCopy[i].options, sudoku[i].options)
	}

	return sudokuCopy
}

func (sudoku Sudoku) print() {
	fmt.Println("Sudoku:")

	for row := 0; row < 9; row++ {
		for _, cell := range sudoku.getRow(row) {
			fmt.Printf("%d", (*cell).number)
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
            optionsString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(cells.options)), delim), "[]")
            fmt.Printf("%v\t", optionsString)
        }
        fmt.Println()
        if (i == 8) {
            fmt.Println("__________________________________________________________________________________________")
        }
    }
}

func (sudoku *Sudoku) getColumn(column int) Cell9Collection {
	cells := Cell9Collection{}
	i := 0
	for j := 0; j < len(sudoku); j++ {
		if (sudoku[j].column == column) {
			cells[i] = &sudoku[j]
			i++
		}
	}
	return cells
}

func (sudoku *Sudoku) getRow(row int) Cell9Collection{
	cells := Cell9Collection{}
	i := 0
	for j := 0; j < len(sudoku); j++ {
		if (sudoku[j].row == row) {
			cells[i] = &sudoku[j]
			i++
		}
	}
	return cells
}

func (sudoku *Sudoku) getSquare(square int) Cell9Collection{
	cells := Cell9Collection{}
	i := 0
	for j := 0; j < len(sudoku); j++ {
		if (sudoku[j].square == square) {
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

func (cells Cell9Collection) valid() bool {
	numbers := []int{}
	for _, cell := range cells {
		number := (*cell).number
		if contains(numbers, number) {
			return false
		}
		if number != 0 {
			numbers = append(numbers, number)
		}
		if cell.number == 0 && len(cell.options) == 0 {
			return false
		}
	}
	return true
}

func (cells Cell9Collection) contains(id int) bool {
	for _, cell := range cells {
		if cell.id == id {
			return true
		}
	}
	return false
}

func (sudoku Sudoku) done() bool {
	for i := 0; i < len(sudoku); i++ {
		if sudoku[i].number == 0 {
			fmt.Printf("Not done: sudoku[row:%d][column:%d] == 0\n", sudoku[i].row, sudoku[i].column)
			return false
		}
	}
	return true
}

func (sudoku *Sudoku) updateOptions() bool {
    fmt.Println("Update options")
    updated := false

	// When number != 0, set options to empty array
	for i := 0; i < len(sudoku); i++ {
		if sudoku[i].number != 0 && len(sudoku[i].options) != 0 {
			sudoku[i].options = []int{}
            updated = true
		}
	}

	// Update options for each row, column and square
	for i := 0; i < 9; i++ {
		row := sudoku.getRow(i)
		column := sudoku.getColumn(i)
		square := sudoku.getSquare(i)

		updatedRow := sudoku.updateOptionsForCells(row)
		updatedColumn := sudoku.updateOptionsForCells(column)
		updatedSquare := sudoku.updateOptionsForCells(square)
        if updatedRow || updatedColumn || updatedSquare {
            updated = true
        }
	}

	// Use method unique rectangle to eliminate options
    // Docs: https://www.learn-sudoku.com/unique-rectangle.html
	for c1 := 0; c1 < len(sudoku); c1++  {
        if len(sudoku[c1].options) != 2 {
            // skip if cell does not have exactly 2 options
            continue
        }
		for c2 := c1 + 1; c2 < len(sudoku); c2++ {
			if len(sudoku[c2].options) != 2 {
                // skip if cell does not have exactly 2 options
                continue
            }
            for c3 := c2 + 1; c3 < len(sudoku); c3++ {
                if len(sudoku[c3].options) != 2 {
                    // skip if cell does not have exactly 2 options
                    continue
                }
				if c1 == c2 || c2 == c3 || c1 == c3 {
					// skip if cells (partially) overlap
					continue
				}
				if sudoku[c1].square != sudoku[c2].square && sudoku[c2].square != sudoku[c3].square && sudoku[c1].square != sudoku[c3].square {
					// 2 cells have to be in the same square
					continue
				}
				if sudoku[c1].row != sudoku[c2].row && sudoku[c2].row != sudoku[c3].row && sudoku[c1].row != sudoku[c3].row {
					// 2 cells have to be in the same row
					continue
				}
				if sudoku[c1].column != sudoku[c2].column && sudoku[c2].column != sudoku[c3].column && sudoku[c1].column != sudoku[c3].column {
					// 2 cells have to be in the same column
					continue
				}
				if !equals(sudoku[c1].options, sudoku[c2].options) || !equals(sudoku[c2].options, sudoku[c3].options) || !equals(sudoku[c1].options, sudoku[c3].options) {
					// skip if cells do not have the same options 
					continue
				}
				c4 := c1 + c3 - c2
				if !subset(sudoku[c4].options, sudoku[c1].options) {
					// skip if c4 does not contain all options from c1 (note: c1 has the same options as c2 and c3, see above)
					continue
				}

				// Unique rectangle found! Remove c1.options from c4.options
				sudoku[c4].removeOptions(sudoku[c1].options)
                updated = true		
			}
		}
	}

	// Use method X wing to eliminate options
	// Source: https://www.learn-sudoku.com/x-wing.html
	for c1 := 0; c1 < len(sudoku); c1++ {
		for c2 := c1+1; c2 < len(sudoku); c2++ {
			for c3 := c2+1; c3 < len(sudoku); c3++ {
				for c4 := c3+1; c4 < len(sudoku); c4++ {
					if sudoku[c1].row != sudoku[c2].row || sudoku[c3].row != sudoku[c4].row {
						continue
					}
					if sudoku[c1].column != sudoku[c3].column || sudoku[c2].column != sudoku[c4].column {
						continue
					}
					intersect := intersectSlices(sudoku[c1].options, sudoku[c2].options, sudoku[c3].options, sudoku[c4].options)
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
							if sudoku[c5].row != sudoku[c1].row && sudoku[c5].row != sudoku[c3].row {
								continue
							} 
							if contains(sudoku[c5].options, option) {
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
							if sudoku[c5].column != sudoku[c1].column && sudoku[c5].column != sudoku[c2].column {
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
							if sudoku[c5].column != sudoku[c1].column && sudoku[c5].column != sudoku[c2].column {
								continue
							} 
							if contains(sudoku[c5].options, option) {
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
							if sudoku[c5].row != sudoku[c1].row && sudoku[c5].row != sudoku[c3].row {
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

	// Use method XY wing to eliminate options
	// Source: https://www.learn-sudoku.com/xy-wing.html
	for c1 := 0; c1 < len(sudoku); c1++ {
		if len(sudoku[c1].options) != 2 {
			// skip if c1 does not exactly have 2 options
			continue
		}
		for c2 := 0; c2 < len(sudoku); c2++ {
			if len(sudoku[c2].options) != 2 {
				// skip if c2 does not exactly have 2 options
				continue
			}
			if equals(sudoku[c1].options, sudoku[c2].options) {
				// skip if c2 has exactly the same options as c1
				continue
			}
			if !sudoku.affect(c1, c2) {
				// skip if c1 and c2 do not affect each other
				continue
			} 
			intersectC2 := intersect(sudoku[c2].options, sudoku[c1].options)
			exceptC2 := except(sudoku[c2].options, sudoku[c1].options)
			if len(intersectC2) != 1 || len(exceptC2) != 1 {
				continue
			}
			for c3 := 0; c3 < len(sudoku); c3++ {
				if len(sudoku[c3].options) != 2 {
					// skip if c3 does not exactly have 2 options
					continue
				}
				if equals(sudoku[c1].options, sudoku[c3].options) {
					// skip if c3 has exactly the same options as c1
					continue
				}
				if !sudoku.affect(c1, c3) {
					// skip if c1 and c3 do not affect eachother
					continue
				} 
				intersectC3 := intersect(sudoku[c3].options, sudoku[c1].options)
				exceptC3 := except(sudoku[c3].options, sudoku[c1].options)
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

	// Use method XYZ wing to eliminate options
	// Source: https://www.sudokuwiki.org/XYZ_Wing
	for c1 := 0; c1 < len(sudoku); c1++ {
		if len(sudoku[c1].options) != 3 {
			continue
		}
		for c2 := 0; c2 < len(sudoku); c2++ {
			if len(sudoku[c2].options) != 2 {
				continue
			}
			if !sudoku.affect(c1, c2) {
				continue
			}
			if !subset(sudoku[c1].options, sudoku[c2].options) {
				continue
			}
			intersectC2 := intersect(sudoku[c2].options, sudoku[c1].options)
			if len(intersectC2) != 2 {
				continue
			}
			for c3 := 0; c3 < len(sudoku); c3++ {
				if len(sudoku[c3].options) != 2 {
					continue
				}
				if equals(sudoku[c1].options, sudoku[c3].options) {
					continue
				}
				if !sudoku.affect(c1, c3) {
					continue
				}
				if !subset(sudoku[c1].options, sudoku[c2].options) {
					continue
				}
				intersectC3 := intersect(sudoku[c3].options, sudoku[c1].options)
				if len(intersectC3) != 2 {
					continue
				}
				intersectC2C3 := intersect(intersectC2, intersectC3)
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

func (sudoku *Sudoku) updateOptionsForCells(cells Cell9Collection) bool {
	updated := false

	// When number != 0, remove the option from the other cells 
	for c1 := 0; c1 < len(cells); c1++ {
		for c2 := 0; c2 < len(cells); c2++ {
			numberC2 := (*cells[c2]).number 
			if c1 == c2 || numberC2 == 0 {
				continue
			}
            removed := cells[c1].removeOption(numberC2)
            if removed {
                updated = true
            }
		}
	}

	// Eliminate options by subsets
	// When cell C1 has X options, and there exist X-1 other cells (c2) with a subset of X, then the X options can be removed from the remaining cells (c3).
	// Example 1: when cell 1 and 2 both have options (1,2), then options (1,2) can be removed from cells 3,...,9.
	// Example 2: when cell 1 has options (1,2,3), cell 2 has options (1,2) and cell 3 has options (2,3), then options (1,2,3) can be removed from cells 4,...,9.
	for c1 := 0; c1 < 9; c1++ {
		optionsC1 := (*cells[c1]).options
		if len(optionsC1) == 0 {
			continue
		}
		subsetCells := []int{} // = cells with a subset of options (from c1) 
		for c2 := 0; c2 < 9; c2++ {
			if subset(optionsC1, (*cells[c2]).options) {
				subsetCells = append(subsetCells, c2)
			}
		}
		if len(optionsC1) != len(subsetCells) {
			// Skip if the number of subset cells is not exactly equal to the number of options from c1.
			continue
		}
		for c3 := 0; c3 < 9; c3++ {
			// Remove optionsC1 from all cells except the "subset cells"
			if contains(subsetCells, c3) {
				continue
			}
			removed := cells[c3].removeOptions(optionsC1)
			if removed {
				updated = true
			}
		}
	}

	// Eliminate options by "9cells" (9cells could be a row, column or square)
	// When number X is missing in "9cells" and the options which include X have exactly 1 row, 1 column or 1 square,
	// X can be eliminated for the other cells in the same row, column or square.
	// For example, consider row 1 where number 2 has to be in either cell 0,1,2, which all fall in square 1.
	// Then number 2 can be eliminated for the other cells in square 1.
	for number := 1; number <= 9; number++ {
		found := false
		rows := []int{}
		columns := []int{}
		squares := []int{}
		for c1 := 0; c1 < len(cells); c1++ {
			if cells[c1].number == number {
				found = true
				break
			}
			if !contains(cells[c1].options, number) {
				continue
			}
			if !contains(rows, cells[c1].row) {
				rows = append(rows, cells[c1].row)
			}
			if !contains(columns, cells[c1].column) {
				columns = append(columns, cells[c1].column)
			}
			if !contains(squares, cells[c1].square) {
				squares = append(squares, cells[c1].square)
			}
		}
		if found {
			continue
		}

		if len(rows) == 1 {
			row := sudoku.getRow(rows[0])
			for c2 := 0; c2 < len(row); c2++ {
				if cells.contains(row[c2].id) {
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
				if cells.contains(col[c2].id) {
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
				if cells.contains(square[c2].id) {
					continue
				}
				removed := square[c2].removeOption(number)
				if removed {
					updated = true
				}
			}
		}
	}

	return updated
}

func (sudoku *Sudoku) updateNumbers() bool {
    fmt.Println("updateNumbers()")
	updated := false

	// Set number when the cell only has 1 option
	for i := 0; i < len(sudoku); i++ {
		if len(sudoku[i].options) == 1 {
			sudoku[i].number = sudoku[i].options[0]
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

func (cells Cell9Collection) updateNumbers() bool {
	updated := false

	// Set number when there's only a single cell which options contains the number
	for number := 1; number <= 9; number++ {
		numberFound := false
		potentialCellCount := 0
		lastCellIndex := -1
		for i := 0; i < 9; i++ {
			if (*cells[i]).number == number {
				numberFound = true
				break
			}
			if contains((*cells[i]).options, number) {
				potentialCellCount++
				lastCellIndex = i
			}
		}
		if !numberFound && potentialCellCount == 1 {
			(*cells[lastCellIndex]).number = number
			updated = true
		}
	}
	return updated
}

func equals(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, e := range s2 {
		if !contains(s1, e) {
			return false
		}
	}
	return true
}

func subset(s1 []int, s2 []int) bool {
	if len(s2) > len(s1) || len(s2) == 0 {
		return false
	}
	for _, e := range s2 {
		if !contains(s1, e) {
			return false
		}
	}
	return true
}

func merge(s1 []int, s2 []int) []int {
	s3 := []int{}
	for _, v := range s1 {
		if !contains(s3, v) {
			s3 = append(s3, v)
		}
	}
	for _, v := range s2 {
		if !contains(s3, v) {
			s3 = append(s3, v)
		}
	}
	return s3
}

// source: https://stackoverflow.com/a/10485970/3737152
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

// returns elements both in s1 and s2
func intersect(s1 []int, s2 []int) []int {
	s3 := []int{}
	for _, v := range s1 {
		if contains(s2, v) {
			s3 = append(s3, v)
		}
	}
	return s3
}

func intersectSlices(slices ...[]int) []int {
	s := []int{}
	for _, slice := range slices {
		intersect := true
		for _, v := range slice {
			for _, otherSlice := range slices {
				if !contains(otherSlice, v) {
					intersect = false
				}
			}
			if intersect && !contains(s, v) {
				s = append(s, v)
			}
		}
	}
	return s
}

// returns element in s1, but not in s2
func except(s1 []int, s2 []int) []int {
	s3 := []int{}
	for _, v := range s1 {
		if !contains(s2, v) {
			s3 = append(s3, v)
		}
	}
	return s3
}

// returns if cell c1 and c2 affect each other
func (sudoku *Sudoku) affect(c1 int, c2 int) bool {
	return sudoku[c1].row == sudoku[c2].row || sudoku[c1].column == sudoku[c2].column || sudoku[c1].square == sudoku[c2].square
}

func (cell *Cell) removeOption(value int) bool {
	for i, option := range (*cell).options {
		if option == value {
			(*cell).options = append((*cell).options[:i], (*cell).options[i+1:]...)
			return true
		}
	}
    return false
}

func (cell *Cell) removeOptions(options []int) bool {
	removed := false
	for _, option := range options {
		r := cell.removeOption(option)
		if r {
			removed = true
		}
	}
	return removed
}