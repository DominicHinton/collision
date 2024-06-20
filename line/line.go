package line

import "collision/point"

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
