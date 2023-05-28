package generic

// CountDeltas provides a function for calculating how many deltas of a given
// size exist in a given range for any kind of numeric value. That is, this is
// equal to the mathematical calculation of floor((a-z)/d).
func CountDeltas[T Number](a, z, delta T) int {
	// ensure we are checking the right order
	if a > z {
		a, z = z, a
	}

	// ensure we have a positive delta
	if delta < 0 {
		delta = -delta
	}

	// works because int == floor for integers > 0
	return int((z - a) / delta)
}
