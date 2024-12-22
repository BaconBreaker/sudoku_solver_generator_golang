package main

import "testing"

// I know test coverage isn't exctensive. But i just can't be bothered
// to write a full test suite on a for-fun project.

// tests that every example in examples can be parsed, printed, and parsed again.
func TestSudokuParsePrint(t *testing.T) {
	var file_paths []string = []string{
		"examples/sudoku_2_solutions.txt",
		"examples/sudoku_3_solutions.txt",
		"examples/sudoku_almost_solved.txt",
		"examples/sudoku_difficult.txt",
		"examples/sudoku_easy.txt",
		"examples/sudoku_no_solution.txt",
	}
	for _, f_path := range file_paths {
		board := read_sudoku_from_file(f_path)
		board_str := convert_sudoku_to_string(board)
		board2 := parse_sudoku(board_str)
		if board != board2 {
			t.Fatalf("example %s is parsed to board_str which doesn't match original input", f_path)
		}
	}
}
