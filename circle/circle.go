package circle

import (
	"collision/line"
	"collision/point"
)

// Circle - defined by centre and radius
type Circle struct {
	centre point.Point // point at centre of circle
	radius float32     // radius of circle
}

// NewCircle - returns a new circle. If radius is provided as a negative, then abs value is assigned
func NewCircle(x, y, radius float32) Circle {
	return NewCircleFromPoint(point.Point{X: x, Y: y}, radius)
}

// NewCircle - returns a new circle. If radius is provided as a negative, then abs value is assigned
func NewCircleFromPoint(p point.Point, radius float32) Circle {
	return Circle{
		centre: p,
		radius: point.Abs(radius),
	}
}

// GetCentreAndRadius - returns a point indicating centre and a float32 indicating readius of a circle
func (c Circle) GetCentreAndRadius() (centre point.Point, radius float32) {
	return c.centre, c.radius
}

// ContainsPoint - returns boolean indicating whether a point is inside a circle
func (c Circle) ContainsPoint(p point.Point) bool {
	distanceFromCentre := c.centre.Distance(p)
	return distanceFromCentre <= c.radius
}

// CircumferenceTouchesPoint - returns boolean indicating whether a point lies on circumference of circle
func (c Circle) CircumferenceTouchesPoint(p point.Point) bool {
	distanceFromCentre := c.centre.Distance(p)
	return point.Abs(distanceFromCentre-c.radius) < point.Delta
}

// CirclesIntersect - returns boolean indicating whether two circles intersect
// no reference to delta, so floating point error possible if circles intersect at one point
func (c Circle) CirclesIntersect(d Circle) bool {
	sumOfRadii := c.radius + d.radius
	distanceBetweenCentres := c.centre.Distance(d.centre)
	return distanceBetweenCentres <= sumOfRadii
}

// InstersectsLineSegment - returns boolean indicating whether circle intersects line segment
func (c Circle) InstersectsLineSegment(ls line.LineSegment) bool {
	// check if start or end of line segment are in circle as this would imply intersection
	if c.ContainsPoint(ls.Start) || c.ContainsPoint(ls.End) {
		return true
	}
	// if neither are inside, find closest point on line (not segment) and check this
	return c.intersectsLineSegmentByCheckingClosestPoint(ls)

}

// intersectsLineSegmentByCheckingClosestPoint - find closest point on line segment to the circle and then check
// whether the closest point satisfies line intersecting circle
func (c Circle) intersectsLineSegmentByCheckingClosestPoint(ls line.LineSegment) bool {
	length := ls.Length()
	dotProduct := (((c.centre.X-ls.Start.X)*(ls.End.X-ls.Start.X) + (c.centre.Y-ls.Start.Y)*(ls.End.Y-ls.Start.Y)) / (length * length))

	xClosest := ls.Start.X + (dotProduct * (ls.End.X - ls.Start.X))
	yClosest := ls.Start.Y + (dotProduct * (ls.End.Y - ls.Start.Y))

	closestPoint := point.Point{X: xClosest, Y: yClosest}

	// check if closest point is in circle - if it doesn't, then they cannot intersect
	if !c.ContainsPoint(closestPoint) {
		return false
	}

	// if at this point, then check if closest point is on line segment
	return ls.HasPoint(closestPoint)
}
