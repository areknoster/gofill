package geom2d

type Segment struct{
	A,B Point
}

func NewSegment(A,B Point) Segment {
	return Segment{A,B}
}

func (s Segment) ToVector() Vector{
	return VecBetweenPoints(s.A, s.B)
}

func (s Segment) MiddlePoint() Point{
	return s.A.MoveByVector(s.ToVector().TimesScalar(0.5))
}

type SegmentSet []Segment

func (ss SegmentSet) ShapeFromMidpoints() Shape {
	vs := make([]Point, len(ss))
	for i, s := range ss {
		vs[i] = s.MiddlePoint()
	}
	return Shape{vs}
}

func (s Segment) MoveByVector(v Vector) Segment {
	return Segment{s.A.MoveByVector(v), s.B.MoveByVector(v)}
}
