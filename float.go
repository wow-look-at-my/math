package math

// Numeric is a constraint for all numeric types.
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

// Float is a constraint for floating-point types.
type Float interface {
	~float32 | ~float64
}
