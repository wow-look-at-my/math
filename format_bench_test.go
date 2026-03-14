package math

import "testing"

// Generic vector benchmarks

func BenchmarkTVec2Add(b *testing.B) {
	a := TVec2[float32]{X: 1, Y: 2}
	c := TVec2[float32]{X: 3, Y: 4}
	for b.Loop() {
		a = a.Add(c)
	}
	_ = a
}

func BenchmarkTVec2Sub(b *testing.B) {
	a := TVec2[float32]{X: 1, Y: 2}
	c := TVec2[float32]{X: 3, Y: 4}
	for b.Loop() {
		a = a.Sub(c)
	}
	_ = a
}

func BenchmarkTVec2Scale(b *testing.B) {
	a := TVec2[float32]{X: 1, Y: 2}
	for b.Loop() {
		a = a.Scale(2)
	}
	_ = a
}

func BenchmarkTVec2Dot(b *testing.B) {
	a := TVec2[float32]{X: 1, Y: 2}
	c := TVec2[float32]{X: 3, Y: 4}
	var r float32
	for b.Loop() {
		r = a.Dot(c)
	}
	_ = r
}

func BenchmarkTVec2LenSq(b *testing.B) {
	a := TVec2[float32]{X: 3, Y: 4}
	var r float32
	for b.Loop() {
		r = a.LenSq()
	}
	_ = r
}

func BenchmarkTVec2Eq(b *testing.B) {
	a := TVec2[float32]{X: 1, Y: 2}
	c := TVec2[float32]{X: 1, Y: 2}
	var r bool
	for b.Loop() {
		r = a.Eq(c)
	}
	_ = r
}

func BenchmarkTVec3Add(b *testing.B) {
	a := TVec3[float64]{X: 1, Y: 2, Z: 3}
	c := TVec3[float64]{X: 4, Y: 5, Z: 6}
	for b.Loop() {
		a = a.Add(c)
	}
	_ = a
}

func BenchmarkTVec3Sub(b *testing.B) {
	a := TVec3[float64]{X: 1, Y: 2, Z: 3}
	c := TVec3[float64]{X: 4, Y: 5, Z: 6}
	for b.Loop() {
		a = a.Sub(c)
	}
	_ = a
}

func BenchmarkTVec3Scale(b *testing.B) {
	a := TVec3[float64]{X: 1, Y: 2, Z: 3}
	for b.Loop() {
		a = a.Scale(2)
	}
	_ = a
}

func BenchmarkTVec3Dot(b *testing.B) {
	a := TVec3[float64]{X: 1, Y: 2, Z: 3}
	c := TVec3[float64]{X: 4, Y: 5, Z: 6}
	var r float64
	for b.Loop() {
		r = a.Dot(c)
	}
	_ = r
}

func BenchmarkTVec3Cross(b *testing.B) {
	a := TVec3[float64]{X: 1, Y: 2, Z: 3}
	c := TVec3[float64]{X: 4, Y: 5, Z: 6}
	for b.Loop() {
		a = a.Cross(c)
	}
	_ = a
}

func BenchmarkTVec3LenSq(b *testing.B) {
	a := TVec3[float64]{X: 1, Y: 2, Z: 3}
	var r float64
	for b.Loop() {
		r = a.LenSq()
	}
	_ = r
}

func BenchmarkTVec3Eq(b *testing.B) {
	a := TVec3[float64]{X: 1, Y: 2, Z: 3}
	c := TVec3[float64]{X: 1, Y: 2, Z: 3}
	var r bool
	for b.Loop() {
		r = a.Eq(c)
	}
	_ = r
}

func BenchmarkTVec3XY(b *testing.B) {
	a := TVec3[float64]{X: 1, Y: 2, Z: 3}
	v := a.XY()
	for b.Loop() {
		v = a.XY()
	}
	_ = v
}

func BenchmarkTVec4Add(b *testing.B) {
	a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	c := TVec4[int32]{X: 5, Y: 6, Z: 7, W: 8}
	for b.Loop() {
		a = a.Add(c)
	}
	_ = a
}

func BenchmarkTVec4Sub(b *testing.B) {
	a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	c := TVec4[int32]{X: 5, Y: 6, Z: 7, W: 8}
	for b.Loop() {
		a = a.Sub(c)
	}
	_ = a
}

func BenchmarkTVec4Scale(b *testing.B) {
	a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	for b.Loop() {
		a = a.Scale(2)
	}
	_ = a
}

func BenchmarkTVec4Dot(b *testing.B) {
	a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	c := TVec4[int32]{X: 5, Y: 6, Z: 7, W: 8}
	var r int32
	for b.Loop() {
		r = a.Dot(c)
	}
	_ = r
}

func BenchmarkTVec4LenSq(b *testing.B) {
	a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	var r int32
	for b.Loop() {
		r = a.LenSq()
	}
	_ = r
}

func BenchmarkTVec4Eq(b *testing.B) {
	a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	c := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	var r bool
	for b.Loop() {
		r = a.Eq(c)
	}
	_ = r
}

func BenchmarkTVec4XY(b *testing.B) {
	a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	v := a.XY()
	for b.Loop() {
		v = a.XY()
	}
	_ = v
}

func BenchmarkTVec4XYZ(b *testing.B) {
	a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	v := a.XYZ()
	for b.Loop() {
		v = a.XYZ()
	}
	_ = v
}

// Generic matrix benchmarks

func BenchmarkTMat2Add(b *testing.B) {
	a := NewTMat2[float32](1, 2, 3, 4)
	c := NewTMat2[float32](5, 6, 7, 8)
	for b.Loop() {
		a = a.Add(c)
	}
	_ = a
}

