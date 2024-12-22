# sudoku_solver_generator_golang

This was my first time coding in go, it does not utilize any goroutines but i still learned the go syntax and types, as well as testing in go.

## Solver
Example of running solver with visualizatoin
```console
go run . -task "solve" -v 1 -file_path examples/sudoku_difficult.txt
```
which should give ytou something like

<img src="resources/solve.gif" width="250" height="300" />

## Generator
Example of running generator
```console
go run . -task "generate" -n_filled_cells 25
```
