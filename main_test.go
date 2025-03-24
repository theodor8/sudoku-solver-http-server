package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestParse(t *testing.T) {
    type testCase struct {
        // input params
        input string

        // expected values
        expected board
        err error
    }

    t.Run("parse", func(t *testing.T) {
        tests := []testCase{
            {input:"900420030006000500300060800702090000630000058000080207007040009008000300060057001",
             expected:board{9,0,0,4,2,0,0,3,0,0,0,6,0,0,0,5,0,0,3,0,0,0,6,0,8,0,0,7,0,2,0,9,0,0,0,0,6,3,0,0,0,0,0,5,8,0,0,0,0,8,0,2,0,7,0,0,7,0,4,0,0,0,9,0,0,8,0,0,0,3,0,0,0,6,0,0,5,7,0,0,1},
            },
            {input:"003020600900305001001806400008102900700000008006708200002609500800203009005010300",
            expected:board{0,0,3,0,2,0,6,0,0,9,0,0,3,0,5,0,0,1,0,0,1,8,0,6,4,0,0,0,0,8,1,0,2,9,0,0,7,0,0,0,0,0,0,0,8,0,0,6,7,0,8,2,0,0,0,0,2,6,0,9,5,0,0,8,0,0,2,0,3,0,0,9,0,0,5,0,1,0,3,0,0},
            },
        }
        for _, test := range tests {
            actual, err := parse(test.input)
            assert.NoError(t, err)
            assert.Equal(t, test.expected, actual)
        }
    })
}

func TestSolve(t *testing.T) {
    type testCase struct {
        // input params
        input board

        // expected values
        expected board
        err error
    }

    t.Run("solve", func(t *testing.T) {
        tests := []testCase{
            {input:board{0,0,3,0,2,0,6,0,0,9,0,0,3,0,5,0,0,1,0,0,1,8,0,6,4,0,0,0,0,8,1,0,2,9,0,0,7,0,0,0,0,0,0,0,8,0,0,6,7,0,8,2,0,0,0,0,2,6,0,9,5,0,0,8,0,0,2,0,3,0,0,9,0,0,5,0,1,0,3,0,0},
             expected:board{4,8,3,9,2,1,6,5,7,9,6,7,3,4,5,8,2,1,2,5,1,8,7,6,4,9,3,5,4,8,1,3,2,9,7,6,7,2,9,5,6,4,1,3,8,1,3,6,7,9,8,2,4,5,3,7,2,6,8,9,5,1,4,8,1,4,2,5,3,7,6,9,6,9,5,4,1,7,3,8,2},
            },
        }
        for _, test := range tests {
            test.input.solve()
            assert.True(t, test.input.valid())
            assert.Equal(t, test.expected, test.input)
        }
    })

}
