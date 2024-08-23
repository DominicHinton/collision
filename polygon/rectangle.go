package polygon

import (
	"collision/line"
	"collision/point"
	"fmt"
)

// XYRectangle - defined by four points, this is strictly defined as a rectangle where the edges
// are all horizontal or vertical in XY space (i.e. if a valid XYRectangle is rotated by a non-integer
// multiple of 90 degrees in XY space then it is no longer valid). Furthermore, the rectangle will take
// form ABCD where A is the bottom left vertex, B top left, C top right and D the bottom right.
// This is to facilate very fast calculation of interactions between XYRectangles. The XYPolygon struct
// can be used to represent all rectangles.
type XYRectangle struct {
	Vertices [4]point.Point     // array of vertices - must be present
	Edges    []line.LineSegment // array of edges
}

// ensure interface implemented
var _ Polygon = &XYRectangle{}

// ValidatePolygon - validate that XYRectangle conforms to required specification
func (r *XYRectangle) ValidatePolygon() error {
	return r.validateXYRectangle()
}

// ContainsPoint - boolean indicating whether or not an XYRectangle contains a point
func (r *XYRectangle) ContainsPoint(p point.Point) bool {
	minX, maxX, minY, maxY, err := point.GetMinMax(r.Vertices[:])
	switch {
	case err != nil:
		return false
	case minX > p.X:
		return false
	case maxX < p.X:
		return false
	case minY > p.Y:
		return false
	case maxY < p.Y:
		return false
	default:
		return true
	}
}

// IntersectsLineSegment - array of intersection points and boolean indicating whether an XYRectangle intersects a line
func (r *XYRectangle) IntersectsLineSegment(l line.LineSegment) ([]point.Point, bool) {
	hit := false
	intersections := make([]point.Point, 0, 2) // a line segment touching two vertices would hit all four lines, so de-duplication is required
	if len(r.Edges) != 4 {
		r.PopulateEdges()
	}
	for _, edge := range r.Edges {
		newIntersection, ok := edge.IntersectsLineSegment(l)
		if ok {
			hit = true
			duplicate := false
			for _, recordedIntersection := range intersections {
				if newIntersection.AreTouching(recordedIntersection) {
					duplicate = true
					break
				}
			}
			if !duplicate {
				intersections = append(intersections, newIntersection)
			}
		}
	}
	return intersections, hit
}

// PopulateEdges - populate the edges of an XYRectange
func (r *XYRectangle) PopulateEdges() {
	r.Edges = make([]line.LineSegment, 4)
	a, b, c, d := r.Vertices[0], r.Vertices[1], r.Vertices[2], r.Vertices[3]
	ab := line.LineSegment{Start: a, End: b}
	bc := line.LineSegment{Start: b, End: c}
	cd := line.LineSegment{Start: c, End: d}
	da := line.LineSegment{Start: d, End: a}
	r.Edges[0] = ab
	r.Edges[1] = bc
	r.Edges[2] = cd
	r.Edges[3] = da
}

// validateXYRectangle - validate XYRectangle
func (r *XYRectangle) validateXYRectangle() error {
	vertices := make([]point.Point, 0, 4)
	for _, v := range r.Vertices {
		vertices = append(vertices, v)
	}
	if len(vertices) != 4 {
		RectangleDimensionError(len(vertices))
	}
	// loop through points to confirm none are the same point
	for i := 0; i < 4; i++ {
		f := vertices[i]
		for j := i + i; j < 4; j++ {
			g := vertices[j]
			if f.AreTouching(g) {
				PointsAreTouchingError(f, g)
			}
		}
	}
	// get min and max x and y values
	minX, maxX, minY, maxY, err := point.GetMinMax(vertices)
	if err != nil {
		return err
	}
	// loop through vertices to ensure min and max x and y values are only values contained in array
	// and each value is used twice
	minPoint := point.Point{X: minX, Y: minY}
	maxPoint := point.Point{X: maxX, Y: maxY}
	if minPoint.SameX(maxPoint) || minPoint.SameY(maxPoint) {
		return SharedMinMaxError(minPoint, maxPoint)
	}
	minXCount, maxXCount, minYCount, maxYCount := 0, 0, 0, 0
	for _, p := range vertices {
		xfound, yFound := false, false
		// check x coordinate
		if p.SameX(minPoint) {
			xfound = true
			minXCount++
		} else if p.SameX(maxPoint) {
			xfound = true
			maxXCount++
		}
		// check y coordinate
		if p.SameY(minPoint) {
			yFound = true
			minYCount++
		} else if p.SameY(maxPoint) {
			yFound = true
			maxYCount++
		}
		if !xfound || !yFound {
			PointNotOnMaxOrMinXYRectangleError(p, minPoint, maxPoint)
		}
	}
	fmt.Println(minXCount, maxXCount, minYCount, maxYCount)
	if minXCount != 2 || maxXCount != 2 || minYCount != 2 || maxYCount != 2 {
		return TwoMinTwoMaxRequiredError(vertices)
	}
	return nil
}

