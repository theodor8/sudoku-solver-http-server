package solver

import (
    "errors"
    "slices"
)


type board []uint8


func parse(str string) (board, error) {
    if (len(str) != 81) {
        return nil, errors.New("invalid input length")
    }
    var b board = make([]uint8, 81)
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


func (b board) rowContains(r, v uint8) bool {
    return slices.Contains(b[rctoi(r, 0) : rctoi(r, 9)], v)
}

func (b board) colContains(c, v uint8) bool {
    for ; c < 81; c += 9 {
        if b[c] == v {
            return true
        }
    }
    return false
}

func (b board) boxContains(i, v uint8) bool {
    bi := i / 3 * 3 * 9 + i % 3 * 3
    for j := range uint8(9) {
        if b[bi + j / 3 * 9 + j % 3] == v {
            return true
        }
    }
    return false
}


func (b board) moveValid(i, v uint8) bool {
    r, c := itorc(i)
    bi := r / 3 * 3 + c / 3
    return !b.rowContains(r, v) && !b.colContains(c, v) && !b.boxContains(bi, v)
}

func (b board) getRow(r uint8) []uint8 {
    return slices.Clone(b[rctoi(r, 0) : rctoi(r, 9)])
}

func (b board) getCol(c uint8) []uint8 {
    col := make([]uint8, 9)
    for i := range uint8(9) {
        col[i] = b[rctoi(i, c)]
    }
    return col
}

func (b board) getBox(i uint8) []uint8 {
    box := make([]uint8, 9)
    r := i / 3 * 3
    c := i % 3 * 3
    for i := range uint8(9) {
        box[i] = b[rctoi(r + i / 3, c + i % 3)]
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
func (b board) valid() bool {
    for i := range uint8(9) {
        if !spliceValid(b.getRow(i)) ||
           !spliceValid(b.getCol(i)) ||
           !spliceValid(b.getBox(i)) {
            return false
        }
    }
    return true
}


func (b board) backtrack() (uint, error) {
    unknowns := make([]uint8, 0, 81) // indices of unknowns
    for i, v := range b {
        if v == 0 {
            unknowns = append(unknowns, uint8(i))
        }
    }
    var cycles uint = 0
    var unknownsIndex uint8 = 0
    var boardIndex uint8 = unknowns[unknownsIndex]
    for {
        foundValidTry := false
        for try := b[boardIndex] + 1; try <= 9; try++ {
            if b.moveValid(boardIndex, try) {
                b[boardIndex] = try
                foundValidTry = true
                break
            }
        }
        if foundValidTry {
            unknownsIndex++
            if unknownsIndex >= uint8(len(unknowns)) {
                break
            }
            boardIndex = unknowns[unknownsIndex]
        } else {
            if unknownsIndex == 0 {
                return cycles, errors.New("unsolvable")
            }
            b[boardIndex] = 0
            unknownsIndex--
            boardIndex = unknowns[unknownsIndex]
        }
        cycles++
    }
    return cycles, nil
}

func (b board) StringFormatted() string {
    var str string
    for i, v := range b {
        str += string(v + '0')
        if i != 0 && i % 8 == 0 {
            str += "\n"
        }
    }
    return str
}

func (b board) String() string {
    var str string
    for _, v := range b {
        str += string(v + '0')
    }
    return str
}


func Solve(boardString string) (string, uint, error) {

    board, err := parse(boardString)
    if err != nil {
        return "", 0, err
    }

    if !board.valid() {
        return "", 0, errors.New("board not valid")
    }

    cycles, err := board.backtrack()
    if err != nil {
        return "", cycles, err
    }

    if slices.Contains(board, 0) {
        return "", cycles, errors.New("board not solved after solve (should not happen)")
    }

    if !board.valid() {
        return "", cycles, errors.New("board not valid after solve (should not happen)")
    }

    return board.String(), cycles, nil
}

