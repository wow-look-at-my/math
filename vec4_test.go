package math

import "testing"

func TestNewVec4(t *testing.T) {
	v := NewVec4(1, 2, 3, 4)
	if v.X() != 1 || v.Y() != 2 || v.Z() != 3 || v.W() != 4 {
		t.Errorf("NewVec4(1,2,3,4) = %v, want {1,2,3,4}", v)
	}
}

func TestVec4Add(t *testing.T) {
	a := NewVec4(1, 2, 3, 4)
	b := NewVec4(5, 6, 7, 8)
	got := a.Add(b)
	want := NewVec4(6, 8, 10, 12)
	if !got.Eq(want) {
		t.Errorf("Add: got %v, want %v", got, want)
	}
}

func TestVec4Sub(t *testing.T) {
	a := NewVec4(5, 7, 9, 11)
	b := NewVec4(1, 2, 3, 4)
	got := a.Sub(b)
	want := NewVec4(4, 5, 6, 7)
	if !got.Eq(want) {
		t.Errorf("Sub: got %v, want %v", got, want)
	}
}

func TestVec4Scale(t *testing.T) {
	v := NewVec4(1, 2, 3, 4)
	got := v.Scale(2)
	want := NewVec4(2, 4, 6, 8)
	if !got.Eq(want) {
		t.Errorf("Scale: got %v, want %v", got, want)
	}
}

func TestVec4Dot(t *testing.T) {
	a := NewVec4(1, 2, 3, 4)
	b := NewVec4(5, 6, 7, 8)
	got := a.Dot(b)
	var want float32 = 70
	if got != want {
		t.Errorf("Dot: got %v, want %v", got, want)
	}
}

func TestVec4Len(t *testing.T) {
	v := NewVec4(1, 2, 2, 4)
	got := v.Len()
	if got != 5 {
		t.Errorf("Len: got %v, want 5", got)
	}
}

func TestVec4LenSq(t *testing.T) {
	v := NewVec4(1, 2, 2, 4)
	got := v.LenSq()
	var want float32 = 25
	if got != want {
		t.Errorf("LenSq: got %v, want %v", got, want)
	}
}

func TestVec4Normalize(t *testing.T) {
	v := NewVec4(0, 0, 0, 5)
	got := v.Normalize()
	want := NewVec4(0, 0, 0, 1)
	if !got.ApproxEq(want, 1e-6) {
		t.Errorf("Normalize: got %v, want %v", got, want)
	}
}

func TestVec4NormalizeZero(t *testing.T) {
	v := NewVec4(0, 0, 0, 0)
	got := v.Normalize()
	want := NewVec4(0, 0, 0, 0)
	if !got.Eq(want) {
		t.Errorf("Normalize zero: got %v, want %v", got, want)
	}
}

func TestVec4Lerp(t *testing.T) {
	a := NewVec4(0, 0, 0, 0)
	b := NewVec4(10, 20, 30, 40)
	got := a.Lerp(b, 0.5)
	want := NewVec4(5, 10, 15, 20)
	if !got.Eq(want) {
		t.Errorf("Lerp: got %v, want %v", got, want)
	}
}

func TestVec4Dist(t *testing.T) {
	a := NewVec4(0, 0, 0, 0)
	b := NewVec4(1, 2, 2, 4)
	got := a.Dist(b)
	if got != 5 {
		t.Errorf("Dist: got %v, want 5", got)
	}
}

func TestVec4Eq(t *testing.T) {
	a := NewVec4(1, 2, 3, 4)
	b := NewVec4(1, 2, 3, 4)
	c := NewVec4(1, 2, 3, 5)
	if !a.Eq(b) {
		t.Error("Eq: equal vectors should be equal")
	}
	if a.Eq(c) {
		t.Error("Eq: different vectors should not be equal")
	}
}

func TestVec4ApproxEq(t *testing.T) {
	a := NewVec4(1.0, 2.0, 3.0, 4.0)
	b := NewVec4(1.0001, 2.0001, 3.0001, 4.0001)
	if !a.ApproxEq(b, 0.001) {
		t.Error("ApproxEq: should be approximately equal")
	}
	if a.ApproxEq(b, 0.00001) {
		t.Error("ApproxEq: should not be approximately equal with tight epsilon")
	}
}

func TestVec4XYZ(t *testing.T) {
	v := NewVec4(1, 2, 3, 4)
	got := v.XYZ()
	want := NewVec3(1, 2, 3)
	if !got.Eq(want) {
		t.Errorf("XYZ: got %v, want %v", got, want)
	}
}

func TestVec4XY(t *testing.T) {
	v := NewVec4(1, 2, 3, 4)
	got := v.XY()
	want := NewVec2(1, 2)
	if !got.Eq(want) {
		t.Errorf("XY: got %v, want %v", got, want)
	}
}
