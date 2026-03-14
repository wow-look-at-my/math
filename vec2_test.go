package math

import (
	"testing"
	"github.com/wow-look-at-my/testify/assert"
)

func TestNewVec2(t *testing.T) {
	v := NewVec2(1, 2)
	assert.False(t, v.X() != 1 || v.Y() != 2)

}

func TestVec2Add(t *testing.T) {
	a := NewVec2(1, 2)
	b := NewVec2(3, 4)
	got := a.Add(b)
	want := NewVec2(4, 6)
	assert.True(t, got.Eq(want))

}

func TestVec2Sub(t *testing.T) {
	a := NewVec2(5, 7)
	b := NewVec2(2, 3)
	got := a.Sub(b)
	want := NewVec2(3, 4)
	assert.True(t, got.Eq(want))

}

func TestVec2Scale(t *testing.T) {
	v := NewVec2(2, 3)
	got := v.Scale(2)
	want := NewVec2(4, 6)
	assert.True(t, got.Eq(want))

}

func TestVec2Dot(t *testing.T) {
	a := NewVec2(1, 2)
	b := NewVec2(3, 4)
	got := a.Dot(b)
	assert.Equal(t, float32(11), got)

}

func TestVec2Len(t *testing.T) {
	v := NewVec2(3, 4)
	got := v.Len()
	assert.Equal(t, float32(5), got)

}

func TestVec2LenSq(t *testing.T) {
	v := NewVec2(3, 4)
	got := v.LenSq()
	assert.Equal(t, float32(25), got)

}

func TestVec2Normalize(t *testing.T) {
	v := NewVec2(3, 4)
	got := v.Normalize()
	want := NewVec2(0.6, 0.8)
	assert.True(t, got.ApproxEq(want, 1e-6))

}

func TestVec2NormalizeZero(t *testing.T) {
	v := NewVec2(0, 0)
	got := v.Normalize()
	want := NewVec2(0, 0)
	assert.True(t, got.Eq(want))

}

func TestVec2Lerp(t *testing.T) {
	a := NewVec2(0, 0)
	b := NewVec2(10, 20)
	got := a.Lerp(b, 0.5)
	want := NewVec2(5, 10)
	assert.True(t, got.Eq(want))

}

func TestVec2Dist(t *testing.T) {
	a := NewVec2(0, 0)
	b := NewVec2(3, 4)
	got := a.Dist(b)
	assert.Equal(t, float32(5), got)

}

func TestVec2Eq(t *testing.T) {
	a := NewVec2(1, 2)
	b := NewVec2(1, 2)
	c := NewVec2(1, 3)
	assert.True(t, a.Eq(b))

	assert.False(t, a.Eq(c))

}

func TestVec2ApproxEq(t *testing.T) {
	a := NewVec2(1.0, 2.0)
	b := NewVec2(1.0001, 2.0001)
	assert.True(t, a.ApproxEq(b, 0.001))

	assert.False(t, a.ApproxEq(b, 0.00001))

}
