package circle

import (
	"collision/line"
	"collision/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewCircle - test that NewCircle behaves as expected
func TestNewCircle(t *testing.T) {
	circle := NewCircle(1, 2, 3)
	assert.Equal(t, circle.centre.X, float32(1), "x coordinate of centre should be one")
	assert.Equal(t, circle.centre.Y, float32(2), "y coordinate of centre should be two")
	assert.Equal(t, circle.radius, float32(3), "radius of circle should be three")

	circle = NewCircle(1, 2, -3)
	assert.Equal(t, circle.centre.X, float32(1), "x coordinate of centre should be one")
	assert.Equal(t, circle.centre.Y, float32(2), "y coordinate of centre should be two")
	assert.Equal(t, circle.radius, float32(3), "radius of circle should be three")
}

// TestNewCircleFromPoint - test that NewCircleFromPoint behaves as expected
func TestNewCircleFromPoint(t *testing.T) {
	p := point.Point{X: 1, Y: 2}
	circle := NewCircleFromPoint(p, 3)
	assert.Equal(t, circle.centre.X, float32(1), "x coordinate of centre should be one")
	assert.Equal(t, circle.centre.Y, float32(2), "y coordinate of centre should be two")
	assert.Equal(t, circle.radius, float32(3), "radius of circle should be three")

	circle = NewCircleFromPoint(p, -3)
	assert.Equal(t, circle.centre.X, float32(1), "x coordinate of centre should be one")
	assert.Equal(t, circle.centre.Y, float32(2), "y coordinate of centre should be two")
	assert.Equal(t, circle.radius, float32(3), "radius of circle should be three")
}

// TestGetCentreAndRadius - test that GetCentreAndRadius behaves as expected
func TestGetCentreAndRadius(t *testing.T) {
	circle := Circle{point.Point{X: 1, Y: 2}, 3}
	centre, radius := circle.GetCentreAndRadius()
	assert.Equal(t, centre.X, float32(1), "x coordinate of centre should be one")
	assert.Equal(t, centre.Y, float32(2), "y coordinate of centre should be two")
	assert.Equal(t, radius, float32(3), "radius of circle should be three")
}

// TestContainsPoint - test that ContainsPoint behaves as expected
func TestContainsPoint(t *testing.T) {
	var testCircle Circle = Circle{centre: point.Point{X: 1, Y: 1}, radius: 5}
	var testPoint point.Point = point.Point{X: 1, Y: 1}

	ok := testCircle.ContainsPoint(testPoint)
	assert.True(t, ok, "point should be inside circle")

	testPoint = point.Point{X: 1, Y: 5.999}
	ok = testCircle.ContainsPoint(testPoint)
	assert.True(t, ok, "point should be inside circle")

	testPoint = point.Point{X: 1, Y: 6.001}
	ok = testCircle.ContainsPoint(testPoint)
	assert.False(t, ok, "point should be outside circle")

	testPoint = point.Point{X: 5.999, Y: 1}
	ok = testCircle.ContainsPoint(testPoint)
	assert.True(t, ok, "point should be inside circle")

	testPoint = point.Point{X: 6.001, Y: 1}
	ok = testCircle.ContainsPoint(testPoint)
	assert.False(t, ok, "point should be outside circle")

	testPoint = point.Point{X: -2.999, Y: 4}
	ok = testCircle.ContainsPoint(testPoint)
	assert.True(t, ok, "point should be inside circle")

	testPoint = point.Point{X: -3.001, Y: 4}
	ok = testCircle.ContainsPoint(testPoint)
	assert.False(t, ok, "point should be outside circle")
}

// TestCircumferenceTouchesPoint - test that CircumferenceTouchesPoint behaves as expected
func TestCircumferenceTouchesPoint(t *testing.T) {
	var testCircle Circle = Circle{centre: point.Point{X: 0, Y: 0}, radius: 5}
	var testPoint point.Point

	testPoint = point.Point{X: -5.0000015, Y: 0}
	ok := testCircle.CircumferenceTouchesPoint(testPoint)
	assert.False(t, ok, "circumference should not touch point")

	testPoint = point.Point{X: -5.0000011, Y: 0}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.True(t, ok, "circumference should touch point")

	testPoint = point.Point{X: -5.0000009, Y: 0}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.True(t, ok, "circumference should touch point")

	testPoint = point.Point{X: 0, Y: 0}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.False(t, ok, "circumference should not touch point")

	testPoint = point.Point{X: 5.0000009, Y: 0}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.True(t, ok, "circumference should touch point")

	testPoint = point.Point{X: 5.0000011, Y: 0}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.True(t, ok, "circumference should touch point")

	testPoint = point.Point{X: 5.0000015, Y: 0}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.False(t, ok, "circumference should not touch point")

	testPoint = point.Point{X: 0, Y: -5.0000015}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.False(t, ok, "circumference should not touch point")

	testPoint = point.Point{X: 0, Y: -5.0000011}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.True(t, ok, "circumference should touch point")

	testPoint = point.Point{X: 0, Y: -5.0000009}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.True(t, ok, "circumference should touch point")

	testPoint = point.Point{X: 0, Y: 5.0000009}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.True(t, ok, "circumference should touch point")

	testPoint = point.Point{X: 0, Y: -5.0000011}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.True(t, ok, "circumference should touch point")

	testPoint = point.Point{X: 0, Y: 5.0000015}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.False(t, ok, "circumference should not touch point")

	testPoint = point.Point{X: 2.9999995, Y: 4}
	ok = testCircle.CircumferenceTouchesPoint(testPoint)
	assert.True(t, ok, "circumference should touch point")
}

// TestCirclesIntersect - test that CirclesIntersect behaves as expected
func TestCirclesIntersect(t *testing.T) {
	c := Circle{centre: point.Point{X: 0, Y: 0}, radius: 10}
	d := Circle{centre: point.Point{X: 0, Y: 0}, radius: 10}

	ok := c.CirclesIntersect(d)
	assert.True(t, ok, "circles should intersect")

	d = Circle{centre: point.Point{X: 10, Y: 0}, radius: 10}
	ok = c.CirclesIntersect(d)
	assert.True(t, ok, "circles should intersect")

	d = Circle{centre: point.Point{X: 20, Y: 0}, radius: 10}
	ok = c.CirclesIntersect(d)
	assert.True(t, ok, "circles should intersect")

	d = Circle{centre: point.Point{X: 20.0001, Y: 0}, radius: 10}
	ok = c.CirclesIntersect(d)
	assert.False(t, ok, "circles should not intersect")
}

// TestIntersectsLineSegment - check that IntersectsLineSegment behaves as expected
func TestIntersectsLineSegment(t *testing.T) {
	c := Circle{centre: point.Point{X: 0, Y: 0}, radius: 10}

	// start and end are centre of circle
	ls := line.LineSegment{Start: point.Point{X: 0, Y: 0}, End: point.Point{X: 0, Y: 0}}
	ok := c.InstersectsLineSegment(ls)
	assert.True(t, ok, "line segment should intersect circle")

	// end is centre of circle, start is outside
	ls = line.LineSegment{Start: point.Point{X: 1000, Y: 50}, End: point.Point{X: 0, Y: 0}}
	ok = c.InstersectsLineSegment(ls)
	assert.True(t, ok, "line segment should intersect circle")

	// end is in circle, start is outside
	ls = line.LineSegment{Start: point.Point{X: 1000, Y: 50}, End: point.Point{X: 5, Y: 5}}
	ok = c.InstersectsLineSegment(ls)
	assert.True(t, ok, "line segment should intersect circle")

	// end is just inside circle
	ls = line.LineSegment{Start: point.Point{X: 1000, Y: 50}, End: point.Point{X: 9.999, Y: 0}}
	ok = c.InstersectsLineSegment(ls)
	assert.True(t, ok, "line segment should intersect circle")

	// end lies on circumference
	ls = line.LineSegment{Start: point.Point{X: 1000, Y: 50}, End: point.Point{X: 10, Y: 0}}
	ok = c.InstersectsLineSegment(ls)
	assert.True(t, ok, "line segment should intersect circle")

	// end lies just outside circle, start is outside
	ls = line.LineSegment{Start: point.Point{X: 1000, Y: 50}, End: point.Point{X: 10.0001, Y: 0}}
	ok = c.InstersectsLineSegment(ls)
	assert.False(t, ok, "line segment should not intersect circle")

	// start and end outside circle, line segment should pass through circumference
	ls = line.LineSegment{Start: point.Point{X: 10, Y: 50}, End: point.Point{X: 10, Y: -100}}
	ok = c.InstersectsLineSegment(ls)
	assert.True(t, ok, "line segment should intersect circle")

	// very near miss
	ls = line.LineSegment{Start: point.Point{X: 10, Y: 50}, End: point.Point{X: 10.0001, Y: -100}}
	ok = c.InstersectsLineSegment(ls)
	assert.False(t, ok, "line segment should not intersect circle")

	// only just hits
	ls = line.LineSegment{Start: point.Point{X: 10, Y: 50}, End: point.Point{X: 9.999, Y: -100}}
	ok = c.InstersectsLineSegment(ls)
	assert.True(t, ok, "line segment should intersect circle")

	// only just hits
	ls = line.LineSegment{Start: point.Point{X: 1000, Y: 10}, End: point.Point{X: -48, Y: 9.999}}
	ok = c.InstersectsLineSegment(ls)
	assert.True(t, ok, "line segment should intersect circle")

	// line segment through circumference
	ls = line.LineSegment{Start: point.Point{X: 1000, Y: 10}, End: point.Point{X: -48, Y: 10}}
	ok = c.InstersectsLineSegment(ls)
	assert.True(t, ok, "line segment should intersect circle")

	// near miss
	ls = line.LineSegment{Start: point.Point{X: 1000, Y: 10.0001}, End: point.Point{X: -48, Y: 10}}
	ok = c.InstersectsLineSegment(ls)
	assert.False(t, ok, "line segment should not intersect circle")
}
