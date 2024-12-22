package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var verbose int
	var file_path string
	var task string
	var n_filled_cells int
	var save_path string

	flag.IntVar(&verbose, "v", 0, "verbosity level. value must be 0 or 1")
	flag.StringVar(
		&file_path,
		"file_path",
		"",
		"path to txt file containing sudoku puzzle. If not set it will promt you in the terminal",
	)
	flag.StringVar(
		&task,
		"task",
		"generate",
		"Which task to solve. Can be either 'generate' or 'solve'",
	)
	flag.IntVar(&n_filled_cells, "n_filled_cells", 30, "number of filled cells in generated sudoku puzzle")
	flag.StringVar(&save_path, "save_path", "", "save solution(s) or puzzle to this file")

	flag.Parse()

	if task == "solve" {
		var board sudoku
		if file_path == "" {
			board = read_sudoku_from_terminal()
		} else {
			board = read_sudoku_from_file(file_path)
		}

		if !check_sudoku_rules(board) {
			board_str := convert_sudoku_to_string(board)
			fmt.Println("Invalid board entered, got the board")
			fmt.Println(board_str)
			return
		}

		if verbose == 1 {
			board_str := convert_sudoku_to_string(board)
			fmt.Println("Got the following sudoku")
			fmt.Println(board_str)
		}

		solutions := solve_backtracking(board, true, verbose)

		if verbose == 1 && len(solutions) != 0 {
			// Move curser in terminal only if we need to overwrite
			// the previously printed boards
			fmt.Print("\033[14A")
		}

		if len(solutions) == 0 {
			fmt.Println("No solutions were found")
		} else if len(solutions) == 1 {
			var solution_str string = convert_sudoku_to_string(solutions[0])
			if save_path != "" {
				// write the whole body at once
				err := os.WriteFile(save_path, []byte(solution_str), 0644)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Unique solution with the following values")
					fmt.Println(solution_str)
				}
			}
		} else {
			fmt.Printf("%d solutions found. Heres the list.\n", len(solutions))
			for _, sol_board := range solutions {
				var solution_str string = convert_sudoku_to_string(sol_board)
				fmt.Println(solution_str)
			}
		}
	}

	if task == "generate" {
		board := generate_sudoku_puzzle(n_filled_cells)
		board_str := convert_sudoku_to_string(board)
		if save_path == "" {
			fmt.Println("generated the following sudoku")
			fmt.Println(board_str)
		} else {
			// write the whole body at once
			err := os.WriteFile(save_path, []byte(board_str), 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}
