package sudoku_tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sudoku [9][9]int

func parse_sudoku(board_str string) sudoku {
	board_str = strings.ReplaceAll(board_str, " ", "")
	board_str = strings.ReplaceAll(board_str, "\n", "")
	board_str = strings.ReplaceAll(board_str, "-", "")
	board_str = strings.ReplaceAll(board_str, "|", "")

	board_char_list := strings.Split(board_str, "")

	board := sudoku{}

	var board_i int = 0
	var board_j int = 0
	for char_i := 0; char_i < 81; char_i++ {
		val, _ := strconv.Atoi(board_char_list[char_i])
		board[board_i][board_j] = val
		board_i = board_i + board_j/8
		board_j = (board_j + 1) % 9
	}

	return board
}

func read_sudoku_from_terminal() sudoku {
	reader := bufio.NewReader(os.Stdin)

	input_contents := ""

	for i := 0; i < 9; i++ {
		text, _ := reader.ReadString('\n')
		input_contents = input_contents + text
	}

	return parse_sudoku(input_contents)
}

func read_sudoku_from_file(file_path string) sudoku {
	b, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Print(err)
		return sudoku{}
	}
	file_contents := string(b)

	return parse_sudoku(file_contents)
}

func convert_sudoku_to_string(board sudoku) string {
	var board_string string = ""

	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			board_string = board_string + strings.Repeat("-", 25) + "\n"
		}
		board_string = board_string + "| "
		for j := 0; j < 9; j++ {
			board_string = board_string + strconv.Itoa(board[i][j]) + " "
			if (j+1)%3 == 0 {
				board_string = board_string + "| "
			}
		}
		board_string = board_string + "\n"

	}
	board_string = board_string + strings.Repeat("-", 25) + "\n"
	return board_string
}
