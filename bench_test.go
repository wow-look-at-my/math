package math

import "testing"

// Vec2 benchmarks

func BenchmarkVec2Add(b *testing.B) {
	a := NewVec2(1, 2)
	c := NewVec2(3, 4)
	for b.Loop() {
		a = a.Add(c)
	}
	_ = a
}

func BenchmarkVec2Dot(b *testing.B) {
	a := NewVec2(1, 2)
	c := NewVec2(3, 4)
	var r float32
	for b.Loop() {
		r = a.Dot(c)
	}
	_ = r
}

func BenchmarkVec2Normalize(b *testing.B) {
	v := NewVec2(3, 4)
	for b.Loop() {
		v = v.Normalize()
	}
	_ = v
}

func BenchmarkVec2Len(b *testing.B) {
	v := NewVec2(3, 4)
	var r float32
	for b.Loop() {
		r = v.Len()
	}
	_ = r
}

// Vec3 benchmarks

func BenchmarkVec3Add(b *testing.B) {
	a := NewVec3(1, 2, 3)
	c := NewVec3(4, 5, 6)
	for b.Loop() {
		a = a.Add(c)
	}
	_ = a
}

func BenchmarkVec3Dot(b *testing.B) {
	a := NewVec3(1, 2, 3)
	c := NewVec3(4, 5, 6)
	var r float32
	for b.Loop() {
		r = a.Dot(c)
	}
	_ = r
}

func BenchmarkVec3Cross(b *testing.B) {
	a := NewVec3(1, 2, 3)
	c := NewVec3(4, 5, 6)
	for b.Loop() {
		a = a.Cross(c)
	}
	_ = a
}

func BenchmarkVec3Normalize(b *testing.B) {
	v := NewVec3(1, 2, 3)
	for b.Loop() {
		v = v.Normalize()
	}
	_ = v
}

// Vec4 benchmarks

func BenchmarkVec4Add(b *testing.B) {
	a := NewVec4(1, 2, 3, 4)
	c := NewVec4(5, 6, 7, 8)
	for b.Loop() {
		a = a.Add(c)
	}
	_ = a
}

func BenchmarkVec4Dot(b *testing.B) {
	a := NewVec4(1, 2, 3, 4)
	c := NewVec4(5, 6, 7, 8)
	var r float32
	for b.Loop() {
		r = a.Dot(c)
	}
	_ = r
}

func BenchmarkVec4Normalize(b *testing.B) {
	v := NewVec4(1, 2, 3, 4)
	for b.Loop() {
		v = v.Normalize()
	}
	_ = v
}

// Mat2 benchmarks

func BenchmarkMat2Mul(b *testing.B) {
	a := NewMat2(1, 2, 3, 4)
	c := NewMat2(5, 6, 7, 8)
	for b.Loop() {
		a = a.Mul(c)
	}
	_ = a
}

func BenchmarkMat2MulVec2(b *testing.B) {
	m := NewMat2(1, 2, 3, 4)
	v := NewVec2(5, 6)
	for b.Loop() {
		v = m.MulVec2(v)
	}
	_ = v
}

func BenchmarkMat2Inverse(b *testing.B) {
	m := NewMat2(1, 2, 3, 4)
	for b.Loop() {
		m = m.Inverse()
	}
	_ = m
}

func BenchmarkMat2Det(b *testing.B) {
	m := NewMat2(1, 2, 3, 4)
	var r float32
	for b.Loop() {
		r = m.Det()
	}
	_ = r
}

// Mat3 benchmarks

func BenchmarkMat3Mul(b *testing.B) {
	a := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	c := NewMat3(9, 8, 7, 6, 5, 4, 3, 2, 1)
	for b.Loop() {
		a = a.Mul(c)
	}
	_ = a
}

func BenchmarkMat3MulVec3(b *testing.B) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	v := NewVec3(1, 2, 3)
	for b.Loop() {
		v = m.MulVec3(v)
	}
	_ = v
}

func BenchmarkMat3Inverse(b *testing.B) {
	m := NewMat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	for b.Loop() {
		m = m.Inverse()
	}
	_ = m
}

func BenchmarkMat3Det(b *testing.B) {
	m := NewMat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	var r float32
	for b.Loop() {
		r = m.Det()
	}
	_ = r
}

func BenchmarkMat3Transpose(b *testing.B) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	for b.Loop() {
		m = m.Transpose()
	}
	_ = m
}

// Mat4 benchmarks

func BenchmarkMat4Mul(b *testing.B) {
	a := Mat4Identity()
	c := Translation(NewVec3(1, 2, 3))
	for b.Loop() {
		a = a.Mul(c)
	}
	_ = a
}

func BenchmarkMat4MulVec4(b *testing.B) {
	m := Translation(NewVec3(1, 2, 3))
	v := NewVec4(1, 2, 3, 1)
	for b.Loop() {
		v = m.MulVec4(v)
	}
	_ = v
}

func BenchmarkMat4Inverse(b *testing.B) {
	m := NewMat4(
		1, 0, 0, 1,
		0, 2, 0, 0,
		0, 0, 3, 0,
		0, 0, 0, 1,
	)
	for b.Loop() {
		m = m.Inverse()
	}
	_ = m
}

func BenchmarkMat4Det(b *testing.B) {
	m := Translation(NewVec3(1, 2, 3))
	var r float32
	for b.Loop() {
		r = m.Det()
	}
	_ = r
}

func BenchmarkMat4Transpose(b *testing.B) {
	m := Mat4Identity()
	for b.Loop() {
		m = m.Transpose()
	}
	_ = m
}

// Transform benchmarks

func BenchmarkTranslation(b *testing.B) {
	v := NewVec3(1, 2, 3)
	var m Mat4
	for b.Loop() {
		m = Translation(v)
	}
	_ = m
}

func BenchmarkLookAt(b *testing.B) {
	eye := NewVec3(0, 0, 5)
	center := NewVec3(0, 0, 0)
	up := NewVec3(0, 1, 0)
	var m Mat4
	for b.Loop() {
		m = LookAt(eye, center, up)
	}
	_ = m
}

func BenchmarkPerspective(b *testing.B) {
	var m Mat4
	for b.Loop() {
		m = Perspective(0.785, 1.6, 0.1, 100)
	}
	_ = m
}
