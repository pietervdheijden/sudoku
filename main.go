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
	fmt.Println("Let's solve this Sudoku!")

	sudoku := Sudoku{}
    
	fmt.Println("First, read the Sudoku from file")
	// sudoku.read("sudoku_level1.txt")
	// sudoku.read("sudoku_level6_1.txt")
	// sudoku.read("sudoku_level6_2.txt") // edition 225, puzzle 77
	// sudoku.read("sudoku_level6_3.txt") // edition 225, puzzle 79
	// sudoku.read("sudoku_level6_4.txt") // edition 225, puzzle 83
	// sudoku.read("sudoku_level8_1.txt") // edition 134, puzzle 26
	// sudoku.read("sudoku_level8_2.txt") // edition 134, puzzle 28
	// sudoku.read("sudoku_level8_3.txt") // edition 134, puzzle 30
	sudoku.read("sudoku_level9_1.txt") // edition 134, puzzle 7
	// sudoku.read("sudoku_level9_2.txt") // edition 134, puzzle 17
	// sudoku.read("sudoku_level9_3.txt") // edition 134, puzzle 25
	// sudoku.read("sudoku_level9_3_wip.txt") // edition 134, puzzle 25
	// sudoku.read("sudoku_level9_4.txt") // edition 134, puzzle 27
	// sudoku.read("sudoku_level9_4_wip.txt") // edition 134, puzzle 27
	// sudoku.read("sudoku_level9_xwing.txt") // edition 134, techniques, x-wing
	// sudoku.read("sudoku_level9_ywing.txt") // edition 134, techniques, y-wing
	fmt.Println(sudoku)

	i := 0
	for {
		i++
		fmt.Printf("######### Iteration %d ##########\n", i)
		sudoku.print()
		if !sudoku.valid() {
			panic("Sudoku is not valid!")
		}
		sudoku.updateOptions()
		updated := sudoku.updateNumbers()
		sudoku.updateOptions() // todo?
		if !updated {
			sudoku.print()
			
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

func (sudoku *Sudoku) getCell(column, row int) *Cell {
	for i := 0; i < len(sudoku); i++ {
		cell := sudoku[i]
		if cell.column == column && cell.row == row {
			return &cell
		}
	}
	panic("ERR! Could not find cell!")
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

func (sudoku *Sudoku) updateOptions() {
	// When number != 0, set options to empty array
	for i := 0; i < len(sudoku); i++ {
		if sudoku[i].number != 0 {
			sudoku[i].options = []int{}
		}
	}

	// Update options for each row, column and square
	for i := 0; i < 9; i++ {
		row := sudoku.getRow(i)
		column := sudoku.getColumn(i)
		square := sudoku.getSquare(i)

		row.updateOptions()
		column.updateOptions()
		square.updateOptions()
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
					row[c2].removeOption(number)
				}
			}
			if len(cols) == 1 {
				col := sudoku.getColumn(cols[0])
				for c2 := 0; c2 < 9; c2++ {
					if (col[c2].square == i) {
						continue
					}
					col[c2].removeOption(number)
				}
			}
		}
	}

	// TODO: implement unique square
	// TODO: rename to unique rectangle
	tuples := []int{}
	for i := 0; i < len(sudoku); i++ {
		if len(sudoku[i].options) == 2 {
			tuples = append(tuples, i)
		}
	}
	fmt.Println("tuples unique square", tuples)
	for _, c1 := range tuples {
		for _, c2 := range tuples {
			if c2 < c1 {
				continue
			}
			for _, c3 := range tuples {
				if c3 < c2 {
					continue
				}
				if c1 == c2 || c2 == c3 || c1 == c3 {
					// skip if cells overlap
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
			}
		}
	}

	// TODO: implement X wing
	// TODO: implement XY wing
	// TODO: implement XYZ wing
	
}

func (cells Cell9Collection) updateOptions() {
	// When number != 0, remove the option from the other cells 
	for c1 := 0; c1 < len(cells); c1++ {
		for c2 := 0; c2 < len(cells); c2++ {
			numberC2 := (*cells[c2]).number 
			if c1 == c2 || numberC2 == 0 {
				continue
			}
			cells[c1].removeOption(numberC2)
		}
	}

	// When x cells have the same x options, remove the options from the other cells
	// For example, when cell 1 and 2 both have options 4 and 6, none of the other cells are allowed to have option 4 or 6
	for c1 := 0; c1 < 9; c1++ {
		options := (*cells[c1]).options
		if len(options) == 0 {
			continue
		}
		count := 0
		for c2 := 0; c2 < 9; c2++ {
			if equals(options, (*cells[c2]).options) {
				count++
			}
		}
		if len(options) != count {
			continue
		}
		for c3 := 0; c3 < 9; c3++ {
			if equals(options, (*cells[c3]).options) {
				continue
			}
			for _, number := range options {
				cells[c3].removeOption(number)
			}
		}
	}

	// TODO: refactor, consider merging with equals() check above
	// Check subsets
	for c1 := 0; c1 < 9; c1++ {
		options := (*cells[c1]).options
		if len(options) == 0 {
			continue
		}
		count := 0
		numbersFound := []int{}
		for c2 := 0; c2 < 9; c2++ {
			optionsC2 := (*cells[c2]).options
			if subset(options, optionsC2) {
				count++
				for _, o := range optionsC2 {
					if !contains(numbersFound, o) {
						numbersFound = append(numbersFound, o)
					}	
				}
			}
		}
		if len(options) != count || !equals(options, numbersFound) {
			continue;
		}
		for c3 := 0; c3 < 9; c3++ {
			if subset(options, (*cells[c3]).options) {
				continue
			}
			for _, number := range options {
				cells[c3].removeOption(number)
			}
		}
	}
}

func (sudoku *Sudoku) updateNumbers() bool {
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

// source: https://stackoverflow.com/a/10485970/3737152
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func (cell *Cell) removeOption(value int) {
	for i, option := range (*cell).options {
		if option == value {
			(*cell).options = append((*cell).options[:i], (*cell).options[i+1:]...)
			return
		}
	}
}

func (cell *Cell) removeOptions(options []int) {
	for _, option := range options {
		cell.removeOption(option)
	}
}