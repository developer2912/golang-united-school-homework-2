package square

import (
	"math"
	"testing"
)

type (
	input struct {
		sideLen  float64
		sidesNum int
	}
	Test struct {
		input  input
		output float64
	}
)

const (
	float64EqualityThreshold = 1e-9
	angle                    = math.Pi / 3
	power                    = 2.0
)

var tests = []Test{
	{input: input{sideLen: 101.435, sidesNum: SidesTriangle}, output: 4455.293334946315},
	{input: input{sideLen: 321.435, sidesNum: SidesTriangle}, output: 44739.07120976212},

	{input: input{sideLen: 2.0, sidesNum: SidesSquare}, output: 4.0},
	{input: input{sideLen: 3.0, sidesNum: SidesSquare}, output: 9.0},

	{input: input{sideLen: 3.0, sidesNum: SidesCircle}, output: 28.274333882308138},
	{input: input{sideLen: 16.32, sidesNum: SidesCircle}, output: 836.7393271794741},

	{input: input{sideLen: 16.32, sidesNum: 1}, output: 0.0},
	{input: input{sideLen: 16.32, sidesNum: 2}, output: 0.0},
	{input: input{sideLen: 16.32, sidesNum: 5}, output: 0.0},
}

func TestCalcSquare(t *testing.T) {
	for _, test := range tests {
		if result := CalcSquare(test.input.sideLen, test.input.sidesNum); math.Abs(result-test.output) > float64EqualityThreshold {
			t.Errorf("Wrong answer: your output: %.6f test.output: %.6f", result, test.output)
		}
	}
}
