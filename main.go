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
	// sudoku.read("sudoku1.txt")
	sudoku.read("sudoku2.txt")
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
		sudoku.updateNumbers()
		// if !updated {
		// 	sudoku.print()
		// 	panic("Couldn't find an option to move. Sudoku is too hard for the algorithm!")
		// }
		if sudoku.done() {
			break
		}
	}

	fmt.Println("######################")
	fmt.Println("Done! Sudoku is solved!")
	sudoku.print()
}

type Cell struct {
	id int
	x int
	y int
	square int
	number int
	options []int
}
type Sudoku [81]Cell
type Cell9Collection [9]*Cell

// TODO: consider reverse read (so 0,0 is left bottom)
func (sudoku *Sudoku) read(fileLocation string) {
	file, err := os.Open(fileLocation)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 0
	y := 0
    for scanner.Scan() {
		for x, v := range strings.Split(scanner.Text(), "") {
			number, err := strconv.Atoi(v)
			if err != nil {
				// ... handle error
				panic(err)
			}
			options := []int{1,2,3,4,5,6,7,8,9}
			if number != 0 {
				options = []int{}
			}
			square := (y/3)*3 + x/3 // (y/3)*3 != y, e.g. (4/3)*3=3
			sudoku[id] = Cell{
				id: id,
				x: x,
				y: y,
				square: square,
				number: number,
				options: options,	
			}
			id++
		}
		y++
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}



func (sudoku Sudoku) print() {
	fmt.Println("Sudoku:")

	for y := 0; y < 9; y++ {
		for _, cell := range sudoku.getCellsByY(y) {
			fmt.Printf("%d", (*cell).number)
		}
		fmt.Printf("\n")
	}
}

func (sudoku *Sudoku) getCellsByX(x int) Cell9Collection {
	cells := Cell9Collection{}
	i := 0
	for j := 0; j < len(sudoku); j++ {
		if (sudoku[j].x == x) {
			cells[i] = &sudoku[j]
			i++
		}
	}
	return cells
}

func (sudoku *Sudoku) getCellsByY(y int) Cell9Collection{
	cells := Cell9Collection{}
	i := 0
	for j := 0; j < len(sudoku); j++ {
		if (sudoku[j].y == y) {
			cells[i] = &sudoku[j]
			i++
		}
	}
	return cells
}

func (sudoku *Sudoku) getCellsBySquare(square int) Cell9Collection{
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

func (sudoku *Sudoku) getCell(x, y int) *Cell {
	for i := 0; i < len(sudoku); i++ {
		cell := sudoku[i]
		if cell.x == x && cell.y == y {
			return &cell
		}
	}
	panic("ERR! Could not find cell!")
}

func (sudoku Sudoku) valid() bool {
	for i := 0; i < 9; i++ {
		cellsY := sudoku.getCellsByY(i)
		cellsX := sudoku.getCellsByX(i)
		cellsSquare := sudoku.getCellsBySquare(i)

		if !cellsY.valid() || !cellsX.valid() || !cellsSquare.valid() {
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
			fmt.Printf("Not done: sudoku[%d][%d] == 0\n", sudoku[i].y, sudoku[i].x)
			return false
		}
	}
	return true
}

func (sudoku *Sudoku) updateOptions() {
	for i := 0; i < 9; i++ {
		cellsY := sudoku.getCellsByY(i)
		cellsX := sudoku.getCellsByX(i)
		cellsSquare := sudoku.getCellsBySquare(i)

		cellsY.updateOptions()
		cellsX.updateOptions()
		cellsSquare.updateOptions()
	}
}

func (cells Cell9Collection) updateOptions() {
	// remove filled in numbers from other cells options
	for c1 := 0; c1 < len(cells); c1++ {
		for c2 := 0; c2 < len(cells); c2++ {
			numberC2 := (*cells[c2]).number 
			if c1 == c2 || numberC2 == 0 {
				continue
			}
			cells[c1].removeOption(numberC2)
		}
	}

	// remove duplicate options
	count := map[string][]int{}
	for i := 0; i < 9; i++ {
		options := (*cells[i]).options
		if len(options) == 0 {
			continue
		}
		key := sliceToString(options, ",")
		count[key] = append(count[key], i)
	}
	for key, indexes := range count {
		options := stringToSlice(key, ",")
		if len(options) != len(indexes) || len(options) == 1 {
			continue
		}
		for _, number := range options {
			for c2 := 0; c2 < 9; c2++ {
				if contains(indexes, c2) {
					continue
				}
				cells[c2].removeOption(number)
			}
		}
	}

	// check missing options
	missingNumbers := []int{}
	for number := 1; number <= 9; number++ {
		found := false
		for i := 0; i < 9; i++ {
			if (*cells[i]).number == number {
				found = true
				break
			}
		}
		if !found {
			missingNumbers = append(missingNumbers, number)
		}
	}
	for i := 0; i < len(missingNumbers); i++ {
		missingNumber := missingNumbers[i]
		optionsCount := 0
		for j := 0; j < 9; j++ {
			if contains((*cells[j]).options, missingNumber) {
				optionsCount++
			}
		}
		if (optionsCount == 1) {
			for k := 0; k < 9; k++ {
				if contains((*cells[k]).options, missingNumber) {
					(*cells[k]).number = missingNumber
					(*cells[k]).options = []int{}
				}
			}
		}
	}
}

// TODO: Merge with updateOptions?
func (sudoku *Sudoku) updateNumbers() bool {
	updated := false
	for i := 0; i < len(sudoku); i++ {
		if len(sudoku[i].options) == 1 {
			sudoku[i].number = sudoku[i].options[0]
			sudoku[i].options = []int{}
			updated = true
		}
	}
	return updated
}

func sliceToString(slice []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), delim), "[]")
}

func stringToSlice(input string, delim string) []int {
	slice := []int{}
	for _, char := range strings.Split(input, delim) {
		value, err := strconv.Atoi(char)
		if err != nil {
			// ... handle error
			panic(err)
		}
		slice = append(slice, value)
	}
	return slice
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