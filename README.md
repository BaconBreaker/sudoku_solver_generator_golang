# sudoku_solver_generator_golang

## Solver
Example of running solver with visualizatoin
```console
go run . -task "solve" -v 1 -file_path examples/sudoku_difficult.txt
```
which should give ytou something like
![Alt Text](resources/solve.gif)

## Generator
Example of running generator
```console
go run . -task "generate" -n_filled_cells 25
```
