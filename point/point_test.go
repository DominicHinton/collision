package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDelta - assertion that Delta is the value assumed by rest of testing suite
func TestDelta(t *testing.T) {
	var assumedDelta float32 = 0.000001
	assert.Equal(t, assumedDelta, Delta, "the value of Delta is not the value expected by test suite")
}

// TestAbs - check expected behaviour of abs function
func TestAbs(t *testing.T) {
	var positive float32 = 3.3
	var negative float32 = -3.3
	var zero float32 = 0

	positiveAbs := Abs(positive)
	assert.Equal(t, positiveAbs, positive, "expect value of positive number to be unchanged")

	negativeAbs := Abs(negative)
	assert.Equal(t, negativeAbs, positive, "expect  negative number to changed to positive of same abs value")

	zeroAbs := Abs(zero)
	assert.Equal(t, zeroAbs, zero, "expect zero to be unchanged by abs function")
}

// TestAreWithinStatedDelta - check expected behaviour of areWithinStatedDelta
func TestAreWithinStatedDelta(t *testing.T) {
	var a float32 = 3.55
	var b float32 = 3.56
	var delta float32 = 1

	ok := AreWithinStatedDelta(a, b, delta)
	assert.True(t, ok, "3.55 and 3.56 are within delta of 1.00")

	delta = 0.1
	ok = AreWithinStatedDelta(a, b, delta)
	assert.True(t, ok, "3.55 and 3.56 are within delta of 0.1")

	delta = 0.01001
	ok = AreWithinStatedDelta(a, b, delta)
	assert.True(t, ok, "3.55 and 3.56 are within delta of 0.1001")

	delta = 0.00999
	ok = AreWithinStatedDelta(a, b, delta)
	assert.False(t, ok, "3.55 and 3.56 are not within delta of 0.0099")

	delta = 0.001
	ok = AreWithinStatedDelta(a, b, delta)
	assert.False(t, ok, "3.55 and 3.56 are not within delta of 0.001")
}

// TestAreWithinGlobalDelta - check functioning of areWithinGlobalDelta
func TestAreWithinGlobalDelta(t *testing.T) {
	var b float32 = 3.56
	var c float32 = 3.56001
	var d float32 = 3.560002
	var e float32 = 3.5600011
	var f float32 = 3.560001
	var g float32 = 3.5600019
	ok := AreWithinGlobalDelta(b, c)
	assert.False(t, ok, "values should not be within delta")

	ok = AreWithinGlobalDelta(b, d)
	assert.False(t, ok, "values should not be within delta")

	ok = AreWithinGlobalDelta(b, e)
	assert.False(t, ok, "values should not be within delta")

	ok = AreWithinGlobalDelta(b, f)
	assert.True(t, ok, "values should be within delta")

	ok = AreWithinGlobalDelta(f, g)
	assert.True(t, ok, "values should be within delta")
}

// TestAreWithinEasyDelta - check functioning of areWithinGlobalDelta
func TestAreWithinEasyDelta(t *testing.T) {
	var b float32 = 3.56
	var c float32 = 3.561
	var d float32 = 3.5602
	var e float32 = 3.56011
	var f float32 = 3.560009
	var g float32 = 3.560019
	ok := AreWithinEasyDelta(b, c)
	assert.False(t, ok, "values should not be within delta")

	ok = AreWithinEasyDelta(b, d)
	assert.False(t, ok, "values should not be within delta")

	ok = AreWithinEasyDelta(b, e)
	assert.False(t, ok, "values should not be within delta")

	ok = AreWithinEasyDelta(b, f)
	assert.True(t, ok, "values should be within delta")

	ok = AreWithinEasyDelta(f, g)
	assert.True(t, ok, "values should be within delta")
}

// TestSameX - test that SameX behaves as expected
func TestSameX(t *testing.T) {
	var a Point = Point{X: 1, Y: 1}
	var b Point = Point{X: 1, Y: 0}
	var c Point = Point{X: 0.999, Y: 1}

	ok := a.SameX(b)
	assert.True(t, ok, "a and b should share same x value")

	ok = a.SameX(c)
	assert.False(t, ok, "a and c should not share same x value")
}

