package math

// TVec2 is a 2-component vector.
type TVec2[T Float] [2]T

// Vec2 is a TVec2 of float32.
type Vec2 = TVec2[float32]

func NewVec2(x, y float32) Vec2 {
	return Vec2{x, y}
}

func (v TVec2[T]) X() T { return v[0] }
func (v TVec2[T]) Y() T { return v[1] }

func (a TVec2[T]) Add(b TVec2[T]) TVec2[T] {
	return TVec2[T]{a[0] + b[0], a[1] + b[1]}
}

func (a TVec2[T]) Sub(b TVec2[T]) TVec2[T] {
	return TVec2[T]{a[0] - b[0], a[1] - b[1]}
}

func (v TVec2[T]) Scale(s T) TVec2[T] {
	return TVec2[T]{v[0] * s, v[1] * s}
}

func (a TVec2[T]) Dot(b TVec2[T]) T {
	return a[0]*b[0] + a[1]*b[1]
}

func (v TVec2[T]) Len() T {
	return sqrt(v.Dot(v))
}

func (v TVec2[T]) LenSq() T {
	return v.Dot(v)
}

func (v TVec2[T]) Normalize() TVec2[T] {
	l := v.Len()
	if l == 0 {
		return TVec2[T]{}
	}
	return v.Scale(1.0 / l)
}

func (a TVec2[T]) Lerp(b TVec2[T], t T) TVec2[T] {
	return a.Add(b.Sub(a).Scale(t))
}

func (a TVec2[T]) Dist(b TVec2[T]) T {
	return a.Sub(b).Len()
}

func (a TVec2[T]) Eq(b TVec2[T]) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func (a TVec2[T]) ApproxEq(b TVec2[T], eps T) bool {
	return abs(a[0]-b[0]) <= eps && abs(a[1]-b[1]) <= eps
}