func BenchmarkTMat2Sub(b *testing.B) {
	a := NewTMat2[float32](1, 2, 3, 4)
	c := NewTMat2[float32](5, 6, 7, 8)
	for b.Loop() {
		a = a.Sub(c)
	}
	_ = a
}

func BenchmarkTMat2Scale(b *testing.B) {
	a := NewTMat2[float32](1, 2, 3, 4)
	for b.Loop() {
		a = a.Scale(2)
	}
	_ = a
}

func BenchmarkTMat2Mul(b *testing.B) {
	a := NewTMat2[float32](1, 2, 3, 4)
	c := NewTMat2[float32](5, 6, 7, 8)
	for b.Loop() {
		a = a.Mul(c)
	}
	_ = a
}

func BenchmarkTMat2MulVec2(b *testing.B) {
	m := NewTMat2[float32](1, 2, 3, 4)
	v := TVec2[float32]{X: 5, Y: 6}
	for b.Loop() {
		v = m.MulVec2(v)
	}
	_ = v
}

func BenchmarkTMat2Transpose(b *testing.B) {
	m := NewTMat2[float32](1, 2, 3, 4)
	for b.Loop() {
		m = m.Transpose()
	}
	_ = m
}

func BenchmarkTMat2Det(b *testing.B) {
	m := NewTMat2[float32](1, 2, 3, 4)
	var r float32
	for b.Loop() {
		r = m.Det()
	}
	_ = r
}

func BenchmarkTMat2Eq(b *testing.B) {
	a := NewTMat2[float32](1, 2, 3, 4)
	c := NewTMat2[float32](1, 2, 3, 4)
	var r bool
	for b.Loop() {
		r = a.Eq(c)
	}
	_ = r
}

func BenchmarkTMat3Add(b *testing.B) {
	a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	c := NewTMat3[float64](9, 8, 7, 6, 5, 4, 3, 2, 1)
	for b.Loop() {
		a = a.Add(c)
	}
	_ = a
}

func BenchmarkTMat3Sub(b *testing.B) {
	a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	c := NewTMat3[float64](9, 8, 7, 6, 5, 4, 3, 2, 1)
	for b.Loop() {
		a = a.Sub(c)
	}
	_ = a
}

func BenchmarkTMat3Scale(b *testing.B) {
	a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	for b.Loop() {
		a = a.Scale(2)
	}
	_ = a
}

func BenchmarkTMat3Mul(b *testing.B) {
	a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	c := NewTMat3[float64](9, 8, 7, 6, 5, 4, 3, 2, 1)
	for b.Loop() {
		a = a.Mul(c)
	}
	_ = a
}

func BenchmarkTMat3MulVec3(b *testing.B) {
	m := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	v := TVec3[float64]{X: 1, Y: 2, Z: 3}
	for b.Loop() {
		v = m.MulVec3(v)
	}
	_ = v
}

func BenchmarkTMat3Transpose(b *testing.B) {
	m := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	for b.Loop() {
		m = m.Transpose()
	}
	_ = m
}

func BenchmarkTMat3Det(b *testing.B) {
	m := NewTMat3[float64](1, 2, 3, 0, 1, 4, 5, 6, 0)
	var r float64
	for b.Loop() {
		r = m.Det()
	}
	_ = r
}

func BenchmarkTMat3Eq(b *testing.B) {
	a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	c := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	var r bool
	for b.Loop() {
		r = a.Eq(c)
	}
	_ = r
}

func BenchmarkTMat4Add(b *testing.B) {
	a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	c := NewTMat4[int32](16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	for b.Loop() {
		a = a.Add(c)
	}
	_ = a
}

func BenchmarkTMat4Sub(b *testing.B) {
	a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	c := NewTMat4[int32](16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	for b.Loop() {
		a = a.Sub(c)
	}
	_ = a
}

func BenchmarkTMat4Scale(b *testing.B) {
	a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	for b.Loop() {
		a = a.Scale(2)
	}
	_ = a
}

func BenchmarkTMat4Mul(b *testing.B) {
	a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	c := NewTMat4[int32](16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	for b.Loop() {
		a = a.Mul(c)
	}
	_ = a
}

func BenchmarkTMat4MulVec4(b *testing.B) {
	m := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	v := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	for b.Loop() {
		v = m.MulVec4(v)
	}
	_ = v
}

func BenchmarkTMat4Transpose(b *testing.B) {
	m := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	for b.Loop() {
		m = m.Transpose()
	}
	_ = m
}

func BenchmarkTMat4Det(b *testing.B) {
	m := NewTMat4[int32](1, 0, 0, 1, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 1)
	var r int32
	for b.Loop() {
		r = m.Det()
	}
	_ = r
}

func BenchmarkTMat4Eq(b *testing.B) {
	a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	c := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	var r bool
	for b.Loop() {
		r = a.Eq(c)
	}
	_ = r
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

func BenchmarkScaling(b *testing.B) {
	v := NewVec3(2, 3, 4)
	var m Mat4
	for b.Loop() {
		m = Scaling(v)
	}
	_ = m
}

func BenchmarkRotationX(b *testing.B) {
	var m Mat4
	for b.Loop() {
		m = RotationX(0.785)
	}
	_ = m
}

func BenchmarkRotationY(b *testing.B) {
	var m Mat4
	for b.Loop() {
		m = RotationY(0.785)
	}
	_ = m
}

func BenchmarkRotationZ(b *testing.B) {
	var m Mat4
	for b.Loop() {
		m = RotationZ(0.785)
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

func BenchmarkOrtho(b *testing.B) {
	var m Mat4
	for b.Loop() {
		m = Ortho(-1, 1, -1, 1, 0.1, 100)
	}
	_ = m
}