// TestSameY - test that SameY behaves as expected
func TestSameY(t *testing.T) {
	var a Point = Point{X: 1, Y: 1}
	var b Point = Point{X: 0, Y: 1}
	var c Point = Point{X: 1, Y: 0.999}

	ok := a.SameY(b)
	assert.True(t, ok, "a and b should share same y value")

	ok = a.SameY(c)
	assert.False(t, ok, "a and c should not share same y value")
}

// TestAreTouching - test that AreTouching behaves as expected
func TestAreTouching(t *testing.T) {
	var a Point = Point{X: 1, Y: 1}
	var b Point = Point{X: 1, Y: 1}
	var c Point = Point{X: 0.9999995, Y: 0.9999995}
	var d Point = Point{X: 1.000001, Y: 1.000001}
	var e Point = Point{X: 0.999995, Y: 0.999995}
	var f Point = Point{X: 1.00001, Y: 1.00001}
	var g Point = Point{X: 2, Y: 2}
	var h Point = Point{X: 1 + 0.9*Delta, Y: 1}
	var i Point = Point{X: 1, Y: 1 + 0.9*Delta}
	var j Point = Point{X: 1 + 1.1*Delta, Y: 1}
	var k Point = Point{X: 1, Y: 1 + 1.1*Delta}

	ok := a.AreTouching(b)
	assert.True(t, ok, "a and b should be touching")

	ok = a.AreTouching(c)
	assert.True(t, ok, "a and c should be touching")

	ok = a.AreTouching(d)
	assert.True(t, ok, "a and d should be touching")

	ok = a.AreTouching(e)
	assert.False(t, ok, "a and e should not be touching")

	ok = a.AreTouching(f)
	assert.False(t, ok, "a and f should not be touching")

	ok = a.AreTouching(g)
	assert.False(t, ok, "a and g should not be touching")

	ok = a.AreTouching(h)
	assert.True(t, ok, "a and h should be touching")

	ok = a.AreTouching(i)
	assert.True(t, ok, "a and i should be touching")

	ok = a.AreTouching(j)
	assert.False(t, ok, "a and j should not be touching")

	ok = a.AreTouching(k)
	assert.False(t, ok, "a and k should not be touching")
}

// TestXDistance - test that X Distance behaves as expected
func TestXDistance(t *testing.T) {
	var a Point = Point{1, 1}
	var b Point = Point{11, 5}
	var c Point = Point{-3, -8}

	dist := a.XDistance(b)
	assert.Equal(t, float32(-10), dist)

	dist = b.XDistance(a)
	assert.Equal(t, float32(10), dist)

	dist = a.XDistance(c)
	assert.Equal(t, float32(4), dist)

	dist = c.XDistance(a)
	assert.Equal(t, float32(-4), dist)
}

// TestYDistance - test that Y Distance behaves as expected
func TestYDistance(t *testing.T) {
	var a Point = Point{1, 1}
	var b Point = Point{11, 5}
	var c Point = Point{-3, -8}

	dist := a.YDistance(b)
	assert.Equal(t, float32(-4), dist)

	dist = b.YDistance(a)
	assert.Equal(t, float32(4), dist)

	dist = a.YDistance(c)
	assert.Equal(t, float32(9), dist)

	dist = c.YDistance(a)
	assert.Equal(t, float32(-9), dist)
}

// TestDistance - test that Distance behaves as expected
func TestDistance(t *testing.T) {
	var a Point = Point{0, 0}
	var b Point = Point{3, 4}
	var c Point = Point{-5, -12}

	dist := a.Distance(b)
	assert.Equal(t, float32(5), dist)

	dist = b.Distance(a)
	assert.Equal(t, float32(5), dist)

	dist = a.Distance(c)
	assert.Equal(t, float32(13), dist)

	dist = c.Distance(a)
	assert.Equal(t, float32(13), dist)
}
