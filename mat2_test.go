package math

import (
	"testing"

	"github.com/wow-look-at-my/testify/assert"
)

func TestNewMat2(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	assert.False(t, m.At(0, 0) != 1 || m.At(0, 1) != 2 || m.At(1, 0) != 3 || m.At(1, 1) != 4)
}

func TestMat2Identity(t *testing.T) {
	m := Mat2Identity()
	assert.False(t, m.At(0, 0) != 1 || m.At(0, 1) != 0 || m.At(1, 0) != 0 || m.At(1, 1) != 1)
}

func TestMat2Col(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Col(0)
	want := NewVec2(1, 3)
	assert.True(t, got.Eq(want))

	got = m.Col(1)
	want = NewVec2(2, 4)
	assert.True(t, got.Eq(want))
}

func TestMat2Row(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Row(0)
	want := NewVec2(1, 2)
	assert.True(t, got.Eq(want))

	got = m.Row(1)
	want = NewVec2(3, 4)
	assert.True(t, got.Eq(want))
}

func TestMat2Add(t *testing.T) {
	a := NewMat2(1, 2, 3, 4)
	b := NewMat2(5, 6, 7, 8)
	got := a.Add(b)
	want := NewMat2(6, 8, 10, 12)
	assert.Equal(t, want, got)
}

func TestMat2Sub(t *testing.T) {
	a := NewMat2(5, 6, 7, 8)
	b := NewMat2(1, 2, 3, 4)
	got := a.Sub(b)
	want := NewMat2(4, 4, 4, 4)
	assert.Equal(t, want, got)
}

func TestMat2Scale(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Scale(2)
	want := NewMat2(2, 4, 6, 8)
	assert.Equal(t, want, got)
}

func TestMat2Mul(t *testing.T) {
	a := NewMat2(1, 2, 3, 4)
	b := NewMat2(5, 6, 7, 8)
	got := a.Mul(b)
	want := NewMat2(19, 22, 43, 50)
	assert.Equal(t, want, got)
}

func TestMat2MulIdentity(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Mul(Mat2Identity())
	assert.Equal(t, m, got)
}

func TestMat2MulVec2(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	v := NewVec2(5, 6)
	got := m.MulVec2(v)
	want := NewVec2(17, 39)
	assert.True(t, got.Eq(want))
}

func TestMat2Transpose(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Transpose()
	want := NewMat2(1, 3, 2, 4)
	assert.Equal(t, want, got)
}

func TestMat2Det(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Det()
	var want float32 = -2
	assert.Equal(t, want, got)
}

func TestMat2Inverse(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	inv := m.Inverse()
	product := m.Mul(inv)
	identity := Mat2Identity()
	for r := 0; r < 2; r++ {
		for c := 0; c < 2; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			assert.False(t, diff > 1e-5 || diff < -1e-5)
		}
	}
}

func TestMat2InverseSingular(t *testing.T) {
	m := NewMat2(1, 2, 2, 4)
	got := m.Inverse()
	zero := Mat2{}
	assert.Equal(t, zero, got)
}
