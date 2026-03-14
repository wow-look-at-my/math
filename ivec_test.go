package math

import (
	"testing"

	"github.com/wow-look-at-my/testify/assert"
)

// ---------------------------------------------------------------------------
// I32Vec2 — full method coverage
// ---------------------------------------------------------------------------

func TestI32Vec2Constructor(t *testing.T) {
	v := NewI32Vec2(3, 4)
	assert.Equal(t, int32(3), v.X)
	assert.Equal(t, int32(4), v.Y)
}

func TestI32Vec2Add(t *testing.T) {
	got := I32Vec2{X: 1, Y: 2}.Add(I32Vec2{X: 3, Y: 4})
	assert.True(t, got.Eq(I32Vec2{X: 4, Y: 6}))
}

func TestI32Vec2Sub(t *testing.T) {
	got := I32Vec2{X: 5, Y: 7}.Sub(I32Vec2{X: 2, Y: 3})
	assert.True(t, got.Eq(I32Vec2{X: 3, Y: 4}))
}

func TestI32Vec2Scale(t *testing.T) {
	got := I32Vec2{X: 2, Y: 3}.Scale(3)
	assert.True(t, got.Eq(I32Vec2{X: 6, Y: 9}))
}

func TestI32Vec2Dot(t *testing.T) {
	got := I32Vec2{X: 1, Y: 2}.Dot(I32Vec2{X: 3, Y: 4})
	assert.Equal(t, int32(11), got)
}

func TestI32Vec2LenSq(t *testing.T) {
	assert.Equal(t, int32(25), I32Vec2{X: 3, Y: 4}.LenSq())
}

func TestI32Vec2Len(t *testing.T) {
	assert.Equal(t, float32(5), I32Vec2{X: 3, Y: 4}.Len())
}

func TestI32Vec2Normalize(t *testing.T) {
	got := I32Vec2{X: 3, Y: 4}.Normalize()
	assert.True(t, got.ApproxEq(NewVec2(0.6, 0.8), 1e-6))
}

func TestI32Vec2NormalizeZero(t *testing.T) {
	got := I32Vec2{}.Normalize()
	assert.True(t, got.Eq(Vec2{}))
}

func TestI32Vec2Dist(t *testing.T) {
	got := I32Vec2{X: 0, Y: 0}.Dist(I32Vec2{X: 3, Y: 4})
	assert.Equal(t, float32(5), got)
}

func TestI32Vec2Lerp(t *testing.T) {
	got := I32Vec2{X: 0, Y: 0}.Lerp(I32Vec2{X: 10, Y: 20}, 0.5)
	assert.True(t, got.Eq(NewVec2(5, 10)))
}

func TestI32Vec2Eq(t *testing.T) {
	a := I32Vec2{X: 1, Y: 2}
	assert.True(t, a.Eq(I32Vec2{X: 1, Y: 2}))
	assert.False(t, a.Eq(I32Vec2{X: 1, Y: 3}))
}

func TestI32Vec2Float32(t *testing.T) {
	got := I32Vec2{X: 3, Y: 4}.Float32()
	assert.True(t, got.Eq(NewVec2(3, 4)))
}

// ---------------------------------------------------------------------------
// I32Vec3 — full method coverage
// ---------------------------------------------------------------------------

func TestI32Vec3Constructor(t *testing.T) {
	v := NewI32Vec3(1, 2, 3)
	assert.Equal(t, int32(1), v.X)
	assert.Equal(t, int32(2), v.Y)
	assert.Equal(t, int32(3), v.Z)
}

func TestI32Vec3Add(t *testing.T) {
	got := I32Vec3{X: 1, Y: 2, Z: 3}.Add(I32Vec3{X: 4, Y: 5, Z: 6})
	assert.True(t, got.Eq(I32Vec3{X: 5, Y: 7, Z: 9}))
}

func TestI32Vec3Sub(t *testing.T) {
	got := I32Vec3{X: 5, Y: 7, Z: 9}.Sub(I32Vec3{X: 1, Y: 2, Z: 3})
	assert.True(t, got.Eq(I32Vec3{X: 4, Y: 5, Z: 6}))
}

func TestI32Vec3Scale(t *testing.T) {
	got := I32Vec3{X: 1, Y: 2, Z: 3}.Scale(3)
	assert.True(t, got.Eq(I32Vec3{X: 3, Y: 6, Z: 9}))
}

func TestI32Vec3Dot(t *testing.T) {
	got := I32Vec3{X: 1, Y: 2, Z: 3}.Dot(I32Vec3{X: 4, Y: 5, Z: 6})
	assert.Equal(t, int32(32), got)
}

