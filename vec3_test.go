package math

import (
	"testing"
	"github.com/wow-look-at-my/testify/assert"
)

func TestNewVec3(t *testing.T) {
	v := NewVec3(1, 2, 3)
	assert.False(t, v.X() != 1 || v.Y() != 2 || v.Z() != 3)

}

func TestVec3Add(t *testing.T) {
	a := NewVec3(1, 2, 3)
	b := NewVec3(4, 5, 6)
	got := a.Add(b)
	want := NewVec3(5, 7, 9)
	assert.True(t, got.Eq(want))

}

func TestVec3Sub(t *testing.T) {
	a := NewVec3(5, 7, 9)
	b := NewVec3(1, 2, 3)
	got := a.Sub(b)
	want := NewVec3(4, 5, 6)
	assert.True(t, got.Eq(want))

}

func TestVec3Scale(t *testing.T) {
	v := NewVec3(1, 2, 3)
	got := v.Scale(3)
	want := NewVec3(3, 6, 9)
	assert.True(t, got.Eq(want))

}

func TestVec3Dot(t *testing.T) {
	a := NewVec3(1, 2, 3)
	b := NewVec3(4, 5, 6)
	got := a.Dot(b)
	var want float32 = 32
	assert.Equal(t, want, got)

}

func TestVec3Cross(t *testing.T) {
	a := NewVec3(1, 0, 0)
	b := NewVec3(0, 1, 0)
	got := a.Cross(b)
	want := NewVec3(0, 0, 1)
	assert.True(t, got.Eq(want))

}

func TestVec3CrossAnticommutative(t *testing.T) {
	a := NewVec3(2, 3, 4)
	b := NewVec3(5, 6, 7)
	ab := a.Cross(b)
	ba := b.Cross(a)
	assert.True(t, ab.ApproxEq(ba.Scale(-1), 1e-6))

}

func TestVec3Len(t *testing.T) {
	v := NewVec3(2, 3, 6)
	got := v.Len()
	assert.Equal(t, float32(7), got)

}

func TestVec3LenSq(t *testing.T) {
	v := NewVec3(2, 3, 6)
	got := v.LenSq()
	var want float32 = 49
	assert.Equal(t, want, got)

}

func TestVec3Normalize(t *testing.T) {
	v := NewVec3(0, 0, 5)
	got := v.Normalize()
	want := NewVec3(0, 0, 1)
	assert.True(t, got.ApproxEq(want, 1e-6))

}

func TestVec3NormalizeZero(t *testing.T) {
	v := NewVec3(0, 0, 0)
	got := v.Normalize()
	want := NewVec3(0, 0, 0)
	assert.True(t, got.Eq(want))

}

func TestVec3Lerp(t *testing.T) {
	a := NewVec3(0, 0, 0)
	b := NewVec3(10, 20, 30)
	got := a.Lerp(b, 0.25)
	want := NewVec3(2.5, 5, 7.5)
	assert.True(t, got.ApproxEq(want, 1e-6))

}

func TestVec3Dist(t *testing.T) {
	a := NewVec3(0, 0, 0)
	b := NewVec3(2, 3, 6)
	got := a.Dist(b)
	assert.Equal(t, float32(7), got)

}

func TestVec3Eq(t *testing.T) {
	a := NewVec3(1, 2, 3)
	b := NewVec3(1, 2, 3)
	c := NewVec3(1, 2, 4)
	assert.True(t, a.Eq(b))

	assert.False(t, a.Eq(c))

}

func TestVec3ApproxEq(t *testing.T) {
	a := NewVec3(1.0, 2.0, 3.0)
	b := NewVec3(1.0001, 2.0001, 3.0001)
	assert.True(t, a.ApproxEq(b, 0.001))

	assert.False(t, a.ApproxEq(b, 0.00001))

}

func TestVec3XY(t *testing.T) {
	v := NewVec3(1, 2, 3)
	got := v.XY()
	want := NewVec2(1, 2)
	assert.True(t, got.Eq(want))

}
