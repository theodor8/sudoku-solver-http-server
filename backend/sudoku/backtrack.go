package sudoku

import (
    "context"
    "slices"
)

// 0 for all, 1 for 1, ...
func (g grid) backtrack(ctx context.Context, numSolutions uint8) []grid {
    gr := slices.Clone(g)
    unknowns := make([]uint8, 0, 81) // indices of unknowns
    for i, v := range gr {
        if v == 0 {
            unknowns = append(unknowns, uint8(i))
        }
    }
    if len(unknowns) == 0 {
        return []grid{slices.Clone(gr)}
    }
    var solutions []grid = make([]grid, 0, 1)
    var unknownsIndex uint8 = 0
    var gridIndex uint8 = unknowns[unknownsIndex]
    for {
        select {
        case <-ctx.Done(): 
            return nil
        default:
        }
        foundValidTry := false
        for try := gr[gridIndex] + 1; try <= 9; try++ {
            if gr.moveValid(gridIndex, try) {
                gr[gridIndex] = try
                foundValidTry = true
                break
            }
        }
        if foundValidTry {
            if unknownsIndex < uint8(len(unknowns)) - 1 {
                unknownsIndex++
                gridIndex = unknowns[unknownsIndex]
            } else {
                solutions = append(solutions, slices.Clone(gr))
                if numSolutions != 0 && uint8(len(solutions)) == numSolutions {
                    break
                }
            }
        } else {
            gr[gridIndex] = 0
            if unknownsIndex == 0 {
                break
            }
            unknownsIndex--
            gridIndex = unknowns[unknownsIndex]
        }
    }
    return solutions
}