func TestI32Vec3Cross(t *testing.T) {
	got := I32Vec3{X: 1, Y: 0, Z: 0}.Cross(I32Vec3{X: 0, Y: 1, Z: 0})
	assert.True(t, got.Eq(I32Vec3{X: 0, Y: 0, Z: 1}))
}

func TestI32Vec3CrossAnticommutative(t *testing.T) {
	a := I32Vec3{X: 2, Y: 3, Z: 4}
	b := I32Vec3{X: 5, Y: 6, Z: 7}
	ab := a.Cross(b)
	ba := b.Cross(a)
	assert.True(t, ab.Eq(ba.Scale(-1)))
}

func TestI32Vec3LenSq(t *testing.T) {
	assert.Equal(t, int32(49), I32Vec3{X: 2, Y: 3, Z: 6}.LenSq())
}

func TestI32Vec3Len(t *testing.T) {
	assert.Equal(t, float32(7), I32Vec3{X: 2, Y: 3, Z: 6}.Len())
}

func TestI32Vec3Normalize(t *testing.T) {
	got := I32Vec3{X: 0, Y: 0, Z: 5}.Normalize()
	assert.True(t, got.ApproxEq(NewVec3(0, 0, 1), 1e-6))
}

func TestI32Vec3NormalizeZero(t *testing.T) {
	got := I32Vec3{}.Normalize()
	assert.True(t, got.Eq(Vec3{}))
}

func TestI32Vec3Dist(t *testing.T) {
	got := I32Vec3{}.Dist(I32Vec3{X: 2, Y: 3, Z: 6})
	assert.Equal(t, float32(7), got)
}

func TestI32Vec3Lerp(t *testing.T) {
	got := I32Vec3{}.Lerp(I32Vec3{X: 10, Y: 20, Z: 30}, 0.25)
	assert.True(t, got.ApproxEq(NewVec3(2.5, 5, 7.5), 1e-6))
}

func TestI32Vec3Eq(t *testing.T) {
	a := I32Vec3{X: 1, Y: 2, Z: 3}
	assert.True(t, a.Eq(I32Vec3{X: 1, Y: 2, Z: 3}))
	assert.False(t, a.Eq(I32Vec3{X: 1, Y: 2, Z: 4}))
}

func TestI32Vec3Float32(t *testing.T) {
	got := I32Vec3{X: 1, Y: 2, Z: 3}.Float32()
	assert.True(t, got.Eq(NewVec3(1, 2, 3)))
}

func TestI32Vec3XY(t *testing.T) {
	got := I32Vec3{X: 1, Y: 2, Z: 3}.XY()
	assert.True(t, got.Eq(I32Vec2{X: 1, Y: 2}))
}

// ---------------------------------------------------------------------------
// I32Vec4 — full method coverage
// ---------------------------------------------------------------------------

func TestI32Vec4Constructor(t *testing.T) {
	v := NewI32Vec4(1, 2, 3, 4)
	assert.Equal(t, int32(1), v.X)
	assert.Equal(t, int32(4), v.W)
}

func TestI32Vec4Add(t *testing.T) {
	got := I32Vec4{X: 1, Y: 2, Z: 3, W: 4}.Add(I32Vec4{X: 5, Y: 6, Z: 7, W: 8})
	assert.True(t, got.Eq(I32Vec4{X: 6, Y: 8, Z: 10, W: 12}))
}

func TestI32Vec4Sub(t *testing.T) {
	got := I32Vec4{X: 5, Y: 7, Z: 9, W: 11}.Sub(I32Vec4{X: 1, Y: 2, Z: 3, W: 4})
	assert.True(t, got.Eq(I32Vec4{X: 4, Y: 5, Z: 6, W: 7}))
}

func TestI32Vec4Scale(t *testing.T) {
	got := I32Vec4{X: 1, Y: 2, Z: 3, W: 4}.Scale(2)
	assert.True(t, got.Eq(I32Vec4{X: 2, Y: 4, Z: 6, W: 8}))
}

func TestI32Vec4Dot(t *testing.T) {
	got := I32Vec4{X: 1, Y: 2, Z: 3, W: 4}.Dot(I32Vec4{X: 5, Y: 6, Z: 7, W: 8})
	assert.Equal(t, int32(70), got)
}

func TestI32Vec4LenSq(t *testing.T) {
	assert.Equal(t, int32(25), I32Vec4{X: 1, Y: 2, Z: 2, W: 4}.LenSq())
}

func TestI32Vec4Len(t *testing.T) {
	assert.Equal(t, float32(5), I32Vec4{X: 1, Y: 2, Z: 2, W: 4}.Len())
}

func TestI32Vec4Normalize(t *testing.T) {
	got := I32Vec4{X: 0, Y: 0, Z: 0, W: 5}.Normalize()
	assert.True(t, got.ApproxEq(NewVec4(0, 0, 0, 1), 1e-6))
}

