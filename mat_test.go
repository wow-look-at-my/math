package math

import (
	"testing"
)

func mat4ApproxEq(a, b Mat4, e float32) bool {
	for i := range a {
		if abs(a[i]-b[i]) > e {
			return false
		}
	}
	return true
}

func mat3ApproxEq(a, b Mat3, e float32) bool {
	for i := range a {
		if abs(a[i]-b[i]) > e {
			return false
		}
	}
	return true
}

func mat2ApproxEq(a, b Mat2, e float32) bool {
	for i := range a {
		if abs(a[i]-b[i]) > e {
			return false
		}
	}
	return true
}

func TestMat2Identity(t *testing.T) {
	m := Mat2Identity()
	v := NewVec2(3, 7)
	got := m.MulVec2(v)
	if !got.ApproxEq(v, eps) {
		t.Errorf("Identity * v: got %v", got)
	}
}

func TestMat2MulInverse(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	inv := m.Inverse()
	product := m.Mul(inv)
	if !mat2ApproxEq(product, Mat2Identity(), eps) {
		t.Errorf("M * M^-1 != I: got %v", product)
	}
}

func TestMat2Det(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	if got := m.Det(); !approx(got, -2, eps) {
		t.Errorf("Det: got %v", got)
	}
}

func TestMat3Identity(t *testing.T) {
	m := Mat3Identity()
	v := NewVec3(1, 2, 3)
	got := m.MulVec3(v)
	if !got.ApproxEq(v, eps) {
		t.Errorf("Identity * v: got %v", got)
	}
}

func TestMat3MulInverse(t *testing.T) {
	m := NewMat3(
		1, 0, 2,
		0, 1, 1,
		1, 1, 0,
	)
	inv := m.Inverse()
	product := m.Mul(inv)
	if !mat3ApproxEq(product, Mat3Identity(), eps) {
		t.Errorf("M * M^-1 != I: got %v", product)
	}
}

func TestMat4Identity(t *testing.T) {
	m := Mat4Identity()
	v := NewVec4(1, 2, 3, 4)
	got := m.MulVec4(v)
	if !got.ApproxEq(v, eps) {
		t.Errorf("Identity * v: got %v", got)
	}
}

func TestMat4MulIdentity(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	got := m.Mul(Mat4Identity())
	if !mat4ApproxEq(got, m, eps) {
		t.Errorf("M * I != M: got %v", got)
	}
}

func TestMat4Inverse(t *testing.T) {
	m := NewMat4(
		1, 0, 0, 1,
		0, 2, 1, 0,
		0, 1, 2, 0,
		1, 0, 0, 2,
	)
	inv := m.Inverse()
	product := m.Mul(inv)
	if !mat4ApproxEq(product, Mat4Identity(), 1e-5) {
		t.Errorf("M * M^-1 != I: got %v", product)
	}
}

func TestMat4Transpose(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	tt := m.Transpose().Transpose()
	if !mat4ApproxEq(tt, m, eps) {
		t.Errorf("Transpose(Transpose(M)) != M")
	}
}

func TestTranslation(t *testing.T) {
	m := Translation(NewVec3(10, 20, 30))
	v := NewVec4(1, 2, 3, 1)
	got := m.MulVec4(v)
	expected := NewVec4(11, 22, 33, 1)
	if !got.ApproxEq(expected, eps) {
		t.Errorf("Translation: got %v, want %v", got, expected)
	}
}

func TestScaling(t *testing.T) {
	m := Scaling(NewVec3(2, 3, 4))
	v := NewVec4(1, 1, 1, 1)
	got := m.MulVec4(v)
	expected := NewVec4(2, 3, 4, 1)
	if !got.ApproxEq(expected, eps) {
		t.Errorf("Scaling: got %v, want %v", got, expected)
	}
}

func TestMat4ExtractMat3(t *testing.T) {
	m := Mat4Identity()
	m3 := m.Mat3()
	if !mat3ApproxEq(m3, Mat3Identity(), eps) {
		t.Errorf("Mat4.Mat3: got %v", m3)
	}
}
