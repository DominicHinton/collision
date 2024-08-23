package line

import (
	"collision/point"
)

// LineSegment - defined by start and end points
type LineSegment struct {
	Start point.Point // start point
	End   point.Point // end point
}

// Length - get length of line segment
func (ls LineSegment) Length() float32 {
	return ls.Start.Distance(ls.End)
}

// IsVertical - boolean indicating vertical line.
// This uses global delta to check near equality of y coords.
// Real life very near vertical line segments will return true.
// Two points touching within global delta (zero length line) will return false.
func (ls LineSegment) IsVertical() bool {
	return ls.Start.SameX(ls.End) && !ls.Start.SameY(ls.End)
}

// IsHorizontal - boolean indicating horizontal line.
// This uses global delta to check near equality of x coords.
// Real life very near horizonal line segments will return true.
// Two points touching within global delta (zerolength line) will return false.
func (ls LineSegment) IsHorizontal() bool {
	return ls.Start.SameY(ls.End) && !ls.Start.SameX(ls.End)
}

// HasPoint - returns boolean indicating whether point lies on segment within delta
func (ls LineSegment) HasPoint(p point.Point) bool {
	distanceFromPToStart := p.Distance(ls.Start)
	distanceFromPToEnd := p.Distance(ls.End)
	totalDistance := distanceFromPToStart + distanceFromPToEnd
	return point.AreWithinEasyDelta(totalDistance, ls.Length())
}

// IntersectsLineSegment - returns boolean indicating whether two line segments meet.
// Also returns coordinates of intersection if true and Point(0,0) if false.
func (ls LineSegment) IntersectsLineSegment(secondLineSegment LineSegment) (point.Point, bool) {
	x1, x2, x3, x4 := ls.Start.X, ls.End.X, secondLineSegment.Start.X, secondLineSegment.End.X
	y1, y2, y3, y4 := ls.Start.Y, ls.End.Y, secondLineSegment.Start.Y, secondLineSegment.End.Y

	// get relative distance to intersection point
	denominator := ((y4-y3)*(x2-x1) - (x4-x3)*(y2-y1))

	// if denominator is zero, then line segments are parallel
	if point.AreWithinGlobalDelta(denominator, 0) {
		// if either end of second line segment lies on first line segment, then the end is an intersection point
		if ls.HasPoint(secondLineSegment.Start) {
			return secondLineSegment.Start, true
		}
		if ls.HasPoint(secondLineSegment.End) {
			return secondLineSegment.End, true
		}
		// if either end of first line segment lines on second line segment, then the end is an intersection point
		if secondLineSegment.HasPoint(ls.Start) {
			return ls.Start, true
		}
		if secondLineSegment.HasPoint(ls.End) {
			return ls.End, true
		}

		// if none of the four ends lie on the other line segment, then the two parallel lines don't overlap
		return point.Point{}, false
	}
	numeratorA := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3))
	numeratorB := ((x2-x1)*(y1-y3) - (y2-y1)*(x1-x3))

	uA := numeratorA / denominator
	uB := numeratorB / denominator

	// check that both relative distances are within range zero to one to ensure intersection point
	// lies on each line segment. If this check is a false then lines don't intersect
	intersects := (uA >= 0 && uA <= 1 && uB >= 0 && uB <= 1)
	if !intersects {
		return point.Point{}, false
	}

	// find coordinates of intersection point
	intersectionX := x1 + (uA * (x2 - x1))
	intersectionY := y1 + (uA * (y2 - y1))

	return point.Point{X: intersectionX, Y: intersectionY}, true
}
