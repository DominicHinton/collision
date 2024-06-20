package point

// Delta - used for determining near equality of two floats
// if | float a - float b | < delta then floats are near equal
const Delta float32 = 0.000001

const EasyDelta float32 = 100 * Delta

// Abs - absolute value of float32. If n < 0, returns -n. Else returns n
func Abs(n float32) float32 {
	if n < 0 {
		return -n
	}
	return n
}

// AreWithinStatedDelta - determine whether two float values a and b are within delta of eachother
func AreWithinStatedDelta(a, b, delta float32) bool {
	return Abs(a-b) < delta
}

// AreWithinGlobalDelta - determine whether two float values a and b are within global delta of eachother
func AreWithinGlobalDelta(a, b float32) bool {
	return AreWithinStatedDelta(a, b, Delta)
}

// AreWithinEasyDelta - determine whether two float values a and b are within global delta of eachother
func AreWithinEasyDelta(a, b float32) bool {
	return AreWithinStatedDelta(a, b, EasyDelta)
}
