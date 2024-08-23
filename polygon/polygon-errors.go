package polygon

import (
	"collision/line"
	"collision/point"
	"fmt"
)

func DimensionError(dimensions int) error {
	return fmt.Errorf("require at least 3 points, the array of points provided was of length %d", dimensions)
}

func IntersectionError(ls1, ls2 line.LineSegment, pt point.Point) error {
	return fmt.Errorf("polygon lines intersect: line %#v and line %#v intersect at point %#v", ls1, ls2, pt)
}

func RectangleDimensionError(dimensions int) error {
	return fmt.Errorf("four corners required to instantiate this rectangle, the array of points provided was of length %d", dimensions)
}

func OppositeCornersXYRectangleDimensionError(dimensions int) error {
	return fmt.Errorf("instantiating rectangle from two opposite corners requires exactly two points to be passed, the array of points was of length %d", dimensions)
}

func OppositeCornersXYRectangleSameXError(x float32) error {
	return fmt.Errorf("both corners passed shared the same x value: %v", x)
}

func OppositeCornersXYRectangleSameYError(y float32) error {
	return fmt.Errorf("both corners passed shared the same y value: %v", y)
}

func PointsAreTouchingError(a, b point.Point) error {
	return fmt.Errorf("two of the four points passed are the same point - a: %#v, b: %#v", a, b)
}

func SharedMinMaxError(min, max point.Point) error {
	return fmt.Errorf("min point %#v and max point %#v share and x or y value", min, max)
}

func PointNotOnMaxOrMinXYRectangleError(p, min, max point.Point) error {
	return fmt.Errorf("found point: %#v not sharing a value with min %#v or max %#v", p, min, max)
}

func TwoMinTwoMaxRequiredError(vertices []point.Point) error {
	return fmt.Errorf("two min and two max x and y values exactly were not for vertices: %#v", vertices)
}

func VertexCountError(vertex point.Point, hitCount int) error {
	return fmt.Errorf("vertex %#v in output rectangle was found %d times in input vertices when it should have appeared once - not a valid XYRectangle", vertex, hitCount)
}
