package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
)

func TestGenDir(t *testing.T) {
	dir := genDir()
	if dir == "" {
		t.Fatal("genDir returned empty string")
	}
	if _, err := os.Stat(filepath.Join(dir, "tvec.go.tmpl")); err != nil {
		t.Fatalf("genDir does not contain expected template files: %v", err)
	}
}

func TestGenerateFile(t *testing.T) {
	tmpl := template.Must(template.New("test").Parse("package main\n\nvar x = {{.N}}\n"))

	outFile := filepath.Join(t.TempDir(), "test.go")

	if err := generateFile(tmpl, outFile, struct{ N int }{N: 42}); err != nil {
		t.Fatalf("generateFile failed: %v", err)
	}

	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("reading generated file: %v", err)
	}
	if !strings.Contains(string(data), "var x = 42") {
		t.Fatalf("generated file does not contain expected content: %s", data)
	}
}

func TestGenerateFileFormatError(t *testing.T) {
	tmpl := template.Must(template.New("test").Parse("not valid go {{.N}}\n"))

	err := generateFile(tmpl, filepath.Join(t.TempDir(), "test.go"), struct{ N int }{N: 42})
	if err == nil {
		t.Fatal("expected error for invalid go source")
	}
}

func TestFuncMapSeq(t *testing.T) {
	fn := funcMap["seq"].(func(int) []int)
	got := fn(4)
	if len(got) != 4 {
		t.Fatalf("seq(4) returned %d elements", len(got))
	}
	for i, v := range got {
		if v != i {
			t.Fatalf("seq(4)[%d] = %d", i, v)
		}
	}
}

func TestFuncMapComponent(t *testing.T) {
	fn := funcMap["component"].(func(int) string)
	expected := []string{"X", "Y", "Z", "W"}
	for i, want := range expected {
		if got := fn(i); got != want {
			t.Fatalf("component(%d) = %q, want %q", i, got, want)
		}
	}
}

func TestFuncMapParam(t *testing.T) {
	fn := funcMap["param"].(func(int) string)
	expected := []string{"x", "y", "z", "w"}
	for i, want := range expected {
		if got := fn(i); got != want {
			t.Fatalf("param(%d) = %q, want %q", i, got, want)
		}
	}
}

func TestFuncMapArithmetic(t *testing.T) {
	add := funcMap["add"].(func(int, int) int)
	sub := funcMap["sub"].(func(int, int) int)
	mul := funcMap["mul"].(func(int, int) int)

	if got := add(3, 4); got != 7 {
		t.Fatalf("add(3,4) = %d", got)
	}
	if got := sub(10, 3); got != 7 {
		t.Fatalf("sub(10,3) = %d", got)
	}
	if got := mul(3, 4); got != 12 {
		t.Fatalf("mul(3,4) = %d", got)
	}
}

func TestFuncMapLast(t *testing.T) {
	fn := funcMap["last"].(func(int, int) bool)
	if !fn(3, 4) {
		t.Fatal("last(3,4) should be true")
	}
	if fn(2, 4) {
		t.Fatal("last(2,4) should be false")
	}
}

func TestFuncMapIntList(t *testing.T) {
	fn := funcMap["intList"].(func(int, int) string)
	if got := fn(3, 1); got != "1, 2, 3" {
		t.Fatalf("intList(3,1) = %q", got)
	}
}

func TestFuncMapInvMat(t *testing.T) {
	fn := funcMap["invMat"].(func(int) string)
	if got := fn(2); got != "1, 2, 3, 4" {
		t.Fatalf("invMat(2) = %q", got)
	}
	if got := fn(3); got != "1, 2, 3, 0, 1, 4, 5, 6, 0" {
		t.Fatalf("invMat(3) = %q", got)
	}
	if got := fn(4); got != "1, 0, 0, 1, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 1" {
		t.Fatalf("invMat(4) = %q", got)
	}
	if got := fn(5); got != "" {
		t.Fatalf("invMat(5) = %q", got)
	}
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
			if err != nil {
				t.Fatalf("failed to parse template %s: %v", name, err)
			}
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
	if err := generateFile(vecTmpl, outFile, vecData{N: 2}); err != nil {
		t.Fatalf("failed to generate vec2.go: %v", err)
	}

	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("reading generated file: %v", err)
	}
	if !strings.Contains(string(data), "TVec2") {
		t.Fatal("generated vec2.go does not contain TVec2")
	}
}