func TestI32Vec4NormalizeZero(t *testing.T) {
	got := I32Vec4{}.Normalize()
	assert.True(t, got.Eq(Vec4{}))
}

func TestI32Vec4Dist(t *testing.T) {
	got := I32Vec4{}.Dist(I32Vec4{X: 1, Y: 2, Z: 2, W: 4})
	assert.Equal(t, float32(5), got)
}

func TestI32Vec4Lerp(t *testing.T) {
	got := I32Vec4{}.Lerp(I32Vec4{X: 10, Y: 20, Z: 30, W: 40}, 0.5)
	assert.True(t, got.Eq(NewVec4(5, 10, 15, 20)))
}

func TestI32Vec4Eq(t *testing.T) {
	a := I32Vec4{X: 1, Y: 2, Z: 3, W: 4}
	assert.True(t, a.Eq(I32Vec4{X: 1, Y: 2, Z: 3, W: 4}))
	assert.False(t, a.Eq(I32Vec4{X: 1, Y: 2, Z: 3, W: 5}))
}

func TestI32Vec4Float32(t *testing.T) {
	got := I32Vec4{X: 1, Y: 2, Z: 3, W: 4}.Float32()
	assert.True(t, got.Eq(NewVec4(1, 2, 3, 4)))
}

func TestI32Vec4XYZ(t *testing.T) {
	got := I32Vec4{X: 1, Y: 2, Z: 3, W: 4}.XYZ()
	assert.True(t, got.Eq(I32Vec3{X: 1, Y: 2, Z: 3}))
}

func TestI32Vec4XY(t *testing.T) {
	got := I32Vec4{X: 1, Y: 2, Z: 3, W: 4}.XY()
	assert.True(t, got.Eq(I32Vec2{X: 1, Y: 2}))
}

// ---------------------------------------------------------------------------
// I32Mat2 — full method coverage
// ---------------------------------------------------------------------------

func TestI32Mat2Constructor(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	assert.Equal(t, int32(1), m.At(0, 0))
	assert.Equal(t, int32(2), m.At(0, 1))
	assert.Equal(t, int32(3), m.At(1, 0))
	assert.Equal(t, int32(4), m.At(1, 1))
}

func TestI32Mat2Identity(t *testing.T) {
	m := I32Mat2Identity()
	assert.Equal(t, int32(1), m.At(0, 0))
	assert.Equal(t, int32(0), m.At(0, 1))
	assert.Equal(t, int32(0), m.At(1, 0))
	assert.Equal(t, int32(1), m.At(1, 1))
}

func TestI32Mat2Col(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	assert.True(t, m.Col(0).Eq(I32Vec2{X: 1, Y: 3}))
	assert.True(t, m.Col(1).Eq(I32Vec2{X: 2, Y: 4}))
}

func TestI32Mat2Row(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	assert.True(t, m.Row(0).Eq(I32Vec2{X: 1, Y: 2}))
	assert.True(t, m.Row(1).Eq(I32Vec2{X: 3, Y: 4}))
}

func TestI32Mat2Add(t *testing.T) {
	a := NewI32Mat2(1, 2, 3, 4)
	b := NewI32Mat2(5, 6, 7, 8)
	got := a.Add(b)
	assert.True(t, got.Eq(NewI32Mat2(6, 8, 10, 12)))
}

func TestI32Mat2Sub(t *testing.T) {
	a := NewI32Mat2(5, 6, 7, 8)
	b := NewI32Mat2(1, 2, 3, 4)
	got := a.Sub(b)
	assert.True(t, got.Eq(NewI32Mat2(4, 4, 4, 4)))
}

func TestI32Mat2Scale(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	got := m.Scale(2)
	assert.True(t, got.Eq(NewI32Mat2(2, 4, 6, 8)))
}

func TestI32Mat2Mul(t *testing.T) {
	a := NewI32Mat2(1, 2, 3, 4)
	b := NewI32Mat2(5, 6, 7, 8)
	got := a.Mul(b)
	assert.True(t, got.Eq(NewI32Mat2(19, 22, 43, 50)))
}

func TestI32Mat2MulIdentity(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	got := m.Mul(I32Mat2Identity())
	assert.True(t, got.Eq(m))
}

func TestI32Mat2MulVec2(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	v := NewI32Vec2(5, 6)
	got := m.MulVec2(v)
	assert.True(t, got.Eq(NewI32Vec2(17, 39)))
}

func TestI32Mat2Transpose(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	got := m.Transpose()
	assert.True(t, got.Eq(NewI32Mat2(1, 3, 2, 4)))
}

