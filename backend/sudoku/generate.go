package sudoku

import (
	"context"
	"errors"
	"math/rand/v2"
	"time"
)



func createFilledGrid(ctx context.Context, rand *rand.Rand) grid {
    var grid grid = make([]uint8, 81)
    var i uint8 = 0
    for i < 81 {
        values := rand.Perm(9)
        for _, v := range values {
            value := uint8(v + 1)
            if !grid.moveValid(i, value) {
                continue
            }
            grid[i] = value
            solutions := grid.backtrack(ctx, 1)
            if solutions == nil { // timeout
                return nil
            }
            if len(solutions) != 0 {
                break
            }
            grid[i] = 0
        }
        i++
    }
    return grid
}

func generate(ctx context.Context, rand *rand.Rand, unknowns uint8) grid {
    grid := createFilledGrid(ctx, rand)
    if grid == nil {
        return nil // timeout
    }
    for i, gridIndex := range rand.Perm(81) {
        if uint8(i) >= unknowns {
            break
        }
        removed := grid[gridIndex]
        grid[gridIndex] = 0
        solutions := grid.backtrack(ctx, 2)
        if solutions == nil {
            return nil // timeout
        }
        if len(solutions) > 1 {
            // more than 1 solution --> put back, go to next index
            grid[gridIndex] = removed
        }
    }
    return grid
}

func Generate(rand *rand.Rand, unknowns uint8) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
    grid := generate(ctx, rand, unknowns)
    if grid == nil {
        return "", errors.New("generate timed out")
    }
    return grid.string(), nil
}
