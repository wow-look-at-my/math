package math

import "math"

func abs[T Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func sqrt[T Float](x T) T {
	return T(math.Sqrt(float64(x)))
}

func cos[T Float](x T) T {
	return T(math.Cos(float64(x)))
}

func sin[T Float](x T) T {
	return T(math.Sin(float64(x)))
}

func tan[T Float](x T) T {
	return T(math.Tan(float64(x)))
}

// Clamp returns x clamped to [min, max].
func Clamp[T Float](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
