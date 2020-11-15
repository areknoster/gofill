package geom

import "math"

type Point struct {
	X, Y float64
}

func (v Point) MoveByVector(vec Vector) Point {
	return Point{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
	}
}


var NanPoint = Point{math.NaN(), math.NaN()}