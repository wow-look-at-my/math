package math

import "testing"

func BenchmarkGenericOps(b *testing.B) {
	va2 := TVec2[float32]{X: 1, Y: 2}
	vb2 := TVec2[float32]{X: 3, Y: 4}
	va3 := TVec3[float64]{X: 1, Y: 2, Z: 3}
	vb3 := TVec3[float64]{X: 4, Y: 5, Z: 6}
	va4 := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	vb4 := TVec4[int32]{X: 5, Y: 6, Z: 7, W: 8}
	ma2 := NewTMat2[float32](1, 2, 3, 4)
	mb2 := NewTMat2[float32](5, 6, 7, 8)
	ma3 := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	mb3 := NewTMat3[float64](9, 8, 7, 6, 5, 4, 3, 2, 1)
	ma4 := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	mb4 := NewTMat4[int32](16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	var vd2 float32
	var vd3 float64
	var vd4 int32
	var veq2, veq3, veq4 bool
	var meq2, meq3, meq4 bool
	for b.Loop() {
		va2 = va2.Add(vb2)
		va2 = va2.Sub(vb2)
		va2 = va2.Scale(2)
		vd2 = va2.Dot(vb2)
		vd2 = va2.LenSq()
		veq2 = va2.Eq(vb2)
		va3 = va3.Add(vb3)
		va3 = va3.Sub(vb3)
		va3 = va3.Scale(2)
		vd3 = va3.Dot(vb3)
		vd3 = va3.LenSq()
		va3 = va3.Cross(vb3)
		veq3 = va3.Eq(vb3)
		_ = va3.XY()
		va4 = va4.Add(vb4)
		va4 = va4.Sub(vb4)
		va4 = va4.Scale(2)
		vd4 = va4.Dot(vb4)
		vd4 = va4.LenSq()
		veq4 = va4.Eq(vb4)
		_ = va4.XY()
		_ = va4.XYZ()
		ma2 = ma2.Add(mb2)
		ma2 = ma2.Sub(mb2)
		ma2 = ma2.Scale(2)
		ma2 = ma2.Transpose()
		meq2 = ma2.Eq(mb2)
		ma3 = ma3.Add(mb3)
		ma3 = ma3.Sub(mb3)
		ma3 = ma3.Scale(2)
		ma3 = ma3.Transpose()
		meq3 = ma3.Eq(mb3)
		ma4 = ma4.Add(mb4)
		ma4 = ma4.Sub(mb4)
		ma4 = ma4.Scale(2)
		ma4 = ma4.Transpose()
		meq4 = ma4.Eq(mb4)
	}
	_ = va2
	_ = vd2
	_ = veq2
	_ = va3
	_ = vd3
	_ = veq3
	_ = va4
	_ = vd4
	_ = veq4
	_ = ma2
	_ = meq2
	_ = ma3
	_ = meq3
	_ = ma4
	_ = meq4
}

func BenchmarkGenericMul(b *testing.B) {
	ma2 := NewTMat2[float32](1, 2, 3, 4)
	mb2 := NewTMat2[float32](5, 6, 7, 8)
	v2 := TVec2[float32]{X: 1, Y: 2}
	ma3 := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
	mb3 := NewTMat3[float64](9, 8, 7, 6, 5, 4, 3, 2, 1)
	v3 := TVec3[float64]{X: 1, Y: 2, Z: 3}
	ma4 := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	mb4 := NewTMat4[int32](16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	v4 := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
	for b.Loop() {
		ma2 = ma2.Mul(mb2)
		v2 = ma2.MulVec2(v2)
		ma3 = ma3.Mul(mb3)
		v3 = ma3.MulVec3(v3)
		ma4 = ma4.Mul(mb4)
		v4 = ma4.MulVec4(v4)
	}
	_ = ma2
	_ = v2
	_ = ma3
	_ = v3
	_ = ma4
	_ = v4
}

func BenchmarkGenericDet(b *testing.B) {
	m2 := NewTMat2[float32](1, 2, 3, 4)
	m3 := NewTMat3[float64](1, 2, 3, 0, 1, 4, 5, 6, 0)
	m4 := NewTMat4[int32](1, 0, 0, 1, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 1)
	var d2 float32
	var d3 float64
	var d4 int32
	for b.Loop() {
		d2 = m2.Det()
		d3 = m3.Det()
		d4 = m4.Det()
	}
	_ = d2
	_ = d3
	_ = d4
}

func BenchmarkTransform(b *testing.B) {
	tv := NewVec3(1, 2, 3)
	sv := NewVec3(2, 3, 4)
	eye := NewVec3(0, 0, 5)
	center := NewVec3(0, 0, 0)
	up := NewVec3(0, 1, 0)
	var m Mat4
	for b.Loop() {
		m = Translation(tv)
		m = Scaling(sv)
		m = RotationX(0.785)
		m = RotationY(0.785)
		m = RotationZ(0.785)
		m = LookAt(eye, center, up)
		m = Perspective(0.785, 1.6, 0.1, 100)
		m = Ortho(-1, 1, -1, 1, 0.1, 100)
	}
	_ = m
}
