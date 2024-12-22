package sudoku_tools

import "math/rand"

// Removes element at index i from slice
func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func generate_sudoku_puzzle(n_filled_cells int) sudoku {
	board_empty := sudoku{}                                //init empty board
	solutions := solve_backtracking(board_empty, false, 0) //generate random solved board
	board := solutions[0]
	var list_of_filled_cells []int = make([]int, 81)
	for i := 0; i < 81; i++ {
		list_of_filled_cells[i] = i
	}

	for n_filled_cells < len(list_of_filled_cells) {
		ind := rand.Intn(len(list_of_filled_cells))
		ind_of_elm := list_of_filled_cells[ind]
		ind_i := ind_of_elm / 9
		ind_j := ind_of_elm % 9
		new_board := sudoku{}
		copy(new_board[:], board[:])
		new_board[ind_i][ind_j] = 0
		new_solutions := solve_backtracking(new_board, true, 0)
		if len(new_solutions) == 1 {
			board = new_board
			list_of_filled_cells = remove(list_of_filled_cells, ind)
		}
	}

	return board
}

// Difficulties:
// Extreme: 23 filled cells
// Master: 25 filled cells
// Expert: 30 filled cells
// Hard: 30 filled cells?
// Medium: 38 filled cells
// Easy: 38 filled cells
