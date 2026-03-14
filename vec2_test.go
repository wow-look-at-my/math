package math

import "testing"

func TestNewVec2(t *testing.T) {
	v := NewVec2(1, 2)
	if v.X() != 1 || v.Y() != 2 {
		t.Errorf("NewVec2(1,2) = %v, want {1,2}", v)
	}
}

func TestVec2Add(t *testing.T) {
	a := NewVec2(1, 2)
	b := NewVec2(3, 4)
	got := a.Add(b)
	want := NewVec2(4, 6)
	if !got.Eq(want) {
		t.Errorf("Add: got %v, want %v", got, want)
	}
}

func TestVec2Sub(t *testing.T) {
	a := NewVec2(5, 7)
	b := NewVec2(2, 3)
	got := a.Sub(b)
	want := NewVec2(3, 4)
	if !got.Eq(want) {
		t.Errorf("Sub: got %v, want %v", got, want)
	}
}

func TestVec2Scale(t *testing.T) {
	v := NewVec2(2, 3)
	got := v.Scale(2)
	want := NewVec2(4, 6)
	if !got.Eq(want) {
		t.Errorf("Scale: got %v, want %v", got, want)
	}
}

func TestVec2Dot(t *testing.T) {
	a := NewVec2(1, 2)
	b := NewVec2(3, 4)
	got := a.Dot(b)
	if got != 11 {
		t.Errorf("Dot: got %v, want 11", got)
	}
}

func TestVec2Len(t *testing.T) {
	v := NewVec2(3, 4)
	got := v.Len()
	if got != 5 {
		t.Errorf("Len: got %v, want 5", got)
	}
}

func TestVec2LenSq(t *testing.T) {
	v := NewVec2(3, 4)
	got := v.LenSq()
	if got != 25 {
		t.Errorf("LenSq: got %v, want 25", got)
	}
}

func TestVec2Normalize(t *testing.T) {
	v := NewVec2(3, 4)
	got := v.Normalize()
	want := NewVec2(0.6, 0.8)
	if !got.ApproxEq(want, 1e-6) {
		t.Errorf("Normalize: got %v, want %v", got, want)
	}
}

func TestVec2NormalizeZero(t *testing.T) {
	v := NewVec2(0, 0)
	got := v.Normalize()
	want := NewVec2(0, 0)
	if !got.Eq(want) {
		t.Errorf("Normalize zero: got %v, want %v", got, want)
	}
}

func TestVec2Lerp(t *testing.T) {
	a := NewVec2(0, 0)
	b := NewVec2(10, 20)
	got := a.Lerp(b, 0.5)
	want := NewVec2(5, 10)
	if !got.Eq(want) {
		t.Errorf("Lerp: got %v, want %v", got, want)
	}
}

func TestVec2Dist(t *testing.T) {
	a := NewVec2(0, 0)
	b := NewVec2(3, 4)
	got := a.Dist(b)
	if got != 5 {
		t.Errorf("Dist: got %v, want 5", got)
	}
}

func TestVec2Eq(t *testing.T) {
	a := NewVec2(1, 2)
	b := NewVec2(1, 2)
	c := NewVec2(1, 3)
	if !a.Eq(b) {
		t.Error("Eq: equal vectors should be equal")
	}
	if a.Eq(c) {
		t.Error("Eq: different vectors should not be equal")
	}
}

func TestVec2ApproxEq(t *testing.T) {
	a := NewVec2(1.0, 2.0)
	b := NewVec2(1.0001, 2.0001)
	if !a.ApproxEq(b, 0.001) {
		t.Error("ApproxEq: should be approximately equal")
	}
	if a.ApproxEq(b, 0.00001) {
		t.Error("ApproxEq: should not be approximately equal with tight epsilon")
	}
}
