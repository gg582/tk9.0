// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tk9.0 is an idiomatic Go wrapper for [libtk9.0].
//
// Parts of the documentation are copied and/or modified from [TkDocs], see the
// LICENSE-TKDOCS file for details.
//
// [TkDocs]: https://tkdocs.com/about.html
// [libtk9.0]: https://pkg.go.dev/modernc.org/libtk9.0
package tk9_0 // import "modernc.org/tk9.0"

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	// "strings"
	"sync"
	"sync/atomic"

	"github.com/evilsocket/islazy/zip"
	lib "modernc.org/libtk9.0"
	tklib "modernc.org/libtk9.0/library"
	tcl "modernc.org/tcl9.0"
)

var (
	id atomic.Int32

	tk     *Tk
	tkErr  error
	tkOnce sync.Once

	tclDir string
	tkDir  string
)

// stdlib returns the path to the Tk standard library or an error, if any. It
// once creates a temporary directory where the standard library is written.
// Subsequent calls to stdlib share the same temporary directory.
//
// stdlib is safe for concurrent access by multiple goroutines.
func stdlib() (dir string, err error) {
	if dir, err = os.MkdirTemp("", "tk-library-"); err != nil {
		return "", err
	}

	fn := filepath.Join(dir, "library.zip")
	if err = os.WriteFile(fn, []byte(tklib.Zip), 0600); err != nil {
		return
	}

	if _, err = zip.Unzip(fn, dir); err != nil {
		return
	}

	return filepath.Join(dir, "library"), nil
}

// newInterp is like [tcl.NewInterp] but additionally initializes the Tk
// subsystem.
//
// [tcl.NewInterp]: https://pkg.go.dev/modernc.org/tcl9.0#NewInterp
func newInterp(tclvars map[string]string) (r *tcl.Interp, err error) {
	if r, err = tcl.NewInterp(tclvars); err != nil {
		return nil, err
	}

	if rc := lib.XTk_Init(r.TLS(), r.Handle()); rc != lib.TCL_OK {
		r.Close()
		return nil, fmt.Errorf("failed to initialize the Tk subsystem")
	}

	return r, nil
}

// Finalize releases all resources held, if any. Finalize is intended to be
// called on process shutdown only.
func Finalize() (err error) {
	if tk != nil {
		err = tk.in.Close()
		tk = nil
	}
	for _, v := range []string{tclDir, tkDir} {
		err = errors.Join(err, os.RemoveAll(v))
	}
	return err
}

// Tk represents the main window of an application. It has an associated Tcl
// interpreter.
type Tk struct {
	*Window
	in *tcl.Interp

	trace bool
}

func (tk *Tk) eval(s string) (r string, err error) {
	if tk.trace {
		defer func() {
			fmt.Fprintf(os.Stderr, "%s ->(r=%v err=%v)", s, r, err)
		}()
	}
	return tk.in.Eval(s, tcl.EvalGlobal)
}

// Initialize performs package initialization and returns a *Tk or error, if
// any.
//
// The returned value is a singleton. Calls to Initialize() are idempotent and
// all return the same (instance, error) tuple.
func Initialize() (r *Tk, err error) {
	tkOnce.Do(func() {
		if tclDir, tkErr = tcl.Stdlib(); err != nil {
			return
		}

		if tkDir, tkErr = stdlib(); tkErr != nil {
			return
		}

		var in *tcl.Interp
		if in, tkErr = newInterp(map[string]string{
			"tcl_library": tclDir,
			"tk_library":  tkDir,
		}); tkErr != nil {
			return
		}

		tk = &Tk{
			Window: &Window{},
			in:     in,
		}
	})
	return tk, tkErr
}

// Window represents a Tk window/widget.
type Window struct {
	fpath string
}

func (w *Window) path() (r string) {
	if r = w.fpath; r == "" {
		r = "."
	}
	return r
}

// func (w *Window) newChild(nm string, opts ...Opt) (*Window, error) {
// 	cls := strings.Replace(nm, "ttk_", "ttk::", 1)
// 	if c := nm[len(nm)-1]; c >= '0' && c <= '9' {
// 		nm += "_"
// 	}
// 	path := fmt.Sprintf("%s.%s%v", w.path(), nm, id.Add(1))
// 	r, err := tk.eval(fmt.Sprintf("%s %s", cls, path))
// 	if err != nil {
// 		return nil, fmt.Errorf("%v %v", err, r)
// 	}
//
// 	return &Window{fpath: r}, nil
// }
//
// // Opt represents a set of options.
// type Opt map[any]any
