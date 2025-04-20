package sudoku

import (
	"errors"
	"slices"
)


type grid []uint8



func parse(str string) (grid, error) {
    if (len(str) != 81) {
        return nil, errors.New("invalid input length")
    }
    var g grid = make([]uint8, 81)
    for i, v := range str {
        if v < '0' || v > '9' {
            return nil, errors.New("invalid char")
        }
        g[i] = uint8(v - '0')
    }
    return g, nil
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
    gi := i / 3 * 3 * 9 + i % 3 * 3
    for j := range uint8(9) {
        if g[gi + j / 3 * 9 + j % 3] == v {
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


func (g grid) string() string {
    var str string
    for _, v := range g {
        str += string(v + '0')
    }
    return str
}


func (g grid) solutionValid(solution grid) error {
    if slices.Contains(solution, 0) {
        return errors.New("solution contains 0 after solve (should not happen)")
    }
    if !solution.valid() {
        return errors.New("solution not valid after solve (should not happen)")
    }
    for i := range g {
        if g[i] != 0 && g[i] != solution[i] {
            return errors.New("solution not matching with grid knowns (should not happen)")
        }
    }
    return nil
}

func IsValid(gridString string) bool {
    grid, err := parse(gridString)
    return err == nil && grid.valid()
}
