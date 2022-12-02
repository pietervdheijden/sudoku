package sudokuV1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pietervdheijden/sudoku/services"
)

type sudokuApiModel struct {
	Sudoku [81]int `json:"sudoku" binding:"required"`
}

func toApiModel(sudoku services.Sudoku) sudokuApiModel {
	sudokuApiModel := sudokuApiModel{};
	for cell := 0; cell < 81; cell++ {
		sudokuApiModel.Sudoku[cell] = sudoku[cell].Number;
	}
	return sudokuApiModel
}

func (input sudokuApiModel) toServiceModel() services.Sudoku {
	fmt.Printf("Read Sudoku from array: %v\n", input)
	sudoku := services.Sudoku{}
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			id := row * 9 + column
			number := input.Sudoku[id]
			options := []int{1,2,3,4,5,6,7,8,9}
			if number != 0 {
				options = []int{}
			}
			square := (row/3)*3 + column/3 // (x/3)*3 != x, e.g. (4/3)*3=3
			sudoku[id] = services.Cell{
				Id: id,
				Row: row,
				Column: column,
				Square: square,
				Number: number,
				Options: options,	
			}
		}
	}
	return sudoku
}

func Solve(c *gin.Context) {
	var sudokuApiModel sudokuApiModel

	if err := c.BindJSON(&sudokuApiModel); err != nil {
		fmt.Println(err)
		return
	}

	sudoku := sudokuApiModel.toServiceModel();
	solved := sudoku.Solve()

	if solved {
		c.JSON(http.StatusOK, toApiModel(sudoku))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Sudoku could not be solved!"})
	}
}

func Hint(c *gin.Context) {
	// TODO: implement endpoint
	// Hint which cell could be solved next (incl. technique and related cells)

	c.JSON(http.StatusOK, gin.H{"task": "hint"})	
}

func Check(c *gin.Context) {
	// TODO: implement endpoint
	// Check if Sudoku is solvable.

	c.JSON(http.StatusOK, gin.H{"task": "check"})
}

func Options(c *gin.Context) {
	// TODO: implement endpoint
	// Show the options (/candidates) foreach cell.

	c.JSON(http.StatusOK, gin.H{"task": "options"})
}