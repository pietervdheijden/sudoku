package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
)

func main() {
    // solve("sudoku_level1.txt")
	// solve("sudoku_level6_1.txt") // edition 225, puzzle 75
	// solve("sudoku_level6_2.txt") // edition 225, puzzle 77
	// solve("sudoku_level6_3.txt") // edition 225, puzzle 79
	// solve("sudoku_level6_4.txt") // edition 225, puzzle 83
	// solve("sudoku_level8_1.txt") // edition 134, puzzle 26
	// solve("sudoku_level8_2.txt") // edition 134, puzzle 28
	// solve("sudoku_level8_3.txt") // edition 134, puzzle 30
	solve("sudoku_level9_1.txt") // edition 134, puzzle 7
	// solve("sudoku_level9_2.txt") // edition 134, puzzle 17
	// solve("sudoku_level9_3.txt") // edition 134, puzzle 25
	// solve("sudoku_level9_4.txt") // edition 134, puzzle 27
	// solve("sudoku_level9_xwing.txt") // edition 134, techniques, x-wing
	// solve("sudoku_level9_ywing.txt") // edition 134, techniques, y-wing
}

func solve(puzzle string) {
	sudoku := Sudoku{}
    sudoku.read(puzzle)
	i := 0
	for {
		i++
		fmt.Printf("######### Iteration %d ##########\n", i)
		sudoku.print()
		if !sudoku.valid() {
			panic("Sudoku is not valid!")
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
		if !updated {
			sudoku.print()
            sudoku.printOptions()
			panic("Couldn't find an option to move. Sudoku is too hard for the algorithm!")
		}
		if sudoku.done() {
			break
		}
	}

	fmt.Println("######################")
	fmt.Println("Done! Sudoku is solved!")
	sudoku.print()
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

func (cells Cell9Collection) valid() bool{
	numbers := []int{}
	for _, cell := range cells {
		number := (*cell).number
		if contains(numbers, number) {
			return false
		}
		if number != 0 {
			numbers = append(numbers, number)
		}
	}
	return true
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

		updatedRow := row.updateOptions()
		updatedColumn := column.updateOptions()
		updatedSquare := square.updateOptions()
        if updatedRow || updatedColumn || updatedSquare {
            updated = true
        }
    }

	// TODO: foreach square, check all rows and cols
	// TODO: refactor
	// TODO: also run for cols and rows (in (cells Cell9Collection) updateOptions()?)
	for i := 0; i < 9; i++ {
		square := sudoku.getSquare(i)
		for number := 1; number <= 9; number++ {
			found := false
			rows := []int{}
			cols := []int{}
			for c1 := 0; c1 < 9; c1++ {
				if square[c1].number == number {
					found = true
					break
				}
				if !contains(square[c1].options, number) {
					continue
				}
				if !contains(rows, square[c1].row) {
					rows = append(rows, square[c1].row)
				}
				if !contains(cols, square[c1].column) {
					cols = append(cols, square[c1].column)
				}
			}
			if found {
				continue
			}
			if len(rows) == 1 {
				row := sudoku.getRow(rows[0])
				for c2 := 0; c2 < 9; c2++ {
					if (row[c2].square == i) {
						continue
					}
					removed := row[c2].removeOption(number)
                    if removed {
                        updated = true
                    }
				}
			}
			if len(cols) == 1 {
				col := sudoku.getColumn(cols[0])
				for c2 := 0; c2 < 9; c2++ {
					if (col[c2].square == i) {
						continue
					}
					removed := col[c2].removeOption(number)
                    if removed {
                        updated = true
                    }
				}
			}
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

	// TODO: implement X wing
	// TODO: implement XY wing
	// TODO: implement XYZ wing

    return updated
}

func (cells Cell9Collection) updateOptions() bool {
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