package solver

import (
    "errors"
    "slices"
)


type grid []uint8


func parse(str string) (grid, error) {
    if (len(str) != 81) {
        return nil, errors.New("invalid input length")
    }
    var b grid = make([]uint8, 81)
    for i, v := range str {
        if v < '0' || v > '9' {
            return nil, errors.New("invalid char")
        }
        b[i] = uint8(v - '0')
    }
    return b, nil
}

func itorc(i uint8) (uint8, uint8) {
    return i / 9, i % 9
}

func rctoi(r, c uint8) uint8 {
    return r * 9 + c
}


func (g grid) rowContains(r, v uint8) bool {
    return slices.Contains(g[rctoi(r, 0) : rctoi(r, 9)], v)
}

func (g grid) colContains(c, v uint8) bool {
    for ; c < 81; c += 9 {
        if g[c] == v {
            return true
        }
    }
    return false
}

func (g grid) boxContains(i, v uint8) bool {
    bi := i / 3 * 3 * 9 + i % 3 * 3
    for j := range uint8(9) {
        if g[bi + j / 3 * 9 + j % 3] == v {
            return true
        }
    }
    return false
}


func (g grid) moveValid(i, v uint8) bool {
    r, c := itorc(i)
    bi := r / 3 * 3 + c / 3
    return !g.rowContains(r, v) && !g.colContains(c, v) && !g.boxContains(bi, v)
}

func (g grid) getRow(r uint8) []uint8 {
    return slices.Clone(g[rctoi(r, 0) : rctoi(r, 9)])
}

func (g grid) getCol(c uint8) []uint8 {
    col := make([]uint8, 9)
    for i := range uint8(9) {
        col[i] = g[rctoi(i, c)]
    }
    return col
}

func (g grid) getBox(i uint8) []uint8 {
    box := make([]uint8, 9)
    r := i / 3 * 3
    c := i % 3 * 3
    for i := range uint8(9) {
        box[i] = g[rctoi(r + i / 3, c + i % 3)]
    }
    return box
}

func spliceValid(s []uint8) bool {
    has := make([]bool, 9)
    for _, v := range s {
        if v == 0 {
            continue
        }
        if has[v - 1] {
            return false
        }
        has[v - 1] = true
    }
    return true
}
func (g grid) valid() bool {
    for i := range uint8(9) {
        if !spliceValid(g.getRow(i)) ||
           !spliceValid(g.getCol(i)) ||
           !spliceValid(g.getBox(i)) {
            return false
        }
    }
    return true
}


func (g grid) backtrack() (uint, error) {
    unknowns := make([]uint8, 0, 81) // indices of unknowns
    for i, v := range g {
        if v == 0 {
            unknowns = append(unknowns, uint8(i))
        }
    }
    var cycles uint = 0
    var unknownsIndex uint8 = 0
    var gridIndex uint8 = unknowns[unknownsIndex]
    for {
        foundValidTry := false
        for try := g[gridIndex] + 1; try <= 9; try++ {
            if g.moveValid(gridIndex, try) {
                g[gridIndex] = try
                foundValidTry = true
                break
            }
        }
        if foundValidTry {
            unknownsIndex++
            if unknownsIndex >= uint8(len(unknowns)) {
                break
            }
            gridIndex = unknowns[unknownsIndex]
        } else {
            if unknownsIndex == 0 {
                return cycles, errors.New("unsolvable")
            }
            g[gridIndex] = 0
            unknownsIndex--
            gridIndex = unknowns[unknownsIndex]
        }
        cycles++
    }
    return cycles, nil
}

func (g grid) stringFormatted() string {
    var str string
    for i, v := range g {
        str += string(v + '0')
        if i != 0 && i % 8 == 0 {
            str += "\n"
        }
    }
    return str
}

func (g grid) string() string {
    var str string
    for _, v := range g {
        str += string(v + '0')
    }
    return str
}


func Solve(gridString string) (string, uint, error) {

    grid, err := parse(gridString)
    if err != nil {
        return "", 0, err
    }

    if !grid.valid() {
        return "", 0, errors.New("sudoku not valid")
    }

    cycles, err := grid.backtrack()
    if err != nil {
        return "", cycles, err
    }

    if slices.Contains(grid, 0) {
        return "", cycles, errors.New("sudoku not solved after solve (should not happen)")
    }

    if !grid.valid() {
        return "", cycles, errors.New("sudoku not valid after solve (should not happen)")
    }

    return grid.string(), cycles, nil
}

func IsValid(gridString string) bool {
    grid, err := parse(gridString)
    return err == nil && grid.valid()
}

func GenerateGrid() string {
    return "not implemented yet"
}