func TestI32Mat2Det(t *testing.T) {
	assert.Equal(t, int32(-2), NewI32Mat2(1, 2, 3, 4).Det())
}

func TestI32Mat2Inverse(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	inv := m.Inverse() // returns Mat2 (float32)
	product := m.Float32().Mul(inv)
	identity := Mat2Identity()
	for r := 0; r < 2; r++ {
		for c := 0; c < 2; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			assert.False(t, diff > 1e-5 || diff < -1e-5)
		}
	}
}

func TestI32Mat2InverseSingular(t *testing.T) {
	m := NewI32Mat2(1, 2, 2, 4)
	got := m.Inverse()
	assert.Equal(t, Mat2{}, got)
}

func TestI32Mat2Eq(t *testing.T) {
	a := NewI32Mat2(1, 2, 3, 4)
	assert.True(t, a.Eq(NewI32Mat2(1, 2, 3, 4)))
	assert.False(t, a.Eq(NewI32Mat2(1, 2, 3, 5)))
}

func TestI32Mat2Float32(t *testing.T) {
	got := NewI32Mat2(1, 2, 3, 4).Float32()
	assert.Equal(t, NewMat2(1, 2, 3, 4), got)
}

// ---------------------------------------------------------------------------
// I32Mat3 — full method coverage
// ---------------------------------------------------------------------------

func TestI32Mat3Constructor(t *testing.T) {
	m := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assert.Equal(t, int32(1), m.At(0, 0))
	assert.Equal(t, int32(3), m.At(0, 2))
	assert.Equal(t, int32(7), m.At(2, 0))
	assert.Equal(t, int32(9), m.At(2, 2))
}

func TestI32Mat3Identity(t *testing.T) {
	m := I32Mat3Identity()
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			var want int32
			if r == c {
				want = 1
			}
			assert.Equal(t, want, m.At(r, c))
		}
	}
}

func TestI32Mat3Col(t *testing.T) {
	m := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assert.True(t, m.Col(1).Eq(I32Vec3{X: 2, Y: 5, Z: 8}))
}

func TestI32Mat3Row(t *testing.T) {
	m := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assert.True(t, m.Row(1).Eq(I32Vec3{X: 4, Y: 5, Z: 6}))
}

func TestI32Mat3Add(t *testing.T) {
	a := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := NewI32Mat3(9, 8, 7, 6, 5, 4, 3, 2, 1)
	got := a.Add(b)
	assert.True(t, got.Eq(NewI32Mat3(10, 10, 10, 10, 10, 10, 10, 10, 10)))
}

func TestI32Mat3Sub(t *testing.T) {
	a := NewI32Mat3(9, 8, 7, 6, 5, 4, 3, 2, 1)
	b := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := a.Sub(b)
	assert.True(t, got.Eq(NewI32Mat3(8, 6, 4, 2, 0, -2, -4, -6, -8)))
}

func TestI32Mat3Scale(t *testing.T) {
	m := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Scale(2)
	assert.True(t, got.Eq(NewI32Mat3(2, 4, 6, 8, 10, 12, 14, 16, 18)))
}

func TestI32Mat3Mul(t *testing.T) {
	a := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := a.Mul(I32Mat3Identity())
	assert.True(t, got.Eq(a))
}

func TestI32Mat3MulNonTrivial(t *testing.T) {
	a := NewI32Mat3(1, 2, 0, 0, 1, 1, 1, 0, 1)
	b := NewI32Mat3(1, 0, 1, 0, 1, 0, 1, 1, 0)
	got := a.Mul(b)
	assert.True(t, got.Eq(NewI32Mat3(1, 2, 1, 1, 2, 0, 2, 1, 1)))
}

func TestI32Mat3MulVec3(t *testing.T) {
	m := I32Mat3Identity()
	v := I32Vec3{X: 3, Y: 4, Z: 5}
	got := m.MulVec3(v)
	assert.True(t, got.Eq(v))
}

func TestI32Mat3Transpose(t *testing.T) {
	m := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Transpose()
	assert.True(t, got.Eq(NewI32Mat3(1, 4, 7, 2, 5, 8, 3, 6, 9)))
}

func TestI32Mat3Det(t *testing.T) {
	m := NewI32Mat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	assert.Equal(t, int32(1), m.Det())
}

func TestI32Mat3Inverse(t *testing.T) {
	m := NewI32Mat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	inv := m.Inverse() // returns Mat3
	product := m.Float32().Mul(inv)
	identity := Mat3Identity()
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			assert.False(t, diff > 1e-4 || diff < -1e-4)
		}
	}
}

func TestI32Mat3InverseSingular(t *testing.T) {
	m := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Inverse()
	assert.Equal(t, Mat3{}, got)
}

