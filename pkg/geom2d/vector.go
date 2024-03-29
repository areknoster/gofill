package geom2d

import (
	"math"
)

type Vector struct {
	X, Y float64
}

func VecBetweenPoints(v1, v2 Point) Vector {
	return Vector{v2.X - v1.X, v2.Y - v1.Y}
}

func VecSqEuclidNormSq(v Vector) float64 {
	return v.X*v.X + v.Y*v.Y
}

func SqDistBetweenPoints(a, b Point) float64 {
	return VecSqEuclidNormSq(VecBetweenPoints(a, b))
}

func DotVec(v0, v1 Vector) float64 {
	return v0.X*v1.X + v0.Y*v1.Y
}

var ZeroVector = Vector{0, 0}
var WrongPoint = Point{math.NaN(), math.NaN()}
var WrongVector = Vector{math.NaN(), math.NaN()}

func ClosestPoint(P, A, B Point) Point {
	vecAB := VecBetweenPoints(A, B)
	vecAP := VecBetweenPoints(A, P)
	normK := DotVec(vecAB, vecAP)
	normL := DotVec(vecAB, vecAB)
	slope := -normK / normL

	f := func(s float64) Point {
		return Point{
			X: (1.0-slope)*A.X + slope*B.X - P.X,
			Y: (1.0-slope)*A.Y + slope*B.Y - P.Y,
		}
	}

	if slope < 0 || slope > 1 {
		return WrongPoint
	}
	return f(slope)

}

func (v Vector) TimesScalar(s float64) Vector {
	return Vector{s * v.X, s * v.Y}
}

func (v Vector) DeltaXY() float64 {
	return v.X / v.Y
}

func (v Vector) Add(va Vector) Vector {
	return Vector{v.X + va.X, v.Y + va.Y}
}
func (v Vector) SqNorm() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.SqNorm())
}

func AngleToVecor(angle, length float64) Vector {
	return Vector{
		X: length * math.Cos(angle),
		Y: length * math.Sin(angle),
	}
}

