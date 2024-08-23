package polygon_test

import (
	"collision/line"
	"collision/point"
	"collision/polygon"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewValidatedXYRectangleFromOppositeVertices - test behaviour of NewValidatedXYRectangleFromOppositeVertices
func TestNewValidatedXYRectangleFromOppositeVertices(t *testing.T) {
	p := getTestPoints(4)
	// unit square a and c vertices
	vertices := []point.Point{p[0][0], p[1][1]}
	r, err := polygon.NewValidatedXYRectangleFromOppositeVertices(vertices)
	assert.Nil(t, err, "unit rectangle from opposite vertices, no errors expected")
	assert.Equal(t, r.Vertices[0], p[0][0], "first corner should be (0,0)")
	assert.Equal(t, r.Vertices[1], p[0][1], "second corner should be (0,1)")
	assert.Equal(t, r.Vertices[2], p[1][1], "third corner should be (1,1)")
	assert.Equal(t, r.Vertices[3], p[1][0], "fourth corner should be (1,0)")
	// unit square c and a vertices
	vertices = []point.Point{p[1][1], p[0][0]}
	r2, err := polygon.NewValidatedXYRectangleFromOppositeVertices(vertices)
	assert.Nil(t, err, "unit rectangle from opposite vertices, no errors expected")
	assert.Equal(t, r, r2, "expect reversal of points to result in same rectangle")
	// unit square b and d vertices
	vertices = []point.Point{p[0][1], p[1][0]}
	r2, err = polygon.NewValidatedXYRectangleFromOppositeVertices(vertices)
	assert.Nil(t, err, "unit rectangle from opposite vertices, no errors expected")
	assert.Equal(t, r, r2, "expect use of other pair of opposite points to result in same rectangle")
	// empty vertices should result in error
	vertices = []point.Point{}
	_, err = polygon.NewValidatedXYRectangleFromOppositeVertices(vertices)
	assert.EqualError(t, err, "instantiating rectangle from two opposite corners requires exactly two points to be passed, the array of points was of length 0", "wrong error encountered")
	// 1 vertex should result in error
	vertices = []point.Point{p[2][2]}
	_, err = polygon.NewValidatedXYRectangleFromOppositeVertices(vertices)
	assert.EqualError(t, err, "instantiating rectangle from two opposite corners requires exactly two points to be passed, the array of points was of length 1", "wrong error encountered")
	// 3 vertices should result in error
	vertices = []point.Point{p[2][2], p[1][1], p[0][0]}
	_, err = polygon.NewValidatedXYRectangleFromOppositeVertices(vertices)
	assert.EqualError(t, err, "instantiating rectangle from two opposite corners requires exactly two points to be passed, the array of points was of length 3", "wrong error encountered")
	// vertices with shared x should result in error
	vertices = []point.Point{p[0][1], p[0][0]}
	_, err = polygon.NewValidatedXYRectangleFromOppositeVertices(vertices)
	assert.EqualError(t, err, "both corners passed shared the same x value: 0", "wrong error encountered")
	// vertices with shared y should result in error
	vertices = []point.Point{p[0][2], p[3][2]}
	_, err = polygon.NewValidatedXYRectangleFromOppositeVertices(vertices)
	assert.EqualError(t, err, "both corners passed shared the same y value: 2", "wrong error encountered")
}

// TestNewValidatedXYRectangleFrom4Points - test function of NewValidatedXYRectangleFrom4Points
func TestNewValidatedXYRectangleFrom4Points(t *testing.T) {
	p := getTestPoints(4)
	// unit rectangle with points in expected order
	vertices := []point.Point{p[0][0], p[0][1], p[1][1], p[1][0]}
	r, err := polygon.NewValidatedXYRectangleFrom4Points(vertices)
	assert.Nil(t, err, "unit rectangle from four vertices in order, no errors expected")
	assert.Equal(t, r.Vertices[0], p[0][0], "first corner should be (0,0)")
	assert.Equal(t, r.Vertices[1], p[0][1], "second corner should be (0,1)")
	assert.Equal(t, r.Vertices[2], p[1][1], "third corner should be (1,1)")
	assert.Equal(t, r.Vertices[3], p[1][0], "fourth corner should be (1,0)")
	// unit rectangle with points in unexpected order
	vertices = []point.Point{p[1][1], p[0][1], p[0][0], p[1][0]}
	r2, err := polygon.NewValidatedXYRectangleFrom4Points(vertices)
	assert.Nil(t, err, "unit rectangle from four vertices out of order, no errors expected")
	assert.Equal(t, r, r2, "order of points passed to constructor should not affect rectangle")
	// 5 vertices is too many
	vertices = []point.Point{p[0][0], p[0][1], p[1][1], p[1][0], p[2][2]}
	_, err = polygon.NewValidatedXYRectangleFrom4Points(vertices)
	assert.EqualError(t, err, "four corners required to instantiate this rectangle, the array of points provided was of length 5", "unexpected error received")
	// 3 vertices is too few
	vertices = []point.Point{p[0][0], p[0][1], p[1][1]}
	_, err = polygon.NewValidatedXYRectangleFrom4Points(vertices)
	assert.EqualError(t, err, "four corners required to instantiate this rectangle, the array of points provided was of length 3", "unexpected error received")
	// same point twice should result in error
	vertices = []point.Point{p[0][0], p[0][1], p[1][1], p[0][0]}
	_, err = polygon.NewValidatedXYRectangleFrom4Points(vertices)
	assert.EqualError(t, err, "two of the four points passed are the same point - a: point.Point{X:0, Y:0}, b: point.Point{X:0, Y:0}", "unexpected error received")
	// four points that do not form a rectangle should result in an error
	vertices = []point.Point{p[0][0], p[0][1], p[1][1], p[2][2]}
	_, err = polygon.NewValidatedXYRectangleFrom4Points(vertices)
	assert.EqualError(t, err, "vertex point.Point{X:0, Y:2} in output rectangle was found 0 times in input vertices when it should have appeared once - not a valid XYRectangle", "unexpected error received")
}

// TestContainsPoint - check the function of ContainsPoint
func TestContainsPoint(t *testing.T) {
	p := getTestPoints(5)
	vertices := [4]point.Point{p[1][1], p[1][3], p[3][3], p[3][1]}
	r := &polygon.XYRectangle{Vertices: vertices}
	ok := r.ContainsPoint(p[1][1])
	assert.True(t, ok, "rectangle should contain each vertex")
	ok = r.ContainsPoint(p[1][3])
	assert.True(t, ok, "rectangle should contain each vertex")
	ok = r.ContainsPoint(p[3][1])
	assert.True(t, ok, "rectangle should contain each vertex")
	ok = r.ContainsPoint(p[3][3])
	assert.True(t, ok, "rectangle should contain each vertex")
	ok = r.ContainsPoint(p[2][2])
	assert.True(t, ok, "rectangle should centre")
	ok = r.ContainsPoint(p[0][2])
	assert.False(t, ok, "x value should be too low")
	ok = r.ContainsPoint(p[4][2])
	assert.False(t, ok, "x value should be too high")
	ok = r.ContainsPoint(p[2][0])
	assert.False(t, ok, "y value should be too low")
	ok = r.ContainsPoint(p[2][4])
	assert.False(t, ok, "y value should be too high")
}

// TestIntersectsLineSegment - test function of IntersectsLineSegment
func TestIntersectsLineSegment(t *testing.T) {
	p := getTestPoints(6)
	vertices := [4]point.Point{p[1][1], p[1][4], p[4][4], p[4][1]}
	r := &polygon.XYRectangle{Vertices: vertices}
	// line segment below rectangle
	ls := line.LineSegment{Start: p[0][0], End: p[5][0]}
	points, hit := r.IntersectsLineSegment(ls)
	assert.Len(t, points, 0, "expect zero length intersection point list")
	assert.False(t, hit, "expect miss here")
	// line segment above rectangle
	ls = line.LineSegment{Start: p[0][5], End: p[5][5]}
	points, hit = r.IntersectsLineSegment(ls)
	assert.Len(t, points, 0, "expect zero length intersection point list")
	assert.False(t, hit, "expect miss here")
	// line segment left of rectangle
	ls = line.LineSegment{Start: p[0][0], End: p[0][5]}
	points, hit = r.IntersectsLineSegment(ls)
	assert.Len(t, points, 0, "expect zero length intersection point list")
	assert.False(t, hit, "expect miss here")
	// line segment right of rectangle
	ls = line.LineSegment{Start: p[5][0], End: p[5][5]}
	points, hit = r.IntersectsLineSegment(ls)
	assert.Len(t, points, 0, "expect zero length intersection point list")
	assert.False(t, hit, "expect miss here")
	// line segment with both ends inside rectangle
	ls = line.LineSegment{Start: p[2][2], End: p[3][3]}
	points, hit = r.IntersectsLineSegment(ls)
	assert.Len(t, points, 0, "expect zero length intersection point list")
	assert.False(t, hit, "expect miss here")
	// line segment = bottom edge
	ls = line.LineSegment{Start: p[1][1], End: p[4][1]}
	points, hit = r.IntersectsLineSegment(ls)
	assert.Len(t, points, 2, "expect two length intersection point list")
	assert.True(t, hit, "expect hit here")
	assert.Equal(t, points[0], p[1][1], "expect hit at (1,1)")
	assert.Equal(t, points[1], p[4][1], "expect hit at (4,1)")
	// line segment = top edge
	ls = line.LineSegment{Start: p[1][4], End: p[4][4]}
	points, hit = r.IntersectsLineSegment(ls)
	assert.Len(t, points, 2, "expect two length intersection point list")
	assert.True(t, hit, "expect hit here")
	assert.Equal(t, points[0], p[1][4], "expect hit at (1,4)")
	assert.Equal(t, points[1], p[4][4], "expect hit at (4,4)")
	// line segment = left edge
	ls = line.LineSegment{Start: p[1][1], End: p[1][4]}
	points, hit = r.IntersectsLineSegment(ls)
	assert.Len(t, points, 2, "expect two length intersection point list")
	assert.True(t, hit, "expect hit here")
	assert.Equal(t, points[0], p[1][1], "expect hit at (1,1)")
	assert.Equal(t, points[1], p[1][4], "expect hit at (1,4)")
	// line segment = right edge
	ls = line.LineSegment{Start: p[1][1], End: p[1][4]}
	points, hit = r.IntersectsLineSegment(ls)
	assert.Len(t, points, 2, "expect two length intersection point list")
	assert.True(t, hit, "expect hit here")
	assert.Equal(t, points[0], p[1][1], "expect hit at (1,1)")
	assert.Equal(t, points[1], p[1][4], "expect hit at (1,4)")
	// FURTHER TEST CASES NEEDED HERE
}
