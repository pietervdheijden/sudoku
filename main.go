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

	// if []int{1,2} == []int{1,2} {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }

	fmt.Println("Let's solve this Sudoku!")

	fmt.Println("First, read the Sudoku from file")
	sudoku := read("sudoku2.txt")
	fmt.Println(sudoku)

	i := 0
	for {
		i++
		fmt.Printf("######### Iteration %d ##########\n", i)
		sudoku.print()

		if !sudoku.valid() {
			panic("Sudoku is not valid!")
		}

		options := sudoku.calculateOptions()
		fmt.Println("Options:")
		for _, option := range options {
			fmt.Println(option)
		}

		// move to func? e.g. solve()?
		moved := false
		for row := 0; row < len(sudoku); row++ {
			for col := 0; col < len(sudoku[row]); col++ {
				if len(options[row][col]) == 1 {
					sudoku[row][col] = options[row][col][0]
					moved = true
				}
			}
		}

		if !moved {
			sudoku.print()
			panic("Couldn't find an option to move. Sudoku is too hard for the algorithm!")
		}


		if sudoku.done() {
			break
		}
		// break;
	}

	fmt.Println("######################")
	fmt.Println("Done! Sudoku is solved!")
	sudoku.print()
}

type Sudoku [9][9]int

func read(fileLocation string) Sudoku {
	file, err := os.Open(fileLocation)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	sudoku := Sudoku{}
    scanner := bufio.NewScanner(file)
	rowIdx := 0
    for scanner.Scan() {
		row := [9]int{}
		for i, v := range strings.Split(scanner.Text(), "") {
			row[i], err = strconv.Atoi(v)
			if err != nil {
				// ... handle error
				panic(err)
			}
		}
		sudoku[rowIdx] = row
		rowIdx++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return sudoku
}



func (sudoku Sudoku) print() {
	fmt.Println("Sudoku:")

	for row := 0; row < len(sudoku); row++ {
		// fmt.Printf("\t")
		for column := 0; column < len(sudoku[row]); column++ {
			fmt.Printf("%d", sudoku[row][column])
		}
		fmt.Printf("\n")
	}
}

func (sudoku Sudoku) valid() bool {
	// Check rows
	for row := 0; row < len(sudoku); row++ {
		numbers := []int{}
		for col := 0; col < len(sudoku[row]); col++ {
			number := sudoku[row][col]
			if contains(numbers, number) {
				return false
			}
			if number != 0 {
				numbers = append(numbers, sudoku[row][col])
			}
		}
	}

	// Check columns
	for col := 0; col < len(sudoku); col++ {
		numbers := []int{}
		for row := 0; row < len(sudoku[col]); row++ {
			number := sudoku[row][col]
			if contains(numbers, number) {
				return false
			}
			if number != 0 {
				numbers = append(numbers, sudoku[row][col])
			}
		}
	}

	// Check squares
	squares := sudoku.getSquares()
	for _, row := range squares {
		for _, col := range row {
			numbers := []int{}
			for _, number := range col {
				if contains(numbers, number) {
					return false
				}
				if number != 0 {
					numbers = append(numbers, number)
				}
			}
		}
	}

	return true
}

func (sudoku Sudoku) getSquares() [3][3][]int {
	squares := [3][3][]int{}
	for row := 0; row < len(sudoku); row++ {
		for col := 0; col < len(sudoku[row]); col++ {
			squares[row / 3][col / 3] = append(squares[row / 3][col / 3], sudoku[row][col])
		}
	}
	return squares
}

func (sudoku Sudoku) getSquare(row, col int) []int {
	return sudoku.getSquares()[row / 3][col / 3]
}

func (sudoku Sudoku) done() bool {
	for row := 0; row < len(sudoku); row++ {
		for column := 0; column < len(sudoku[row]); column++ {
			if sudoku[row][column] == 0 {
				fmt.Printf("Not done: sudoku[%d][%d] == 0\n", row, column)
				return false
			}
		}
	}
	return true
}

func (sudoku Sudoku) calculateOptions() [9][9][]int {
	allOptions := [9][9][]int{}
	for row := 0; row < len(sudoku); row++ {
		for column := 0; column < len(sudoku[row]); column++ {
			if sudoku[row][column] != 0 {
				continue
			}

			// Init options
			// Start with all options, and then eliminate.
			options := []int{1,2,3,4,5,6,7,8,9}

			// Check row
			for otherColumn := 0; otherColumn < len(sudoku[row]); otherColumn++ {
				remove(&options, sudoku[row][otherColumn])
			}

			// Check column
			for otherRow := 0; otherRow < len(sudoku); otherRow++ {
				remove(&options, sudoku[otherRow][column])
			}

			// Check square
			square := sudoku.getSquare(row, column)
			for _, number := range square {
				remove(&options, number)
			}

			allOptions[row][column] = options
		}
	}

	// Check duplicate tuples - rows
	for row := 0; row < len(allOptions); row++ {
		// Group indexes by option
		count := map[string][]int{}
		for col := 0; col < len(allOptions[row]); col++ {
			if len(allOptions[row][col]) == 0 {
				continue
			}
			key := sliceToString(allOptions[row][col], ",")
			count[key] = append(count[key], col)
		}
		for option, indexes := range count {
			slice := stringToSlice(option, ",")
			if len(slice) != len(indexes) || len(slice) == 1 {
				continue
			}
			for _, number := range slice {
				for col := 0; col < len(allOptions[row]); col++ {
					if contains(indexes, col) {
						continue
					}
					remove(&allOptions[row][col], number)
				}
			}
		}
	}

	// search duplicate options - cols
	for col := 0; col < len(allOptions); col++ {
		// Group indexes by option
		count := map[string][]int{}
		for row := 0; row < len(allOptions[col]); row++ {
			if len(allOptions[row][col]) == 0 {
				continue
			}
			key := sliceToString(allOptions[row][col], ",")
			count[key] = append(count[key], row)
		}
		for option, indexes := range count {
			slice := stringToSlice(option, ",")
			if len(slice) != len(indexes) || len(slice) == 1 {
				continue
			}
			for _, number := range slice {
				for row := 0; row < len(allOptions[col]); row++ {
					if contains(indexes, row) {
						continue
					}
					remove(&allOptions[row][col], number)
				}
			}
		}
	}

	// search duplicate options - squares
	// TODO
	// for _, square := range sudoku.getSquares() {
	// 	options := []int{}
	// 	for row := 0; row < len(square); row++ {
	// 		for col := 0; col < len(square[row]); col++ {
	// 			options = append(options, square[row][col])
	// 		}
	// 	}

	// 	count := map[string][]int{}
	// 	for i := 0; i < len(options[col]); row++ {
	// 		if len(allOptions[row][col]) == 0 {
	// 			continue
	// 		}
	// 		key := sliceToString(allOptions[row][col], ",")
	// 		count[key] = append(count[key], row)
	// 	}
	// }


	return allOptions
}

// func (slice []int) toString2() string {
// 	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), delim), "[]")
// }

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

func remove(slice *[]int, value int) {
	for i, v := range *slice {
		if v == value {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return
		}
	}
}

// source: https://stackoverflow.com/a/37335777/3737152
// func remove(slice []int, s int) []int {
//     return append(slice[:s], slice[s+1:]...)
// }