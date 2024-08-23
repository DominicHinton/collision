package polygon_test

import (
	"collision/line"
	"collision/point"
	"collision/polygon"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// getTestPoints - returns points in array for ease of testing
func getTestPoints(size int) [][]point.Point {
	testPoints := make([][]point.Point, size)
	for i := 0; i < size; i++ {
		testPoints[i] = make([]point.Point, size)
	}
	for i := 0; i < size; i++ {
		x := float32(i)
		for j := 0; j < size; j++ {
			y := float32(j)
			testPoints[i][j] = point.Point{X: x, Y: y}
		}
	}
	return testPoints
}

// TestPopulateEdges - test correct functioning of PopulateEdges
func TestPopulateEdges(t *testing.T) {
	a := point.Point{X: 0, Y: 0}
	b := point.Point{X: 0, Y: 10}
	c := point.Point{X: 10, Y: 10}
	d := point.Point{X: 10, Y: 0}

	p := polygon.XYPolygon{
		Vertices: []point.Point{a, b, c, d},
	}
	expectedEdges := []line.LineSegment{{Start: a, End: b}, {Start: b, End: c}, {Start: c, End: d}, {Start: d, End: a}}
	p.PopulateEdges()
	for i, edge := range p.Edges {
		fmt.Printf("%d : expectedEdge = %#v, actualEdge = %#v\n\n", i, expectedEdges[i], edge)
		assert.Equal(t, expectedEdges[i], edge)
	}
}

// TestValidatePolygon - test ValidatePolygon behaves as expected
func TestValidatePolygon(t *testing.T) {
	c := getTestPoints(10)

	// one point should not allow valid polygon to be constructed
	p := polygon.XYPolygon{Vertices: []point.Point{c[0][0]}}
	nilLineSegment := line.LineSegment{}
	nilPoint := point.Point{}
	ls1, ls2, pt, err := p.ValidateXYPolygon()
	expectedErr := "require at least 3 points, the array of points provided was of length 1"
	assert.EqualError(t, err, expectedErr, "one vertex polgon should return dimensionality error")
	assert.Equal(t, nilLineSegment, ls1, "expected nil line segment")
	assert.Equal(t, nilLineSegment, ls2, "expected nil line segment")
	assert.Equal(t, nilPoint, pt, "expected nil point")

	// two points should not allow valid polygon to be constructed
	p = polygon.XYPolygon{Vertices: []point.Point{c[0][0], c[1][0]}}
	ls1, ls2, pt, err = p.ValidateXYPolygon()
	expectedErr = "require at least 3 points, the array of points provided was of length 2"
	assert.EqualError(t, err, expectedErr, "one vertex polgon should return dimensionality error")
	assert.Equal(t, nilLineSegment, ls1, "expected nil line segment")
	assert.Equal(t, nilLineSegment, ls2, "expected nil line segment")
	assert.Equal(t, nilPoint, pt, "expected nil point")

	// IN CURRENT IMPLEMENTATION, THREE POINTS SHOULD ALWAYS VALIDATE. THIS MAY BE ALTERED LATER.
	p = polygon.XYPolygon{Vertices: []point.Point{c[0][0], c[1][0], c[1][1]}}
	ls1, ls2, pt, err = p.ValidateXYPolygon()
	assert.Nil(t, err, "expect triangle to be validated")
	assert.Equal(t, nilLineSegment, ls1, "expected nil line segment")
	assert.Equal(t, nilLineSegment, ls2, "expected nil line segment")
	assert.Equal(t, nilPoint, pt, "expected nil point")

	// unit square should be valid
	p = polygon.XYPolygon{Vertices: []point.Point{c[0][0], c[1][0], c[1][1], c[0][1]}}
	ls1, ls2, pt, err = p.ValidateXYPolygon()
	assert.Nil(t, err, "expect unit square to be validated")
	assert.Equal(t, nilLineSegment, ls1, "expected nil line segment")
	assert.Equal(t, nilLineSegment, ls2, "expected nil line segment")
	assert.Equal(t, nilPoint, pt, "expected nil point")

	// quadrilateral with one intersection should fail
	p = polygon.XYPolygon{Vertices: []point.Point{c[0][1], c[1][1], c[1][2], c[0][0]}}
	ls1, ls2, pt, err = p.ValidateXYPolygon()
	assert.EqualError(t, err, "polygon lines intersect: line line.LineSegment{Start:point.Point{X:0, Y:1}, End:point.Point{X:1, Y:1}} and line line.LineSegment{Start:point.Point{X:1, Y:2}, End:point.Point{X:0, Y:0}} intersect at point point.Point{X:0.5, Y:1}", "unexpected error received")
	assert.Equal(t, line.LineSegment{Start: point.Point{X: 0, Y: 1}, End: point.Point{X: 1, Y: 1}}, ls1, "unexepcted line segment returned")
	assert.Equal(t, line.LineSegment{Start: point.Point{X: 1, Y: 2}, End: point.Point{X: 0, Y: 0}}, ls2, "unexepcted line segment returned")
	assert.Equal(t, point.Point{X: 0.5, Y: 1}, pt, "unexpected intersection point returned")
}
