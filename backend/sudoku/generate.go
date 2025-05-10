package sudoku

import "math/rand/v2"



func createFilledGrid(rand *rand.Rand) grid {
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
            if len(grid.backtrack(1)) != 0 {
                break
            }
            grid[i] = 0
        }
        i++
    }
    return grid
}

func generate(rand *rand.Rand, unknowns uint8) grid {
    grid := createFilledGrid(rand)
    for i, gridIndex := range rand.Perm(81) {
        if uint8(i) >= unknowns {
            break
        }
        removed := grid[gridIndex]
        grid[gridIndex] = 0
        if len(grid.backtrack(2)) > 1 {
            // more than 1 solution --> put back, go to next index
            grid[gridIndex] = removed
        }
    }
    return grid
}

func Generate(rand *rand.Rand, unknowns uint8) string {
    return generate(rand, unknowns).string()
}
