package sudoku_tools

import "testing"

// I again choose to be lazy and just test my generator quicktest style.

// Generate 20 sudokus and verify that they have unique solution
func TestSudokuGenerator(t *testing.T) {
	for i := 0; i < 20; i++ {
		board := generate_sudoku_puzzle(30)
		solution_list := solve_backtracking(board, true, 0)
		if len(solution_list) != 1 {
			t.Fatalf("number of solutions is %d when it should be 1", len(solution_list))
		}
		if !check_sudoku_rules(board) {
			t.Fatal("Values in solution break sudoku rules")
		}
	}

}
