package point

import (
	"errors"
	"math"
)

// Point - represents a point in 2D space
type Point struct {
	X float32 // position in x dimension
	Y float32 // position in y dimension
}

// SameX - determine if two points have same x value to within value of Delta
func (A Point) SameX(B Point) bool {
	return AreWithinGlobalDelta(A.X, B.X)
}

// SameY - determine if two points have same y value to within value of Delta
func (A Point) SameY(B Point) bool {
	return AreWithinGlobalDelta(A.Y, B.Y)
}

// AreTouching - determine if two points are touching to within value of Delta
func (A Point) AreTouching(B Point) bool {
	return A.SameX(B) && A.SameY(B)
}

// XDistance - separation between points A and B along x dimension
func (A Point) XDistance(B Point) float32 {
	return A.X - B.X
}

// YDistance - separation between points A and B along y dimension
func (A Point) YDistance(B Point) float32 {
	return A.Y - B.Y
}

// Distance - distance between two points
func (A Point) Distance(B Point) float32 {
	xDistance := float64(A.XDistance(B))
	yDistance := float64(A.YDistance(B))
	return float32(math.Sqrt((xDistance * xDistance) + (yDistance * yDistance)))
}

// GetMinMax - return maximum and minimum x and y values from an array of points
func GetMinMax(points []Point) (minX, maxX, minY, maxY float32, err error) {
	if len(points) == 0 {
		err = errors.New("empty array of points cannot result in minimum and maximum values")
		return
	}
	// instantiate min and max values from first entry in array
	minX, maxX = points[0].X, points[0].X
	minY, maxY = points[0].Y, points[0].Y
	// loop through remaining points
	for i := 1; i < len(points); i++ {
		p := points[i]
		// set min or max X if necessary
		if p.X < minX {
			minX = p.X
		} else if p.X > maxX {
			maxX = p.X
		}
		// set min or max Y if necessary
		if p.Y < minY {
			minY = p.Y
		} else if p.Y > maxY {
			maxY = p.Y
		}
	}
	return
}
