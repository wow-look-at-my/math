package main

import (
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/wow-look-at-my/testify/require"
)

func TestGenDir(t *testing.T) {
	dir := genDir()
	require.NotEqual(t, "", dir)
	_, err := os.Stat(filepath.Join(dir, "tvec.go.tmpl"))
	require.NoError(t, err)
}

func TestGenerateFile(t *testing.T) {
	tmpl := template.Must(template.New("test").Parse("package main\n\nvar x = {{.N}}\n"))

	outFile := filepath.Join(t.TempDir(), "test.go")

	require.NoError(t, generateFile(tmpl, outFile, struct{ N int }{N: 42}))

	data, err := os.ReadFile(outFile)
	require.NoError(t, err)
	require.Contains(t, string(data), "var x = 42")
}

func TestGenerateFileFormatError(t *testing.T) {
	tmpl := template.Must(template.New("test").Parse("not valid go {{.N}}\n"))

	err := generateFile(tmpl, filepath.Join(t.TempDir(), "test.go"), struct{ N int }{N: 42})
	require.Error(t, err)
}

func TestFuncMapSeq(t *testing.T) {
	fn := funcMap["seq"].(func(int) []int)
	got := fn(4)
	require.Equal(t, 4, len(got))
	for i, v := range got {
		require.Equal(t, i, v)
	}
}

func TestFuncMapComponent(t *testing.T) {
	fn := funcMap["component"].(func(int) string)
	expected := []string{"X", "Y", "Z", "W"}
	for i, want := range expected {
		require.Equal(t, want, fn(i))
	}
}

func TestFuncMapParam(t *testing.T) {
	fn := funcMap["param"].(func(int) string)
	expected := []string{"x", "y", "z", "w"}
	for i, want := range expected {
		require.Equal(t, want, fn(i))
	}
}

func TestFuncMapArithmetic(t *testing.T) {
	add := funcMap["add"].(func(int, int) int)
	sub := funcMap["sub"].(func(int, int) int)
	mul := funcMap["mul"].(func(int, int) int)

	require.Equal(t, 7, add(3, 4))
	require.Equal(t, 7, sub(10, 3))
	require.Equal(t, 12, mul(3, 4))
}

func TestFuncMapLast(t *testing.T) {
	fn := funcMap["last"].(func(int, int) bool)
	require.True(t, fn(3, 4))
	require.False(t, fn(2, 4))
}

func TestFuncMapIntList(t *testing.T) {
	fn := funcMap["intList"].(func(int, int) string)
	require.Equal(t, "1, 2, 3", fn(3, 1))
}

func TestFuncMapInvMat(t *testing.T) {
	fn := funcMap["invMat"].(func(int) string)
	require.Equal(t, "1, 2, 3, 4", fn(2))
	require.Equal(t, "1, 2, 3, 0, 1, 4, 5, 6, 0", fn(3))
	require.Equal(t, "1, 0, 0, 1, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 1", fn(4))
	require.Equal(t, "", fn(5))
}

func TestTemplatesParse(t *testing.T) {
	dir := genDir()
	templates := []string{
		"tvec.go.tmpl", "tmat.go.tmpl", "tquat.go.tmpl",
		"fvec.go.tmpl", "fmat.go.tmpl", "fquat.go.tmpl",
		"ivec.go.tmpl", "imat.go.tmpl",
		"itest.go.tmpl", "ftest.go.tmpl",
		"bench.go.tmpl",
	}
	for _, name := range templates {
		t.Run(name, func(t *testing.T) {
			_, err := template.New(name).Funcs(funcMap).ParseFiles(filepath.Join(dir, name))
			require.NoError(t, err)
		})
	}
}