func TestGenerateMat(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	matTmpl := template.Must(
		template.New("tmat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "tmat.go.tmpl")),
	)
	outFile := filepath.Join(tmpDir, "mat3.go")
	if err := generateFile(matTmpl, outFile, matData{N: 3, N2: 9}); err != nil {
		t.Fatalf("failed to generate mat3.go: %v", err)
	}

	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("reading generated file: %v", err)
	}
	if !strings.Contains(string(data), "TMat3") {
		t.Fatal("generated mat3.go does not contain TMat3")
	}
}

func TestGenerateQuat(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	quatTmpl := template.Must(
		template.New("tquat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "tquat.go.tmpl")),
	)
	outFile := filepath.Join(tmpDir, "quat.go")
	if err := generateFile(quatTmpl, outFile, nil); err != nil {
		t.Fatalf("failed to generate quat.go: %v", err)
	}

	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("reading generated file: %v", err)
	}
	if !strings.Contains(string(data), "TQuat") {
		t.Fatal("generated quat.go does not contain TQuat")
	}
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
			if err := generateFile(fvecTmpl, outFile, ivecData{
				N: n, Prefix: ft.Prefix, ElemType: ft.ElemType,
			}); err != nil {
				t.Fatalf("failed to generate %svec%d.go: %v", ft.FilePrefix, n, err)
			}

			outFile = filepath.Join(tmpDir, ft.FilePrefix+"mat.go")
			if err := generateFile(fmatTmpl, outFile, imatData{
				N: n, N2: n * n, Prefix: ft.Prefix, ElemType: ft.ElemType,
			}); err != nil {
				t.Fatalf("failed to generate %smat%d.go: %v", ft.FilePrefix, n, err)
			}
		}

		outFile := filepath.Join(tmpDir, ft.FilePrefix+"quat.go")
		if err := generateFile(fquatTmpl, outFile, ivecData{
			N: 4, Prefix: ft.Prefix, ElemType: ft.ElemType,
		}); err != nil {
			t.Fatalf("failed to generate %squat.go: %v", ft.FilePrefix, err)
		}
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
		if err := generateFile(ivecTmpl, outFile, ivecData{
			N: 3, Prefix: it.Prefix, ElemType: it.GoType,
		}); err != nil {
			t.Fatalf("failed to generate %s vec3: %v", it.Prefix, err)
		}

		outFile = filepath.Join(tmpDir, "mat.go")
		if err := generateFile(imatTmpl, outFile, imatData{
			N: 3, N2: 9, Prefix: it.Prefix, ElemType: it.GoType,
		}); err != nil {
			t.Fatalf("failed to generate %s mat3: %v", it.Prefix, err)
		}
	}
}

func TestGenerateTests(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	ftestTmpl := template.Must(
		template.New("ftest.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "ftest.go.tmpl")),
	)
	outFile := filepath.Join(tmpDir, "f_test.go")
	if err := generateFile(ftestTmpl, outFile, itestData{
		Prefix: "", ElemType: "float32",
	}); err != nil {
		t.Fatalf("failed to generate f_test.go: %v", err)
	}

	itestTmpl := template.Must(
		template.New("itest.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "itest.go.tmpl")),
	)
	outFile = filepath.Join(tmpDir, "i32_test.go")
	if err := generateFile(itestTmpl, outFile, itestData{
		Prefix: "I32", ElemType: "int32", Signed: true,
	}); err != nil {
		t.Fatalf("failed to generate i32_test.go: %v", err)
	}
}

func TestGenerateBenchmarks(t *testing.T) {
	dir := genDir()
	tmpDir := t.TempDir()

	benchTmpl := template.Must(
		template.New("bench.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "bench.go.tmpl")),
	)
	outFile := filepath.Join(tmpDir, "f32_bench_test.go")
	if err := generateFile(benchTmpl, outFile, benchData{
		Name: "F32", Prefix: "", ElemType: "float32", LenType: "float32",
		LerpT: "0.5", HasVecApproxEq: true, HasMatApproxEq: true,
		HasFloat32: false, HasMat3: true, Epsilon: "1e-6",
	}); err != nil {
		t.Fatalf("failed to generate f32_bench_test.go: %v", err)
	}
}