// NewValidatedXYRectangleFrom4Points - returns pointer to validated XYRectangle if points list is valid, error otherwise
func NewValidatedXYRectangleFrom4Points(vertices []point.Point) (*XYRectangle, error) {
	if len(vertices) != 4 {
		return &XYRectangle{}, RectangleDimensionError(len(vertices))
	}
	// ensure no vertices are duplicates
	for i := 0; i < 4; i++ {
		p1 := vertices[i]
		for j := i + 1; j < 4; j++ {
			p2 := vertices[j]
			if p1.AreTouching(p2) {
				return &XYRectangle{}, PointsAreTouchingError(p1, p2)
			}
		}
	}

	minX, maxX, minY, maxY, err := point.GetMinMax(vertices)
	if err != nil {
		return &XYRectangle{}, err
	}

	// get a rectangle from min and max x y values
	r := newRectangleFromMinMax(minX, maxX, minY, maxY)

	// validate that rectangle corners match the vertices provided
	hitCount := []int{0, 0, 0, 0}
	// each of the four produced vertices must match exactly one of the vertices passed to the function
	for i, vertexOut := range r.Vertices {
		for _, vertexIn := range vertices {
			if vertexOut.AreTouching(vertexIn) {
				hitCount[i]++
			}
		}
	}
	for i, count := range hitCount {
		if count != 1 {
			return &XYRectangle{}, VertexCountError(r.Vertices[i], count)
		}
	}

	// check this is a valid rectangle
	err = r.validateXYRectangle()
	if err != nil {
		return &XYRectangle{}, err
	}

	return r, nil
}

// NewValidatedXYRectangleFromOppositeVertices - returns pointer to validated XYRectangle if the two vertices share neither an X or a Y value
func NewValidatedXYRectangleFromOppositeVertices(vertices []point.Point) (*XYRectangle, error) {
	if len(vertices) != 2 {
		return &XYRectangle{}, OppositeCornersXYRectangleDimensionError(len(vertices))
	}
	f, g := vertices[0], vertices[1]
	if f.SameX(g) {
		return &XYRectangle{}, OppositeCornersXYRectangleSameXError(f.X)
	}
	if f.SameY(g) {
		return &XYRectangle{}, OppositeCornersXYRectangleSameYError(f.Y)
	}
	minX, maxX, minY, maxY, err := point.GetMinMax(vertices)
	if err != nil {
		return &XYRectangle{}, err
	}

	return newRectangleFromMinMax(minX, maxX, minY, maxY), nil
}

// newRectangleFromMinMax - given miniumum and maximum x and y, return XYRectangle pointer
func newRectangleFromMinMax(minX, maxX, minY, maxY float32) *XYRectangle {
	a := point.Point{X: minX, Y: minY}
	b := point.Point{X: minX, Y: maxY}
	c := point.Point{X: maxX, Y: maxY}
	d := point.Point{X: maxX, Y: minY}

	return &XYRectangle{Vertices: [4]point.Point{a, b, c, d}}
}
