package math

import "testing"

func TestNewVec3(t *testing.T) {
	v := NewVec3(1, 2, 3)
	if v.X() != 1 || v.Y() != 2 || v.Z() != 3 {
		t.Errorf("NewVec3(1,2,3) = %v, want {1,2,3}", v)
	}
}

func TestVec3Add(t *testing.T) {
	a := NewVec3(1, 2, 3)
	b := NewVec3(4, 5, 6)
	got := a.Add(b)
	want := NewVec3(5, 7, 9)
	if !got.Eq(want) {
		t.Errorf("Add: got %v, want %v", got, want)
	}
}

func TestVec3Sub(t *testing.T) {
	a := NewVec3(5, 7, 9)
	b := NewVec3(1, 2, 3)
	got := a.Sub(b)
	want := NewVec3(4, 5, 6)
	if !got.Eq(want) {
		t.Errorf("Sub: got %v, want %v", got, want)
	}
}

func TestVec3Scale(t *testing.T) {
	v := NewVec3(1, 2, 3)
	got := v.Scale(3)
	want := NewVec3(3, 6, 9)
	if !got.Eq(want) {
		t.Errorf("Scale: got %v, want %v", got, want)
	}
}

func TestVec3Dot(t *testing.T) {
	a := NewVec3(1, 2, 3)
	b := NewVec3(4, 5, 6)
	got := a.Dot(b)
	var want float32 = 32
	if got != want {
		t.Errorf("Dot: got %v, want %v", got, want)
	}
}

func TestVec3Cross(t *testing.T) {
	a := NewVec3(1, 0, 0)
	b := NewVec3(0, 1, 0)
	got := a.Cross(b)
	want := NewVec3(0, 0, 1)
	if !got.Eq(want) {
		t.Errorf("Cross: got %v, want %v", got, want)
	}
}

func TestVec3CrossAnticommutative(t *testing.T) {
	a := NewVec3(2, 3, 4)
	b := NewVec3(5, 6, 7)
	ab := a.Cross(b)
	ba := b.Cross(a)
	if !ab.ApproxEq(ba.Scale(-1), 1e-6) {
		t.Errorf("Cross should be anticommutative: a×b=%v, b×a=%v", ab, ba)
	}
}

func TestVec3Len(t *testing.T) {
	v := NewVec3(2, 3, 6)
	got := v.Len()
	if got != 7 {
		t.Errorf("Len: got %v, want 7", got)
	}
}

func TestVec3LenSq(t *testing.T) {
	v := NewVec3(2, 3, 6)
	got := v.LenSq()
	var want float32 = 49
	if got != want {
		t.Errorf("LenSq: got %v, want %v", got, want)
	}
}

func TestVec3Normalize(t *testing.T) {
	v := NewVec3(0, 0, 5)
	got := v.Normalize()
	want := NewVec3(0, 0, 1)
	if !got.ApproxEq(want, 1e-6) {
		t.Errorf("Normalize: got %v, want %v", got, want)
	}
}

func TestVec3NormalizeZero(t *testing.T) {
	v := NewVec3(0, 0, 0)
	got := v.Normalize()
	want := NewVec3(0, 0, 0)
	if !got.Eq(want) {
		t.Errorf("Normalize zero: got %v, want %v", got, want)
	}
}

func TestVec3Lerp(t *testing.T) {
	a := NewVec3(0, 0, 0)
	b := NewVec3(10, 20, 30)
	got := a.Lerp(b, 0.25)
	want := NewVec3(2.5, 5, 7.5)
	if !got.ApproxEq(want, 1e-6) {
		t.Errorf("Lerp: got %v, want %v", got, want)
	}
}

func TestVec3Dist(t *testing.T) {
	a := NewVec3(0, 0, 0)
	b := NewVec3(2, 3, 6)
	got := a.Dist(b)
	if got != 7 {
		t.Errorf("Dist: got %v, want 7", got)
	}
}

func TestVec3Eq(t *testing.T) {
	a := NewVec3(1, 2, 3)
	b := NewVec3(1, 2, 3)
	c := NewVec3(1, 2, 4)
	if !a.Eq(b) {
		t.Error("Eq: equal vectors should be equal")
	}
	if a.Eq(c) {
		t.Error("Eq: different vectors should not be equal")
	}
}

func TestVec3ApproxEq(t *testing.T) {
	a := NewVec3(1.0, 2.0, 3.0)
	b := NewVec3(1.0001, 2.0001, 3.0001)
	if !a.ApproxEq(b, 0.001) {
		t.Error("ApproxEq: should be approximately equal")
	}
	if a.ApproxEq(b, 0.00001) {
		t.Error("ApproxEq: should not be approximately equal with tight epsilon")
	}
}

func TestVec3XY(t *testing.T) {
	v := NewVec3(1, 2, 3)
	got := v.XY()
	want := NewVec2(1, 2)
	if !got.Eq(want) {
		t.Errorf("XY: got %v, want %v", got, want)
	}
}
