package math

import (
	"testing"

	"github.com/wow-look-at-my/testify/assert"
)

// I32Vec2 tests

func TestI32Vec2Add(t *testing.T) {
	a := I32Vec2{X: 1, Y: 2}
	b := I32Vec2{X: 3, Y: 4}
	got := a.Add(b)
	want := I32Vec2{X: 4, Y: 6}
	assert.True(t, got.Eq(want))
}

func TestI32Vec2Sub(t *testing.T) {
	a := I32Vec2{X: 5, Y: 7}
	b := I32Vec2{X: 2, Y: 3}
	got := a.Sub(b)
	want := I32Vec2{X: 3, Y: 4}
	assert.True(t, got.Eq(want))
}

func TestI32Vec2Scale(t *testing.T) {
	v := I32Vec2{X: 2, Y: 3}
	got := v.Scale(3)
	want := I32Vec2{X: 6, Y: 9}
	assert.True(t, got.Eq(want))
}

func TestI32Vec2Dot(t *testing.T) {
	a := I32Vec2{X: 1, Y: 2}
	b := I32Vec2{X: 3, Y: 4}
	got := a.Dot(b)
	assert.Equal(t, int32(11), got)
}

func TestI32Vec2LenSq(t *testing.T) {
	v := I32Vec2{X: 3, Y: 4}
	got := v.LenSq()
	assert.Equal(t, int32(25), got)
}

func TestI32Vec2Len(t *testing.T) {
	v := I32Vec2{X: 3, Y: 4}
	got := v.Len()
	assert.Equal(t, float32(5), got)
}

func TestI32Vec2Normalize(t *testing.T) {
	v := I32Vec2{X: 3, Y: 4}
	got := v.Normalize()
	want := NewVec2(0.6, 0.8)
	assert.True(t, got.ApproxEq(want, 1e-6))
}

func TestI32Vec2NormalizeZero(t *testing.T) {
	v := I32Vec2{X: 0, Y: 0}
	got := v.Normalize()
	want := Vec2{}
	assert.True(t, got.Eq(want))
}

func TestI32Vec2Dist(t *testing.T) {
	a := I32Vec2{X: 0, Y: 0}
	b := I32Vec2{X: 3, Y: 4}
	got := a.Dist(b)
	assert.Equal(t, float32(5), got)
}

func TestI32Vec2Lerp(t *testing.T) {
	a := I32Vec2{X: 0, Y: 0}
	b := I32Vec2{X: 10, Y: 20}
	got := a.Lerp(b, 0.5)
	want := NewVec2(5, 10)
	assert.True(t, got.Eq(want))
}

func TestI32Vec2Float32(t *testing.T) {
	v := I32Vec2{X: 3, Y: 4}
	got := v.Float32()
	want := NewVec2(3, 4)
	assert.True(t, got.Eq(want))
}

// I32Vec3 tests

func TestI32Vec3Cross(t *testing.T) {
	a := I32Vec3{X: 1, Y: 0, Z: 0}
	b := I32Vec3{X: 0, Y: 1, Z: 0}
	got := a.Cross(b)
	want := I32Vec3{X: 0, Y: 0, Z: 1}
	assert.True(t, got.Eq(want))
}

func TestI32Vec3Normalize(t *testing.T) {
	v := I32Vec3{X: 0, Y: 0, Z: 5}
	got := v.Normalize()
	want := NewVec3(0, 0, 1)
	assert.True(t, got.ApproxEq(want, 1e-6))
}

func TestI32Vec3XY(t *testing.T) {
	v := I32Vec3{X: 1, Y: 2, Z: 3}
	got := v.XY()
	want := I32Vec2{X: 1, Y: 2}
	assert.True(t, got.Eq(want))
}

// U32Vec2 tests

func TestU32Vec2Add(t *testing.T) {
	a := U32Vec2{X: 1, Y: 2}
	b := U32Vec2{X: 3, Y: 4}
	got := a.Add(b)
	want := U32Vec2{X: 4, Y: 6}
	assert.True(t, got.Eq(want))
}

func TestU32Vec2Normalize(t *testing.T) {
	v := U32Vec2{X: 3, Y: 4}
	got := v.Normalize()
	want := NewVec2(0.6, 0.8)
	assert.True(t, got.ApproxEq(want, 1e-6))
}

// U8Vec3 tests

func TestU8Vec3Add(t *testing.T) {
	a := U8Vec3{X: 1, Y: 2, Z: 3}
	b := U8Vec3{X: 4, Y: 5, Z: 6}
	got := a.Add(b)
	want := U8Vec3{X: 5, Y: 7, Z: 9}
	assert.True(t, got.Eq(want))
}

// Generic TVec tests

func TestTVec2IntAdd(t *testing.T) {
	a := TVec2[int]{X: 1, Y: 2}
	b := TVec2[int]{X: 3, Y: 4}
	got := a.Add(b)
	want := TVec2[int]{X: 4, Y: 6}
	assert.True(t, got.Eq(want))
}

func TestTVec3IntDot(t *testing.T) {
	a := TVec3[int]{X: 1, Y: 2, Z: 3}
	b := TVec3[int]{X: 4, Y: 5, Z: 6}
	got := a.Dot(b)
	assert.Equal(t, 32, got)
}

// I32Mat2 tests

func TestI32Mat2Mul(t *testing.T) {
	a := NewI32Mat2(1, 2, 3, 4)
	b := NewI32Mat2(5, 6, 7, 8)
	got := a.Mul(b)
	want := NewI32Mat2(19, 22, 43, 50)
	assert.True(t, got.Eq(want))
}

func TestI32Mat2Det(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	got := m.Det()
	assert.Equal(t, int32(-2), got)
}

func TestI32Mat2Inverse(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	inv := m.Inverse()
	// Inverse of I32Mat2 returns Mat2 (float32)
	mf := m.Float32()
	product := mf.Mul(inv)
	identity := Mat2Identity()
	for r := 0; r < 2; r++ {
		for c := 0; c < 2; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			assert.False(t, diff > 1e-5 || diff < -1e-5)
		}
	}
}

func TestI32Mat2MulVec2(t *testing.T) {
	m := NewI32Mat2(1, 2, 3, 4)
	v := NewI32Vec2(5, 6)
	got := m.MulVec2(v)
	want := NewI32Vec2(17, 39)
	assert.True(t, got.Eq(want))
}

func TestNewI32Vec2(t *testing.T) {
	v := NewI32Vec2(3, 4)
	assert.Equal(t, int32(3), v.X)
	assert.Equal(t, int32(4), v.Y)
}
