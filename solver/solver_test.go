package solver

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSolve(t *testing.T) {
    type testCase struct {
        // input params
        input string

        // expected values
        expected []string
        err error
    }

    t.Run("Solve", func(t *testing.T) {
        tests := []testCase{
            // test cases from https://sudokusolver.app/
            {
                input:"003020600900305001001806400008102900700000008006708200002609500800203009005010300",
                expected:[]string{"483921657967345821251876493548132976729564138136798245372689514814253769695417382"},
            },
            {
                input:"604023000000000000009085000010006800083009706000300200300060129860100007000000400",
                expected:[]string{"654723981738941652129685374217456893583219746496378215345867129862194537971532468"},
            },
            {
                input:"004006302800003000000040800013070900008000630502000000000308500900000010037150006",
                expected:[]string{"174986352826513749359742861613875924748291635592634178461328597985467213237159486"},
            },
            {
                input:"019000000370000000046700380000060000003001600760090405000005060000106047030002900",
                expected:[]string{"819354726372618594546729381194567832253841679768293415427985163985136247631472958"},
            },
            {
                input:"217840900080000000300000058000000000108007094574900080000000002400062000030009007",
                expected:[]string{"217845963985376421346291758693184275128657394574923186769438512451762839832519647"},
            },
            {
                input:"000980000000200500080370002270000890000000003000802060002700150097108600006000900",
                expected:[]string{"425981376731246589689375412273614895168597243954832761342769158597128634816453927"},
            },
            // {
            //     input:"000000000000003085001020000000507000004000100090000000500000073002010000000040009",
            //     expected:"987654321246173985351928746128537694634892157795461832519286473472319568863745219",
            // },
            {
                input:"030005401000037000000000027004060705300009040800000000920600000600400590078502000",
                expected:[]string{"732985461146237958589146327294361785357829146861754239925613874613478592478592613"},
            },
            {
                input:"400070900009000000682540130008750001006090002000063040800000000020007000030410506",
                expected:[]string{"413276985759381264682549137348752691176894352295163748861925473524637819937418526"},
            },
            {
                input:"020000000409060080503000002800400900050000600034076800700048003300002078090700040",
                expected:[]string{"621384759479265381583917462862453917157829634934176825715648293346592178298731546",
                                  "621384759479265381583917462862453917157829634934176825716548293345692178298731546",
                                  "621387459479265381583914762862453917157829634934176825715648293346592178298731546",
                                  "621387459479265381583914762862453917157829634934176825716548293345692178298731546"},
            },
            {
                input:"003904208000300004006000005020030001460002003008090600080209000095001000000070080",
                expected:[]string{"153964278872315964946728135729436851461582793538197642387259416695841327214673589"},
            },
            {
                input:"479532816352681794861794523218946357597823461634175982725369148986417235143258679",
                expected:[]string{"479532816352681794861794523218946357597823461634175982725369148986417235143258679"},
            },
            {
                input:"485916372726453981913278654197685423864327195352149867538761249641592738279834516",
                expected:[]string{"485916372726453981913278654197685423864327195352149867538761249641592738279834516"},
            },
            {
                input:"000090030007024000200610080570386120080000090020001006060408019018007503342100000",
                expected:[]string{},
            },
        }
        for _, test := range tests {
            actual, err := Solve(test.input)
            assert.NoError(t, err)
            assert.Equal(t, len(test.expected), len(actual))
            for _, solution := range actual {
                assert.Contains(t, test.expected, solution)
            }
        }
    })
}

func TestGenerate(t *testing.T) {
    t.Run("Generate", func(t *testing.T) {
        r := rand.New(rand.NewPCG(1, 2))
        for range 10 {
            generated := Generate(r)
            solutions, err := Solve(generated)
            assert.NoError(t, err)
            assert.Equal(t, 1, len(solutions))
            assert.NotEqual(t, generated, solutions[0])
        }
    })
}