func TestI32Mat3Eq(t *testing.T) {
	a := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assert.True(t, a.Eq(NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9)))
	assert.False(t, a.Eq(I32Mat3Identity()))
}

func TestI32Mat3Float32(t *testing.T) {
	got := NewI32Mat3(1, 2, 3, 4, 5, 6, 7, 8, 9).Float32()
	assert.Equal(t, NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9), got)
}

// ---------------------------------------------------------------------------
// I32Mat4 — full method coverage
// ---------------------------------------------------------------------------

func TestI32Mat4Constructor(t *testing.T) {
	m := NewI32Mat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	assert.Equal(t, int32(1), m.At(0, 0))
	assert.Equal(t, int32(4), m.At(0, 3))
	assert.Equal(t, int32(13), m.At(3, 0))
	assert.Equal(t, int32(16), m.At(3, 3))
}

func TestI32Mat4Identity(t *testing.T) {
	m := I32Mat4Identity()
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			var want int32
			if r == c {
				want = 1
			}
			assert.Equal(t, want, m.At(r, c))
		}
	}
}

func TestI32Mat4Col(t *testing.T) {
	m := NewI32Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	got := m.Col(2)
	assert.True(t, got.Eq(I32Vec4{X: 3, Y: 7, Z: 11, W: 15}))
}

func TestI32Mat4Row(t *testing.T) {
	m := NewI32Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	got := m.Row(2)
	assert.True(t, got.Eq(I32Vec4{X: 9, Y: 10, Z: 11, W: 12}))
}

func TestI32Mat4Add(t *testing.T) {
	a := I32Mat4Identity()
	got := a.Add(a)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			assert.Equal(t, a.At(r, c)*2, got.At(r, c))
		}
	}
}

func TestI32Mat4Sub(t *testing.T) {
	a := I32Mat4Identity()
	got := a.Sub(a)
	assert.True(t, got.Eq(I32Mat4{}))
}

func TestI32Mat4Scale(t *testing.T) {
	m := I32Mat4Identity()
	got := m.Scale(3)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			assert.Equal(t, m.At(r, c)*3, got.At(r, c))
		}
	}
}

func TestI32Mat4Mul(t *testing.T) {
	m := NewI32Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	got := m.Mul(I32Mat4Identity())
	assert.True(t, got.Eq(m))
}

func TestI32Mat4MulVec4(t *testing.T) {
	m := I32Mat4Identity()
	v := I32Vec4{X: 1, Y: 2, Z: 3, W: 4}
	got := m.MulVec4(v)
	assert.True(t, got.Eq(v))
}

func TestI32Mat4Transpose(t *testing.T) {
	m := NewI32Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	got := m.Transpose()
	want := NewI32Mat4(1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15, 4, 8, 12, 16)
	assert.True(t, got.Eq(want))
}

func TestI32Mat4TransposeDouble(t *testing.T) {
	m := NewI32Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	assert.True(t, m.Transpose().Transpose().Eq(m))
}

func TestI32Mat4Det(t *testing.T) {
	m := NewI32Mat4(1, 0, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 4)
	assert.Equal(t, int32(24), m.Det())
}

func TestI32Mat4Inverse(t *testing.T) {
	m := NewI32Mat4(1, 0, 0, 1, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 1)
	inv := m.Inverse() // returns Mat4
	product := m.Float32().Mul(inv)
	identity := Mat4Identity()
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			assert.False(t, diff > 1e-4 || diff < -1e-4)
		}
	}
}

func TestI32Mat4InverseSingular(t *testing.T) {
	m := NewI32Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	got := m.Inverse()
	assert.Equal(t, Mat4{}, got)
}

func TestI32Mat4Mat3(t *testing.T) {
	m := NewI32Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	got := m.Mat3()
	assert.True(t, got.Eq(NewI32Mat3(1, 2, 3, 5, 6, 7, 9, 10, 11)))
}

func TestI32Mat4Eq(t *testing.T) {
	a := I32Mat4Identity()
	assert.True(t, a.Eq(I32Mat4Identity()))
	assert.False(t, a.Eq(I32Mat4{}))
}

func TestI32Mat4Float32(t *testing.T) {
	got := I32Mat4Identity().Float32()
	assert.Equal(t, Mat4Identity(), got)
}

// ---------------------------------------------------------------------------
// I8 — basic coverage (one per vec/mat size)
// ---------------------------------------------------------------------------

func TestI8Vec2Add(t *testing.T) {
	got := I8Vec2{X: 1, Y: 2}.Add(I8Vec2{X: 3, Y: 4})
	assert.True(t, got.Eq(I8Vec2{X: 4, Y: 6}))
}

