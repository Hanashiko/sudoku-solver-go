package main 

import (
	"fmt"
)

func print_grid(grid [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
}

func find_empty_location(grid [9][9]int, position *[2]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if (grid[row][column] == 0) {
				position[0] = row
				position[1] = column
				return true
			}
		}
	}
	return false
}

func used_in_row(grid [9][9]int, row int, number int) bool {
	for i := 0; i < 9; i++ {
		if (grid[row][i] == number) {
			return true
		}
	}
	return false
}

func used_in_column(grid [9][9]int, column int, number int) bool {
	for i := 0; i < 9; i++ {
		if (grid[i][column] == number) {
			return true
		}
	}
	return false
}

func used_in_box(grid [9][9]int, row int, column int, number int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (grid[i + row][j + column] == number) {
				return true
			}
		}
	}
	return false
}

func is_safe(grid [9][9]int, row int, column int, number int) bool {
	return (!used_in_row(grid, row, number) &&
		   (!used_in_column(grid, column, number) &&
		   (!used_in_box(grid, row - row % 3, 
						 column - column % 3, number))))
}
func solve(grid *[9][9]int) bool {
	var position [2]int = [2]int{0, 0}

	if (!find_empty_location(*grid, &position)) {
		return true
	}

	var row int = position[0]
	var column int = position[1]

	for number := 1; number < 10; number++ {
		if (is_safe(*grid, row, column, number)) {
			grid[row][column] = number
			if solve(grid) {
				return true
			}
			grid[row][column] = 0
		}
	}
	return false
}

func main() {
	var grid [9][9]int = [9][9]int{
        {3, 0, 6, 5, 0, 8, 4, 0, 0},
        {5, 2, 0, 0, 0, 0, 0, 0, 0},
        {0, 8, 7, 0, 0, 0, 0, 3, 1},
        {0, 0, 3, 0, 1, 0, 0, 8, 0},
        {9, 0, 0, 8, 6, 3, 0, 0, 5},
        {0, 5, 0, 0, 9, 0, 6, 0, 0},
        {1, 3, 0, 0, 0, 0, 2, 5, 0},
        {0, 0, 0, 0, 0, 0, 0, 7, 4},
        {0, 0, 5, 2, 0, 6, 3, 0, 0}}
	grid = [9][9]int{
        {0, 0, 4, 6, 7, 2, 0, 0, 0},
        {5, 0, 0, 8, 0, 0, 0, 9, 6},
        {0, 6, 3, 0, 4, 0, 0, 0, 8},
        {3, 8, 2, 1, 0, 0, 9, 6, 0},
        {4, 7, 5, 0, 0, 0, 1, 0, 0},
        {9, 1, 0, 2, 0, 4, 5, 0, 0},
        {0, 0, 0, 0, 0, 0, 0, 2, 9},
        {0, 0, 1, 0, 0, 0, 7, 4, 3},
        {2, 0, 0, 0, 6, 3, 8, 0, 1}}
	grid = [9][9]int{
        {0, 0, 0, 0, 0, 0, 0, 2, 3},
        {0, 0, 0, 0, 0, 0, 7, 5, 0},
        {0, 5, 0, 4, 8, 0, 0, 0, 0},
        {0, 0, 4, 0, 0, 9, 0, 0, 0},
        {1, 0, 0, 0, 6, 7, 0, 0, 2},
        {0, 6, 0, 0, 0, 0, 0, 8, 0},
        {0, 3, 0, 2, 0, 0, 0, 4, 0},
        {0, 4, 0, 1, 0, 0, 0, 0, 5},
        {8, 0, 0, 0, 0, 5, 6, 0, 0}}
	grid = [9][9]int{
		{0, 0, 3, 0, 4, 5, 0, 0, 0},
		{0, 5, 0, 0, 0, 0, 0, 0, 0},
		{0, 4, 0, 0, 6, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 3, 1, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 6, 0},
		{8, 0, 1, 9, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 0, 0, 0, 0, 0, 6, 5, 4},
		{3, 9, 0, 0, 0, 0, 0, 0, 2}}		
	if solve(&grid) {
		print_grid(grid)
	} else {
		fmt.Println("No solution exists")
	}
}