package failures

import (
	"encoding/json"
	"gotypes/lessons/reference/conversions"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Каждый неверный пример компилируется отдельно: причина отказа проверяется,
// а обычная сборка проекта остаётся рабочей.
func TestCompileErrors(t *testing.T) {
	raw, err := os.ReadFile("testdata/cases.json")
	if err != nil {
		t.Fatal(err)
	}
	var cases []struct {
		File  string
		Error string
	}
	if err := json.Unmarshal(raw, &cases); err != nil {
		t.Fatal(err)
	}
	_, file, _, _ := runtime.Caller(0)
	root := filepath.Dir(filepath.Dir(filepath.Dir(file)))
	for _, tc := range cases {
		t.Run(tc.File, func(t *testing.T) {
			dir := t.TempDir()
			source, err := os.ReadFile(filepath.Join("testdata", tc.File))
			if err != nil {
				t.Fatal(err)
			}
			module := "module invalid\n\ngo 1.23.0\nrequire gotypes v0.0.0\nreplace gotypes => " + filepath.ToSlash(root) + "\n"
			for name, data := range map[string][]byte{"go.mod": []byte(module), "invalid.go": source} {
				if err := os.WriteFile(filepath.Join(dir, name), data, 0600); err != nil {
					t.Fatal(err)
				}
			}
			cmd := exec.Command(filepath.Join(runtime.GOROOT(), "bin", "go"), "build", ".")
			cmd.Dir = dir
			cmd.Env = append(os.Environ(), "GOWORK=off", "GOTOOLCHAIN=local")
			output, err := cmd.CombinedOutput()
			if err == nil || !strings.Contains(string(output), tc.Error) {
				t.Fatalf("expected %q; err=%v\n%s", tc.Error, err, output)
			}
		})
	}
}

func TestExpectedPanics(t *testing.T) {
	cases := map[string]func(){
		"assertion":            func() { var x interface{} = 42; _ = x.(int64) },
		"array pointer length": func() { s := []int{1}; _ = (*[2]int)(s) },
		"array value length":   func() { s := []int{1}; _ = [2]int(s) },
		"interface comparison": func() { conversions.Equal[any]([]int{1}, []int{1}) },
		"nil callback":         func() { var fn func(int) bool; fn(0) },
	}
	for name, run := range cases {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("expected panic")
				}
			}()
			run()
		})
	}
}
