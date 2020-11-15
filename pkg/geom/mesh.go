package geom

type Mesh struct {
	points               PointsSet
	horizontal, vertical int
}

func NewMesh(horizontal, vertical int) Mesh{
	m := Mesh{
		points:     make([]Point, horizontal * vertical),
		horizontal: horizontal,
		vertical:   vertical,
	}
	for i := 0; i < horizontal; i++{
		for j := 0; j < vertical; j++{
			m.points[m.getPointIndex(i,j)] = Point{
				X: float64(i)/float64(horizontal - 1),
				Y: float64(i)/float64(vertical - 1),
			}
		}
	}
	return m
}

// returns -1 if does not belong to mesh
func (m Mesh) getPointIndex(x, y int) int {
	if x >= m.horizontal || y >= m.vertical {
		return -1
	}
	return x + y*m.horizontal
}

func (m Mesh) GetSegments() SegmentSet {
	v, h := m.vertical, m.horizontal
	segments := make([]Segment, 3*v*h-2*(v+h)+1)
	s := 0
	for i := 0; i < m.horizontal; i++ {
		for j := 0; j < m.vertical; j++ {
			//add right
			if -1 != m.getPointIndex(i+1, j) {
				segments[s] = NewSegment(
					m.points[m.getPointIndex(i, j)],
					m.points[m.getPointIndex(i+1, j)])
				s++
			}
			//add up-right
			if -1 != m.getPointIndex(i+1, j+1) {
				segments[s] = NewSegment(
					m.points[m.getPointIndex(i, j)],
					m.points[m.getPointIndex(i+1, j+1)])
				s++
			}
			//add up
			if -1 != m.getPointIndex(i+1, j+1) {
				segments[s] = NewSegment(
					m.points[m.getPointIndex(i, j)],
					m.points[m.getPointIndex(i, j+1)])
				s++
			}
		}
	}
	return segments
}
