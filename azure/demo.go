//go:build ignore

// $ go run -tags=tk.eval demo.go &
package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	. "modernc.org/tk9.0"
)

//go:embed Azure-ttk-theme-main.zip
var zip []byte

func fail(rc int, x any) {
	fmt.Fprintln(os.Stderr, x)
	os.Exit(rc)
}

func mustEval(s string, args ...any) (r string) {
	r, err := Eval(fmt.Sprintf(s, args...))
	if err != nil {
		fail(1, err)
	}

	return r
}

func main() {
	root, err := os.UserCacheDir()
	if err != nil {
		fail(1, err)
	}

	// Example path, but something like this could be used.
	dir := filepath.Join(root, "modernc.org", "tk9.0.0", "themes", "azure")
	if err := os.MkdirAll(dir, 0770); err != nil {
		fail(1, err)
	}

	pth := filepath.Join(dir, "azure-ttk-theme.zip")
	if err := os.WriteFile(pth, zip, 0660); err != nil {
		fail(1, err)
	}

	// Same story about example path.
	mount := "/themes/azure"
	mustEval("zipfs mount %s %s", pth, mount)
	mustEval("source //zipfs:%s/Azure-ttk-theme-main/azure.tcl", mount)
	Pack(TButton(Txt("Light"), Command(func() { mustEval("set_theme light") })),
		TButton(Txt("Dark"), Command(func() { mustEval("set_theme dark") })),
		TExit(),
		Pady("2m"), Ipady("1m"))
	App.Wait()
}