func TestI8Vec2Normalize(t *testing.T) {
	got := I8Vec2{X: 3, Y: 4}.Normalize()
	assert.True(t, got.ApproxEq(NewVec2(0.6, 0.8), 1e-6))
}

func TestI8Vec3Cross(t *testing.T) {
	got := I8Vec3{X: 1, Y: 0, Z: 0}.Cross(I8Vec3{X: 0, Y: 1, Z: 0})
	assert.True(t, got.Eq(I8Vec3{X: 0, Y: 0, Z: 1}))
}

func TestI8Vec4Len(t *testing.T) {
	assert.Equal(t, float32(5), I8Vec4{X: 1, Y: 2, Z: 2, W: 4}.Len())
}

func TestI8Mat2Mul(t *testing.T) {
	a := NewI8Mat2(1, 2, 3, 4)
	b := NewI8Mat2(5, 6, 7, 8)
	got := a.Mul(b)
	assert.True(t, got.Eq(NewI8Mat2(19, 22, 43, 50)))
}

func TestI8Mat3Det(t *testing.T) {
	m := NewI8Mat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	assert.Equal(t, int8(1), m.Det())
}

func TestI8Mat4MulIdentity(t *testing.T) {
	m := NewI8Mat4(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)
	got := m.Mul(I8Mat4Identity())
	assert.True(t, got.Eq(m))
}

// ---------------------------------------------------------------------------
// I16 — basic coverage
// ---------------------------------------------------------------------------

func TestI16Vec2Add(t *testing.T) {
	got := I16Vec2{X: 100, Y: 200}.Add(I16Vec2{X: 300, Y: 400})
	assert.True(t, got.Eq(I16Vec2{X: 400, Y: 600}))
}

func TestI16Vec3Normalize(t *testing.T) {
	got := I16Vec3{X: 0, Y: 0, Z: 5}.Normalize()
	assert.True(t, got.ApproxEq(NewVec3(0, 0, 1), 1e-6))
}

func TestI16Vec4Dot(t *testing.T) {
	got := I16Vec4{X: 1, Y: 2, Z: 3, W: 4}.Dot(I16Vec4{X: 5, Y: 6, Z: 7, W: 8})
	assert.Equal(t, int16(70), got)
}

func TestI16Mat2Inverse(t *testing.T) {
	m := NewI16Mat2(1, 2, 3, 4)
	inv := m.Inverse()
	product := m.Float32().Mul(inv)
	identity := Mat2Identity()
	for r := 0; r < 2; r++ {
		for c := 0; c < 2; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			assert.False(t, diff > 1e-5 || diff < -1e-5)
		}
	}
}

func TestI16Mat3Mul(t *testing.T) {
	a := NewI16Mat3(1, 2, 0, 0, 1, 1, 1, 0, 1)
	b := NewI16Mat3(1, 0, 1, 0, 1, 0, 1, 1, 0)
	got := a.Mul(b)
	assert.True(t, got.Eq(NewI16Mat3(1, 2, 1, 1, 2, 0, 2, 1, 1)))
}

func TestI16Mat4Det(t *testing.T) {
	m := NewI16Mat4(1, 0, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 4)
	assert.Equal(t, int16(24), m.Det())
}

// ---------------------------------------------------------------------------
// I64 — basic coverage
// ---------------------------------------------------------------------------

func TestI64Vec2Add(t *testing.T) {
	got := I64Vec2{X: 1000000, Y: 2000000}.Add(I64Vec2{X: 3000000, Y: 4000000})
	assert.True(t, got.Eq(I64Vec2{X: 4000000, Y: 6000000}))
}

func TestI64Vec3Cross(t *testing.T) {
	got := I64Vec3{X: 1, Y: 0, Z: 0}.Cross(I64Vec3{X: 0, Y: 1, Z: 0})
	assert.True(t, got.Eq(I64Vec3{X: 0, Y: 0, Z: 1}))
}

func TestI64Vec4Normalize(t *testing.T) {
	got := I64Vec4{X: 0, Y: 0, Z: 0, W: 10}.Normalize()
	assert.True(t, got.ApproxEq(NewVec4(0, 0, 0, 1), 1e-6))
}

func TestI64Mat2Mul(t *testing.T) {
	a := NewI64Mat2(1, 2, 3, 4)
	b := NewI64Mat2(5, 6, 7, 8)
	assert.True(t, a.Mul(b).Eq(NewI64Mat2(19, 22, 43, 50)))
}

func TestI64Mat3Inverse(t *testing.T) {
	m := NewI64Mat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	inv := m.Inverse()
	product := m.Float32().Mul(inv)
	identity := Mat3Identity()
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			assert.False(t, diff > 1e-4 || diff < -1e-4)
		}
	}
}

