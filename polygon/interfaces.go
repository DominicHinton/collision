package polygon

// Polygon - interface
type Polygon interface {
	// ValidatePolygon - verify validity of concrete polygon implementation
	ValidatePolygon() (validationErr error)
}
