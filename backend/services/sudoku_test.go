package services

import (
	"strconv"
	"strings"
	"testing"
)

// TODO: create more granular tests per technique

func TestSolve_Level1(t *testing.T) {
	testSolve(t, "000920000040851000256003091100085409098730162000200530007060900900002680080090054")
}

// edition 225, puzzle 7
func TestSolve_Level6_1(t *testing.T) {
	testSolve(t, "900041000800620001000005006050307000004000060781000500003200007000000400072000000")
}

// edition 225, puzzle 75
func TestSolve_Level6_2(t *testing.T) {
	testSolve(t, "005200000600810500802503010004080051000007000000004360000090020008000040329000000")
}

// edition 225, puzzle 77
func TestSolve_Level6_3(t *testing.T) {
	testSolve(t, "005390070060800000000602500007000061630700000940000000000020900003010080800503200")
}

// edition 225, puzzle 79
func TestSolve_Level6_4(t *testing.T) {
	testSolve(t, "000008020070000408500000000200000054600050000900010600001803000869000007000004090")
}

// edition 225, puzzle 83
func TestSolve_Level6_5(t *testing.T) {
	testSolve(t, "200370000500000000000001260009017030010000050600200080080000070005000024093060500")
}

// edition 134, puzzle 26
func TestSolve_Level8_1(t *testing.T) {
	testSolve(t, "001000200009006081870021000010009000750000032000500010000930046630200100007000300")
}

// edition 134, puzzle 28
func TestSolve_Level8_2(t *testing.T) {
	testSolve(t, "000000900004090756007000204203104090000708000070902308905000400632070800008000000")
}

// edition 134, puzzle 30
func TestSolve_Level8_3(t *testing.T) {
	testSolve(t, "000082005500000070000310002000007540760000098083100000300041000090000004800670000")
}

// edition 134, puzzle 17 (contradiction / forcing chain)
func TestSolve_Level9_1(t *testing.T) {
	testSolve(t, "007300605060000000908006001004900702000010000102005300200600504000000030305009100")
}

// edition 134, puzzle 25 (contradiction)
func TestSolve_Level9_2(t *testing.T) {
	testSolve(t, "000006090001070004005021070050004018004000300180500060040130600500080900090200000")
}

// edition 134, puzzle 27 (XY wing + XYZ wing)
func TestSolve_Level9_3(t *testing.T) {
	testSolve(t, "023090008901004000000007001008000039000020000490000100600500000000600704800030250")
}

// edition 134, puzzle 29 (contradiction)
func TestSolve_Level9_4(t *testing.T) {
	testSolve(t, "126000790004000001000003000908060070000307000060010802000200000800000900092000416")
}

// edition 134, puzzle 31 (X wing)
func TestSolve_Level9_5(t *testing.T) {
	testSolve(t, "000208000003050060900007080200003009085000210700500004070300001040090300000402000")
}

// edition 134, puzzle 33 (contradiction)
func TestSolve_Level9_6(t *testing.T) {
	testSolve(t, "000006004900407500800900031015000003000060000400000120530002008004603005200800000")
}

// edition 104, puzzle 9
func TestSolve_Level9_7(t *testing.T) {
	testSolve(t, "503100002100043900000000700060009020005070600080400050009000000002380006800002304")
}

// edition 104, puzzle 11
func TestSolve_Level9_8(t *testing.T) {
	testSolve(t, "009700510410000000600093000800006700500402001006300002000830004000000075034005600")
}

// edition 104, puzzle 13 (contradiction)
func TestSolve_Level9_9(t *testing.T) {
	testSolve(t, "000489050040100200000000078007024903000000000604590800780000000003008090060372000")
}

// edition 104, puzzle 15
func TestSolve_Level9_10(t *testing.T) {
	testSolve(t, "000010060000976005000300140573000900800000007004000283048007000600598000050020000")
}

// edition 104, puzzle 17
func TestSolve_Level9_11(t *testing.T) {
	testSolve(t, "000090007000200004002340190008003001090070050600500400075024800400001000300050000")
}

// edition 104, puzzle 19 (contradiction)
func TestSolve_Level9_12(t *testing.T) {
	testSolve(t, "000000040024000900080107630000071000097605480000480000013208050006000810070000000")
}

func testSolve(t *testing.T, puzzle string) {
	sudoku := readSudoku("000000040024000900080107630000071000097605480000480000013208050006000810070000000")
	solved := sudoku.Solve()
	if ! solved {
		t.Fatal("Sudoku could not be solved!")
	}
}

func readSudoku(input string) Sudoku {
	// Convert input from string to []int
	strs := strings.Split(input, "")
	cells := make([]int, len(strs))
	for i := range cells {
        cells[i], _ = strconv.Atoi(strs[i])
    }    

	// Parse sudoku
	sudoku := Sudoku{}
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			id := row * 9 + column
			number := cells[id]
			options := []int{1,2,3,4,5,6,7,8,9}
			if number != 0 {
				options = []int{}
			}
			square := (row/3)*3 + column/3 // (x/3)*3 != x, e.g. (4/3)*3=3
			sudoku[id] = Cell{
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
