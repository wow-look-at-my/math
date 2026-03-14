package math

import (
	stdmath "math"
	"testing"
)

func TestNewMat4(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	if m.At(0, 0) != 1 || m.At(0, 3) != 4 || m.At(3, 0) != 13 || m.At(3, 3) != 16 {
		t.Errorf("NewMat4 element access incorrect")
	}
}

func TestMat4Identity(t *testing.T) {
	m := Mat4Identity()
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			var want float32
			if r == c {
				want = 1
			}
			if m.At(r, c) != want {
				t.Errorf("Identity at (%d,%d): got %v, want %v", r, c, m.At(r, c), want)
			}
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
	if !got.Eq(want) {
		t.Errorf("Col(2): got %v, want %v", got, want)
	}
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
	if !got.Eq(want) {
		t.Errorf("Row(2): got %v, want %v", got, want)
	}
}

func TestMat4Add(t *testing.T) {
	a := Mat4Identity()
	b := Mat4Identity()
	got := a.Add(b)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			want := a.At(r, c) * 2
			if got.At(r, c) != want {
				t.Errorf("Add at (%d,%d): got %v, want %v", r, c, got.At(r, c), want)
			}
		}
	}
}

func TestMat4Sub(t *testing.T) {
	a := Mat4Identity()
	got := a.Sub(a)
	zero := TMat4[float32]{}
	if got != zero {
		t.Errorf("Sub self: got %v, want zero", got)
	}
}

func TestMat4Scale(t *testing.T) {
	m := Mat4Identity()
	got := m.Scale(3)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			want := m.At(r, c) * 3
			if got.At(r, c) != want {
				t.Errorf("Scale at (%d,%d): got %v, want %v", r, c, got.At(r, c), want)
			}
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
	if got != m {
		t.Errorf("Mul identity: got %v, want %v", got, m)
	}
}

func TestMat4MulVec4(t *testing.T) {
	m := Mat4Identity()
	v := NewVec4(1, 2, 3, 4)
	got := m.MulVec4(v)
	if !got.Eq(v) {
		t.Errorf("MulVec4 identity: got %v, want %v", got, v)
	}
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
	if got != want {
		t.Errorf("Transpose: got %v, want %v", got, want)
	}
}

func TestMat4TransposeDouble(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	)
	got := m.Transpose().Transpose()
	if got != m {
		t.Errorf("Double transpose should equal original")
	}
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
	if got != want {
		t.Errorf("Det diagonal: got %v, want %v", got, want)
	}
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
			if diff > 1e-4 || diff < -1e-4 {
				t.Errorf("M*M^-1 != I at (%d,%d): got %v", r, c, product.At(r, c))
			}
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
	if got != zero {
		t.Errorf("Inverse of singular matrix: got %v, want zero", got)
	}
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
	if got != want {
		t.Errorf("Mat3: got %v, want %v", got, want)
	}
}

func TestTranslation(t *testing.T) {
	m := Translation(NewVec3(2, 3, 4))
	v := NewVec4(1, 1, 1, 1)
	got := m.MulVec4(v)
	want := NewVec4(3, 4, 5, 1)
	if !got.Eq(want) {
		t.Errorf("Translation: got %v, want %v", got, want)
	}
}

func TestScaling(t *testing.T) {
	m := Scaling(NewVec3(2, 3, 4))
	v := NewVec4(1, 1, 1, 1)
	got := m.MulVec4(v)
	want := NewVec4(2, 3, 4, 1)
	if !got.Eq(want) {
		t.Errorf("Scaling: got %v, want %v", got, want)
	}
}

func TestRotationX(t *testing.T) {
	m := RotationX(float32(stdmath.Pi / 2))
	v := NewVec4(0, 1, 0, 1)
	got := m.MulVec4(v)
	want := NewVec4(0, 0, 1, 1)
	if !got.ApproxEq(want, 1e-5) {
		t.Errorf("RotationX: got %v, want %v", got, want)
	}
}

func TestRotationY(t *testing.T) {
	m := RotationY(float32(stdmath.Pi / 2))
	v := NewVec4(1, 0, 0, 1)
	got := m.MulVec4(v)
	want := NewVec4(0, 0, -1, 1)
	if !got.ApproxEq(want, 1e-5) {
		t.Errorf("RotationY: got %v, want %v", got, want)
	}
}

func TestRotationZ(t *testing.T) {
	m := RotationZ(float32(stdmath.Pi / 2))
	v := NewVec4(1, 0, 0, 1)
	got := m.MulVec4(v)
	want := NewVec4(0, 1, 0, 1)
	if !got.ApproxEq(want, 1e-5) {
		t.Errorf("RotationZ: got %v, want %v", got, want)
	}
}

func TestLookAt(t *testing.T) {
	eye := NewVec3(0, 0, 5)
	center := NewVec3(0, 0, 0)
	up := NewVec3(0, 1, 0)
	m := LookAt(eye, center, up)
	// Looking down -Z, the origin should map to (0,0,-5,1) in view space
	v := m.MulVec4(NewVec4(0, 0, 0, 1))
	if !v.ApproxEq(NewVec4(0, 0, -5, 1), 1e-5) {
		t.Errorf("LookAt: origin in view space = %v, want (0,0,-5,1)", v)
	}
}

func TestPerspective(t *testing.T) {
	m := Perspective(float32(stdmath.Pi/4), 1.0, 0.1, 100.0)
	// A point on the near plane center should map to z=-1 after perspective divide
	v := m.MulVec4(NewVec4(0, 0, -0.1, 1))
	ndc := v.Scale(1.0 / v.W())
	if ndc.Z() < -1.01 || ndc.Z() > -0.99 {
		t.Errorf("Perspective near plane: NDC z = %v, want ~-1", ndc.Z())
	}
}

func TestOrtho(t *testing.T) {
	m := Ortho(-1, 1, -1, 1, -1, 1)
	// Center should map to origin
	v := m.MulVec4(NewVec4(0, 0, 0, 1))
	want := NewVec4(0, 0, 0, 1)
	if !v.ApproxEq(want, 1e-5) {
		t.Errorf("Ortho center: got %v, want %v", v, want)
	}
}

func TestUtilClamp(t *testing.T) {
	if Clamp[float32](5, 0, 10) != 5 {
		t.Error("Clamp: value in range should be unchanged")
	}
	if Clamp[float32](-1, 0, 10) != 0 {
		t.Error("Clamp: value below min should be clamped to min")
	}
	if Clamp[float32](15, 0, 10) != 10 {
		t.Error("Clamp: value above max should be clamped to max")
	}
}
