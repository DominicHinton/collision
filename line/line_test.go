package line

import (
	"collision/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLength - test that Length behaves as expected
func TestLength(t *testing.T) {
	ls := LineSegment{Start: point.Point{X: 0, Y: 0}, End: point.Point{X: 3, Y: 4}}
	assert.Equal(t, float32(5), ls.Length(), "length should be 5")

	ls = LineSegment{Start: point.Point{X: -5, Y: 0}, End: point.Point{X: 0, Y: 12}}
	assert.Equal(t, float32(13), ls.Length(), "length should be 13")
}

// TestIsVertical - test that IsVertical behaves as expected
func TestIsVertical(t *testing.T) {
	ls := LineSegment{Start: point.Point{X: 0, Y: 0}, End: point.Point{X: 0, Y: 4}}

	ok := ls.IsVertical()
	assert.True(t, ok, "should be vertical")

	ls = LineSegment{Start: point.Point{X: -100, Y: 0}, End: point.Point{X: -100.000001, Y: 4}}
	ok = ls.IsVertical()
	assert.True(t, ok, "should be vertical")

	ls = LineSegment{Start: point.Point{X: -100, Y: 0}, End: point.Point{X: -100.00001, Y: 4}}
	ok = ls.IsVertical()
	assert.False(t, ok, "should not be vertical")
}

// TestIsHorizontal - test that IsHorizontal  behaves as expected
func TestIsHorizontal(t *testing.T) {
	ls := LineSegment{Start: point.Point{X: 0, Y: 0}, End: point.Point{X: 4, Y: 0}}

	ok := ls.IsHorizontal()
	assert.True(t, ok, "should be horizontal")

	ls = LineSegment{Start: point.Point{X: 0, Y: -100}, End: point.Point{X: 4, Y: -100.000001}}
	ok = ls.IsHorizontal()
	assert.True(t, ok, "should be horizontal")

	ls = LineSegment{Start: point.Point{X: 0, Y: -100}, End: point.Point{X: -4, Y: -100.00001}}
	ok = ls.IsHorizontal()
	assert.False(t, ok, "should not be horizontal")
}

// TestHasPoint - test that HasPoint behaves as expected
func TestHasPoint(t *testing.T) {
	ls := LineSegment{Start: point.Point{X: 0, Y: 0}, End: point.Point{X: 50, Y: -100}}

	p := point.Point{X: 0, Y: 0}
	ok := ls.HasPoint(p)
	assert.True(t, ok, "point should be on line")

	p = point.Point{X: 50, Y: -100}
	ok = ls.HasPoint(p)
	assert.True(t, ok, "point should be on line")

	p = point.Point{X: 25, Y: -50}
	ok = ls.HasPoint(p)
	assert.True(t, ok, "point should be on line")

	p = point.Point{X: 5, Y: -10}
	ok = ls.HasPoint(p)
	assert.True(t, ok, "point should be on line")

	p = point.Point{X: 5.04, Y: -10}
	ok = ls.HasPoint(p)
	assert.True(t, ok, "point should be on line")

	p = point.Point{X: 5.06, Y: -10}
	ok = ls.HasPoint(p)
	assert.False(t, ok, "point should be on line")
}

// TestIntersectsLineSegment - test for correct function of IntersectsLineSegment
func TestIntersectsLineSegment(t *testing.T) {
	// check that a line intersecting with itself will return line.Start and true
	ls1 := LineSegment{Start: point.Point{X: 0, Y: 0}, End: point.Point{X: 10, Y: 10}}
	intersectionPoint, ok := ls1.IntersectsLineSegment(ls1)
	assert.True(t, ok)
	expectedPoint := point.Point{X: 0, Y: 0}
	assert.Equal(t, expectedPoint, intersectionPoint, "lines should intersect at (0,0)")

	// check an obvious intersection
	ls2 := LineSegment{Start: point.Point{X: 0, Y: 1}, End: point.Point{X: 1, Y: 0}}
	intersectionPoint, ok = ls1.IntersectsLineSegment(ls2)
	assert.True(t, ok)
	expectedPoint = point.Point{X: 0.5, Y: 0.5}
	assert.Equal(t, expectedPoint, intersectionPoint, "lines should intersect at (0.5, 0.5)")

	// two parallel line segments with no intersection
	ls2 = LineSegment{Start: point.Point{X: 0, Y: 1}, End: point.Point{X: 10, Y: 11}}
	intersectionPoint, ok = ls1.IntersectsLineSegment(ls2)
	assert.False(t, ok)
	expectedPoint = point.Point{X: 0, Y: 0}
	assert.Equal(t, expectedPoint, intersectionPoint, "lines should not intersect")

	// two non-parallel lines with no intersection
	ls2 = LineSegment{Start: point.Point{X: -6, Y: -9}, End: point.Point{X: -2, Y: -10}}
	intersectionPoint, ok = ls1.IntersectsLineSegment(ls2)
	assert.False(t, ok)
	expectedPoint = point.Point{X: 0, Y: 0}
	assert.Equal(t, expectedPoint, intersectionPoint, "lines should not intersect")

}
