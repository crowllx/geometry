package geometry

func min(a, b float32) float32 {
	if a <= b {
		return a
	}

	return b
}

func max(a, b float32) float32 {
	if a >= b {
		return a
	}

	return b
}