func TestGenerateVec(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	vecTmpl := template.Must(
		template.New("tvec.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "tvec.go.tmpl")),
	)
	outFile := filepath.Join(tmpDir, "vec2.go")
	require.NoError(t, generateFile(vecTmpl, outFile, vecData{N: 2}))

	data, err := os.ReadFile(outFile)
	require.NoError(t, err)
	require.Contains(t, string(data), "TVec2")
}

func TestGenerateMat(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	matTmpl := template.Must(
		template.New("tmat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "tmat.go.tmpl")),
	)
	outFile := filepath.Join(tmpDir, "mat3.go")
	require.NoError(t, generateFile(matTmpl, outFile, matData{N: 3, N2: 9}))

	data, err := os.ReadFile(outFile)
	require.NoError(t, err)
	require.Contains(t, string(data), "TMat3")
}

func TestGenerateQuat(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	quatTmpl := template.Must(
		template.New("tquat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "tquat.go.tmpl")),
	)
	outFile := filepath.Join(tmpDir, "quat.go")
	require.NoError(t, generateFile(quatTmpl, outFile, nil))

	data, err := os.ReadFile(outFile)
	require.NoError(t, err)
	require.Contains(t, string(data), "TQuat")
}

func TestGenerateConcreteFloatTypes(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	fvecTmpl := template.Must(
		template.New("fvec.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "fvec.go.tmpl")),
	)
	fmatTmpl := template.Must(
		template.New("fmat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "fmat.go.tmpl")),
	)
	fquatTmpl := template.Must(
		template.New("fquat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "fquat.go.tmpl")),
	)

	floatTypes := []struct {
		Prefix     string
		ElemType   string
		FilePrefix string
	}{
		{"", "float32", "f"},
		{"D", "float64", "d"},
	}

	for _, ft := range floatTypes {
		for _, n := range []int{2, 3, 4} {
			outFile := filepath.Join(tmpDir, ft.FilePrefix+"vec.go")
			require.NoError(t, generateFile(fvecTmpl, outFile, ivecData{
				N: n, Prefix: ft.Prefix, ElemType: ft.ElemType,
			}))

			outFile = filepath.Join(tmpDir, ft.FilePrefix+"mat.go")
			require.NoError(t, generateFile(fmatTmpl, outFile, imatData{
				N: n, N2: n * n, Prefix: ft.Prefix, ElemType: ft.ElemType,
			}))
		}

		outFile := filepath.Join(tmpDir, ft.FilePrefix+"quat.go")
		require.NoError(t, generateFile(fquatTmpl, outFile, ivecData{
			N: 4, Prefix: ft.Prefix, ElemType: ft.ElemType,
		}))
	}
}

func TestGenerateIntegerTypes(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	ivecTmpl := template.Must(
		template.New("ivec.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "ivec.go.tmpl")),
	)
	imatTmpl := template.Must(
		template.New("imat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "imat.go.tmpl")),
	)

	intTypes := []struct {
		Prefix string
		GoType string
	}{
		{"I32", "int32"},
		{"U32", "uint32"},
	}

	for _, it := range intTypes {
		outFile := filepath.Join(tmpDir, "vec.go")
		require.NoError(t, generateFile(ivecTmpl, outFile, ivecData{
			N: 3, Prefix: it.Prefix, ElemType: it.GoType,
		}))

		outFile = filepath.Join(tmpDir, "mat.go")
		require.NoError(t, generateFile(imatTmpl, outFile, imatData{
			N: 3, N2: 9, Prefix: it.Prefix, ElemType: it.GoType,
		}))
	}
}

func TestGenerateTests(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	ftestTmpl := template.Must(
		template.New("ftest.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "ftest.go.tmpl")),
	)
	outFile := filepath.Join(tmpDir, "f_test.go")
	require.NoError(t, generateFile(ftestTmpl, outFile, itestData{
		Prefix: "", ElemType: "float32",
	}))

	itestTmpl := template.Must(
		template.New("itest.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "itest.go.tmpl")),
	)
	outFile = filepath.Join(tmpDir, "i32_test.go")
	require.NoError(t, generateFile(itestTmpl, outFile, itestData{
		Prefix: "I32", ElemType: "int32", Signed: true,
	}))
}

func TestGenerateBenchmarks(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	benchTmpl := template.Must(
		template.New("bench.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "bench.go.tmpl")),
	)
	outFile := filepath.Join(tmpDir, "f32_bench_test.go")
	require.NoError(t, generateFile(benchTmpl, outFile, benchData{
		Name: "F32", Prefix: "", ElemType: "float32", LenType: "float32",
		LerpT: "0.5", HasVecApproxEq: true, HasMatApproxEq: true,
		HasFloat32: false, HasMat3: true, Epsilon: "1e-6",
	}))
}
