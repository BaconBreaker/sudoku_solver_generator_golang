package main

import "testing"

// Note that I am lazy with writing tests. I only check that certain puzzles
// produce the right amount of solutions and that the values are legal using
// an untested function. ¯\_(ツ)_/¯
// I am not testing all subroutines or the correctness of the puzzle solutions.

// Easy puzzle with unique solution
func TestSudokuEasy(t *testing.T) {
	board := read_sudoku_from_file("./examples/sudoku_easy.txt")
	solution_list := solve_backtracking(board, true, 0)
	if len(solution_list) != 1 {
		t.Fatalf("len(solution_list) = %d when it should be 1", len(solution_list))
	}
	if !check_sudoku_rules(board) {
		t.Fatal("Values in solution break sudoku rules")
	}
}

func TestSudokuDifficult(t *testing.T) {
	board := read_sudoku_from_file("./examples/sudoku_difficult.txt")
	solution_list := solve_backtracking(board, true, 0)
	if len(solution_list) != 1 {
		t.Fatalf("len(solution_list) = %d when it should be 1", len(solution_list))
	}
	if !check_sudoku_rules(board) {
		t.Fatal("Values in solution break sudoku rules")
	}
}

func TestSudokuAlmostSolved(t *testing.T) {
	board := read_sudoku_from_file("./examples/sudoku_almost_solved.txt")
	solution_list := solve_backtracking(board, true, 0)
	if len(solution_list) != 1 {
		t.Fatalf("len(solution_list) = %d when it should be 1", len(solution_list))
	}
	if !check_sudoku_rules(board) {
		t.Fatal("Values in solution break sudoku rules")
	}
}

func TestSudoku2Solutions(t *testing.T) {
	board := read_sudoku_from_file("./examples/sudoku_2_solutions.txt")
	solution_list := solve_backtracking(board, true, 0)
	if len(solution_list) != 2 {
		t.Fatalf("len(solution_list) = %d when it should be 2", len(solution_list))
	}
	if !check_sudoku_rules(board) {
		t.Fatal("Values in solution break sudoku rules")
	}
}

func TestSudoku3Solutions(t *testing.T) {
	board := read_sudoku_from_file("./examples/sudoku_3_solutions.txt")
	solution_list := solve_backtracking(board, true, 0)
	if len(solution_list) != 3 {
		t.Fatalf("len(solution_list) = %d when it should be 3", len(solution_list))
	}
	if !check_sudoku_rules(board) {
		t.Fatal("Values in solution break sudoku rules")
	}
}

func TestSudokuNoSolutions(t *testing.T) {
	board := read_sudoku_from_file("./examples/sudoku_no_solution.txt")
	solution_list := solve_backtracking(board, true, 0)
	if len(solution_list) != 0 {
		t.Fatalf("len(solution_list) = %d when it should be 0", len(solution_list))
	}
	if !check_sudoku_rules(board) {
		t.Fatal("Values in solution break sudoku rules")
	}
}
