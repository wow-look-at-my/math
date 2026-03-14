package math

import (
	"testing"
	"github.com/wow-look-at-my/testify/assert"
)

func TestNewVec4(t *testing.T) {
	v := NewVec4(1, 2, 3, 4)
	assert.False(t, v.X() != 1 || v.Y() != 2 || v.Z() != 3 || v.W() != 4)

}

func TestVec4Add(t *testing.T) {
	a := NewVec4(1, 2, 3, 4)
	b := NewVec4(5, 6, 7, 8)
	got := a.Add(b)
	want := NewVec4(6, 8, 10, 12)
	assert.True(t, got.Eq(want))

}

func TestVec4Sub(t *testing.T) {
	a := NewVec4(5, 7, 9, 11)
	b := NewVec4(1, 2, 3, 4)
	got := a.Sub(b)
	want := NewVec4(4, 5, 6, 7)
	assert.True(t, got.Eq(want))

}

func TestVec4Scale(t *testing.T) {
	v := NewVec4(1, 2, 3, 4)
	got := v.Scale(2)
	want := NewVec4(2, 4, 6, 8)
	assert.True(t, got.Eq(want))

}

func TestVec4Dot(t *testing.T) {
	a := NewVec4(1, 2, 3, 4)
	b := NewVec4(5, 6, 7, 8)
	got := a.Dot(b)
	var want float32 = 70
	assert.Equal(t, want, got)

}

func TestVec4Len(t *testing.T) {
	v := NewVec4(1, 2, 2, 4)
	got := v.Len()
	assert.Equal(t, float32(5), got)

}

func TestVec4LenSq(t *testing.T) {
	v := NewVec4(1, 2, 2, 4)
	got := v.LenSq()
	var want float32 = 25
	assert.Equal(t, want, got)

}

func TestVec4Normalize(t *testing.T) {
	v := NewVec4(0, 0, 0, 5)
	got := v.Normalize()
	want := NewVec4(0, 0, 0, 1)
	assert.True(t, got.ApproxEq(want, 1e-6))

}

func TestVec4NormalizeZero(t *testing.T) {
	v := NewVec4(0, 0, 0, 0)
	got := v.Normalize()
	want := NewVec4(0, 0, 0, 0)
	assert.True(t, got.Eq(want))

}

func TestVec4Lerp(t *testing.T) {
	a := NewVec4(0, 0, 0, 0)
	b := NewVec4(10, 20, 30, 40)
	got := a.Lerp(b, 0.5)
	want := NewVec4(5, 10, 15, 20)
	assert.True(t, got.Eq(want))

}

func TestVec4Dist(t *testing.T) {
	a := NewVec4(0, 0, 0, 0)
	b := NewVec4(1, 2, 2, 4)
	got := a.Dist(b)
	assert.Equal(t, float32(5), got)

}

func TestVec4Eq(t *testing.T) {
	a := NewVec4(1, 2, 3, 4)
	b := NewVec4(1, 2, 3, 4)
	c := NewVec4(1, 2, 3, 5)
	assert.True(t, a.Eq(b))

	assert.False(t, a.Eq(c))

}

func TestVec4ApproxEq(t *testing.T) {
	a := NewVec4(1.0, 2.0, 3.0, 4.0)
	b := NewVec4(1.0001, 2.0001, 3.0001, 4.0001)
	assert.True(t, a.ApproxEq(b, 0.001))

	assert.False(t, a.ApproxEq(b, 0.00001))

}

func TestVec4XYZ(t *testing.T) {
	v := NewVec4(1, 2, 3, 4)
	got := v.XYZ()
	want := NewVec3(1, 2, 3)
	assert.True(t, got.Eq(want))

}

func TestVec4XY(t *testing.T) {
	v := NewVec4(1, 2, 3, 4)
	got := v.XY()
	want := NewVec2(1, 2)
	assert.True(t, got.Eq(want))

}
