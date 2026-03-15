package main

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"seq": func(n int) []int {
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		return s
	},
	"add":       func(a, b int) int { return a + b },
	"sub":       func(a, b int) int { return a - b },
	"mul":       func(a, b int) int { return a * b },
	"component": func(i int) string { return []string{"X", "Y", "Z", "W"}[i] },
	"param":     func(i int) string { return []string{"x", "y", "z", "w"}[i] },
	"last":      func(i, n int) bool { return i == n-1 },
	"join":      strings.Join,
	"intList": func(count, start int) string {
		parts := make([]string, count)
		for i := range parts {
			parts[i] = fmt.Sprintf("%d", start+i)
		}
		return strings.Join(parts, ", ")
	},
	"invMat": func(n int) string {
		switch n {
		case 2:
			return "1, 2, 3, 4"
		case 3:
			return "1, 2, 3, 0, 1, 4, 5, 6, 0"
		case 4:
			return "1, 0, 0, 1, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 1"
		}
		return ""
	},
}

type vecData struct {
	N int
}

type matData struct {
	N  int
	N2 int
}

type ivecData struct {
	N        int
	Prefix   string
	ElemType string
}

type imatData struct {
	N        int
	N2       int
	Prefix   string
	ElemType string
}

type itestData struct {
	Prefix   string
	ElemType string
	Signed   bool
}

type benchData struct {
	Name           string
	Prefix         string
	ElemType       string
	LenType        string
	LerpT          string
	HasVecApproxEq bool
	HasMatApproxEq bool
	HasFloat32     bool
	HasMat3        bool
	Epsilon        string
}

// genDir returns the directory containing this source file,
// which is where the .tmpl files live.
func genDir() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}

func generateFile(tmpl *template.Template, filename string, data interface{}) error {
	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("execute template for %s: %w", filename, err)
	}

	formatted, err := format.Source([]byte(buf.String()))
	if err != nil {
		return fmt.Errorf("format %s: %w\n\nRaw output:\n%s", filename, err, buf.String())
	}

	if err := os.WriteFile(filename, formatted, 0644); err != nil {
		return fmt.Errorf("write %s: %w", filename, err)
	}

	fmt.Printf("Generated %s\n", filename)
	return nil
}

func run(outDir string) error {
	dir := genDir()

	parseTmpl := func(name string) *template.Template {
		return template.Must(
			template.New(name).Funcs(funcMap).ParseFiles(filepath.Join(dir, name)),
		)
	}

	outPath := func(name string) string {
		return filepath.Join(outDir, name)
	}

	vecTmpl := parseTmpl("tvec.go.tmpl")
	matTmpl := parseTmpl("tmat.go.tmpl")
	ivecTmpl := parseTmpl("ivec.go.tmpl")
	imatTmpl := parseTmpl("imat.go.tmpl")
	itestTmpl := parseTmpl("itest.go.tmpl")
	quatTmpl := parseTmpl("tquat.go.tmpl")

	// Generate generic types.
	for _, n := range []int{2, 3, 4} {
		if err := generateFile(vecTmpl, outPath(fmt.Sprintf("vec%d.go", n)), vecData{N: n}); err != nil {
			return err
		}
	}

	for _, n := range []int{2, 3, 4} {
		if err := generateFile(matTmpl, outPath(fmt.Sprintf("mat%d.go", n)), matData{N: n, N2: n * n}); err != nil {
			return err
		}
	}

	if err := generateFile(quatTmpl, outPath("quat.go"), nil); err != nil {
		return err
	}

	// Generate floating-point concrete types (float32 and float64).
	fvecTmpl := parseTmpl("fvec.go.tmpl")
	fmatTmpl := parseTmpl("fmat.go.tmpl")
	ftestTmpl := parseTmpl("ftest.go.tmpl")
	fquatTmpl := parseTmpl("fquat.go.tmpl")

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
			if err := generateFile(fvecTmpl, outPath(fmt.Sprintf("%svec%d.go", ft.FilePrefix, n)), ivecData{
				N:        n,
				Prefix:   ft.Prefix,
				ElemType: ft.ElemType,
			}); err != nil {
				return err
			}
			if err := generateFile(fmatTmpl, outPath(fmt.Sprintf("%smat%d.go", ft.FilePrefix, n)), imatData{
				N:        n,
				N2:       n * n,
				Prefix:   ft.Prefix,
				ElemType: ft.ElemType,
			}); err != nil {
				return err
			}
		}

		if err := generateFile(fquatTmpl, outPath(fmt.Sprintf("%squat.go", ft.FilePrefix)), ivecData{
			N:        4,
			Prefix:   ft.Prefix,
			ElemType: ft.ElemType,
		}); err != nil {
			return err
		}

		if err := generateFile(ftestTmpl, outPath(fmt.Sprintf("%s_test.go", ft.FilePrefix)), itestData{
			Prefix:   ft.Prefix,
			ElemType: ft.ElemType,
		}); err != nil {
			return err
		}
	}

	// Generate integer concrete types.
	intTypes := []struct {
		Prefix string
		GoType string
		Signed bool
	}{
		{"I8", "int8", true},
		{"I16", "int16", true},
		{"I32", "int32", true},
		{"I64", "int64", true},
		{"U8", "uint8", false},
		{"U16", "uint16", false},
		{"U32", "uint32", false},
		{"U64", "uint64", false},
	}

	for _, it := range intTypes {
		prefix := strings.ToLower(it.Prefix)
		for _, n := range []int{2, 3, 4} {
			if err := generateFile(ivecTmpl, outPath(fmt.Sprintf("%svec%d.go", prefix, n)), ivecData{
				N:        n,
				Prefix:   it.Prefix,
				ElemType: it.GoType,
			}); err != nil {
				return err
			}

			if err := generateFile(imatTmpl, outPath(fmt.Sprintf("%smat%d.go", prefix, n)), imatData{
				N:        n,
				N2:       n * n,
				Prefix:   it.Prefix,
				ElemType: it.GoType,
			}); err != nil {
				return err
			}
		}

		// Generate test file for this integer type.
		if err := generateFile(itestTmpl, outPath(fmt.Sprintf("%s_test.go", prefix)), itestData{
			Prefix:   it.Prefix,
			ElemType: it.GoType,
			Signed:   it.Signed,
		}); err != nil {
			return err
		}
	}

	// Generate benchmark files.
	benchTmpl := parseTmpl("bench.go.tmpl")

	benchTypes := []benchData{
		{Name: "F32", Prefix: "", ElemType: "float32", LenType: "float32", LerpT: "0.5", HasVecApproxEq: true, HasMatApproxEq: true, HasFloat32: false, HasMat3: true, Epsilon: "1e-6"},
		{Name: "F64", Prefix: "D", ElemType: "float64", LenType: "float64", LerpT: "0.5", HasVecApproxEq: true, HasMatApproxEq: true, HasFloat32: true, HasMat3: true, Epsilon: "1e-10"},
	}

	for _, it := range intTypes {
		benchTypes = append(benchTypes, benchData{
			Name: it.Prefix, Prefix: it.Prefix, ElemType: it.GoType, LenType: "float32", LerpT: "0.5",
			HasVecApproxEq: false, HasMatApproxEq: false, HasFloat32: true, HasMat3: false, Epsilon: "",
		})
	}

	for _, bt := range benchTypes {
		filename := strings.ToLower(bt.Name) + "_bench_test.go"
		if err := generateFile(benchTmpl, outPath(filename), bt); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := run("."); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
