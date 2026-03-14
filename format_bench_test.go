package math

import "testing"

func BenchmarkGenericOps(b *testing.B) {
	b.Run("Vec2/Add", func(b *testing.B) {
		a := TVec2[float32]{X: 1, Y: 2}
		c := TVec2[float32]{X: 3, Y: 4}
		for b.Loop() { a = a.Add(c) }
		_ = a
	})
	b.Run("Vec2/Sub", func(b *testing.B) {
		a := TVec2[float32]{X: 1, Y: 2}
		c := TVec2[float32]{X: 3, Y: 4}
		for b.Loop() { a = a.Sub(c) }
		_ = a
	})
	b.Run("Vec2/Scale", func(b *testing.B) {
		a := TVec2[float32]{X: 1, Y: 2}
		for b.Loop() { a = a.Scale(2) }
		_ = a
	})
	b.Run("Vec2/Dot", func(b *testing.B) {
		a := TVec2[float32]{X: 1, Y: 2}
		c := TVec2[float32]{X: 3, Y: 4}
		var r float32
		for b.Loop() { r = a.Dot(c) }
		_ = r
	})
	b.Run("Vec2/LenSq", func(b *testing.B) {
		a := TVec2[float32]{X: 3, Y: 4}
		var r float32
		for b.Loop() { r = a.LenSq() }
		_ = r
	})
	b.Run("Vec2/Eq", func(b *testing.B) {
		a := TVec2[float32]{X: 1, Y: 2}
		c := TVec2[float32]{X: 1, Y: 2}
		var r bool
		for b.Loop() { r = a.Eq(c) }
		_ = r
	})
	b.Run("Vec3/Add", func(b *testing.B) {
		a := TVec3[float64]{X: 1, Y: 2, Z: 3}
		c := TVec3[float64]{X: 4, Y: 5, Z: 6}
		for b.Loop() { a = a.Add(c) }
		_ = a
	})
	b.Run("Vec3/Sub", func(b *testing.B) {
		a := TVec3[float64]{X: 1, Y: 2, Z: 3}
		c := TVec3[float64]{X: 4, Y: 5, Z: 6}
		for b.Loop() { a = a.Sub(c) }
		_ = a
	})
	b.Run("Vec3/Scale", func(b *testing.B) {
		a := TVec3[float64]{X: 1, Y: 2, Z: 3}
		for b.Loop() { a = a.Scale(2) }
		_ = a
	})
	b.Run("Vec3/Dot", func(b *testing.B) {
		a := TVec3[float64]{X: 1, Y: 2, Z: 3}
		c := TVec3[float64]{X: 4, Y: 5, Z: 6}
		var r float64
		for b.Loop() { r = a.Dot(c) }
		_ = r
	})
	b.Run("Vec3/Cross", func(b *testing.B) {
		a := TVec3[float64]{X: 1, Y: 2, Z: 3}
		c := TVec3[float64]{X: 4, Y: 5, Z: 6}
		for b.Loop() { a = a.Cross(c) }
		_ = a
	})
	b.Run("Vec3/LenSq", func(b *testing.B) {
		a := TVec3[float64]{X: 1, Y: 2, Z: 3}
		var r float64
		for b.Loop() { r = a.LenSq() }
		_ = r
	})
	b.Run("Vec3/Eq", func(b *testing.B) {
		a := TVec3[float64]{X: 1, Y: 2, Z: 3}
		c := TVec3[float64]{X: 1, Y: 2, Z: 3}
		var r bool
		for b.Loop() { r = a.Eq(c) }
		_ = r
	})
	b.Run("Vec3/XY", func(b *testing.B) {
		a := TVec3[float64]{X: 1, Y: 2, Z: 3}
		v := a.XY()
		for b.Loop() { v = a.XY() }
		_ = v
	})
	b.Run("Vec4/Add", func(b *testing.B) {
		a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		c := TVec4[int32]{X: 5, Y: 6, Z: 7, W: 8}
		for b.Loop() { a = a.Add(c) }
		_ = a
	})
	b.Run("Vec4/Sub", func(b *testing.B) {
		a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		c := TVec4[int32]{X: 5, Y: 6, Z: 7, W: 8}
		for b.Loop() { a = a.Sub(c) }
		_ = a
	})
	b.Run("Vec4/Scale", func(b *testing.B) {
		a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		for b.Loop() { a = a.Scale(2) }
		_ = a
	})
	b.Run("Vec4/Dot", func(b *testing.B) {
		a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		c := TVec4[int32]{X: 5, Y: 6, Z: 7, W: 8}
		var r int32
		for b.Loop() { r = a.Dot(c) }
		_ = r
	})
	b.Run("Vec4/LenSq", func(b *testing.B) {
		a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		var r int32
		for b.Loop() { r = a.LenSq() }
		_ = r
	})
	b.Run("Vec4/Eq", func(b *testing.B) {
		a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		c := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		var r bool
		for b.Loop() { r = a.Eq(c) }
		_ = r
	})
	b.Run("Vec4/XY", func(b *testing.B) {
		a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		v := a.XY()
		for b.Loop() { v = a.XY() }
		_ = v
	})
	b.Run("Vec4/XYZ", func(b *testing.B) {
		a := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		v := a.XYZ()
		for b.Loop() { v = a.XYZ() }
		_ = v
	})
	b.Run("Mat2/Add", func(b *testing.B) {
		a := NewTMat2[float32](1, 2, 3, 4)
		c := NewTMat2[float32](5, 6, 7, 8)
		for b.Loop() { a = a.Add(c) }
		_ = a
	})
	b.Run("Mat2/Sub", func(b *testing.B) {
		a := NewTMat2[float32](1, 2, 3, 4)
		c := NewTMat2[float32](5, 6, 7, 8)
		for b.Loop() { a = a.Sub(c) }
		_ = a
	})
	b.Run("Mat2/Scale", func(b *testing.B) {
		a := NewTMat2[float32](1, 2, 3, 4)
		for b.Loop() { a = a.Scale(2) }
		_ = a
	})
	b.Run("Mat2/Transpose", func(b *testing.B) {
		m := NewTMat2[float32](1, 2, 3, 4)
		for b.Loop() { m = m.Transpose() }
		_ = m
	})
	b.Run("Mat2/Eq", func(b *testing.B) {
		a := NewTMat2[float32](1, 2, 3, 4)
		c := NewTMat2[float32](1, 2, 3, 4)
		var r bool
		for b.Loop() { r = a.Eq(c) }
		_ = r
	})
	b.Run("Mat3/Add", func(b *testing.B) {
		a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
		c := NewTMat3[float64](9, 8, 7, 6, 5, 4, 3, 2, 1)
		for b.Loop() { a = a.Add(c) }
		_ = a
	})
	b.Run("Mat3/Sub", func(b *testing.B) {
		a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
		c := NewTMat3[float64](9, 8, 7, 6, 5, 4, 3, 2, 1)
		for b.Loop() { a = a.Sub(c) }
		_ = a
	})
	b.Run("Mat3/Scale", func(b *testing.B) {
		a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
		for b.Loop() { a = a.Scale(2) }
		_ = a
	})
	b.Run("Mat3/Transpose", func(b *testing.B) {
		m := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
		for b.Loop() { m = m.Transpose() }
		_ = m
	})
	b.Run("Mat3/Eq", func(b *testing.B) {
		a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
		c := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
		var r bool
		for b.Loop() { r = a.Eq(c) }
		_ = r
	})
	b.Run("Mat4/Add", func(b *testing.B) {
		a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
		c := NewTMat4[int32](16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
		for b.Loop() { a = a.Add(c) }
		_ = a
	})
	b.Run("Mat4/Sub", func(b *testing.B) {
		a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
		c := NewTMat4[int32](16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
		for b.Loop() { a = a.Sub(c) }
		_ = a
	})
	b.Run("Mat4/Scale", func(b *testing.B) {
		a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
		for b.Loop() { a = a.Scale(2) }
		_ = a
	})
	b.Run("Mat4/Transpose", func(b *testing.B) {
		m := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
		for b.Loop() { m = m.Transpose() }
		_ = m
	})
	b.Run("Mat4/Eq", func(b *testing.B) {
		a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
		c := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
		var r bool
		for b.Loop() { r = a.Eq(c) }
		_ = r
	})
}

func BenchmarkGenericMul(b *testing.B) {
	b.Run("Mat2", func(b *testing.B) {
		a := NewTMat2[float32](1, 2, 3, 4)
		c := NewTMat2[float32](5, 6, 7, 8)
		for b.Loop() { a = a.Mul(c) }
		_ = a
	})
	b.Run("Mat2Vec2", func(b *testing.B) {
		m := NewTMat2[float32](1, 2, 3, 4)
		v := TVec2[float32]{X: 5, Y: 6}
		for b.Loop() { v = m.MulVec2(v) }
		_ = v
	})
	b.Run("Mat3", func(b *testing.B) {
		a := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
		c := NewTMat3[float64](9, 8, 7, 6, 5, 4, 3, 2, 1)
		for b.Loop() { a = a.Mul(c) }
		_ = a
	})
	b.Run("Mat3Vec3", func(b *testing.B) {
		m := NewTMat3[float64](1, 2, 3, 4, 5, 6, 7, 8, 9)
		v := TVec3[float64]{X: 1, Y: 2, Z: 3}
		for b.Loop() { v = m.MulVec3(v) }
		_ = v
	})
	b.Run("Mat4", func(b *testing.B) {
		a := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
		c := NewTMat4[int32](16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
		for b.Loop() { a = a.Mul(c) }
		_ = a
	})
	b.Run("Mat4Vec4", func(b *testing.B) {
		m := NewTMat4[int32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
		v := TVec4[int32]{X: 1, Y: 2, Z: 3, W: 4}
		for b.Loop() { v = m.MulVec4(v) }
		_ = v
	})
}

func BenchmarkGenericDet(b *testing.B) {
	b.Run("Mat2", func(b *testing.B) {
		m := NewTMat2[float32](1, 2, 3, 4)
		var r float32
		for b.Loop() { r = m.Det() }
		_ = r
	})
	b.Run("Mat3", func(b *testing.B) {
		m := NewTMat3[float64](1, 2, 3, 0, 1, 4, 5, 6, 0)
		var r float64
		for b.Loop() { r = m.Det() }
		_ = r
	})
	b.Run("Mat4", func(b *testing.B) {
		m := NewTMat4[int32](1, 0, 0, 1, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 1)
		var r int32
		for b.Loop() { r = m.Det() }
		_ = r
	})
}

// Transform benchmarks

func BenchmarkTransform(b *testing.B) {
	b.Run("Translation", func(b *testing.B) {
		v := NewVec3(1, 2, 3)
		var m Mat4
		for b.Loop() { m = Translation(v) }
		_ = m
	})
	b.Run("Scaling", func(b *testing.B) {
		v := NewVec3(2, 3, 4)
		var m Mat4
		for b.Loop() { m = Scaling(v) }
		_ = m
	})
	b.Run("RotationX", func(b *testing.B) {
		var m Mat4
		for b.Loop() { m = RotationX(0.785) }
		_ = m
	})
	b.Run("RotationY", func(b *testing.B) {
		var m Mat4
		for b.Loop() { m = RotationY(0.785) }
		_ = m
	})
	b.Run("RotationZ", func(b *testing.B) {
		var m Mat4
		for b.Loop() { m = RotationZ(0.785) }
		_ = m
	})
	b.Run("LookAt", func(b *testing.B) {
		eye := NewVec3(0, 0, 5)
		center := NewVec3(0, 0, 0)
		up := NewVec3(0, 1, 0)
		var m Mat4
		for b.Loop() { m = LookAt(eye, center, up) }
		_ = m
	})
	b.Run("Perspective", func(b *testing.B) {
		var m Mat4
		for b.Loop() { m = Perspective(0.785, 1.6, 0.1, 100) }
		_ = m
	})
	b.Run("Ortho", func(b *testing.B) {
		var m Mat4
		for b.Loop() { m = Ortho(-1, 1, -1, 1, 0.1, 100) }
		_ = m
	})
}
