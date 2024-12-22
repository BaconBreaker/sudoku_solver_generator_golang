package sudoku_tools

import (
	"fmt"
	"math/rand"
)

// Helper function that counts the times a specific
// value shows up in array.
func count_values[K comparable](cells []K, value K) int {
	count := 0
	for _, elm := range cells {
		if elm == value {
			count = count + 1
		}
	}
	return count
}

// Helper function that extracts a column from a sudoku type (2D array)
func get_column(board sudoku, column_index int) []int {
	var column = make([]int, 0)
	for _, row := range board {
		column = append(column, row[column_index])
	}
	return column
}

// Helper function that extracts boxes from a sudoku type (2D array)
// given the index of said box (e.g. (0,0) is upper left and (1,1) is center)
// Values are given left-ro-right in reading order.
func get_box(board sudoku, row_index int, col_index int) []int {
	var box = make([]int, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			box = append(box, board[row_index*3+i][col_index*3+j])
		}
	}
	return box
}

// Checks if value in cell doesn't break sudoku rules
func check_cell(board sudoku, row_index int, col_index int) bool {
	var value int = board[row_index][col_index]
	var box [9]int = [9]int(get_box(board, row_index/3, col_index/3))
	var column [9]int = [9]int(get_column(board, col_index))
	var row [9]int = board[row_index]

	if count_values(row[:], value) >= 2 {
		return false
	}
	if count_values(box[:], value) >= 2 {
		return false
	}
	if count_values(column[:], value) >= 2 {
		return false
	}

	return true
}

// Checks if board satisfies standard sudoku rules
// Just checks if every non-zero value is legal
func check_sudoku_rules(board sudoku) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val != 0 {
				if !check_cell(board, i, j) {
					return false
				}
			}
		}
	}
	return true
}

// Solves a single cell by recursing through every possible value in said cell.
func solve_backtracking_step(board sudoku, row_index int, column_index int, solution_list []sudoku, verbose int, return_first bool) []sudoku {
	var new_row_index int
	var new_column_index int
	var new_solutions []sudoku = []sudoku{}

	// Check if we are at the end
	if row_index == 9 && column_index == 0 {
		return []sudoku{board}
	}

	// Check if cell is filled, if so, continue to next cell
	if board[row_index][column_index] != 0 {
		new_row_index = row_index + column_index/8
		new_column_index = (column_index + 1) % 9
		// We don't have to combine with current solution list since it is always empty
		return solve_backtracking_step(board, new_row_index, new_column_index, []sudoku{}, verbose, return_first)
	}

	// Try all possible values in this cell and see if it leads to a solution
	// The order in which the values are tested is random to allow for the solver
	// to be used to generate random sudoku puzzles.
	var new_board sudoku = board

	candidate_values := rand.Perm(9)

	for _, candidate_value := range candidate_values {
		new_board[row_index][column_index] = candidate_value + 1 //candidate value is in range [0, 8]

		if verbose == 1 {
			new_board_str := convert_sudoku_to_string(new_board)
			fmt.Printf("\033[13A%s", new_board_str)
		}

		if check_cell(new_board, row_index, column_index) {
			new_row_index = row_index + column_index/8
			new_column_index = (column_index + 1) % 9
			if return_first {
				solutions := solve_backtracking_step(new_board, new_row_index, new_column_index, []sudoku{}, verbose, return_first)
				if len(solutions) != 0 {
					return solutions
				}
			} else {
				solutions := solve_backtracking_step(new_board, new_row_index, new_column_index, []sudoku{}, verbose, return_first)
				new_solutions = append(new_solutions, solutions...)
			}
		}
	}

	return append(solution_list, new_solutions...)
}

// Entry point for recursive backtracking
func solve_backtracking(board sudoku, return_all bool, verbose int) []sudoku {
	if verbose == 1 {
		// Print first iteration of the board
		fmt.Println("Attempting to solve")
		board_str := convert_sudoku_to_string(board)
		fmt.Print(board_str)
	}
	return solve_backtracking_step(board, 0, 0, []sudoku{}, verbose, !return_all)
}
