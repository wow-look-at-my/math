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

func main() {
	dir := genDir()

	vecTmpl := template.Must(
		template.New("vec.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "vec.go.tmpl")),
	)
	matTmpl := template.Must(
		template.New("mat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "mat.go.tmpl")),
	)
	ivecTmpl := template.Must(
		template.New("ivec.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "ivec.go.tmpl")),
	)
	imatTmpl := template.Must(
		template.New("imat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "imat.go.tmpl")),
	)
	itestTmpl := template.Must(
		template.New("itest.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "itest.go.tmpl")),
	)

	// Generate generic + float concrete types.
	for _, n := range []int{2, 3, 4} {
		if err := generateFile(vecTmpl, fmt.Sprintf("vec%d.go", n), vecData{N: n}); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	for _, n := range []int{2, 3, 4} {
		if err := generateFile(matTmpl, fmt.Sprintf("mat%d.go", n), matData{N: n, N2: n * n}); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	// Generate float64 (double) concrete types.
	fvecTmpl := template.Must(
		template.New("fvec.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "fvec.go.tmpl")),
	)
	fmatTmpl := template.Must(
		template.New("fmat.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "fmat.go.tmpl")),
	)
	ftestTmpl := template.Must(
		template.New("ftest.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "ftest.go.tmpl")),
	)

	for _, n := range []int{2, 3, 4} {
		if err := generateFile(fvecTmpl, fmt.Sprintf("dvec%d.go", n), ivecData{
			N:        n,
			Prefix:   "D",
			ElemType: "float64",
		}); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if err := generateFile(fmatTmpl, fmt.Sprintf("dmat%d.go", n), imatData{
			N:        n,
			N2:       n * n,
			Prefix:   "D",
			ElemType: "float64",
		}); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	if err := generateFile(ftestTmpl, "d_test.go", itestData{
		Prefix:   "D",
		ElemType: "float64",
	}); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
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
			if err := generateFile(ivecTmpl, fmt.Sprintf("%svec%d.go", prefix, n), ivecData{
				N:        n,
				Prefix:   it.Prefix,
				ElemType: it.GoType,
			}); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			if err := generateFile(imatTmpl, fmt.Sprintf("%smat%d.go", prefix, n), imatData{
				N:        n,
				N2:       n * n,
				Prefix:   it.Prefix,
				ElemType: it.GoType,
			}); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		}

		// Generate test file for this integer type.
		if err := generateFile(itestTmpl, fmt.Sprintf("%s_test.go", prefix), itestData{
			Prefix:   it.Prefix,
			ElemType: it.GoType,
			Signed:   it.Signed,
		}); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	// Generate benchmark files.
	benchTmpl := template.Must(
		template.New("bench.go.tmpl").Funcs(funcMap).ParseFiles(filepath.Join(dir, "bench.go.tmpl")),
	)

	benchTypes := []struct {
		Prefix         string
		ElemType       string
		LenType        string
		HasVecApproxEq bool
		HasMatApproxEq bool
		HasFloat32     bool
		HasMat3        bool
		Epsilon        string
		Filename       string
	}{
		{"", "float32", "float32", true, true, false, true, "1e-6", "f32_bench_test.go"},
		{"D", "float64", "float64", true, true, true, true, "1e-10", "d_bench_test.go"},
	}

	for _, it := range intTypes {
		benchTypes = append(benchTypes, struct {
			Prefix         string
			ElemType       string
			LenType        string
			HasVecApproxEq bool
			HasMatApproxEq bool
			HasFloat32     bool
			HasMat3        bool
			Epsilon        string
			Filename       string
		}{it.Prefix, it.GoType, "float32", false, false, true, false, "", strings.ToLower(it.Prefix) + "_bench_test.go"})
	}

	for _, bt := range benchTypes {
		if err := generateFile(benchTmpl, bt.Filename, benchData{
			Prefix:         bt.Prefix,
			ElemType:       bt.ElemType,
			LenType:        bt.LenType,
			LerpT:          "0.5",
			HasVecApproxEq: bt.HasVecApproxEq,
			HasMatApproxEq: bt.HasMatApproxEq,
			HasFloat32:     bt.HasFloat32,
			HasMat3:        bt.HasMat3,
			Epsilon:        bt.Epsilon,
		}); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}
