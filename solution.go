package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesCircle   = iota
	SidesTriangle = iota + 3
	SidesSquare   = iota + 4
)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	const power = 2.0
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, power)
	case SidesTriangle:
		angle := math.Pi / 3
		return math.Pow(sideLen, power) * math.Sin(angle) * 0.5
	case SidesSquare:
		return math.Pow(sideLen, power)
	default:
		return 0
	}
}
