package polygon

import (
	"collision/line"
	"collision/point"
)

// XYPolygon - defined by a series of points, may also contain edges. Validati
type XYPolygon struct {
	Vertices []point.Point      // array of vertices - must be present
	Edges    []line.LineSegment // array of line segments - need not be present, but will be populated if necessary
}

// ensure interface is implemented
var _ Polygon = &XYPolygon{}

// NewValidatedXYPolygon - returns a pointer to a valid XYPolygon, returns error if not a valid XYPolygon
func NewValidatedXYPolygon(vertices []point.Point) (*XYPolygon, error) {
	order := len(vertices)
	if order < 3 {
		return &XYPolygon{}, DimensionError(order)
	}
	p := &XYPolygon{Vertices: vertices}
	p.PopulateEdges()
	_, _, _, err := p.ValidateXYPolygon()
	if err != nil {
		return &XYPolygon{}, err
	}
	return p, nil
}

// PopulateEdges - use polygon vertices to populate edges
func (p *XYPolygon) PopulateEdges() {
	order := len(p.Vertices)
	p.Edges = make([]line.LineSegment, 0, order)
	for i := 0; i < order-1; i++ {
		p.Edges = append(p.Edges, line.LineSegment{Start: p.Vertices[i], End: p.Vertices[i+1]})
	}
	p.Edges = append(p.Edges, line.LineSegment{Start: p.Vertices[order-1], End: p.Vertices[0]})
}

// ValidatePolygon - check that XYPolygon is valid, returning error only
func (p *XYPolygon) ValidatePolygon() error {
	_, _, _, err := p.ValidateXYPolygon()
	return err
}

// ValidateXYPolygon - check that XYPolygon is valid, if the polygon self-intersects, the first pair of lines
// and intersection point found are returned with the error.
func (p *XYPolygon) ValidateXYPolygon() (segment1, segment2 line.LineSegment, intersectionPoint point.Point, validationErr error) {
	// check dimensions
	order := len(p.Vertices)
	if order < 3 {
		validationErr = DimensionError(order)
		return
	}
	if len(p.Edges) != order {
		p.PopulateEdges()
	}

	var intersects bool
	segment1 = p.Edges[0]
	// first edge is dealt with before looping through other line segments as the first and last
	// edges will intersect at first vertex.
	// CONSIDER CASE OF THREE POINTS IN A STRAIGHT LINE CONSTRUED AS TRIANGLE
	for j := 2; j < order-1; j++ {
		segment2 = p.Edges[j]
		intersectionPoint, intersects = segment1.IntersectsLineSegment(segment2)
		if intersects {
			validationErr = IntersectionError(segment1, segment2, intersectionPoint)
			return
		}
	}
	// check for intersection of other edges
	for i := 1; i < order; i++ {
		// select current edge for comparison with remaining edges
		segment1 = p.Edges[i]
		// check for intersection: this must occur at shared vertices so do not check next edge.
		for j := i + 2; j < order; j++ {
			segment2 = p.Edges[j]
			intersectionPoint, intersects = segment1.IntersectsLineSegment(segment2)
			if intersects {
				validationErr = IntersectionError(segment1, segment2, intersectionPoint)
				return
			}
		}
	}
	// if here, no intersections outside shared vertices have been found
	return line.LineSegment{}, line.LineSegment{}, point.Point{}, nil
}
