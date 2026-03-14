package math

import (
	stdmath "math"
	"testing"
	"github.com/wow-look-at-my/testify/assert"
)

func TestNewMat4(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	assert.False(t, m.At(0, 0) != 1 || m.At(0, 3) != 4 || m.At(3, 0) != 13 || m.At(3, 3) != 16)

}

func TestMat4Identity(t *testing.T) {
	m := Mat4Identity()
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			var want float32
			if r == c {
				want = 1
			}
			assert.Equal(t, want, m.At(r, c))

		}
	}
}

func TestMat4Col(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	got := m.Col(2)
	want := NewVec4(3, 7, 11, 15)
	assert.True(t, got.Eq(want))

}

func TestMat4Row(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	got := m.Row(2)
	want := NewVec4(9, 10, 11, 12)
	assert.True(t, got.Eq(want))

}

func TestMat4Add(t *testing.T) {
	a := Mat4Identity()
	b := Mat4Identity()
	got := a.Add(b)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			want := a.At(r, c) * 2
			assert.Equal(t, want, got.At(r, c))

		}
	}
}

func TestMat4Sub(t *testing.T) {
	a := Mat4Identity()
	got := a.Sub(a)
	zero := TMat4[float32]{}
	assert.Equal(t, zero, got)

}

func TestMat4Scale(t *testing.T) {
	m := Mat4Identity()
	got := m.Scale(3)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			want := m.At(r, c) * 3
			assert.Equal(t, want, got.At(r, c))

		}
	}
}

func TestMat4Mul(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	got := m.Mul(Mat4Identity())
	assert.Equal(t, m, got)

}

func TestMat4MulVec4(t *testing.T) {
	m := Mat4Identity()
	v := NewVec4(1, 2, 3, 4)
	got := m.MulVec4(v)
	assert.True(t, got.Eq(v))

}

func TestMat4Transpose(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	got := m.Transpose()
	want := NewMat4(
		1, 5, 9, 13,
		2, 6, 10, 14,
		3, 7, 11, 15,
		4, 8, 12, 16,
	)
	assert.Equal(t, want, got)

}

func TestMat4TransposeDouble(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	got := m.Transpose().Transpose()
	assert.Equal(t, m, got)

}

func TestMat4Det(t *testing.T) {
	m := NewMat4(
		1, 0, 0, 0,
		0, 2, 0, 0,
		0, 0, 3, 0,
		0, 0, 0, 4,
	)
	got := m.Det()
	var want float32 = 24
	assert.Equal(t, want, got)

}

func TestMat4Inverse(t *testing.T) {
	m := NewMat4(
		1, 0, 0, 1,
		0, 2, 0, 0,
		0, 0, 3, 0,
		0, 0, 0, 1,
	)
	inv := m.Inverse()
	product := m.Mul(inv)
	identity := Mat4Identity()
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			assert.False(t, diff > 1e-4 || diff < -1e-4)

		}
	}
}

func TestMat4InverseSingular(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	got := m.Inverse()
	zero := TMat4[float32]{}
	assert.Equal(t, zero, got)

}

func TestMat4Mat3(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	got := m.Mat3()
	want := NewMat3(1, 2, 3, 5, 6, 7, 9, 10, 11)
	assert.Equal(t, want, got)

}

func TestTranslation(t *testing.T) {
	m := Translation(NewVec3(2, 3, 4))
	v := NewVec4(1, 1, 1, 1)
	got := m.MulVec4(v)
	want := NewVec4(3, 4, 5, 1)
	assert.True(t, got.Eq(want))

}

func TestScaling(t *testing.T) {
	m := Scaling(NewVec3(2, 3, 4))
	v := NewVec4(1, 1, 1, 1)
	got := m.MulVec4(v)
	want := NewVec4(2, 3, 4, 1)
	assert.True(t, got.Eq(want))

}

func TestRotationX(t *testing.T) {
	m := RotationX(float32(stdmath.Pi / 2))
	v := NewVec4(0, 1, 0, 1)
	got := m.MulVec4(v)
	want := NewVec4(0, 0, 1, 1)
	assert.True(t, got.ApproxEq(want, 1e-5))

}

func TestRotationY(t *testing.T) {
	m := RotationY(float32(stdmath.Pi / 2))
	v := NewVec4(1, 0, 0, 1)
	got := m.MulVec4(v)
	want := NewVec4(0, 0, -1, 1)
	assert.True(t, got.ApproxEq(want, 1e-5))

}

func TestRotationZ(t *testing.T) {
	m := RotationZ(float32(stdmath.Pi / 2))
	v := NewVec4(1, 0, 0, 1)
	got := m.MulVec4(v)
	want := NewVec4(0, 1, 0, 1)
	assert.True(t, got.ApproxEq(want, 1e-5))

}

func TestLookAt(t *testing.T) {
	eye := NewVec3(0, 0, 5)
	center := NewVec3(0, 0, 0)
	up := NewVec3(0, 1, 0)
	m := LookAt(eye, center, up)
	// Looking down -Z, the origin should map to (0,0,-5,1) in view space
	v := m.MulVec4(NewVec4(0, 0, 0, 1))
	assert.True(t, v.ApproxEq(NewVec4(0, 0, -5, 1), 1e-5))

}

func TestPerspective(t *testing.T) {
	m := Perspective(float32(stdmath.Pi/4), 1.0, 0.1, 100.0)
	// A point on the near plane center should map to z=-1 after perspective divide
	v := m.MulVec4(NewVec4(0, 0, -0.1, 1))
	ndc := v.Scale(1.0 / v.W())
	assert.False(t, ndc.Z() < -1.01 || ndc.Z() > -0.99)

}

func TestOrtho(t *testing.T) {
	m := Ortho(-1, 1, -1, 1, -1, 1)
	// Center should map to origin
	v := m.MulVec4(NewVec4(0, 0, 0, 1))
	want := NewVec4(0, 0, 0, 1)
	assert.True(t, v.ApproxEq(want, 1e-5))

}

func TestUtilClamp(t *testing.T) {
	assert.Equal(t, float32(5), Clamp[float32](5, 0, 10))

	assert.Equal(t, float32(0), Clamp[float32](-1, 0, 10))

	assert.Equal(t, float32(10), Clamp[float32](15, 0, 10))

}