func TestI64Mat4Transpose(t *testing.T) {
	m := NewI64Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	assert.True(t, m.Transpose().Transpose().Eq(m))
}

// ---------------------------------------------------------------------------
// U8 — basic coverage
// ---------------------------------------------------------------------------

func TestU8Vec2Add(t *testing.T) {
	got := U8Vec2{X: 1, Y: 2}.Add(U8Vec2{X: 3, Y: 4})
	assert.True(t, got.Eq(U8Vec2{X: 4, Y: 6}))
}

func TestU8Vec2Normalize(t *testing.T) {
	got := U8Vec2{X: 3, Y: 4}.Normalize()
	assert.True(t, got.ApproxEq(NewVec2(0.6, 0.8), 1e-6))
}

func TestU8Vec3Add(t *testing.T) {
	got := U8Vec3{X: 1, Y: 2, Z: 3}.Add(U8Vec3{X: 4, Y: 5, Z: 6})
	assert.True(t, got.Eq(U8Vec3{X: 5, Y: 7, Z: 9}))
}

func TestU8Vec4Len(t *testing.T) {
	assert.Equal(t, float32(5), U8Vec4{X: 1, Y: 2, Z: 2, W: 4}.Len())
}

func TestU8Mat2Mul(t *testing.T) {
	a := NewU8Mat2(1, 2, 3, 4)
	b := NewU8Mat2(5, 6, 7, 8)
	got := a.Mul(b)
	assert.True(t, got.Eq(NewU8Mat2(19, 22, 43, 50)))
}

func TestU8Mat3Identity(t *testing.T) {
	m := U8Mat3Identity()
	assert.Equal(t, uint8(1), m.At(0, 0))
	assert.Equal(t, uint8(0), m.At(0, 1))
}

func TestU8Mat4MulIdentity(t *testing.T) {
	m := U8Mat4Identity()
	assert.True(t, m.Mul(m).Eq(m))
}

// ---------------------------------------------------------------------------
// U16 — basic coverage
// ---------------------------------------------------------------------------

func TestU16Vec2Add(t *testing.T) {
	got := U16Vec2{X: 100, Y: 200}.Add(U16Vec2{X: 300, Y: 400})
	assert.True(t, got.Eq(U16Vec2{X: 400, Y: 600}))
}

func TestU16Vec3Normalize(t *testing.T) {
	got := U16Vec3{X: 0, Y: 0, Z: 5}.Normalize()
	assert.True(t, got.ApproxEq(NewVec3(0, 0, 1), 1e-6))
}

func TestU16Vec4Dot(t *testing.T) {
	got := U16Vec4{X: 1, Y: 2, Z: 3, W: 4}.Dot(U16Vec4{X: 5, Y: 6, Z: 7, W: 8})
	assert.Equal(t, uint16(70), got)
}

func TestU16Mat2Det(t *testing.T) {
	// 4*1 - 2*3 = -2, but uint16 wraps
	m := NewU16Mat2(4, 2, 3, 1)
	// det = 4*1 - 2*3 = 4 - 6 = underflow for uint16
	// Use a case that doesn't underflow
	m2 := NewU16Mat2(3, 1, 1, 2)
	assert.Equal(t, uint16(5), m2.Det())
	_ = m
}

func TestU16Mat3Mul(t *testing.T) {
	m := U16Mat3Identity()
	assert.True(t, m.Mul(m).Eq(m))
}

func TestU16Mat4Transpose(t *testing.T) {
	m := NewU16Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	assert.True(t, m.Transpose().Transpose().Eq(m))
}

// ---------------------------------------------------------------------------
// U32 — basic coverage
// ---------------------------------------------------------------------------

func TestU32Vec2Add(t *testing.T) {
	got := U32Vec2{X: 1, Y: 2}.Add(U32Vec2{X: 3, Y: 4})
	assert.True(t, got.Eq(U32Vec2{X: 4, Y: 6}))
}

func TestU32Vec2Normalize(t *testing.T) {
	got := U32Vec2{X: 3, Y: 4}.Normalize()
	assert.True(t, got.ApproxEq(NewVec2(0.6, 0.8), 1e-6))
}

func TestU32Vec3Cross(t *testing.T) {
	got := U32Vec3{X: 1, Y: 0, Z: 0}.Cross(U32Vec3{X: 0, Y: 1, Z: 0})
	assert.True(t, got.Eq(U32Vec3{X: 0, Y: 0, Z: 1}))
}

func TestU32Vec4Len(t *testing.T) {
	assert.Equal(t, float32(5), U32Vec4{X: 1, Y: 2, Z: 2, W: 4}.Len())
}

