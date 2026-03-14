package math

import "testing"

func TestNewMat3(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	if m.At(0, 0) != 1 || m.At(0, 2) != 3 || m.At(2, 0) != 7 || m.At(2, 2) != 9 {
		t.Errorf("NewMat3 element access incorrect: %v", m)
	}
}

func TestMat3Identity(t *testing.T) {
	m := Mat3Identity()
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
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

func TestMat3Col(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Col(1)
	want := NewVec3(2, 5, 8)
	if !got.Eq(want) {
		t.Errorf("Col(1): got %v, want %v", got, want)
	}
}

func TestMat3Row(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Row(1)
	want := NewVec3(4, 5, 6)
	if !got.Eq(want) {
		t.Errorf("Row(1): got %v, want %v", got, want)
	}
}

func TestMat3Add(t *testing.T) {
	a := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := NewMat3(9, 8, 7, 6, 5, 4, 3, 2, 1)
	got := a.Add(b)
	want := NewMat3(10, 10, 10, 10, 10, 10, 10, 10, 10)
	if got != want {
		t.Errorf("Add: got %v, want %v", got, want)
	}
}

func TestMat3Sub(t *testing.T) {
	a := NewMat3(9, 8, 7, 6, 5, 4, 3, 2, 1)
	b := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := a.Sub(b)
	want := NewMat3(8, 6, 4, 2, 0, -2, -4, -6, -8)
	if got != want {
		t.Errorf("Sub: got %v, want %v", got, want)
	}
}

func TestMat3Scale(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Scale(2)
	want := NewMat3(2, 4, 6, 8, 10, 12, 14, 16, 18)
	if got != want {
		t.Errorf("Scale: got %v, want %v", got, want)
	}
}

func TestMat3Mul(t *testing.T) {
	a := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	id := Mat3Identity()
	got := a.Mul(id)
	if got != a {
		t.Errorf("Mul identity: got %v, want %v", got, a)
	}
}

func TestMat3MulNonTrivial(t *testing.T) {
	a := NewMat3(1, 2, 0, 0, 1, 1, 1, 0, 1)
	b := NewMat3(1, 0, 1, 0, 1, 0, 1, 1, 0)
	got := a.Mul(b)
	want := NewMat3(1, 2, 1, 1, 2, 0, 2, 1, 1)
	if got != want {
		t.Errorf("Mul: got %v, want %v", got, want)
	}
}

func TestMat3MulVec3(t *testing.T) {
	m := NewMat3(1, 0, 0, 0, 1, 0, 0, 0, 1)
	v := NewVec3(3, 4, 5)
	got := m.MulVec3(v)
	if !got.Eq(v) {
		t.Errorf("MulVec3 identity: got %v, want %v", got, v)
	}
}

func TestMat3Transpose(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Transpose()
	want := NewMat3(1, 4, 7, 2, 5, 8, 3, 6, 9)
	if got != want {
		t.Errorf("Transpose: got %v, want %v", got, want)
	}
}

func TestMat3Det(t *testing.T) {
	m := NewMat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	got := m.Det()
	var want float32 = 1
	if got != want {
		t.Errorf("Det: got %v, want %v", got, want)
	}
}

func TestMat3Inverse(t *testing.T) {
	m := NewMat3(1, 2, 3, 0, 1, 4, 5, 6, 0)
	inv := m.Inverse()
	product := m.Mul(inv)
	identity := Mat3Identity()
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			diff := product.At(r, c) - identity.At(r, c)
			if diff > 1e-4 || diff < -1e-4 {
				t.Errorf("M*M^-1 != I at (%d,%d): got %v", r, c, product.At(r, c))
			}
		}
	}
}

func TestMat3InverseSingular(t *testing.T) {
	m := NewMat3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Inverse()
	zero := TMat3[float32]{}
	if got != zero {
		t.Errorf("Inverse of singular matrix: got %v, want zero", got)
	}
}
