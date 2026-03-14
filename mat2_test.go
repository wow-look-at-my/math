package math

import "testing"

func TestNewMat2(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	if m.At(0, 0) != 1 || m.At(0, 1) != 2 || m.At(1, 0) != 3 || m.At(1, 1) != 4 {
		t.Errorf("NewMat2 element access incorrect: %v", m)
	}
}

func TestMat2Identity(t *testing.T) {
	m := Mat2Identity()
	if m.At(0, 0) != 1 || m.At(0, 1) != 0 || m.At(1, 0) != 0 || m.At(1, 1) != 1 {
		t.Errorf("Mat2Identity: got %v", m)
	}
}

func TestMat2Col(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Col(0)
	want := NewVec2(1, 3)
	if !got.Eq(want) {
		t.Errorf("Col(0): got %v, want %v", got, want)
	}
	got = m.Col(1)
	want = NewVec2(2, 4)
	if !got.Eq(want) {
		t.Errorf("Col(1): got %v, want %v", got, want)
	}
}

func TestMat2Row(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Row(0)
	want := NewVec2(1, 2)
	if !got.Eq(want) {
		t.Errorf("Row(0): got %v, want %v", got, want)
	}
	got = m.Row(1)
	want = NewVec2(3, 4)
	if !got.Eq(want) {
		t.Errorf("Row(1): got %v, want %v", got, want)
	}
}

func TestMat2Add(t *testing.T) {
	a := NewMat2(1, 2, 3, 4)
	b := NewMat2(5, 6, 7, 8)
	got := a.Add(b)
	want := NewMat2(6, 8, 10, 12)
	if got != want {
		t.Errorf("Add: got %v, want %v", got, want)
	}
}

func TestMat2Sub(t *testing.T) {
	a := NewMat2(5, 6, 7, 8)
	b := NewMat2(1, 2, 3, 4)
	got := a.Sub(b)
	want := NewMat2(4, 4, 4, 4)
	if got != want {
		t.Errorf("Sub: got %v, want %v", got, want)
	}
}

func TestMat2Scale(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Scale(2)
	want := NewMat2(2, 4, 6, 8)
	if got != want {
		t.Errorf("Scale: got %v, want %v", got, want)
	}
}

func TestMat2Mul(t *testing.T) {
	a := NewMat2(1, 2, 3, 4)
	b := NewMat2(5, 6, 7, 8)
	got := a.Mul(b)
	want := NewMat2(19, 22, 43, 50)
	if got != want {
		t.Errorf("Mul: got %v, want %v", got, want)
	}
}

func TestMat2MulIdentity(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Mul(Mat2Identity())
	if got != m {
		t.Errorf("Mul identity: got %v, want %v", got, m)
	}
}

func TestMat2MulVec2(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	v := NewVec2(5, 6)
	got := m.MulVec2(v)
	want := NewVec2(17, 39)
	if !got.Eq(want) {
		t.Errorf("MulVec2: got %v, want %v", got, want)
	}
}

func TestMat2Transpose(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Transpose()
	want := NewMat2(1, 3, 2, 4)
	if got != want {
		t.Errorf("Transpose: got %v, want %v", got, want)
	}
}

func TestMat2Det(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	got := m.Det()
	var want float32 = -2
	if got != want {
		t.Errorf("Det: got %v, want %v", got, want)
	}
}

func TestMat2Inverse(t *testing.T) {
	m := NewMat2(1, 2, 3, 4)
	inv := m.Inverse()
	product := m.Mul(inv)
	identity := Mat2Identity()
	for r := 0; r < 2; r++ {
		for c := 0; c < 2; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			if diff > 1e-5 || diff < -1e-5 {
				t.Errorf("M*M^-1 != I at (%d,%d): got %v", r, c, product.At(r, c))
			}
		}
	}
}

func TestMat2InverseSingular(t *testing.T) {
	m := NewMat2(1, 2, 2, 4)
	got := m.Inverse()
	zero := TMat2[float32]{}
	if got != zero {
		t.Errorf("Inverse of singular matrix: got %v, want zero", got)
	}
}