func TestU32Mat2Mul(t *testing.T) {
	a := NewU32Mat2(1, 2, 3, 4)
	b := NewU32Mat2(5, 6, 7, 8)
	got := a.Mul(b)
	assert.True(t, got.Eq(NewU32Mat2(19, 22, 43, 50)))
}

func TestU32Mat3MulIdentity(t *testing.T) {
	m := U32Mat3Identity()
	assert.True(t, m.Mul(m).Eq(m))
}

func TestU32Mat4Det(t *testing.T) {
	m := NewU32Mat4(1, 0, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 4)
	assert.Equal(t, uint32(24), m.Det())
}

// ---------------------------------------------------------------------------
// U64 — basic coverage
// ---------------------------------------------------------------------------

func TestU64Vec2Add(t *testing.T) {
	got := U64Vec2{X: 1, Y: 2}.Add(U64Vec2{X: 3, Y: 4})
	assert.True(t, got.Eq(U64Vec2{X: 4, Y: 6}))
}

func TestU64Vec3Normalize(t *testing.T) {
	got := U64Vec3{X: 0, Y: 0, Z: 5}.Normalize()
	assert.True(t, got.ApproxEq(NewVec3(0, 0, 1), 1e-6))
}

func TestU64Vec4Dot(t *testing.T) {
	got := U64Vec4{X: 1, Y: 2, Z: 3, W: 4}.Dot(U64Vec4{X: 5, Y: 6, Z: 7, W: 8})
	assert.Equal(t, uint64(70), got)
}

func TestU64Mat2Mul(t *testing.T) {
	a := NewU64Mat2(1, 2, 3, 4)
	b := NewU64Mat2(5, 6, 7, 8)
	assert.True(t, a.Mul(b).Eq(NewU64Mat2(19, 22, 43, 50)))
}

func TestU64Mat3MulIdentity(t *testing.T) {
	m := U64Mat3Identity()
	assert.True(t, m.Mul(m).Eq(m))
}

func TestU64Mat4Transpose(t *testing.T) {
	m := NewU64Mat4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	assert.True(t, m.Transpose().Transpose().Eq(m))
}

// ---------------------------------------------------------------------------
// Generic TVec — basic ops with various types
// ---------------------------------------------------------------------------

func TestTVec2IntAdd(t *testing.T) {
	a := TVec2[int]{X: 1, Y: 2}
	b := TVec2[int]{X: 3, Y: 4}
	got := a.Add(b)
	assert.True(t, got.Eq(TVec2[int]{X: 4, Y: 6}))
}

func TestTVec2IntDot(t *testing.T) {
	a := TVec2[int]{X: 1, Y: 2}
	b := TVec2[int]{X: 3, Y: 4}
	assert.Equal(t, 11, a.Dot(b))
}

func TestTVec3IntCross(t *testing.T) {
	a := TVec3[int]{X: 1, Y: 0, Z: 0}
	b := TVec3[int]{X: 0, Y: 1, Z: 0}
	got := a.Cross(b)
	assert.True(t, got.Eq(TVec3[int]{X: 0, Y: 0, Z: 1}))
}

func TestTVec3IntDot(t *testing.T) {
	a := TVec3[int]{X: 1, Y: 2, Z: 3}
	b := TVec3[int]{X: 4, Y: 5, Z: 6}
	assert.Equal(t, 32, a.Dot(b))
}

func TestTVec4IntScale(t *testing.T) {
	v := TVec4[int]{X: 1, Y: 2, Z: 3, W: 4}
	got := v.Scale(2)
	assert.True(t, got.Eq(TVec4[int]{X: 2, Y: 4, Z: 6, W: 8}))
}

func TestTVec2Float64(t *testing.T) {
	a := TVec2[float64]{X: 1.5, Y: 2.5}
	b := TVec2[float64]{X: 3.5, Y: 4.5}
	got := a.Add(b)
	assert.True(t, got.Eq(TVec2[float64]{X: 5.0, Y: 7.0}))
}

// ---------------------------------------------------------------------------
// Generic TMat — basic ops
// ---------------------------------------------------------------------------

func TestTMat2IntMul(t *testing.T) {
	a := NewTMat2[int](1, 2, 3, 4)
	b := NewTMat2[int](5, 6, 7, 8)
	got := a.Mul(b)
	assert.True(t, got.Eq(NewTMat2[int](19, 22, 43, 50)))
}

func TestTMat3IntTranspose(t *testing.T) {
	m := NewTMat3[int](1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Transpose().Transpose()
	assert.True(t, got.Eq(m))
}

func TestTMat4IntMulVec4(t *testing.T) {
	m := NewTMat4[int](1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)
	v := TVec4[int]{X: 1, Y: 2, Z: 3, W: 4}
	got := m.MulVec4(v)
	assert.True(t, got.Eq(v))
}
