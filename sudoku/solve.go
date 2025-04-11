package sudoku

import (
	"errors"
)

func Solve(gridString string) ([]string, error) {
	grid, err := parse(gridString)
	if err != nil {
		return nil, err
	}

	if !grid.valid() {
		return nil, errors.New("sudoku not valid")
	}

	solutions := grid.backtrack(0)

	solutionStrings := make([]string, len(solutions))
	for i, solution := range solutions {
		if err := grid.solutionValid(solution); err != nil {
			return nil, err
		}
		solutionStrings[i] = solution.string()
	}

	return solutionStrings, nil
}
