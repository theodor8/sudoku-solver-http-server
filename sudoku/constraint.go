package sudoku

import "slices"


// -1: no possible,  0: more than 1 possible, 1-9: one possible
func onlyOnePossible(possible []bool) int {
    var possibleNum int = -1
    for i := range possible {
        if possible[i] {
            if possibleNum != -1 {
                return 0
            }
            possibleNum = i + 1
        }
    }
    return possibleNum
}


func (g grid) constraint() []grid {
    g = slices.Clone(g)

    possible := make([][]bool, 81)
    for i := range possible {
        possible[i] = make([]bool, 9)
        for j := range possible[i] {
            possible[i][j] = true
        }
    }

    for cell := range 81 {
        if g[cell] == 0 {
            continue
        }

        col := cell % 9
        row := cell / 9
        box := (row/3)*3*9 + col/3

        // set not possible
        for i := range 9 {
            // col
            possible[col+i*9][g[cell]-1] = false

            // row
            possible[row*9+i][g[cell]-1] = false

            // box
            possible[box+i%3*9+i/3][g[cell]-1] = false
        }
    }

    for cell := range 81 {

        possibleNum := onlyOnePossible(possible[cell])
        if possibleNum == -1 {
            // unsolvable (no possible)
            // TODO: constraint unsolvable
        }
        if possibleNum == 0 {
            // more than 1 possible
            continue
        }

        g[cell] = uint8(possibleNum)
    }




    return nil
}
