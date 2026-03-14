package math

import (
	"testing"

	"github.com/wow-look-at-my/testify/assert"
)

// ---------------------------------------------------------------------------
// Generic TVec — basic ops with various types
// ---------------------------------------------------------------------------

func TestTVec2IntAdd(t *testing.T) {
	a := TVec2[int]{X: 1, Y: 2}
	b := TVec2[int]{X: 3, Y: 4}
	got := a.Add(b)
	assert.True(t, got.Eq(TVec2[int]{X: 4, Y: 6}))
}

func TestTVec2IntSub(t *testing.T) {
	a := TVec2[int]{X: 5, Y: 7}
	b := TVec2[int]{X: 2, Y: 3}
	got := a.Sub(b)
	assert.True(t, got.Eq(TVec2[int]{X: 3, Y: 4}))
}

func TestTVec2IntScale(t *testing.T) {
	v := TVec2[int]{X: 2, Y: 3}
	got := v.Scale(3)
	assert.True(t, got.Eq(TVec2[int]{X: 6, Y: 9}))
}

func TestTVec2IntDot(t *testing.T) {
	a := TVec2[int]{X: 1, Y: 2}
	b := TVec2[int]{X: 3, Y: 4}
	assert.Equal(t, 11, a.Dot(b))
}

func TestTVec2IntLenSq(t *testing.T) {
	assert.Equal(t, 25, TVec2[int]{X: 3, Y: 4}.LenSq())
}

func TestTVec2IntEq(t *testing.T) {
	a := TVec2[int]{X: 1, Y: 2}
	assert.True(t, a.Eq(TVec2[int]{X: 1, Y: 2}))
	assert.False(t, a.Eq(TVec2[int]{X: 1, Y: 3}))
}

func TestTVec2Float64(t *testing.T) {
	a := TVec2[float64]{X: 1.5, Y: 2.5}
	b := TVec2[float64]{X: 3.5, Y: 4.5}
	got := a.Add(b)
	assert.True(t, got.Eq(TVec2[float64]{X: 5.0, Y: 7.0}))
}

func TestTVec3IntAdd(t *testing.T) {
	a := TVec3[int]{X: 1, Y: 2, Z: 3}
	b := TVec3[int]{X: 4, Y: 5, Z: 6}
	got := a.Add(b)
	assert.True(t, got.Eq(TVec3[int]{X: 5, Y: 7, Z: 9}))
}

func TestTVec3IntCross(t *testing.T) {
	a := TVec3[int]{X: 1, Y: 0, Z: 0}
	b := TVec3[int]{X: 0, Y: 1, Z: 0}
	got := a.Cross(b)
	assert.True(t, got.Eq(TVec3[int]{X: 0, Y: 0, Z: 1}))
}

func TestTVec3IntDot(t *testing.T) {
	a := TVec3[int]{X: 1, Y: 2, Z: 3}
	b := TVec3[int]{X: 4, Y: 5, Z: 6}
	assert.Equal(t, 32, a.Dot(b))
}

func TestTVec3IntXY(t *testing.T) {
	v := TVec3[int]{X: 1, Y: 2, Z: 3}
	got := v.XY()
	assert.True(t, got.Eq(TVec2[int]{X: 1, Y: 2}))
}

func TestTVec4IntScale(t *testing.T) {
	v := TVec4[int]{X: 1, Y: 2, Z: 3, W: 4}
	got := v.Scale(2)
	assert.True(t, got.Eq(TVec4[int]{X: 2, Y: 4, Z: 6, W: 8}))
}

func TestTVec4IntXYZ(t *testing.T) {
	v := TVec4[int]{X: 1, Y: 2, Z: 3, W: 4}
	got := v.XYZ()
	assert.True(t, got.Eq(TVec3[int]{X: 1, Y: 2, Z: 3}))
}

func TestTVec4IntXY(t *testing.T) {
	v := TVec4[int]{X: 1, Y: 2, Z: 3, W: 4}
	got := v.XY()
	assert.True(t, got.Eq(TVec2[int]{X: 1, Y: 2}))
}

// ---------------------------------------------------------------------------
// Generic TMat — basic ops
// ---------------------------------------------------------------------------

func TestTMat2IntMul(t *testing.T) {
	a := NewTMat2[int](1, 2, 3, 4)
	b := NewTMat2[int](5, 6, 7, 8)
	got := a.Mul(b)
	assert.True(t, got.Eq(NewTMat2[int](19, 22, 43, 50)))
}

func TestTMat2IntCol(t *testing.T) {
	m := NewTMat2[int](1, 2, 3, 4)
	assert.True(t, m.Col(0).Eq(TVec2[int]{X: 1, Y: 3}))
}

func TestTMat2IntRow(t *testing.T) {
	m := NewTMat2[int](1, 2, 3, 4)
	assert.True(t, m.Row(0).Eq(TVec2[int]{X: 1, Y: 2}))
}

func TestTMat2IntDet(t *testing.T) {
	m := NewTMat2[int](1, 2, 3, 4)
	assert.Equal(t, -2, m.Det())
}

func TestTMat3IntTranspose(t *testing.T) {
	m := NewTMat3[int](1, 2, 3, 4, 5, 6, 7, 8, 9)
	got := m.Transpose().Transpose()
	assert.True(t, got.Eq(m))
}

func TestTMat3IntDet(t *testing.T) {
	m := NewTMat3[int](1, 2, 3, 0, 1, 4, 5, 6, 0)
	assert.Equal(t, 1, m.Det())
}

func TestTMat3IntMulVec3(t *testing.T) {
	m := NewTMat3[int](1, 0, 0, 0, 1, 0, 0, 0, 1)
	v := TVec3[int]{X: 3, Y: 4, Z: 5}
	assert.True(t, m.MulVec3(v).Eq(v))
}

func TestTMat4IntMulVec4(t *testing.T) {
	m := NewTMat4[int](1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)
	v := TVec4[int]{X: 1, Y: 2, Z: 3, W: 4}
	got := m.MulVec4(v)
	assert.True(t, got.Eq(v))
}

func TestTMat4IntDet(t *testing.T) {
	m := NewTMat4[int](1, 0, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 4)
	assert.Equal(t, 24, m.Det())
}
