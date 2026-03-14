package math

import (
	"testing"
)

const eps = 1e-6

func TestVec2Basic(t *testing.T) {
	a := NewVec2(1, 2)
	b := NewVec2(3, 4)

	if got := a.Add(b); got != NewVec2(4, 6) {
		t.Errorf("Add: got %v", got)
	}
	if got := a.Sub(b); got != NewVec2(-2, -2) {
		t.Errorf("Sub: got %v", got)
	}
	if got := a.Scale(2); got != NewVec2(2, 4) {
		t.Errorf("Scale: got %v", got)
	}
	if got := a.Dot(b); got != 11 {
		t.Errorf("Dot: got %v", got)
	}
}

func TestVec2Normalize(t *testing.T) {
	v := NewVec2(3, 4)
	n := v.Normalize()
	if !n.ApproxEq(NewVec2(0.6, 0.8), eps) {
		t.Errorf("Normalize: got %v", n)
	}
	if l := n.Len(); !approx(l, 1, eps) {
		t.Errorf("normalized length: got %v", l)
	}
}

func TestVec2Lerp(t *testing.T) {
	a := NewVec2(0, 0)
	b := NewVec2(10, 20)
	got := a.Lerp(b, 0.5)
	if !got.ApproxEq(NewVec2(5, 10), eps) {
		t.Errorf("Lerp: got %v", got)
	}
}

func TestVec3Cross(t *testing.T) {
	x := NewVec3(1, 0, 0)
	y := NewVec3(0, 1, 0)
	z := x.Cross(y)
	if !z.ApproxEq(NewVec3(0, 0, 1), eps) {
		t.Errorf("Cross: got %v", z)
	}
}

func TestVec3Basic(t *testing.T) {
	a := NewVec3(1, 2, 3)
	b := NewVec3(4, 5, 6)

	if got := a.Add(b); got != NewVec3(5, 7, 9) {
		t.Errorf("Add: got %v", got)
	}
	if got := a.Dot(b); got != 32 {
		t.Errorf("Dot: got %v", got)
	}
}

func TestVec4Basic(t *testing.T) {
	a := NewVec4(1, 2, 3, 4)
	b := NewVec4(5, 6, 7, 8)

	if got := a.Add(b); got != NewVec4(6, 8, 10, 12) {
		t.Errorf("Add: got %v", got)
	}
	if got := a.Dot(b); got != 70 {
		t.Errorf("Dot: got %v", got)
	}
}

func TestVec4Swizzle(t *testing.T) {
	v := NewVec4(1, 2, 3, 4)
	if got := v.XYZ(); got != NewVec3(1, 2, 3) {
		t.Errorf("XYZ: got %v", got)
	}
	if got := v.XY(); got != NewVec2(1, 2) {
		t.Errorf("XY: got %v", got)
	}
}

func TestVec2Zero(t *testing.T) {
	v := Vec2{}
	n := v.Normalize()
	if n != (Vec2{}) {
		t.Errorf("Normalize zero: got %v", n)
	}
}

func approx(a, b, e float32) bool {
	return abs(a-b) <= e
}
