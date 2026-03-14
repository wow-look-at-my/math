package math

import (
	"testing"

	"github.com/wow-look-at-my/testify/assert"
)

func TestNewMat3(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	assert.False(t, m.At(0, 0) != 1 || m.At(0, 2) != 3 || m.At(2, 0) != 7 || m.At(2, 2) != 9)
}

func TestMat3Identity(t *testing.T) {
	m := Mat3Identity()
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			var want float32
			if r == c {
				want = 1
			}
			assert.Equal(t, want, m.At(r, c))
		}
	}
}

func TestMat3Col(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Col(1)
	want := NewVec3(2, 5, 8)
	assert.True(t, got.Eq(want))
}

func TestMat3Row(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Row(1)
	want := NewVec3(4, 5, 6)
	assert.True(t, got.Eq(want))
}

func TestMat3Add(t *testing.T) {
	a := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := NewMat3(9, 8, 7, 6, 5, 4, 3, 2, 1)
	got := a.Add(b)
	want := NewMat3(10, 10, 10, 10, 10, 10, 10, 10, 10)
	assert.Equal(t, want, got)
}

func TestMat3Sub(t *testing.T) {
	a := NewMat3(9, 8, 7, 6, 5, 4, 3, 2, 1)
	b := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := a.Sub(b)
	want := NewMat3(8, 6, 4, 2, 0, -2, -4, -6, -8)
	assert.Equal(t, want, got)
}

func TestMat3Scale(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Scale(2)
	want := NewMat3(2, 4, 6, 8, 10, 12, 14, 16, 18)
	assert.Equal(t, want, got)
}

func TestMat3Mul(t *testing.T) {
	a := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	id := Mat3Identity()
	got := a.Mul(id)
	assert.Equal(t, a, got)
}

func TestMat3MulNonTrivial(t *testing.T) {
	a := NewMat3(1, 2, 0, 0, 1, 1, 1, 0, 1)
	b := NewMat3(1, 0, 1, 0, 1, 0, 1, 1, 0)
	got := a.Mul(b)
	want := NewMat3(1, 2, 1, 1, 2, 0, 2, 1, 1)
	assert.Equal(t, want, got)
}

func TestMat3MulVec3(t *testing.T) {
	m := NewMat3(1, 0, 0, 0, 1, 0, 0, 0, 1)
	v := NewVec3(3, 4, 5)
	got := m.MulVec3(v)
	assert.True(t, got.Eq(v))
}

func TestMat3Transpose(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Transpose()
	want := NewMat3(1, 4, 7, 2, 5, 8, 3, 6, 9)
	assert.Equal(t, want, got)
}

func TestMat3Det(t *testing.T) {
	m := NewMat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	got := m.Det()
	var want float32 = 1
	assert.Equal(t, want, got)
}

func TestMat3Inverse(t *testing.T) {
	m := NewMat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	inv := m.Inverse()
	product := m.Mul(inv)
	identity := Mat3Identity()
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			assert.False(t, diff > 1e-4 || diff < -1e-4)
		}
	}
}

func TestMat3InverseSingular(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Inverse()
	zero := Mat3{}
	assert.Equal(t, zero, got)
}
