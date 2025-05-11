package sudoku

import (
	"context"
	"errors"
	"time"
)

func Solve(gridString string) ([]string, error) {
	grid, err := parse(gridString)
	if err != nil {
		return nil, err
	}

	if !grid.valid() {
		return nil, errors.New("sudoku not valid")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	solutions := grid.backtrack(ctx, 0)
	if solutions == nil {
		return nil, errors.New("solve timed out")
	}

	solutionStrings := make([]string, len(solutions))
	for i, solution := range solutions {
		if err := grid.solutionValid(solution); err != nil {
			return nil, err
		}
		solutionStrings[i] = solution.string()
	}

	return solutionStrings, nil
}
