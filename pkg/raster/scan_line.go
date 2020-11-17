package raster

import (
	"fmt"
	"image"
	"math"

	"github.com/areknoster/gofill/pkg/geom2d"
)

type bucketEdge struct {
	yMax, xMin, delta float64
	next              *bucketEdge
}

func (be *bucketEdge) insertInFront(el *bucketEdge) {
	el.next = be.next
	be.next = el
}

func (be *bucketEdge) insertOrdered(newEl *bucketEdge) {
	cur := be
	next := cur.next
	for {
		if next == nil || next.xMin >= newEl.xMin {
			cur.insertInFront(newEl)
			return
		}
		cur = next
		next = next.next
	}
}

func (be *bucketEdge) deleteNext() {
	be.next = be.next.next
}
func (be *bucketEdge) String() string {
	str := ""
	for ; be != nil; be = be.next {
		if be.xMin == -math.MaxFloat64{
			continue
		}
		str += fmt.Sprintf("(%f, %f, %f)", be.yMax, be.xMin, be.delta)
	}
	return str + "\n"
}

type edgeTable struct {
	bucketEdges []*bucketEdge
	offset      int
}

func (et edgeTable) String() string {
	str := "y: (yMax, xMin, delta), ...\n"
	for i, be := range et.bucketEdges {
		str += fmt.Sprintf("%d: %s\n", i+et.offset, be)
	}
	return str
}

func getYRange(shape geom2d.Shape) (float64, float64) {
	min, max := math.MaxFloat64, -math.MaxFloat64
	for _, point := range shape.PointsSet {
		if point.Y < min {
			min = point.Y
		}
		if point.Y > max {
			max = point.Y
		}
	}
	return min, max
}

const EPSILON float64 = 0.0000001
func floatEquals(a, b float64) bool {
	return (a - b) < EPSILON && (b - a) < EPSILON
}

func newETFromShape(shape geom2d.Shape, h int) edgeTable {
	denormalize := func(f float64) int {
		return int(f * float64(h))
	}
	yMinNorm, yMaxNorm := getYRange(shape)
	normRange := yMaxNorm - yMinNorm
	yMinDenorm := denormalize(yMinNorm)
	yMAxDenorm := denormalize(yMaxNorm)
	denormRange := yMAxDenorm - yMinDenorm
	dnr := float64(denormRange) / normRange
	et := edgeTable{
		bucketEdges: make([]*bucketEdge, yMAxDenorm-yMinDenorm+1),
		offset:      yMinDenorm,
	}
	// set sentinels
	for i := range et.bucketEdges {
		et.bucketEdges[i] = &bucketEdge{
			xMin: -math.MaxFloat64,
			next: nil,
		}
	}
	edges := shape.ToEdgeSet()
	be := et.bucketEdges
	for _, edge := range edges {
		if edge.A.Y < edge.B.Y {
			edge.A, edge.B = edge.B, edge.A
		}
		spot := int((edge.B.Y - yMinNorm) * dnr)
		delta := edge.ToVector().DeltaXY()
		if delta == math.Inf(1) {
			continue
		}
		newEl := &bucketEdge{
			yMax:  edge.A.Y,
			xMin:  edge.B.X,
			delta: delta,
			next:  nil,
		}
		// insert in right place
		be[spot].insertOrdered(newEl)
	}
	return et
}

func ScanLine(shape geom2d.Shape, img *image.RGBA, set func(p Pixel)) {
	denormalizeX := func(f float64) int {
		return int(f * float64(img.Rect.Max.X))
	}
	denormalizeY := func(f float64) int {
		return int(f * float64(img.Rect.Max.Y))
	}
	et := newETFromShape(shape, img.Rect.Max.Y)
	//fmt.Print(et)
	aet := &bucketEdge{
		xMin: -math.MaxFloat64,
		next: nil,
	}
	y := et.offset
	for {
		//fmt.Printf("y: %d\n", y)
		bes := et.bucketEdges[y-et.offset]
		next := &bucketEdge{}
		for cur := bes.next; cur != nil; cur = next {
			next = cur.next
			cur.next = nil
			aet.insertOrdered(cur)
		}
		//fmt.Printf("elements added: %s\n", aet)
		// fil spaces
		var first,second *bucketEdge
		for second = aet; second.next != nil; {
			first = second.next
			second = first.next
			if second == nil{
				fmt.Printf("elements added: %s\n", aet)
				break

			}
			x0 := denormalizeX(first.xMin)
			x1 := denormalizeX(second.xMin)
			for ; x0 <= x1; x0++ {
				set(Pixel{x0,y})
			}
		}
		y++
		//remove unused
		for cur := aet; cur.next != nil; {
			if denormalizeY(cur.next.yMax) == y {
				cur.deleteNext()
				continue
			}
			cur = cur.next
		}
		//fmt.Printf("unused removed: %s\n", aet)
		if y-et.offset == len(et.bucketEdges) {
			return
		}
		//update x
		for cur := aet.next; cur != nil; cur = cur.next {
			if cur.delta == 0.0 ||cur.delta == -0.0{
				continue
			}
			cur.xMin += cur.delta / float64(img.Rect.Max.Y)
		}
		//fmt.Printf("x updated: %s\n", aet)

	}
}
